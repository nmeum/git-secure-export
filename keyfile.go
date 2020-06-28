package export

import (
	"crypto/rand"
	"io"
	"os"
)

const (
	KeySize = 32
)

func CreateKey(path string) error {
	var key [KeySize]byte
	_, err := rand.Read(key[:])
	if err != nil {
		return err
	}

	file, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_EXCL, 0600)
	if err != nil {
		return err
	}
	defer file.Close()

	// Write and reset offset for read in ReadKey() function.
	_, err = file.Write(key[:])
	if err != nil {
		return err
	}

	return nil
}

func ReadKey(path string) (*[KeySize]byte, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var key [KeySize]byte
	_, err = io.ReadFull(file, key[:])
	if err != nil {
		return nil, err
	}

	return &key, nil
}
