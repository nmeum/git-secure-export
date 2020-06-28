package cmd

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
)

type CryptFn func(in io.Reader, inLen int64, out io.Writer) (int, error)

func Handle(in io.Reader, out io.Writer, fn CryptFn) error {
	reader := bufio.NewReader(in)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}

		var numBytes uint
		n, _ := fmt.Sscanf(line, "data %d\n", &numBytes)
		if n != 1 {
			fmt.Fprint(out, line)
			continue
		}

		buf := make([]byte, numBytes)
		writer := bytes.NewBuffer(buf)
		n, err = fn(reader, int64(numBytes), writer)
		if err != nil {
			return err
		}

		fmt.Fprintf(out, "data %d\n", n)
		_, err = writer.WriteTo(out)
		if err != nil {
			return err
		}

		// TODO: Handle optional terminating newline
	}

	return nil
}
