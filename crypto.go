package export

import (
	"crypto/rand"
	"errors"
	"golang.org/x/crypto/nacl/secretbox"
	"io"
)

const (
	// From the secretbox example code.
	nonceSize = 24

	// From `go doc secretbox`:
	//  If in doubt, 16KB is a reasonable chunk size.
	maxChunkSize = 16 * 1024
)

func Encrypt(in io.Reader, inLen int64, out io.Writer, key *[32]byte) (int, error) {
	nwritten := 0
	remaining := inLen

	for remaining > 0 {
		// From `go doc secretbox`:
		//   You must use a different nonce for each message you encrypt
		//   with the same key. Since the nonce here is 192 bits long, a
		//   random value provides a sufficiently small probability of
		//   repeats.
		var nonce [nonceSize]byte
		if _, err := io.ReadFull(rand.Reader, nonce[:]); err != nil {
			return nwritten, err
		}

		chunckLen := min(maxChunkSize, remaining)
		remaining -= chunckLen

		plain := make([]byte, chunckLen)
		_, err := io.ReadFull(in, plain)
		if err != nil {
			return nwritten, err
		}

		cipher := secretbox.Seal(nonce[:], plain, &nonce, key)
		n, err := out.Write(cipher)
		if err != nil {
			return nwritten, err
		}

		nwritten += n
	}

	return nwritten, nil
}

func Decrypt(in io.Reader, inLen int64, out io.Writer, key *[32]byte) (int, error) {
	nwritten := 0
	remaining := inLen

	for remaining > nonceSize {
		var nonce [nonceSize]byte
		if _, err := io.ReadFull(in, nonce[:]); err != nil {
			return nwritten, err
		}
		remaining -= nonceSize

		chunckLen := min(maxChunkSize, remaining)
		remaining -= chunckLen

		chunck := make([]byte, chunckLen)
		if _, err := io.ReadFull(in, chunck); err != nil {
			return nwritten, err
		}

		decrypted, ok := secretbox.Open(nil, chunck, &nonce, key)
		if !ok {
			return nwritten, errors.New("decryption error")
		}
		n, err := out.Write(decrypted)
		if err != nil {
			return nwritten, err
		}

		nwritten += n
	}

	return nwritten, nil
}
