package export

import (
	"crypto/rand"
	"io/ioutil"
	"os"
)

const (
	KeySize = 32
)

func createKeyFile(path string) (*os.File, error) {
	var key [KeySize]byte
	_, err := rand.Read(key[:])
	if err != nil {
		return nil, err
	}

	file, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_EXCL, 0600)
	if err != nil {
		return nil, err
	}

	// Write and reset offset for read in ReadKey() function.
	_, err = file.Write(key[:])
	if err != nil {
		file.Close()
		return nil, err
	}
	_, err = file.Seek(0, 0)
	if err != nil {
		file.Close()
		return nil, err
	}

	return file, nil
}

func ReadKey(path string) ([]byte, error) {
	file, err := os.Open(path)
	if os.IsNotExist(err) {
		file, err = createKeyFile(path)
	}
	if err != nil {
		return nil, err
	}
	defer file.Close()

	key, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	return key, nil
}
