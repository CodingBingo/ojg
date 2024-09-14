// Copyright (c) 2023, Peter Ohler, All rights reserved.

package jp_test

import (
	"testing"

	"github.com/CodingBingo/ojg/jp"
	"github.com/CodingBingo/ojg/tt"
)

func TestString(t *testing.T) {
	type Data struct {
		src    string
		expect string
	}
	for i, td := range []*Data{
		{src: "abc", expect: `|abc|`},
		{src: "&", expect: `|&|`},
		{src: "a\tbc", expect: `|a\tbc|`},
		{src: "a<b>c", expect: `|a<b>c|`},
		{src: "a 𝄢 note", expect: `|a 𝄢 note|`},
		{src: "a\u001ec", expect: `|a\u001ec|`},
		{src: "a\u2028b\u2029c", expect: `|a\u2028b\u2029c|`},
		{src: "abc\ufffd", expect: `|abc\ufffd|`},
	} {
		var buf []byte
		buf = jp.AppendString(buf, td.src, '|')
		tt.Equal(t, td.expect, string(buf), i, ": ", td.src)
	}
}
