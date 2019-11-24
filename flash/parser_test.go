package flash_test

import (
	"github.com/find-a-job/flash"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParser(t *testing.T) {
	ParserError := "字符串 Parse went Error"

	assert.Equal(t, flash.Parser("int"), flash.Ident{Type: "int"}, ParserError)
	assert.Equal(t, flash.Parser("string"), flash.Ident{Type: "string"}, ParserError)
	assert.Equal(t, flash.Parser("bool"), flash.Ident{Type: "bool"}, ParserError)

	//@parse: Array<int>
	assert.Equal(
		t, flash.Parser("Array<int>"),
		flash.Ident{
			Type: "Array",
			TypeArgs: []flash.Ident{
				flash.Ident{
					Type: "int",
				},
			},
		},
		ParserError,
	)

	//@parse: Array<Array<int>>
	assert.Equal(
		t, flash.Parser("Array<Array<int>>"),
		flash.Ident{
			Type: "Array",
			TypeArgs: []flash.Ident{
				flash.Ident{
					Type: "Array",
					TypeArgs: []flash.Ident{
						flash.Ident{
							Type: "int",
						},
					},
				},
			},
		},
		ParserError,
	)

	//@parse: Map<string, Map<string, bool>>
	assert.Equal(
		t, flash.Parser("Map<string, Map<string, bool>>"),
		flash.Ident{
			Type: "Map",
			TypeArgs: []flash.Ident{
				flash.Ident{Type: "string"},
				flash.Ident{
					Type: "Map",
					TypeArgs: []flash.Ident{
						flash.Ident{Type: "string"},
						flash.Ident{Type: "bool"},
					},
				},
			},
		},
		ParserError,
	)

	//@parse: Map<Map<string, bool>, string>
	assert.Equal(
		t, flash.Parser("Map<Map<string, bool>, string>"),
		flash.Ident{
			Type: "Map",
			TypeArgs: []flash.Ident{
				flash.Ident{
					Type: "Map",
					TypeArgs: []flash.Ident{
						flash.Ident{Type: "string"},
						flash.Ident{Type: "bool"},
					},
				},
				flash.Ident{Type: "string"},
			},
		},
		ParserError,
	)
}
