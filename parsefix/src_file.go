package parsefix

import (
	"bytes"
)

func newSrcFile(data []byte) *srcFile {
	return &srcFile{
		lines: bytes.SplitAfter(data, []byte("\n")),
	}
}

type srcFile struct {
	lines [][]byte
}

func (src *srcFile) Bytes() []byte {
	return bytes.Join(src.lines, nil)
}
