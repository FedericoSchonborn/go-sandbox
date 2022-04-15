package html_test

import (
	"testing"

	"github.com/FedericoSchonborn/go-sandbox/html/a"
	"github.com/FedericoSchonborn/go-sandbox/html/p"
	"github.com/FedericoSchonborn/go-sandbox/html/span"
)

func Test(t *testing.T) {
	p.New(
		"Here is a ",
		a.New("https://go.dev", "link").
			Target(a.TargetBlank),
		", and a ",
		span.New("span"),
		".",
	)
}
