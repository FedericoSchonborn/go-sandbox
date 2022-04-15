package span

type Element struct {
	children []any
}

func New(children ...any) *Element {
	return &Element{children}
}

func (span *Element) Render() error {
	return nil
}
