package a

type Target string

const (
	TargetBlank  Target = "_blank"
	TargetSelf   Target = "_self"
	TargetParent Target = "_parent"
	TargetTop    Target = "_top"
)

type Element struct {
	href     string
	target   Target
	children []any
}

func New(href string, children ...any) *Element {
	return &Element{href: href, children: children}
}

func (a *Element) Target(target Target) *Element {
	a.target = target
	return a
}

func (a *Element) Render() error {
	return nil
}
