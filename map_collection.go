package main

import (
	"fmt"
)

type MapCollection struct {
	value map[string]interface{}
	BaseCollection
}

func (c MapCollection) Only(keys []string) Collection {
	var (
		d MapCollection
		m = make(map[string]interface{}, 0)
	)

	for _, k := range keys {
		m[k] = c.value[k]
	}
	d.value = m
	d.length = len(m)

	return d
}

func (c MapCollection) Prepend(values ...interface{}) Collection {

	var m = copyMap(c.value)
	m[values[0].(string)] = values[1]

	return MapCollection{m, BaseCollection{length: len(m)}}
}

func (c MapCollection) ToMap() map[string]interface{} {
	return c.value
}

func (c MapCollection) Contains(value interface{}, callback ...interface{}) bool {
	if len(callback) != 0 {
		return callback[0].(func() bool)()
	}

	t := fmt.Sprintf("%T", c.value)
	switch {
	case t == "[]map[string]string":
		return parseContainsParam(c.value, intToString(value))
	default:
		return parseContainsParam(c.value, value)
	}
}

func (c MapCollection) ContainsStrict(value interface{}, callback ...interface{}) bool {
	if len(callback) != 0 {
		return callback[0].(func() bool)()
	}

	return parseContainsParam(c.value, value)
}

func (c MapCollection) Dd() {
	dd(c)
}

func (c MapCollection) Dump() {
	dump(c)
}
