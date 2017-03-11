package main

import (
	"bytes"
)

func main() {
	buf2 := []byte("12")
	buf4 := []byte("3456")
	buf3 := make([]byte, 3)

	var buf bytes.Buffer
	buf.Write(buf2)
	buf.Write(buf4)

	buf.Read(buf3)
	println(string(buf3))
	buf.Read(buf3)
	println(string(buf3))

	buf.Write(buf2)
	buf.Write(buf4)

	buf.Read(buf3)
	println(string(buf3))
	buf.Read(buf3)
	println(string(buf3))
}
