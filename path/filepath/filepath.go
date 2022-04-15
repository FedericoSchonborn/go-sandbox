package filepath

import (
	"io/fs"
	"path/filepath"
	"strings"
)

const (
	Separator     = filepath.Separator
	ListSeparator = filepath.ListSeparator
)

var (
	ErrBadPattern = filepath.ErrBadPattern
	SkipDir       = filepath.SkipDir
)

type WalkFunc func(path FilePath, info fs.FileInfo, err error) error

type FilePath string

func New(path string) FilePath {
	return FilePath(path)
}

func FromSlash(path string) FilePath {
	return New(filepath.FromSlash(path))
}

func Glob(pattern string) (matches []FilePath, err error) {
	vs, err := filepath.Glob(pattern)
	if err != nil {
		return nil, err
	}

	for _, v := range vs {
		matches = append(matches, New(v))
	}

	return matches, nil
}

func (fp FilePath) Abs() (FilePath, error) {
	v, err := filepath.Abs(string(fp))
	return New(v), err
}

func (fp FilePath) Base() string {
	return filepath.Base(string(fp))
}

func (fp FilePath) Clean() FilePath {
	return New(filepath.Clean(string(fp)))
}

func (fp FilePath) Dir() FilePath {
	return New(filepath.Dir(string(fp)))
}

func (fp FilePath) EvalSymlinks() (FilePath, error) {
	v, err := filepath.EvalSymlinks(string(fp))
	return New(v), err
}

func (fp FilePath) Ext() string {
	return filepath.Ext(string(fp))
}

func (fp FilePath) IsAbs() bool {
	return filepath.IsAbs(string(fp))
}

func (fp FilePath) Join(elem ...string) FilePath {
	return New(filepath.Join(append([]string{string(fp)}, elem...)...))
}

func (fp FilePath) Match(pattern string) (matched bool, err error) {
	return filepath.Match(pattern, string(fp))
}

func (fp FilePath) Rel(target FilePath) (FilePath, error) {
	v, err := filepath.Rel(string(fp), string(target))
	return New(v), err
}

func (fp FilePath) Split() (dir FilePath, file string) {
	sdir, sfile := filepath.Split(string(fp))
	return New(sdir), sfile
}

func (fp FilePath) ToSlash() string {
	return filepath.ToSlash(string(fp))
}

func (fp FilePath) VolumeName() string {
	return filepath.VolumeName(string(fp))
}

func (fp FilePath) Walk(fn WalkFunc) error {
	return filepath.Walk(string(fp), func(path string, info fs.FileInfo, err error) error {
		return fn(New(path), info, err)
	})
}

func (fp FilePath) String() string {
	return string(fp)
}

type List []FilePath

func NewList(list string) List {
	fpaths := filepath.SplitList(list)
	if len(fpaths) == 0 {
		return List{}
	}

	v := make(List, len(fpaths))
	for i, fpath := range fpaths {
		v[i] = New(fpath)
	}

	return v
}

func (l List) String() string {
	elems := make([]string, len(l))
	for i, fpath := range l {
		elems[i] = string(fpath)
	}

	return strings.Join(elems, ":")
}
