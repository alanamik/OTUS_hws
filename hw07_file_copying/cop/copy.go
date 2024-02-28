package copy

import (
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	pb "github.com/cheggaaa/pb/v3"
)

var (
	ErrUnsupportedFile       = errors.New("unsupported file")
	ErrOffsetExceedsFileSize = errors.New("offset exceeds file size")
	ErrInvalidPath           = errors.New("wrong path to the file")
	ErrEmptyFile             = errors.New("the file is empty")
	ErrUnknownLength         = errors.New("the length of the file is unknown")
)

func Copy(fromPath, toPath string, offset, limit int64) error {
	in, err := os.Open(fromPath)
	if err != nil {
		return ErrInvalidPath
	}
	defer func() {
		cerr := in.Close()
		if cerr != nil {
			err = cerr
		}
	}()

	r, err := ioutil.ReadFile(fromPath)
	type_ := http.DetectContentType(r)
	if !strings.Contains(type_, "text/plain") {
		return ErrUnsupportedFile
	}

	stat, err := in.Stat()
	size := stat.Size()
	if size == 0 {
		return ErrEmptyFile
	}
	if limit > size {
		limit = size
	}

	out, err := os.Create(toPath)
	if err != nil {
		return ErrInvalidPath
	}
	defer func() {
		cerr := out.Close()
		if err == nil {
			err = cerr
		}
	}()

	bar := pb.Start64(size)
	barReader := bar.NewProxyReader(in)
	defer bar.Finish()

	if offset == 0 && limit == 0 {
		_, err := io.Copy(out, barReader)
		if err != nil {
			return err
		}
		return nil
	}

	if stat.Size() > offset && stat.Size() < limit {
		return ErrOffsetExceedsFileSize
	}

	if offset >= 0 && limit > 0 {
		_, err = in.Seek(offset, io.SeekStart)
		if err != nil {
			return err
		}
		_, err := io.CopyN(out, barReader, limit)
		if err != nil && err != io.EOF {
			return err
		}
		return nil
	}

	err = out.Sync()
	bar.Finish()
	return nil
}
