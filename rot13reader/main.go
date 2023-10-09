package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (rt rot13Reader) Read(p []byte) (int, error) {
	n, err := rt.r.Read(p)
	for i := 0; i < n; i++ {

		if p[i] >= 'A' && p[i] <= 'Z' {
			p[i] = 65 + (((p[i] - 65) + 13) % 26)
		} else if p[i] >= 'a' && p[i] <= 'z' {
			p[i] = 97 + (((p[i] - 97) + 13) % 26)
		}

	}
	return n, err

}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
