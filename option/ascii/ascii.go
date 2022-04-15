package ascii

import (
	"github.com/FedericoSchonborn/go-sandbox/ascii"
	"github.com/FedericoSchonborn/go-sandbox/option"
)

func ToAscii(r rune) option.Option[byte] {
	if !ascii.IsAscii(r) {
		return option.None[byte]()
	}

	return option.Some(byte(r))
}
