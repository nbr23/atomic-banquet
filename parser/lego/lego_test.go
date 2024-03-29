package lego

import (
	"testing"

	"github.com/nbr23/atomic-banquet/parser"
	testsuite "github.com/nbr23/atomic-banquet/utils"
)

func TestLegoParseNew(t *testing.T) {
	testsuite.TestParseSuccess(
		t,
		Lego{},
		&parser.Options{
			OptionsList: parser.OptionsList{
				&parser.Option{
					Flag:  "category",
					Type:  "string",
					Value: "new",
				},
			},
			Parser: Lego{},
		},
		1,
		`^\[\w+\] [0-9]+ - .* - (Available now|Pre-order this item today,) .*$`,
	)
}

func TestLegoParseComingSoon(t *testing.T) {
	testsuite.TestParseSuccess(
		t,
		Lego{},
		&parser.Options{
			OptionsList: parser.OptionsList{
				&parser.Option{
					Flag:  "category",
					Type:  "string",
					Value: "coming-soon",
				},
			},
			Parser: Lego{},
		},
		1,
		`^\[\w+\] [0-9]+ - .* - Coming Soon .*$`,
	)
}
