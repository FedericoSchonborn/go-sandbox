package p

type Element struct {
	children []any
}

func New(children ...any) *Element {
	return &Element{children}
}

func (p *Element) Render() error {
	return nil
}
