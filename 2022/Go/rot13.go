package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func rot13(x byte) byte {
	switch {
	case x >= 65 && x <= 77:
		fallthrough
	case x >= 97 && x <= 109:
		x = x + 13
	case x >= 78 && x <= 90:
		fallthrough
	case x >= 110 && x <= 122:
		x = x - 13
	}
	return x
}

func (self *rot13Reader) Read(buff []byte) (int, error) {
	n, err := self.r.Read(buff)
	for i := 0; i <= n; i++ {
		buff[i] = rot13(buff[i])
	}
	return n, err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
	println()
}
