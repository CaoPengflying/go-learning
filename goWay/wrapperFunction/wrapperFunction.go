package wrapperFunction

import (
	"io"
	"os"
	"strings"
)

func LimitReader(reader io.Reader, n int64) io.Reader {
	return &LimitedReader{reader, n}
}

type LimitedReader struct {
	R io.Reader
	N int64
}

func (l *LimitedReader) Read(p []byte) (n int, err error) {
	return
}

func main() {
	reader := strings.NewReader("some io.Reader stream to be read")
	limitReader := io.LimitReader(reader, 4)
	_, err := io.Copy(os.Stdout, limitReader)
	if err != nil {
		return
	}
}
