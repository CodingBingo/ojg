// Copyright (c) 2023, Peter Ohler, All rights reserved.

package jp_test

import (
	"github.com/CodingBingo/ojg/alt"
)

type entry struct {
	key   string
	value any
}

type ordered struct {
	entries []*entry
}

func (o *ordered) Simplify() any {
	var simple []any
	for _, e := range o.entries {
		simple = append(simple, map[string]any{"key": e.key, "value": alt.Decompose(e.value)})
	}
	return simple
}

func (o *ordered) ValueAtIndex(index int) any {
	if index < 0 || len(o.entries) <= index {
		return nil
	}
	return o.entries[index].value
}

func (o *ordered) ValueForKey(key string) (value any, has bool) {
	for _, e := range o.entries {
		if e.key == key {
			return e.value, true
		}
	}
	return
}

func (o *ordered) SetValueForKey(key string, value any) {
	for _, e := range o.entries {
		if e.key == key {
			e.value = value
			return
		}
	}
	o.entries = append(o.entries, &entry{key: key, value: value})
}

func (o *ordered) RemoveValueForKey(key string) {
	for i, e := range o.entries {
		if e.key == key {
			copy(o.entries[i:], o.entries[i+1:])
			o.entries = o.entries[:len(o.entries)-1]
		}
	}
}

func (o *ordered) SetValueAtIndex(index int, value any) {
	if 0 <= index && index < len(o.entries) {
		o.entries[index].value = value
	}
}

func (o *ordered) RemoveValueAtIndex(index int) {
	if 0 <= index && index < len(o.entries) {
		copy(o.entries[index:], o.entries[index+1:])
		o.entries = o.entries[:len(o.entries)-1]
	}
}

type keyed struct {
	ordered
}

func (o *keyed) Keys() (keys []string) {
	for _, e := range o.entries {
		keys = append(keys, e.key)
	}
	return
}

type indexed struct {
	ordered
}

func (o *indexed) Size() int {
	return len(o.entries)
}

type keydex struct {
	keyed
}

func (o *keydex) Size() int {
	return len(o.entries)
}

func orderedFromSimple(v any) (o any) {
	switch tv := v.(type) {
	case []any:
		ind := indexed{}
		for _, v2 := range tv {
			ind.entries = append(ind.entries, &entry{value: orderedFromSimple(v2)})
		}
		o = &ind
	case map[string]any:
		kd := keyed{}
		for k, v2 := range tv {
			kd.entries = append(kd.entries, &entry{key: k, value: orderedFromSimple(v2)})
		}
		o = &kd
	default:
		o = v
	}
	return
}
