package main

import (
	"flag"
)

var (
	from, to      string
	limit, offset int64
)

func init() {
	flag.StringVar(&from, "from", "", "file to read from")
	flag.StringVar(&to, "to", "", "file to write to")
	flag.Int64Var(&limit, "limit", 0, "limit of bytes to copy")
	flag.Int64Var(&offset, "offset", 0, "offset in input file")
}

func main() {
	//	flag.Parse()
	from := "testdata/input.txt"
	to := "out.txt"
	//limit:= 10
	Copy(from, to, 0, 0)
	//hw07_file_copying.Copy(from, to, offset, limit)
}
