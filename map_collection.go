package collection

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

func (c MapCollection) DiffAssoc(m map[string]interface{}) Collection {
	var d = make(map[string]interface{}, 0)
	for key, value := range m {
		if v, ok := c.value[key]; ok {
			if v != value {
				d[key] = value
			}
		}
	}
	return MapCollection{
		value: d,
	}
}

func (c MapCollection) DiffKeys(m map[string]interface{}) Collection {
	var d = make(map[string]interface{}, 0)
	for key, value := range c.value {
		if _, ok := m[key]; !ok {
			d[key] = value
		}
	}
	return MapCollection{
		value: d,
	}
}

func (c MapCollection) Each(cb func(item, value interface{}) (interface{}, bool)) Collection {
	var d = make(map[string]interface{}, 0)
	var (
		newValue interface{}
		stop     = false
	)
	for key, value := range c.value {
		if !stop {
			newValue, stop = cb(key, value)
			d[key] = newValue
		} else {
			d[key] = value
		}
	}
	return MapCollection{
		value: d,
	}
}

func (c MapCollection) Every(cb CB) bool {
	for key, value := range c.value {
		if !cb(key, value) {
			return false
		}
	}
	return true
}

func (c MapCollection) Except(keys []string) Collection {
	var d = copyMap(c.value)

	for _, key := range keys {
		delete(d, key)
	}
	return MapCollection{
		value: d,
	}
}

func (c MapCollection) FlatMap(cb func(value interface{}) interface{}) Collection {
	var d = make(map[string]interface{}, 0)
	var (
		newValue interface{}
	)
	for key, value := range c.value {
		newValue = cb(value)
		d[key] = newValue
	}
	return MapCollection{
		value: d,
	}
}

func (c MapCollection) Flip() Collection {
	var d = make(map[string]interface{}, 0)
	for key, value := range c.value {
		d[fmt.Sprintf("%v", value)] = key
	}
	return MapCollection{
		value: d,
	}
}

func (c MapCollection) Forget(k string) Collection {
	var d = copyMap(c.value)

	for key := range c.value {
		if key == k {
			delete(d, key)
		}
	}

	return MapCollection{
		value: d,
	}
}

func (c MapCollection) Get(k string, v ...interface{}) interface{} {
	if len(v) > 0 {
		if value, ok := c.value[k]; ok {
			return value
		} else {
			return v[0]
		}
	} else {
		return c.value[k]
	}
}

func (c MapCollection) Has(keys ...string) bool {
	for _, key := range keys {
		exist := false
		for kk := range c.value {
			if key == kk {
				exist = true
				break
			}
		}
		if !exist {
			return false
		}
	}
	return true
}

func (c MapCollection) IntersectByKeys(m map[string]interface{}) Collection {
	var d = make(map[string]interface{}, 0)
	for key, value := range c.value {
		for kk := range m {
			if kk == key {
				d[kk] = value
			}
		}
	}
	return MapCollection{
		value: d,
	}
}

func (c MapCollection) IsEmpty() bool {
	return len(c.value) == 0
}

func (c MapCollection) IsNotEmpty() bool {
	return len(c.value) != 0
}
