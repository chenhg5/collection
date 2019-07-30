package main

import (
	"github.com/magiconair/properties/assert"
	"github.com/shopspring/decimal"
	"strings"
	"testing"
)

func TestStringArrayCollection_Join(t *testing.T) {
	a := []string{"h", "e", "l", "l", "o"}

	assert.Equal(t, Collect(a).Join(""), "hello")
}

var (
	numbers = []int{1, 2, 3, 4, 5, 6, 6, 7, 8, 8}
	foo     = []map[string]interface{}{
		{
			"foo": 10,
		}, {
			"foo": 30,
		}, {
			"foo": 20,
		}, {
			"foo": 40,
		},
	}
)

func TestNumberArrayCollection_Sum(t *testing.T) {
	assert.Equal(t, Collect(numbers).Sum().IntPart(), int64(50))

	var floatTest = []float64{143.66, -14.55}
	//c := floatTest[0] + floatTest[1]
	//fmt.Println(c)

	assert.Equal(t, Collect(floatTest).Sum().String(), "129.11")
}

func TestCollection_Splice(t *testing.T) {
	a := []string{"h", "e", "l", "l", "o"}

	c := Collect(a)
	assert.Equal(t, c.Splice(1, 3).ToStringArray(), []string{"e", "l", "l"})
	assert.Equal(t, c.Splice(1).ToStringArray(), []string{"e", "l", "l", "o"})

	assert.Equal(t, Collect(numbers).Splice(2, 1).ToNumberArray(),
		[]decimal.Decimal{nd(3)})

	assert.Equal(t, Collect(foo).Splice(1, 2).ToMapArray(), []map[string]interface{}{
		{
			"foo": 30,
		}, {
			"foo": 20,
		},
	})
}

func TestCollection_Take(t *testing.T) {
	a := []string{"h", "e", "l", "l", "o"}

	assert.Equal(t, Collect(a).Take(-2).ToStringArray(), []string{"l", "o"})

	assert.Equal(t, Collect(numbers).Take(4).ToNumberArray(),
		[]decimal.Decimal{nd(1), nd(2), nd(3), nd(4)})

	assert.Equal(t, Collect(foo).Take(2).ToMapArray(), []map[string]interface{}{
		{
			"foo": 10,
		}, {
			"foo": 30,
		},
	})
}

func TestCollection_All(t *testing.T) {
	a := []string{"h", "e", "l", "l", "o"}

	assert.Equal(t, Collect(a).All(), []interface{}{"h", "e", "l", "l", "o"})
	assert.Equal(t, len(Collect(numbers).All()), 10)
	assert.Equal(t, Collect(foo).All()[1], map[string]interface{}{"foo": 30})
}

func TestCollection_Mode(t *testing.T) {
	a := []string{"h", "e", "l", "l", "o", "w", "o", "l", "d"}
	foo2 := []map[string]interface{}{
		{
			"foo": 10,
		}, {
			"foo": 30,
		}, {
			"foo": 20,
		}, {
			"foo": 40,
		}, {
			"foo": 40,
		},
	}

	m := Collect(numbers).Mode()
	assert.Equal(t, m[0].(decimal.Decimal).IntPart() == int64(8) ||
		m[0].(decimal.Decimal).IntPart() == int64(6), true)
	assert.Equal(t, m[1].(decimal.Decimal).IntPart() == int64(8) ||
		m[1].(decimal.Decimal).IntPart() == int64(6), true)

	assert.Equal(t, Collect(a).Mode(), []interface{}{"l"})
	assert.Equal(t, Collect(foo2).Mode("foo"), []interface{}{40})
}

func TestCollection_Chunk(t *testing.T) {
	a := []string{"h", "e", "l", "l", "o"}

	assert.Equal(t, Collect(foo).Chunk(2).value[0][0], map[string]interface{}{"foo": 10})
	assert.Equal(t, len(Collect(numbers).Chunk(3).value), 4)
	assert.Equal(t, Collect(a).Chunk(3).value[0][2], "l")
}

