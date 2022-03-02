package ascii

import (
	"github.com/fdschonborn/sandbox/go/ascii"
	"github.com/fdschonborn/sandbox/go/option"
)

func ToAscii(r rune) option.Option[byte] {
	if !ascii.IsAscii(r) {
		return option.None[byte]()
	}

	return option.Some(byte(r))
}