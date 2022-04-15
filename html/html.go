package html

type Element interface {
	Render() error
}

func Render(elem Element) error {
	return elem.Render()
}
