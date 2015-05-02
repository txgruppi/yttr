package yttr

import (
	"errors"
	"io"
	"mime"
	"os"
	"path"
	"strings"
)

var (
	FileNotSupportedErr = errors.New("The file format is not supported")
	InvalidDaysErr      = errors.New("Days must be 1, 4 or 12")
)

type File interface {
	io.Reader
	Name() string
	Type() string
	Size() Size
	Days() Days
	DownloadOnly() Bool
}

func NewFile(reader io.Reader, name, mime string, size int64, days int, download bool) File {
	return &file{
		name:     name,
		mime:     mime,
		size:     size,
		days:     days,
		download: download,
		reader:   reader,
	}
}

func NewFileFromPath(filepath string, days int, download bool) (File, error) {
	if days != 1 && days != 4 && days != 12 {
		return nil, InvalidDaysErr
	}

	info, err := os.Stat(filepath)
	if err != nil {
		return nil, err
	}

	t := mime.TypeByExtension(path.Ext(filepath))
	if t == "" {
		return nil, FileNotSupportedErr
	}

	if strings.IndexRune(t, ';') != -1 {
		t = strings.SplitN(t, ";", 2)[0]
	}

	reader, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}

	return NewFile(reader, info.Name(), t, info.Size(), days, download), nil
}

type file struct {
	name     string
	mime     string
	size     int64
	days     int
	download bool
	reader   io.Reader
}

func (f *file) Size() Size {
	return Size(f.size)
}

func (f *file) Type() string {
	return f.mime
}

func (f *file) Name() string {
	return f.name
}

func (f *file) Days() Days {
	return Days(f.days)
}

func (f *file) DownloadOnly() Bool {
	return Bool(f.download)
}

func (f *file) Read(p []byte) (int, error) {
	return f.reader.Read(p)
}