func TestCollection_Collapse(t *testing.T) {
	a := []string{"h", "e", "l", "l", "o"}

	assert.Equal(t, Collect(foo).Chunk(2).Collapse(), Collect(foo))
	assert.Equal(t, Collect(a).Chunk(3).Collapse(), Collect(a))
	assert.Equal(t, Collect(numbers).Chunk(3).Collapse(), Collect(numbers))
}

func TestCollection_Concat(t *testing.T) {
	testNumbers := []int{1, 2, 3, 4, 5, 6, 6, 7, 8, 8, 9}
	a := []string{"h", "e", "l", "l", "o"}

	assert.Equal(t, len(Collect(foo).Concat(
		[]map[string]interface{}{{"foo": 100}}).ToMapArray()), 5)
	assert.Equal(t, Collect(numbers).Concat(
		[]decimal.Decimal{newDecimalFromInterface(9)}), Collect(testNumbers))
	assert.Equal(t, Collect(a).Concat([]string{"world"}).All()[5], "world")
	assert.Equal(t, Collect(numbers).Chunk(2).Concat(
		[][]interface{}{}).Collapse(), Collect(numbers))
}

func TestCollection_Contains(t *testing.T) {
	a := []string{"2", "3", "4", "5", "6"}

	assert.Equal(t, Collect(foo).Contains(10), true)
	assert.Equal(t, Collect(numbers).Contains(10), false)
	assert.Equal(t, Collect(a).Contains(5), true)
	assert.Equal(t, Collect(a).Contains("5"), true)
	assert.Equal(t, Collect(foo[3]).Contains(map[string]interface{}{"foo": 40}), true)

	c := Collect(numbers)
	value := 10
	assert.Equal(t, c.Contains(value, func() bool {
		for _, v := range c.ToNumberArray() {
			if v.LessThan(newDecimalFromInterface(value)) {
				return true
			}
		}
		return false
	}), true)
}

func TestCollection_ContainsStrict(t *testing.T) {
	a := []string{"h", "e", "l", "l", "o"}

	assert.Equal(t, Collect(foo).Contains(10), true)
	assert.Equal(t, Collect(numbers).Contains(10), false)
	assert.Equal(t, Collect(a).Contains("l"), true)
	assert.Equal(t, Collect(foo[3]).Contains(map[string]interface{}{"foo": 40}), true)

	c := Collect(numbers)
	value := 10
	assert.Equal(t, c.Contains(value, func() bool {
		for _, v := range c.ToNumberArray() {
			if v.GreaterThan(newDecimalFromInterface(value)) {
				return true
			}
		}
		return false
	}), false)
}

func TestCollection_CountBy(t *testing.T) {
	a := []string{"h", "e", "l", "l", "o"}
	assert.Equal(t, Collect(a).CountBy()["l"], 2)
	assert.Equal(t, Collect(numbers).CountBy()[float64(8)], 2)

	c := Collect([]string{"alice@gmail.com", "bob@yahoo.com", "carlos@gmail.com"})
	assert.Equal(t, c.CountBy(func() map[interface{}]int {
		valueCount := make(map[interface{}]int)
		for _, v := range c.ToStringArray() {
			f := strings.Split(v, "@")[1]
			valueCount[f]++
		}
		return valueCount
	}), map[interface{}]int{"gmail.com": 2, "yahoo.com": 1})
}

func TestCollection_CrossJoin(t *testing.T) {
	a := []interface{}{"h", "e", "l", "l", "o"}
	b := []interface{}{1, 2, 3, 4, 5, 6, 6, 7, 8, 8}

	assert.Equal(t, len(Collect(foo).CrossJoin(a, b).value), 200)
	assert.Equal(t, len(Collect(numbers).CrossJoin(a).value), 50)
	assert.Equal(t, Collect(foo).CrossJoin(b, a, b).value[1234][2], "l")
}

func TestCollection_Dd(t *testing.T) {
	a := []interface{}{"h", "e", "l", "l", "o"}

	Collect(foo).Dd()
	Collect(numbers).Dd()
	Collect(a).Dd()
	Collect(foo[2]).Dd()
}

func TestCollection_Dump(t *testing.T) {
	a := []interface{}{"h", "e", "l", "l", "o"}

	Collect(foo).Dump()
	Collect(numbers).Dump()
	Collect(a).Dump()
	Collect(foo[2]).Dump()
}
