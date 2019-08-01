package collection

import (
	"github.com/shopspring/decimal"
	"math"
	"math/rand"
	"time"
)

type NumberArrayCollection struct {
	value []decimal.Decimal
	BaseCollection
}

func (c NumberArrayCollection) Sum(key ...string) decimal.Decimal {

	var sum = decimal.New(0, 0)

	for i := 0; i < len(c.value); i++ {
		sum = sum.Add(c.value[i])
	}

	return sum
}

func (c NumberArrayCollection) Min(key ...string) decimal.Decimal {

	var smallest = decimal.New(0, 0)

	for i := 0; i < len(c.value); i++ {
		if i == 0 {
			smallest = c.value[i]
			continue
		}
		if smallest.GreaterThan(c.value[i]) {
			smallest = c.value[i]
		}
	}

	return smallest
}

func (c NumberArrayCollection) Max(key ...string) decimal.Decimal {

	var biggest = decimal.New(0, 0)

	for i := 0; i < len(c.value); i++ {
		if i == 0 {
			biggest = c.value[i]
			continue
		}
		if biggest.LessThan(c.value[i]) {
			biggest = c.value[i]
		}
	}

	return biggest
}

func (c NumberArrayCollection) Prepend(values ...interface{}) Collection {
	var d NumberArrayCollection

	var n = make([]decimal.Decimal, len(c.value))
	copy(n, c.value)

	d.value = append([]decimal.Decimal{newDecimalFromInterface(values[0])}, n...)
	d.length = len(d.value)

	return d
}

func (c NumberArrayCollection) Splice(index ...int) Collection {

	if len(index) == 1 {
		var n = make([]decimal.Decimal, len(c.value))
		copy(n, c.value)
		n = n[index[0]:]

		return NumberArrayCollection{n, BaseCollection{length: len(n)}}
	} else if len(index) > 1 {
		var n = make([]decimal.Decimal, len(c.value))
		copy(n, c.value)
		n = n[index[0] : index[0]+index[1]]

		return NumberArrayCollection{n, BaseCollection{length: len(n)}}
	} else {
		panic("invalid argument")
	}
}

