package yttr_test

import (
	"io/ioutil"
	"os"
	"path"
	"testing"

	"github.com/txgruppi/yttr"
)

var (
	f            = yttr.NewFile(nil, "testing.txt", "text/plain", 123, 4, true)
	fp yttr.File = nil
)

func TestFileName(t *testing.T) {
	equal(t, "testing.txt", f.Name())
}

func TestFileType(t *testing.T) {
	equal(t, "text/plain", f.Type())
}

func TestFileSize(t *testing.T) {
	equal(t, yttr.Size(123), f.Size())
}

func TestFileDays(t *testing.T) {
	equal(t, yttr.Days(4), f.Days())
}

func TestFileDownloadOnly(t *testing.T) {
	equal(t, yttr.Bool(true), f.DownloadOnly())
}

func TestFileFromPathName(t *testing.T) {
	equal(t, "test.txt", fp.Name())
}

func TestFileFromPathType(t *testing.T) {
	equal(t, "text/plain", fp.Type())
}

func TestFileFromPathSize(t *testing.T) {
	equal(t, yttr.Size(16), fp.Size())
}

func TestFileFromPathDays(t *testing.T) {
	equal(t, yttr.Days(12), fp.Days())
}

func TestFileFromPathDownloadOnly(t *testing.T) {
	equal(t, yttr.Bool(false), fp.DownloadOnly())
}

func TestFileFromPathUnknownMime(t *testing.T) {
	p, err := createFile("something.unknown", "Just another file")
	equal(t, nil, err)

	f, err := yttr.NewFileFromPath(p, 4, false)
	equal(t, nil, f)
	equal(t, yttr.FileNotSupportedErr, err)
}

func TestFileWithInvalidDays(t *testing.T) {
	p, err := createFile("another-test.txt", "One more test")
	equal(t, nil, err)
	f, err := yttr.NewFileFromPath(p, 3, false)
	equal(t, nil, f)
	equal(t, yttr.InvalidDaysErr, err)
}

func TestFileFromPathNotExists(t *testing.T) {
	f, err := yttr.NewFileFromPath("/some/invalid/test/path", 1, false)
	equal(t, nil, f)
	notEqual(t, nil, err)
}

func TestFileFromPathRead(t *testing.T) {
	contents, err := ioutil.ReadAll(fp)
	equal(t, nil, err)
	equal(t, "Just a test file", string(contents))
}

func init() {
	p, err := createFile("test.txt", "Just a test file")
	if err != nil {
		panic(err)
	}

	fp, err = yttr.NewFileFromPath(p, 12, false)
	if err != nil {
		panic(err)
	}
}

func createFile(name, content string) (string, error) {
	tmp := os.TempDir()

	p := path.Join(tmp, name)

	out, err := os.Create(p)
	if err != nil {
		return "", err
	}

	_, err = out.WriteString(content)
	out.Close()
	if err != nil {
		return "", err
	}

	return p, nil
}
