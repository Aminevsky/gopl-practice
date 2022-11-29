package main

import (
	"io"
	"log"
	"os"
	"strings"
)

type NewReader struct{
	base io.Reader
	limit int64
	readBytes int64
}
func (r *NewReader) Read(p []byte) (int, error) {
	if r.readBytes >= r.limit {
		return 0, io.EOF
	}

	n, err := r.base.Read(p[:r.limit])
	r.readBytes += int64(n)
	return n, err
}

func LimitReader(r io.Reader, n int64) io.Reader {
	var reader NewReader
	reader.base = r
	reader.limit = n

	return &reader
}

func main() {
	r := strings.NewReader("some io.Reader stream to be read\n")
	lr := LimitReader(r, 4)

	if _, err := io.Copy(os.Stdout, lr); err != nil {
		log.Fatal(err)
	}

	//res, err := lr.Read(make([]byte, 4))
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//fmt.Println(res)
}
