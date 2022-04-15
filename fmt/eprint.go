package fmt

import (
	"fmt"
	"os"
)

func Eprint(a ...any) (n int, err error) {
	return fmt.Fprint(os.Stderr, a...)
}

func Eprintf(format string, a ...any) (n int, err error) {
	return fmt.Fprintf(os.Stderr, format, a...)
}

func Eprintln(a ...any) (n int, err error) {
	return fmt.Fprintln(os.Stderr, a...)
}

func Eprintfln(format string, a ...any) (n int, err error) {
	return Fprintfln(os.Stderr, format, a...)
}
