package main

import (
	"flag"
	"fmt"
	"io"
	"os"
)

const BUFFER_SIZE = 512

func main() {
	flag.Parse()
	if flag.NArg() != 2 {
		arg0, _ := os.Executable()
		fmt.Println("Usage:", arg0, "[in_file] [out_file]")
		os.Exit(-1)
	}

	in_file_name := flag.Arg(0)
	out_file_name := flag.Arg(1)

	in_file, err := os.OpenFile(in_file_name, os.O_RDONLY, 0)
	if err != nil {
		fmt.Println("Can't open file", in_file_name)
		os.Exit(-2)
	}
	defer in_file.Close()
	out_file, err := os.OpenFile(out_file_name, os.O_CREATE | os.O_TRUNC | os.O_WRONLY, 0777)
	if err != nil {
		fmt.Println("Can't create file", out_file_name)
		os.Exit(-3)
	}
	defer out_file.Close()

	in_buf := make([]byte, BUFFER_SIZE)
	out_buf := make([]byte, BUFFER_SIZE)

	remaining, _ := in_file.Seek(0, io.SeekEnd)
	for remaining > 0 {
		var to_read int64
		if remaining > BUFFER_SIZE {
			to_read = BUFFER_SIZE
		} else {
			to_read = remaining
		}
		in_file.Seek(-to_read, io.SeekCurrent)
		read, _ := in_file.Read(in_buf[:to_read])
		for i := 0; i < read; i++ {
			out_buf[i] = in_buf[read-i-1]
		}
		out_file.Write(out_buf[:read])
		in_file.Seek(-int64(read), io.SeekCurrent)
		remaining = remaining - int64(read)
	}
}
