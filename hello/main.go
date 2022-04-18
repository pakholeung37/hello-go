package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (reader rot13Reader) Read(a []byte) (int, error) {
	n, err := reader.r.Read(a)
	for i := 0; i < n; i++ {
		if a[i] >= 'A' && a[i] <= 'Z' {
			a[i] = 'A' + ((a[i] - 'A' + 13) % 26)
		} else if a[i] >= 'a' && a[i] <= 'z' {
			a[i] = 'a' + ((a[i] - 'a' + 13) % 26)
		}
	}
	return n, err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
