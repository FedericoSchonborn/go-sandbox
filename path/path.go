package path

import (
	"path"
)

var (
	ErrBadPattern = path.ErrBadPattern
)

type Path string

func New(path string) Path {
	return Path(path)
}

func (p Path) Base() string {
	return path.Base(string(p))
}

func (p Path) Clean() Path {
	return New(path.Clean(string(p)))
}

func (p Path) Dir() Path {
	return New(path.Dir(string(p)))
}

func (p Path) Ext() string {
	return path.Ext(string(p))
}

func (p Path) IsAbs() bool {
	return path.IsAbs(string(p))
}

func (p Path) Join(elem ...string) Path {
	return New(path.Join(append([]string{string(p)}, elem...)...))
}

func (p Path) Match(pattern string) (matched bool, err error) {
	return path.Match(pattern, string(p))
}

func (p Path) Split() (dir Path, file string) {
	sdir, sfile := path.Split(string(p))
	return New(sdir), sfile
}

func (p Path) String() string {
	return string(p)
}
