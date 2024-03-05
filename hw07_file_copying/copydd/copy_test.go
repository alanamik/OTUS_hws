package hw07filecopying

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCopy(t *testing.T) {
	os.Chdir("..")
	tmp := os.TempDir()

	t.Run("invalid_from_path", func(t *testing.T) {
		require.EqualError(t, Copy("/test", "out.txt", 0, 0), ErrInvalidPath.Error())
	})
	t.Run("invalid_to_path", func(t *testing.T) {
		require.EqualError(t, Copy("testdata/input.txt", "test/test.txt", 0, 0), ErrInvalidPath.Error())
	})
	t.Run("work in tmp", func(t *testing.T) {
		err := Copy("testdata/input.txt", fmt.Sprintf("%s/%s", tmp, "out.txt"), 0, 0)
		require.FileExists(t, fmt.Sprintf("%s/%s", tmp, "out.txt"))
		require.NoError(t, err)
	})
	t.Run("input file is empty", func(t *testing.T) {
		require.EqualError(t, Copy("testdata/empty.txt", fmt.Sprintf("%s/%s", tmp, "empty.txt"), 0, 0), ErrEmptyFile.Error())
	})
	t.Run("one symbol from the start", func(t *testing.T) {
		testFile, _ := os.Open("testdata/test.txt")
		stat, _ := testFile.Stat()
		testSize := stat.Size()
		tt := testSize - 1
		err := Copy("testdata/test.txt", fmt.Sprintf("%s/%s", tmp, "test1.txt"), 0, testSize-tt)
		require.NoError(t, err)
		out, _ := os.ReadFile(fmt.Sprintf("%s/%s", tmp, "test1.txt"))
		str := string(out)
		require.Equal(t, "0", str)
	})
	t.Run("one symbol from in the end", func(t *testing.T) {
		testFile, _ := os.Open("testdata/test.txt")
		in, _ := testFile.Stat()
		testSize := in.Size()
		fmt.Println(testSize)
		err := Copy("testdata/test.txt", fmt.Sprintf("%s/%s", tmp, "test2.txt"), testSize-1, testSize)
		require.NoError(t, err)
		out, _ := os.ReadFile(fmt.Sprintf("%s/%s", tmp, "test2.txt"))
		str := string(out)
		require.Equal(t, "9", str)
	})
}
