package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (rr rot13Reader) Read(bytes []byte) (int, error) {
	n, err := rr.r.Read(bytes)
	for i := range bytes {
		shift := 0

		if (bytes[i] > 'a' && bytes[i] < 'n') || (bytes[i] > 'A' && bytes[i] < 'N') {
			shift = 13
		}

		if (bytes[i] >= 'n' && bytes[i] <= 'z') || (bytes[i] >= 'N' && bytes[i] <= 'Z') {
			shift = -13
		}
		bytes[i] = bytes[i] + byte(shift)
	}

	return n, err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
