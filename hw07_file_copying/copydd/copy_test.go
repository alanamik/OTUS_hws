package hw07filecopying

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCopy(t *testing.T) {
	t.Run("invalid_from_path", func(t *testing.T) {
		require.EqualError(t, Copy("/test", "", 0, 0), ErrInvalidPath.Error())
	})
	t.Run("invalid_to_path", func(t *testing.T) {
		require.EqualError(t, Copy("testdata/input.txt", "test/test.txt", 0, 0), ErrInvalidPath.Error())
	})
	t.Run("work in tmp", func(t *testing.T) {
		err := Copy("testdata/input.txt", "tmp/test.txt", 0, 0)
		require.FileExists(t, "tmp/test.txt")
		require.NoError(t, err)
	})
	t.Run("input file is empty", func(t *testing.T) {
		require.EqualError(t, Copy("testdata/empty.txt", "tmp/empty.txt", 0, 0), ErrEmptyFile.Error())
	})
	t.Run("one symbol from the start", func(t *testing.T) {
		testFile, _ := os.Open("testdata/test.txt")
		stat, _ := testFile.Stat()
		testSize := stat.Size()
		tt := testSize - 1
		err := Copy("testdata/test.txt", "tmp/test1.txt", 0, testSize-tt)
		require.NoError(t, err)
		out, _ := os.ReadFile("tmp/test1.txt")
		str := string(out)
		require.Equal(t, str, "0")
	})
	t.Run("one symbol from in the end", func(t *testing.T) {
		testFile, _ := os.Open("testdata/test.txt")
		in, _ := testFile.Stat()
		testSize := in.Size()
		fmt.Println(testSize)
		err := Copy("testdata/test.txt", "tmp/test2.txt", testSize-1, testSize)
		require.NoError(t, err)
		out, _ := os.ReadFile("tmp/test2.txt")
		str := string(out)
		require.Equal(t, str, "9")
	})
}