func (c NumberArrayCollection) Take(num int) Collection {
	var d NumberArrayCollection
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

func (c NumberArrayCollection) All() []interface{} {
	s := make([]interface{}, len(c.value))
	for i := 0; i < len(c.value); i++ {
		s[i] = c.value[i]
	}

	return s
}

func (c NumberArrayCollection) Mode(key ...string) []interface{} {
	valueCount := c.CountBy()
	maxCount := 0
	maxValue := make([]interface{}, len(valueCount))
	for v, c := range valueCount {
		switch {
		case c < maxCount:
			continue
		case c == maxCount:
			maxValue = append(maxValue, newDecimalFromInterface(v))
		case c > maxCount:
			maxValue = append([]interface{}{}, newDecimalFromInterface(v))
			maxCount = c
		}
	}
	return maxValue
}

func (c NumberArrayCollection) ToNumberArray() []decimal.Decimal {
	return c.value
}

func (c NumberArrayCollection) ToIntArray() []int {
	var v = make([]int, len(c.value))
	for i, value := range c.value {
		v[i] = int(value.IntPart())
	}
	return v
}

func (c NumberArrayCollection) Chunk(num int) MultiDimensionalArrayCollection {
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

func (c NumberArrayCollection) Concat(value interface{}) Collection {
	return NumberArrayCollection{
		value:          append(c.value, value.([]decimal.Decimal)...),
		BaseCollection: BaseCollection{length: c.length + len(value.([]decimal.Decimal))},
	}
}

func (c NumberArrayCollection) Contains(value interface{}, callback ...interface{}) bool {
	if len(callback) != 0 {
		return callback[0].(func() bool)()
	}

	return containsValue(c.value, value)
}

func (c NumberArrayCollection) ContainsStrict(value interface{}, callback ...interface{}) bool {
	if len(callback) != 0 {
		return callback[0].(func() bool)()
	}

	return containsValue(c.value, value)
}

func (c NumberArrayCollection) CountBy(callback ...interface{}) map[interface{}]int {
	if len(callback) != 0 {
		return callback[0].(func() map[interface{}]int)()
	}

	valueCount := make(map[interface{}]int)
	for _, v := range c.value {
		f, _ := v.Float64()
		valueCount[f]++
	}

	return valueCount
}

func (c NumberArrayCollection) CrossJoin(array ...[]interface{}) MultiDimensionalArrayCollection {
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

func (c NumberArrayCollection) Dd() {
	dd(c)
}

func (c NumberArrayCollection) Dump() {
	dump(c)
}

func (c NumberArrayCollection) Diff(m interface{}) Collection {
	ms := newDecimalArray(m)
	var d = make([]decimal.Decimal, 0)
	for _, value := range c.value {
		exist := false
		for i := 0; i < len(ms); i++ {
			if ms[i].Equal(value) {
				exist = true
				break
			}
		}
		if !exist {
			d = append(d, value)
		}
	}
	return NumberArrayCollection{
		value: d,
	}
}

func (c NumberArrayCollection) Each(cb func(item, value interface{}) (interface{}, bool)) Collection {
	var d = make([]decimal.Decimal, 0)
	var (
		newValue interface{}
		stop     = false
	)
	for key, value := range c.value {
		if !stop {
			newValue, stop = cb(key, value)
			d = append(d, newDecimalFromInterface(newValue))
		} else {
			d = append(d, value)
		}
	}
	return NumberArrayCollection{
		value: d,
	}
}

func (c NumberArrayCollection) Every(cb CB) bool {
	for key, value := range c.value {
		if !cb(key, value) {
			return false
		}
	}
	return true
}

func (c NumberArrayCollection) Filter(cb CB) Collection {
	var d = make([]decimal.Decimal, 0)
	for key, value := range c.value {
		if cb(key, value) {
			d = append(d, value)
		}
	}
	return NumberArrayCollection{
		value: d,
	}
}

func (c NumberArrayCollection) First(cbs ...CB) interface{} {
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

func (c NumberArrayCollection) ForPage(page, size int) Collection {
	var d = make([]decimal.Decimal, len(c.value))
	copy(d, c.value)
	if size > len(d) || size*(page-1) > len(d) {
		return NumberArrayCollection{
			value: d,
		}
	}
	if (page+1)*size > len(d) {
		return NumberArrayCollection{
			value: d[(page-1)*size:],
		}
	} else {
		return NumberArrayCollection{
			value: d[(page-1)*size : (page)*size],
		}
	}
}

func (c NumberArrayCollection) IsEmpty() bool {
	return len(c.value) == 0
}

func (c NumberArrayCollection) IsNotEmpty() bool {
	return len(c.value) != 0
}

func (c NumberArrayCollection) Last(cbs ...CB) interface{} {
	if len(cbs) > 0 {
		var last interface{}
		for key, value := range c.value {
			if cbs[0](key, value) {
				last = value
			}
		}
		return last
	} else {
		if len(c.value) > 0 {
			return c.value[len(c.value)-1]
		} else {
			return nil
		}
	}
}

func (c NumberArrayCollection) Median(key ...string) decimal.Decimal {

	if len(c.value) < 2 {
		return c.value[0]
	}

	var f = make([]decimal.Decimal, len(c.value))
	copy(f, c.value)
	qsort(f, true)
	return f[len(f)/2].Add(f[len(f)/2-1]).Div(nd(2))
}

func (c NumberArrayCollection) Merge(i interface{}) Collection {
	m := newDecimalArray(i)
	var d = make([]decimal.Decimal, len(c.value))
	copy(d, c.value)
	d = append(d, m...)

	return NumberArrayCollection{
		value: d,
	}
}

func (c NumberArrayCollection) Pad(num int, value interface{}) Collection {
	if len(c.value) > num {
		d := make([]decimal.Decimal, len(c.value))
		copy(d, c.value)
		return NumberArrayCollection{
			value: d,
		}
	}
	if num > 0 {
		d := make([]decimal.Decimal, num)
		for i := 0; i < num; i++ {
			if i < len(c.value) {
				d[i] = c.value[i]
			} else {
				d[i] = nd(value)
			}
		}
		return NumberArrayCollection{
			value: d,
		}
	} else {
		d := make([]decimal.Decimal, -num)
		for i := 0; i < -num; i++ {
			if i < -num-len(c.value) {
				d[i] = nd(value)
			} else {
				d[i] = c.value[i]
			}
		}
		return NumberArrayCollection{
			value: d,
		}
	}
}

func (c NumberArrayCollection) Partition(cb PartCB) (Collection, Collection) {
	var d1 = make([]decimal.Decimal, 0)
	var d2 = make([]decimal.Decimal, 0)

	for i := 0; i < len(c.value); i++ {
		if cb(i) {
			d1 = append(d1, c.value[i])
		} else {
			d2 = append(d2, c.value[i])
		}
	}

	return NumberArrayCollection{
		value: d1,
	}, NumberArrayCollection{
		value: d2,
	}
}

func (c NumberArrayCollection) Pop() interface{} {
	last := c.value[len(c.value)-1]
	c.value = c.value[:len(c.value)-1]
	return last
}

func (c NumberArrayCollection) Push(v interface{}) Collection {
	var d = make([]decimal.Decimal, len(c.value)+1)
	for i := 0; i < len(d); i++ {
		if i < len(c.value) {
			d[i] = c.value[i]
		} else {
			d[i] = nd(v)
		}
	}

	return NumberArrayCollection{
		value: d,
	}
}

func (c NumberArrayCollection) Random(num ...int) Collection {
	if len(num) == 0 {
		return BaseCollection{
			value: c.value[rand.Intn(len(c.value))],
		}
	} else {
		if num[0] > len(c.value) {
			panic("wrong num")
		}
		var d = make([]decimal.Decimal, len(c.value))
		copy(d, c.value)
		for i := 0; i < len(c.value)-num[0]; i++ {
			index := rand.Intn(len(d))
			d = append(d[:index], d[index+1:]...)
		}
		return NumberArrayCollection{
			value: d,
		}
	}
}

func (c NumberArrayCollection) Reduce(cb ReduceCB) interface{} {
	var res interface{}

	for i := 0; i < len(c.value); i++ {
		res = cb(res, c.value[i])
	}

	return res
}

func (c NumberArrayCollection) Reject(cb CB) Collection {
	var d = make([]decimal.Decimal, 0)
	for key, value := range c.value {
		if !cb(key, value) {
			d = append(d, value)
		}
	}
	return NumberArrayCollection{
		value: d,
	}
}

func (c NumberArrayCollection) Reverse() Collection {
	var d = make([]decimal.Decimal, len(c.value))
	j := 0
	for i := len(c.value) - 1; i > -1; i-- {
		d[j] = c.value[i]
		j++
	}
	return NumberArrayCollection{
		value: d,
	}
}

func (c NumberArrayCollection) Search(v interface{}) int {
	if cb, ok := v.(CB); ok {
		for i := 0; i < len(c.value); i++ {
			if cb(i, c.value[i]) {
				return i
			}
		}
	} else {
		n := nd(v)
		for i := 0; i < len(c.value); i++ {
			if n == c.value[i] {
				return i
			}
		}
	}
	return -1
}

func (c NumberArrayCollection) Shift() Collection {
	var d = make([]decimal.Decimal, len(c.value))
	copy(d, c.value)
	d = d[1:]
	return NumberArrayCollection{
		value: d,
	}
}

func (c NumberArrayCollection) Shuffle() Collection {
	var d = make([]decimal.Decimal, len(c.value))
	copy(d, c.value)
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(c.value), func(i, j int) { d[i], d[j] = d[j], d[i] })
	return NumberArrayCollection{
		value: d,
	}
}

func (c NumberArrayCollection) Slice(keys ...int) Collection {
	var d = make([]decimal.Decimal, len(c.value))
	copy(d, c.value)
	if len(keys) == 1 {
		return NumberArrayCollection{
			value: d[keys[0]:],
		}
	} else {
		return NumberArrayCollection{
			value: d[keys[0] : keys[0]+keys[1]],
		}
	}
}

func (c NumberArrayCollection) Sort() Collection {
	var d = make([]decimal.Decimal, len(c.value))
	copy(d, c.value)
	qsort(d, true)
	return NumberArrayCollection{
		value: d,
	}
}

func (c NumberArrayCollection) SortByDesc() Collection {
	var d = make([]decimal.Decimal, len(c.value))
	copy(d, c.value)
	qsort(d, false)
	return NumberArrayCollection{
		value: d,
	}
}

func (c NumberArrayCollection) Split(num int) Collection {
	var d = make([][]interface{}, math.Ceil(float64(len(c.value))/float64(num)))

	j := 0
	for i := 0; i < len(c.value); i++ {
		if i%num == 0 {
			if i+num <= len(c.value) {
				d[j] = make([]interface{}, num)
			} else {
				d[j] = make([]interface{}, len(c.value)-i)
			}
			d[j][i%num] = c.value[i]
			j++
		} else {
			d[j][i%num] = c.value[i]
		}
	}

	return MultiDimensionalArrayCollection{
		value: d,
	}
}

func (c NumberArrayCollection) Unique() Collection {
	var d = make([]decimal.Decimal, len(c.value))
	copy(d, c.value)
	x := make([]decimal.Decimal, 0)
	for _, i := range d {
		if len(x) == 0 {
			x = append(x, i)
		} else {
			for k, v := range x {
				if i == v {
					break
				}
				if k == len(x)-1 {
					x = append(x, i)
				}
			}
		}
	}
	return NumberArrayCollection{
		value: x,
	}
}
