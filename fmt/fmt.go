package fmt

import (
	"fmt"
	"strconv"
	"strings"
)

func Format(f fmt.State, verb rune) string {
	var sb strings.Builder
	sb.WriteByte('%')

	if w, ok := f.Width(); ok {
		sb.WriteString(strconv.Itoa(w))
	}

	if p, ok := f.Precision(); ok {
		sb.WriteByte('.')
		sb.WriteString(strconv.Itoa(p))
	}

	if f.Flag('+') {
		sb.WriteByte('+')
	}

	if f.Flag('-') {
		sb.WriteByte('-')
	}

	if f.Flag('#') {
		sb.WriteByte('#')
	}

	if f.Flag(' ') {
		sb.WriteByte(' ')
	}

	if f.Flag('0') {
		sb.WriteByte('0')
	}

	sb.WriteRune(verb)
	return sb.String()
}
