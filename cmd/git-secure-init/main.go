package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	git "github.com/nmeum/git-secure-export"
)

func main() {
	log.SetFlags(log.Lshortfile)

	gitDir, err := git.GetDir()
	if err != nil {
		log.Fatal(err)
	}
	keyPath := filepath.Join(gitDir, "git-secure-key")

	_, err = os.Stat(keyPath)
	if !os.IsNotExist(err) {
		log.Fatalf("key file %q already exists\n", keyPath)
	}
	err = git.CreateKey(keyPath)
	if err != nil {
		log.Fatal("key creation failed:", err)
	}

	fmt.Printf("Initialized symmetric key in %s\n", keyPath)
}
