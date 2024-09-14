// Copyright (c) 2024, Peter Ohler, All rights reserved.
package sen_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/CodingBingo/ojg/jp"
	"github.com/CodingBingo/ojg/pretty"
	"github.com/CodingBingo/ojg/sen"
	"github.com/CodingBingo/ojg/tt"
)

func TestMatch(t *testing.T) {
	var buf []byte
	err := sen.Match([]byte(`{a:1 b:2}`), func(path jp.Expr, data any) {
		buf = append(buf, []byte(fmt.Sprintf("%s: %v", path, pretty.SEN(data)))...)
	}, jp.C("a"))
	tt.Nil(t, err)
	tt.Equal(t, "$.a: 1", string(buf))
}

func TestMatchString(t *testing.T) {
	var buf []byte
	err := sen.MatchString(`{a:1 b:2}`, func(path jp.Expr, data any) {
		buf = append(buf, []byte(fmt.Sprintf("%s: %v", path, pretty.SEN(data)))...)
	}, jp.C("a"))
	tt.Nil(t, err)
	tt.Equal(t, "$.a: 1", string(buf))
}

func TestMatchLoad(t *testing.T) {
	var buf []byte
	err := sen.MatchLoad(strings.NewReader(`{a:1 b:2}`), func(path jp.Expr, data any) {
		buf = append(buf, []byte(fmt.Sprintf("%s: %v", path, pretty.SEN(data)))...)
	}, jp.C("a"))
	tt.Nil(t, err)
	tt.Equal(t, "$.a: 1", string(buf))
}
