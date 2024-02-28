package main

import (
	copy "OTUS_hws/alanamik/hw07_file_copying/cop"
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCopy(t *testing.T) {
	t.Run("invalid_from_path", func(t *testing.T) {
		require.EqualError(t, copy.Copy("/test", "", 0, 0), copy.ErrInvalidPath.Error())
	})
	t.Run("invalid_to_path", func(t *testing.T) {
		require.EqualError(t, copy.Copy("testdata/input.txt", "test/test.txt", 0, 0), copy.ErrInvalidPath.Error())
	})

	t.Run("work in tmp", func(t *testing.T) {
		err := copy.Copy("testdata/input.txt", "tmp/test.txt", 0, 0)
		require.FileExists(t, "tmp/test.txt")
		require.NoError(t, err)
	})

	// image, mp4, etc

	t.Run("unsupported file", func(t *testing.T) {
		require.EqualError(t, copy.Copy("testdata/image.jpg", "out.txt", 0, 0), copy.ErrUnsupportedFile.Error())
	})

	t.Run("input file is empty", func(t *testing.T) {
		require.EqualError(t, copy.Copy("testdata/empty.txt", "tmp/empty.txt", 0, 0), copy.ErrEmptyFile.Error())
	})

	t.Run("one symbol from the start", func(t *testing.T) {

		test_file, _ := os.Open("testdata/test.txt")
		stat, _ := test_file.Stat()
		test_size := stat.Size()
		tt := test_size - 1
		err := copy.Copy("testdata/test.txt", "tmp/test1.txt", 0, test_size-tt)
		require.NoError(t, err)
		out, _ := os.ReadFile("tmp/test1.txt")
		str := string(out)
		require.Equal(t, str, "0")
	})

	t.Run("one symbol from in the end", func(t *testing.T) {

		f, _ := os.Open("testdata/test.txt")
		in, _ := f.Stat()
		test_size := in.Size()
		fmt.Println(test_size)
		err := copy.Copy("testdata/test.txt", "tmp/test2.txt", test_size-1, test_size)
		require.NoError(t, err)
		out, _ := os.ReadFile("tmp/test2.txt")
		str := string(out)
		require.Equal(t, str, "9")

	})
}
