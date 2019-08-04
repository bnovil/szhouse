package main

import (
	"os"
)

func main() {

	// open output file
	fo, err := os.Create("output.txt")
	if err != nil {
		panic(err)
	}
	// close fo on exit and check for its returned error
	defer func() {
		if err := fo.Close(); err != nil {
			panic(err)
		}
	}()

	fo.WriteString("string test")

	// // make a buffer to keep chunks that are read
	// buf := make([]byte, 1024)

	// buf = []byte("test")

	// // // read a chunk
	// // n, err := fi.Read(buf)
	// // if err != nil && err != io.EOF {
	// // 	panic(err)
	// // }

	// // write a chunk
	// if _, err := fo.Write(buf); err != nil {
	// 	panic(err)
	// }

}
