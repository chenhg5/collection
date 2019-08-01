package collection

import (
	"fmt"
	"github.com/magiconair/properties/assert"
	"github.com/shopspring/decimal"
	"reflect"
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
	assert.Equal(t, Collect(floatTest).Sum().String(), "129.11")
}

func ExampleNumberArrayCollection_Sum() {
	var floatTest = []float64{143.66, -14.55}
	fmt.Println(Collect(floatTest).Sum().String())

	// Output: 129.11
}

func TestStringArrayCollection_Splice(t *testing.T) {
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

func TestStringArrayCollection_Take(t *testing.T) {
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

func TestStringArrayCollection_Diff(t *testing.T) {
	a := []string{"h", "e", "l", "l", "o"}
	assert.Equal(t, Collect(a).Diff([]string{"e", "o"}).ToStringArray(), []string{"h", "l", "l"})
	c := []int{3, 4, 5, 6, 7, 8}
	assert.Equal(t, Collect(c).Diff([]int{5, 6}).ToNumberArray(), []decimal.Decimal{nd(3), nd(4), nd(7), nd(8)})
}

func TestMapCollection_DiffAssoc(t *testing.T) {
	a := map[string]interface{}{
		"name": "mike",
		"sex":  1,
	}
	assert.Equal(t, reflect.DeepEqual(Collect(a).DiffAssoc(map[string]interface{}{
		"name": "miker",
	}).ToMap(), map[string]interface{}{
		"name": "miker",
	}), true)
}

func TestMapCollection_DiffKeys(t *testing.T) {
	a := map[string]interface{}{
		"name": "mike",
		"sex":  1,
	}
	assert.Equal(t, reflect.DeepEqual(Collect(a).DiffKeys(map[string]interface{}{
		"name": "miker",
	}).ToMap(), map[string]interface{}{
		"sex": 1,
	}), true)
}

func TestNumberArrayCollection_Each(t *testing.T) {
	a := []int{2, 3, 4, 5, 6, 7}

	assert.Equal(t, Collect(a).Each(func(item, value interface{}) (interface{}, bool) {
		return value.(decimal.Decimal).IntPart() + 2, false
	}).ToIntArray(), []int{4, 5, 6, 7, 8, 9})
}

func TestNumberArrayCollection_Every(t *testing.T) {

	a := []int{2, 3, 4, 5, 6, 7}

	assert.Equal(t, Collect(a).Every(func(item, value interface{}) bool {
		return value.(decimal.Decimal).GreaterThanOrEqual(nd(5))
	}), false)
}

func TestMapCollection_Except(t *testing.T) {
	a := map[string]interface{}{
		"name": "mike",
		"sex":  1,
	}
	assert.Equal(t, reflect.DeepEqual(Collect(a).Except([]string{"name"}).ToMap(), map[string]interface{}{
		"sex": 1,
	}), true)
}

func TestNumberArrayCollection_Filter(t *testing.T) {
	a := []int{2, 3, 4, 5, 6, 7}

	assert.Equal(t, Collect(a).Filter(func(item, value interface{}) bool {
		return value.(decimal.Decimal).IntPart() > 4
	}).ToIntArray(), []int{5, 6, 7})
}

func TestNumberArrayCollection_First(t *testing.T) {
	a := []int{2, 3, 4, 5, 6, 7}

	assert.Equal(t, Collect(a).First(func(item, value interface{}) bool {
		return value.(decimal.Decimal).IntPart() > 4
	}), nd(5))
}

func TestMapArrayCollection_FirstWhere(t *testing.T) {
	a := []map[string]interface{}{
		{"name": "mike", "sex": 0},
		{"name": "Mary", "sex": 1},
		{"name": "Jane", "sex": 1},
	}
	assert.Equal(t, reflect.DeepEqual(Collect(a).FirstWhere("sex", 1), map[string]interface{}{
		"name": "Mary", "sex": 1,
	}), true)
}

func TestMapCollection_FlatMap(t *testing.T) {
	a := map[string]interface{}{
		"name": "mike",
	}
	assert.Equal(t, reflect.DeepEqual(Collect(a).FlatMap(func(value interface{}) interface{} {
		return "user_" + value.(string)
	}).ToMap(), map[string]interface{}{
		"name": "user_mike",
	}), true)
}

func TestMapCollection_Flip(t *testing.T) {
	a := map[string]interface{}{
		"name": "mike",
	}
	assert.Equal(t, reflect.DeepEqual(Collect(a).Flip().ToMap(), map[string]interface{}{
		"mike": "name",
	}), true)
}

func TestMapCollection_Forget(t *testing.T) {
	a := map[string]interface{}{
		"name": "mike",
		"sex":  1,
	}
	assert.Equal(t, reflect.DeepEqual(Collect(a).Forget("name").ToMap(), map[string]interface{}{
		"sex": 1,
	}), true)
}

func TestNumberArrayCollection_ForPage(t *testing.T) {
	a := []int{2, 3, 4, 5, 6, 7}

	assert.Equal(t, Collect(a).ForPage(2, 3).ToIntArray(), []int{5, 6, 7})
}

func TestMapCollection_Get(t *testing.T) {
	a := map[string]interface{}{
		"name": "mike",
	}
	assert.Equal(t, Collect(a).Get("name"), "mike")
}

func TestMapArrayCollection_GroupBy(t *testing.T) {
	a := []map[string]interface{}{
		{"name": "mike", "sex": 0},
		{"name": "Mary", "sex": 1},
		{"name": "Jane", "sex": 1},
	}
	assert.Equal(t, Collect(a).GroupBy("sex").ToMap()["1"], []map[string]interface{}{
		{"name": "Mary", "sex": 1},
		{"name": "Jane", "sex": 1},
	})
}

func TestMapCollection_Has(t *testing.T) {
	a := map[string]interface{}{
		"name": "mike",
	}
	assert.Equal(t, Collect(a).Has("name"), true)
}

func TestMapArrayCollection_Implode(t *testing.T) {
	a := []map[string]interface{}{
		{"name": "mike", "sex": 0},
		{"name": "Mary", "sex": 1},
		{"name": "Jane", "sex": 1},
	}
	assert.Equal(t, Collect(a).Implode("name", "|"), "mike|Mary|Jane")
}

func TestStringArrayCollection_Intersect(t *testing.T) {
	a := []string{"h", "e", "l", "l", "o"}
	assert.Equal(t, Collect(a).Intersect([]string{"e", "l", "l", "f"}).ToStringArray(), []string{"e", "l", "l"})
}

func TestMapCollection_IntersectByKeys(t *testing.T) {
	a := map[string]interface{}{
		"name": "mike",
		"sex":  0,
	}
	assert.Equal(t, reflect.DeepEqual(Collect(a).IntersectByKeys(map[string]interface{}{
		"name": "mike",
		"city": "beijing",
	}).ToMap(), map[string]interface{}{
		"name": "mike",
	}), true)
}

func TestStringArrayCollection_IsEmpty(t *testing.T) {
	a := []string{"h", "e", "l", "l", "o"}
	assert.Equal(t, Collect(a).IsEmpty(), false)
}

func TestStringArrayCollection_IsNotEmpty(t *testing.T) {
	a := []string{"h", "e", "l", "l", "o"}
	assert.Equal(t, Collect(a).IsNotEmpty(), true)
}

func TestMapArrayCollection_KeyBy(t *testing.T) {
	a := []map[string]interface{}{
		{"name": "mike", "sex": 0},
		{"name": "Mary", "sex": 1},
		{"name": "Jane", "sex": 1},
	}
	assert.Equal(t, Collect(a).KeyBy("sex").ToMap()["1"], []map[string]interface{}{
		{"name": "Jane", "sex": 1},
	})
}

func TestMapCollection_Keys(t *testing.T) {
	a := map[string]interface{}{
		"name": "mike",
		"sex":  1,
	}
	assert.Equal(t, Collect(a).Keys().ToStringArray(), []string{"name", "sex"})
}

func TestBaseCollection_Last(t *testing.T) {

}

func TestBaseCollection_MapToGroups(t *testing.T) {

}

func TestBaseCollection_MapWithKeys(t *testing.T) {

}

func TestBaseCollection_Median(t *testing.T) {

}

func TestBaseCollection_Merge(t *testing.T) {

}

func TestBaseCollection_Pad(t *testing.T) {

}

func TestBaseCollection_Partition(t *testing.T) {

}

func TestBaseCollection_Pop(t *testing.T) {

}

func TestBaseCollection_Push(t *testing.T) {

}

func TestBaseCollection_Random(t *testing.T) {

}

func TestBaseCollection_Reduce(t *testing.T) {

}

func TestBaseCollection_Reject(t *testing.T) {

}

func TestBaseCollection_Reverse(t *testing.T) {

}

func TestBaseCollection_Search(t *testing.T) {

}

func TestBaseCollection_Shift(t *testing.T) {

}

func TestBaseCollection_Shuffle(t *testing.T) {

}

func TestBaseCollection_Slice(t *testing.T) {

}

func TestBaseCollection_Sort(t *testing.T) {

}

func TestBaseCollection_SortByDesc(t *testing.T) {

}

func TestBaseCollection_Splice(t *testing.T) {

}

func TestBaseCollection_Split(t *testing.T) {

}

func TestBaseCollection_Unique(t *testing.T) {

}

func TestBaseCollection_WhereIn(t *testing.T) {

}

func TestBaseCollection_WhereNotIn(t *testing.T) {

}

func TestBaseCollection_ToJson(t *testing.T) {

}

func TestBaseCollection_ToNumberArray(t *testing.T) {

}

func TestBaseCollection_ToIntArray(t *testing.T) {

}

func TestBaseCollection_ToStringArray(t *testing.T) {

}

func TestBaseCollection_ToMap(t *testing.T) {

}

func TestBaseCollection_ToMapArray(t *testing.T) {

}

func TestBaseCollection_Where(t *testing.T) {

}
