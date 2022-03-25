package parsefix

import (
	"bytes"
)

type fixerContext struct {
	issue string
	loc   location
	src   *srcFile
}

func (ctx *fixerContext) contains(s string) bool {
	return bytes.Contains(ctx.src.lines[ctx.loc.line], []byte(s))
}

func (ctx *fixerContext) nextNonWhitespaceIs(b byte) bool {
	lineStart := ctx.loc.column
	for _, line := range ctx.src.lines[ctx.loc.line:] {
		for offset := range line[lineStart:] {
			col := offset + lineStart
			switch line[col] {
			// taken from unicode.IsSpace()
			case '\t', '\n', '\v', '\f', '\r', ' ', 0x85, 0xA0:
				continue
			case b:
				return true
			default:
				return false
			}
		}
		lineStart = 0
	}

	return false
}

func (ctx *fixerContext) replace(from, to string) {
	ctx.src.lines[ctx.loc.line] = bytes.Replace(
		ctx.src.lines[ctx.loc.line], []byte(from), []byte(to), 1)
}

func (ctx *fixerContext) insertByte(b byte) {
	line := ctx.src.lines[ctx.loc.line]
	pos := ctx.loc.column

	line = append(line, 0)
	copy(line[pos+1:], line[pos:])
	line[pos] = b

	ctx.src.lines[ctx.loc.line] = line
}
