package hw07filecopying

import (
	"errors"
	"io"
	"os"

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

	if stat.Size() > offset && stat.Size() < limit {
		return ErrOffsetExceedsFileSize
	}

	if offset == 0 && limit == 0 {
		bar := pb.Start64(size)
		barReader := bar.NewProxyReader(in)

		_, err := io.Copy(out, barReader)
		if err != nil {
			return err
		}
		bar.Finish()
		return nil
	}

	if offset >= 0 && limit > 0 {
		ret, err := in.Seek(offset, io.SeekStart)
		if err != nil {
			return err
		}
		bar := pb.Full.Start64(limit - ret)
		barReader := bar.NewProxyReader(in)

		_, err = io.CopyN(out, barReader, limit)
		if err != nil && !errors.Is(err, io.EOF) {
			return err
		}
		bar.Finish()
		return nil
	}
	err = out.Sync()

	return nil
}
