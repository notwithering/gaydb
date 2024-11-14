package gaydb

import (
	"errors"
	"io"
	"os"

	"github.com/vmihailenco/msgpack/v5"
)

// Get opens filename, reads, and then decodes the data to the value pointed to by v.
// If the file doesn't exist, it will transfer to function Put.
func Get(filename string, v interface{}) error {
	file, err := os.OpenFile(filename, os.O_RDONLY, os.ModePerm)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return Put(filename, v)
		}

		return err
	}
	defer file.Close()

	b, err := io.ReadAll(file)
	if err != nil {
		return err
	}

	return msgpack.Unmarshal(b, v)
}

// Put places the encoding v into filename.
// If the file doesn't exist, it will be created.
func Put(filename string, v interface{}) error {
	file, err := os.OpenFile(filename, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, os.ModePerm)
	if err != nil {
		return err
	}
	defer file.Close()

	b, err := msgpack.Marshal(&v)
	if err != nil {
		return err
	}

	_, err = file.Write(b)
	return err
}
