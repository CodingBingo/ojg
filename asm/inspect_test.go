// Copyright (c) 2021, Peter Ohler, All rights reserved.

package asm_test

import (
	"testing"

	"github.com/CodingBingo/ojg/asm"
	"github.com/CodingBingo/ojg/sen"
	"github.com/CodingBingo/ojg/tt"
)

func TestInspect(t *testing.T) {
	p := asm.NewPlan([]any{
		"asm",
		[]any{"inspect", "test", "$"},
	})
	tt.Equal(t, "[asm [inspect test $]]", sen.String(p), "inspect plan simplify")
	fn, _ := p.Args[0].(*asm.Fn)
	tt.NotNil(t, fn)
	tt.Equal(t, "[inspect test $]", fn.String(), "inspect string")
}
