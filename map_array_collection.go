package main

import (
	"fmt"
	"github.com/shopspring/decimal"
	"strconv"
)

type MapArrayCollection struct {
	value []map[string]interface{}
	BaseCollection
}

func (c MapArrayCollection) Value() interface{} {
	return c.value
}

func (c MapArrayCollection) Sum(key ...string) decimal.Decimal {
	var sum = decimal.New(0, 0)

	for i := 0; i < len(c.value); i++ {
		sum = sum.Add(newDecimalFromInterface(c.value[i][key[0]]))
	}

	return sum
}

func (c MapArrayCollection) Min(key ...string) decimal.Decimal {

	var (
		smallest = decimal.New(0, 0)
		number   decimal.Decimal
	)

	for i := 0; i < len(c.value); i++ {
		number = newDecimalFromInterface(c.value[i][key[0]])
		if i == 0 {
			smallest = number
			continue
		}
		if smallest.GreaterThan(number) {
			smallest = number
		}
	}

	return smallest
}

func (c MapArrayCollection) Max(key ...string) decimal.Decimal {

	var (
		biggest = decimal.New(0, 0)
		number  decimal.Decimal
	)

	for i := 0; i < len(c.value); i++ {
		number = newDecimalFromInterface(c.value[i][key[0]])
		if i == 0 {
			biggest = number
			continue
		}
		if biggest.LessThan(number) {
			biggest = number
		}
	}

	return biggest
}

func (c MapArrayCollection) Pluck(key string) Collection {
	var s = make([]interface{}, 0)
	for i := 0; i < len(c.value); i++ {
		s = append(s, c.value[i][key])
	}
	return Collect(s)
}

func (c MapArrayCollection) Prepend(values ...interface{}) Collection {

	var d MapArrayCollection

	var n = make([]map[string]interface{}, len(c.value))
	copy(n, c.value)

	d.value = append([]map[string]interface{}{values[0].(map[string]interface{})}, n...)
	d.length = len(d.value)

	return d
}

func (c MapArrayCollection) Only(keys []string) Collection {
	var d MapArrayCollection

	var ma = make([]map[string]interface{}, 0)
	for _, k := range keys {
		m := make(map[string]interface{}, 0)
		for _, v := range c.value {
			m[k] = v[k]
		}
		ma = append(ma, m)
	}
	d.value = ma
	d.length = len(ma)

	return d
}

func (c MapArrayCollection) Splice(index ...int) Collection {

	if len(index) == 1 {
		var n = make([]map[string]interface{}, len(c.value))
		copy(n, c.value)
		n = n[index[0]:]

		return MapArrayCollection{n, BaseCollection{length: len(n)}}
	} else if len(index) > 1 {
		var n = make([]map[string]interface{}, len(c.value))
		copy(n, c.value)
		n = n[index[0] : index[0]+index[1]]

		return MapArrayCollection{n, BaseCollection{length: len(n)}}
	} else {
		panic("invalid argument")
	}
}

func (c MapArrayCollection) Take(num int) Collection {
	var d MapArrayCollection
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

func (c MapArrayCollection) All() []interface{} {
	s := make([]interface{}, len(c.value))
	for i := 0; i < len(c.value); i++ {
		s[i] = c.value[i]
	}

	return s
}

func (c MapArrayCollection) Mode(key ...string) []interface{} {
	valueCount := make(map[interface{}]int)
	for i := 0; i < c.length; i++ {
		if v, ok := c.value[i][key[0]]; ok {
			valueCount[v]++
		}
	}

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

func (c MapArrayCollection) ToMapArray() []map[string]interface{} {
	return c.value
}

func (c MapArrayCollection) Chunk(num int) MultiDimensionalArrayCollection {
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

func (c MapArrayCollection) Concat(value interface{}) Collection {
	return MapArrayCollection{
		value:          append(c.value, value.([]map[string]interface{})...),
		BaseCollection: BaseCollection{length: c.length + len(value.([]map[string]interface{}))},
	}
}

func (c MapArrayCollection) Contains(value interface{}, callback ...interface{}) bool {
	if len(callback) != 0 {
		return callback[0].(func() bool)()
	}

	t := fmt.Sprintf("%T", c.value)
	switch {
	case t == "[]map[string]string":
		for _, m := range c.value {
			if parseContainsParam(m, intToString(value)) {
				return true
			}
		}
		return false
	default:
		for _, m := range c.value {
			if parseContainsParam(m, value) {
				return true
			}
		}
		return false
	}
}

func parseContainsParam(m map[string]interface{}, value interface{}) bool {
	switch value.(type) {
	case map[string]interface{}:
		return containsKeyValue(m, value.(map[string]interface{}))
	default:
		return containsValue(m, value)
	}
}

func intToString(value interface{}) interface{} {
	switch value.(type) {
	case int:
		return strconv.Itoa(value.(int))
	case int64:
		return strconv.FormatInt(value.(int64), 10)
	default:
		return value
	}
}

func containsValue(m interface{}, value interface{}) bool {
	switch m.(type) {
	case map[string]interface{}:
		for _, v := range m.(map[string]interface{}) {
			if v == value {
				return true
			}
		}
		return false
	case []decimal.Decimal:
		for _, v := range m.([]decimal.Decimal) {
			if v.Equal(newDecimalFromInterface(value)) {
				return true
			}
		}
		return false
	case []string:
		for _, v := range m.([]string) {
			if v == value {
				return true
			}
		}
		return false
	default:
		panic("wrong type")
	}
}

func containsKeyValue(m map[string]interface{}, value map[string]interface{}) bool {
	for k, v := range value {
		if _, ok := m[k]; !ok && m[k] != v {
			return false
		}
	}

	return true
}

func (c MapArrayCollection) ContainsStrict(value interface{}, callback ...interface{}) bool {
	if len(callback) != 0 {
		return callback[0].(func() bool)()
	}

	for _, m := range c.value {
		if parseContainsParam(m, value) {
			return true
		}
	}

	return false
}

func (c MapArrayCollection) CrossJoin(array ...[]interface{}) MultiDimensionalArrayCollection {
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

// vl: length of value
// ai: index of array
// si: index of value's sub-array
func assignmentToValue(value, array [][]interface{}, vl, si, ai, preOffset int) {
	offset := preOffset / len(array[ai])
	times := 0

	for i := 0; i < vl; i++ {
		if i >= preOffset && i%preOffset == 0 {
			times++
		}
		value[i][si] = array[ai][(i-preOffset*times)/offset]
	}

	if ai < len(array)-1 {
		assignmentToValue(value, array, vl, si+1, ai+1, offset)
	}
}

func (c MapArrayCollection) Dd() {
	dd(c)
}

func (c MapArrayCollection) Dump() {
	dump(c)
}
