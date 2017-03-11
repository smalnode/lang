package main

import (
	"bytes"
	"io"
	"os"
)

func Encode(dst, src []byte) int {
	if len(src) == 0 {
		return 0
	}

	for i := 0; i < len(src); i++ {
		b := int(src[i])
		switch {
		case 65 <= b && b <= 90:
			dst[i] = src[i] + 32
		case 97 <= b && b <= 122:
			dst[i] = src[i] - 32
		default:
			dst[i] = src[i]
		}
	}

	return len(src)
}

func Decode(dst, src []byte) int {
	return Encode(dst, src)
}

type encoder struct {
	err  error
	w    io.Writer
	nbuf int
	buf  [1024]byte
}

func NewEncoder(w io.Writer) io.WriteCloser {
	return &encoder{w: w}
}

func (e *encoder) Write(p []byte) (n int, err error) {
	if e.err != nil {
		return 0, e.err
	}

	for len(p) > 0 {
		pl := len(p)
		nl := 1024 - e.nbuf
		if pl > nl {
			Encode(e.buf[e.nbuf:], p[:nl])
			p = p[nl:]
			if _, err := e.w.Write(e.buf[0:1024]); err != nil {
				e.err = err
				return n, err
			}
			n += nl
			e.nbuf = 0
		} else {
			Encode(e.buf[e.nbuf:], p[:pl])
			p = p[pl:]
			e.nbuf = pl
		}
	}
	return n, nil
}

func (e *encoder) Close() error {
	if e.err == nil && e.nbuf > 0 {
		_, e.err = e.w.Write(e.buf[:e.nbuf])
		e.nbuf = 0
	}

	return e.err
}

func main() {
	var outbuf bytes.Buffer
	e := NewEncoder(&outbuf)
	buf := make([]byte, 48)
	for i := 0; i < len(buf); i++ {
		buf[i] = 'A'
	}
	if _, err := e.Write(buf); err != nil {
		panic(err)
	}
	e.Close()

	println(outbuf.Len())
	outbuf.WriteTo(os.Stdout)
}
