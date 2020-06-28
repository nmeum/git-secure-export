package main

import (
	"io"
	"log"
	"os"
	"path/filepath"

	git "github.com/nmeum/git-secure-export"
	cmd "github.com/nmeum/git-secure-export/cmd"
)

var key *[git.KeySize]byte

func cryptFn(in io.Reader, inLen int64, out io.Writer) (int, error) {
	return git.Decrypt(in, inLen, out, key)
}

func main() {
	log.SetFlags(log.Lshortfile)

	gitDir, err := git.GetDir()
	if err != nil {
		log.Fatal(err)
	}

	keyPath := filepath.Join(gitDir, "git-secure-key")
	key, err = git.ReadKey(keyPath, false)
	if err != nil {
		log.Fatal(err)
	}

	if err = cmd.Handle(os.Stdin, os.Stdout, cryptFn); err != nil {
		log.Fatal(err)
	}
}
