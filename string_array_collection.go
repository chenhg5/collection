package main

import (
	"fmt"
)

type StringArrayCollection struct {
	value []string
	BaseCollection
}

func (c StringArrayCollection) Join(delimiter string) string {
	s := ""
	for i := 0; i < len(c.value); i++ {
		if i != len(c.value)-1 {
			s += c.value[i] + delimiter
		} else {
			s += c.value[i]
		}
	}
	return s
}

func (c StringArrayCollection) Combine(value []interface{}) Collection {
	var (
		m      = make(map[string]interface{}, 0)
		length = c.length
		d      MapCollection
	)

	if length > len(value) {
		length = len(value)
	}

	for i := 0; i < length; i++ {
		m[c.value[i]] = value[i]
	}

	d.value = m
	d.length = len(m)

	return d
}

func (c StringArrayCollection) Prepend(values ...interface{}) Collection {

	var d StringArrayCollection

	var n = make([]string, len(c.value))
	copy(n, c.value)

	d.value = append([]string{values[0].(string)}, n...)
	d.length = len(d.value)

	return d
}

func (c StringArrayCollection) Splice(index ...int) Collection {

	if len(index) == 1 {
		var n = make([]string, len(c.value))
		copy(n, c.value)
		n = n[index[0]:]

		return StringArrayCollection{n, BaseCollection{length: len(n)}}
	} else if len(index) > 1 {
		var n = make([]string, len(c.value))
		copy(n, c.value)
		n = n[index[0] : index[0]+index[1]]

		return StringArrayCollection{n, BaseCollection{length: len(n)}}
	} else {
		panic("invalid argument")
	}
}

func (c StringArrayCollection) Take(num int) Collection {
	var d StringArrayCollection
	if num > c.length {
		panic("not enough elements to take")
	}

	if num >= 0 {
		d.value = c.value[:num]
		d.length = num
	} else {
		d.value = c.value[len(c.value)+num:]
		d.length = 0 - num
	}

	return d
}

func (c StringArrayCollection) All() []interface{} {
	s := make([]interface{}, len(c.value))
	for i := 0; i < len(c.value); i++ {
		s[i] = c.value[i]
	}

	return s
}

func (c StringArrayCollection) Mode(key ...string) []interface{} {
	valueCount := c.CountBy()
	maxCount := 0
	maxValue := make([]interface{}, len(valueCount))
	for v, c := range valueCount {
		switch {
		case c < maxCount:
			continue
		case c == maxCount:
			maxValue = append(maxValue, v)
		case c > maxCount:
			maxValue = append([]interface{}{}, v)
			maxCount = c
		}
	}
	return maxValue
}

func (c StringArrayCollection) ToStringArray() []string {
	return c.value
}

func (c StringArrayCollection) Chunk(num int) MultiDimensionalArrayCollection {
	var d MultiDimensionalArrayCollection
	d.length = c.length/num + 1
	d.value = make([][]interface{}, d.length)

	count := 0
	for i := 1; i <= c.length; i++ {
		switch {
		case i == c.length:
			if i%num == 0 {
				d.value[count] = c.All()[i-num:]
				d.value = d.value[:d.length-1]
			} else {
				d.value[count] = c.All()[i-i%num:]
			}
		case i%num != 0 || i < num:
			continue
		default:
			d.value[count] = c.All()[i-num : i]
			count++
		}
	}

	return d
}

func (c StringArrayCollection) Concat(value interface{}) Collection {
	return StringArrayCollection{
		value:          append(c.value, value.([]string)...),
		BaseCollection: BaseCollection{length: c.length + len(value.([]string))},
	}
}

func (c StringArrayCollection) Contains(value interface{}, callback ...interface{}) bool {
	if len(callback) != 0 {
		return callback[0].(func() bool)()
	}

	t := fmt.Sprintf("%T", c.value)
	switch {
	case t == "[]string":
		return containsValue(c.value, intToString(value))
	default:
		return containsValue(c.value, value)
	}
}

func (c StringArrayCollection) ContainsStrict(value interface{}, callback ...interface{}) bool {
	if len(callback) != 0 {
		return callback[0].(func() bool)()
	}

	return containsValue(c.value, value)
}

func (c StringArrayCollection) CountBy(callback ...interface{}) map[interface{}]int {
	if len(callback) != 0 {
		return callback[0].(func() map[interface{}]int)()
	}

	valueCount := make(map[interface{}]int)
	for _, v := range c.value {
		valueCount[v]++
	}

	return valueCount
}

func (c StringArrayCollection) CrossJoin(array ...[]interface{}) MultiDimensionalArrayCollection {
	var d MultiDimensionalArrayCollection

	// A two-dimensional-slice's initial
	length := len(c.value)
	for _, s := range array {
		length *= len(s)
	}
	value := make([][]interface{}, length)
	for i := range value {
		value[i] = make([]interface{}, len(array)+1)
	}

	offset := length / c.length
	for i := 0; i < length; i++ {
		value[i][0] = c.value[i/offset]
	}
	assignmentToValue(value, array, length, 1, 0, offset)

	d.value = value
	d.length = length
	return d
}

func (c StringArrayCollection) Dd() {
	dd(c)
}

func (c StringArrayCollection) Dump() {
	dump(c)
}

func (c StringArrayCollection) Diff(m interface{}) Collection {
	ms := m.([]string)
	var d = make([]string, 0)
	for _, value := range c.value {
		exist := false
		for i := 0; i < len(ms); i++ {
			if ms[i] == value {
				exist = true
				break
			}
		}
		if !exist {
			d = append(d, value)
		}
	}
	return StringArrayCollection{
		value: d,
	}
}

func (c StringArrayCollection) Each(cb func(item, value interface{}) (interface{}, bool)) Collection {
	var d = make([]string, 0)
	var (
		newValue interface{}
		stop     = false
	)
	for key, value := range c.value {
		if !stop {
			newValue, stop = cb(key, value)
			d = append(d, newValue.(string))
		} else {
			d = append(d, value)
		}
	}
	return StringArrayCollection{
		value: d,
	}
}

func (c StringArrayCollection) Every(cb CB) bool {
	for key, value := range c.value {
		if !cb(key, value) {
			return false
		}
	}
	return true
}

func (c StringArrayCollection) Filter(cb CB) Collection {
	var d = make([]string, 0)
	for key, value := range c.value {
		if cb(key, value) {
			d = append(d, value)
		}
	}
	return StringArrayCollection{
		value: d,
	}
}

func (c StringArrayCollection) First(cbs ...CB) interface{} {
	if len(cbs) > 0 {
		for key, value := range c.value {
			if cbs[0](key, value) {
				return value
			}
		}
		return nil
	} else {
		if len(c.value) > 0 {
			return c.value[0]
		} else {
			return nil
		}
	}
}

func (c StringArrayCollection) Intersect(keys []string) Collection {
	var d = make([]string, 0)
	for _, value := range c.value {
		for _, v := range keys {
			if v == value {
				d = append(d, value)
				break
			}
		}
	}
	return StringArrayCollection{
		value: d,
	}
}

func (c StringArrayCollection) IsEmpty() bool {
	return len(c.value) == 0
}

func (c StringArrayCollection) IsNotEmpty() bool {
	return len(c.value) != 0
}