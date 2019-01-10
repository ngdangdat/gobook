package main

import (
	"bytes"
	"fmt"
	"io"
)

var debug = false

func main() {
	testBuf := new(bytes.Buffer)
	var buf *bytes.Buffer
	if debug {
		buf = new(bytes.Buffer)
	}
	buf = testBuf
	// f(buf)
	fmt.Println(*buf)
}

func isNil(out io.Writer) bool {
	if out == nil {
		return true
	}
	return false
}

func f(out io.Writer) {
	if out != nil {
		out.Write([]byte("Done!"))
	}
}
