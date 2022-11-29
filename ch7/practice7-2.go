package main

import (
	"fmt"
	"io"
	"os"
)

type NewWriter struct{
	baseWriter io.Writer
	nBytes int64
}
func (w *NewWriter) Write(p []byte) (int, error) {
	n, error := w.baseWriter.Write(p)
	w.nBytes += int64(n)
	return n, error
}

func CountingWriter(w io.Writer) (io.Writer, *int64) {
	var writer NewWriter
	writer.baseWriter = w

	return &writer, &(writer.nBytes)
}

func main() {
	writer, nBytes := CountingWriter(os.Stdout)
	fmt.Fprintf(writer, "hello %s\n", "test")
	fmt.Println("write bytes:", *nBytes)
}
