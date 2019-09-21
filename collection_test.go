package collection

import (
	"fmt"
	"github.com/magiconair/properties/assert"
	"github.com/shopspring/decimal"
	"reflect"
	"strings"
	"testing"
)

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

func TestStringArrayCollection_Join(t *testing.T) {
	a := []string{"h", "e", "l", "l", "o"}

	assert.Equal(t, Collect(a).Join(""), "hello")
}

func ExampleBaseCollection_Join() {
	a := []string{"h", "e", "l", "l", "o"}
	fmt.Println(Collect(a).Join(""))

	// Output: hello
}

func TestNumberArrayCollection_Sum(t *testing.T) {
	assert.Equal(t, Collect(numbers).Sum().IntPart(), int64(50))

	var floatTest = []float64{143.66, -14.55}
	assert.Equal(t, Collect(floatTest).Sum().String(), "129.11")
}

func ExampleBaseCollection_Sum() {
	var floatTest = []float64{143.66, -14.55}
	fmt.Println(Collect(floatTest).Sum().String())

	// Output: 129.11
}

func TestNumberArrayCollection_Avg(t *testing.T) {
	assert.Equal(t, Collect(numbers).Avg().IntPart(), int64(5))

	var floatTest = []float64{143.66, -14.55}
	assert.Equal(t, Collect(floatTest).Avg().String(), "64.555")
}

func ExampleBaseCollection_Avg() {
	var floatTest = []float64{143.66, -14.55}
	fmt.Println(Collect(floatTest).Avg().String())

	// Output: 64.555
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

func ExampleBaseCollection_Splice() {
	a := []string{"h", "e", "l", "l", "o"}
	fmt.Println(Collect(a).Splice(1, 3).ToStringArray())

	// Output: [e l l]
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

func ExampleBaseCollection_Take() {
	a := []string{"h", "e", "l", "l", "o"}
	fmt.Println(Collect(a).Take(-2).ToStringArray())

	// Output: [l o]
}

func TestCollection_All(t *testing.T) {
	a := []string{"h", "e", "l", "l", "o"}

	assert.Equal(t, Collect(a).All(), []interface{}{"h", "e", "l", "l", "o"})
	assert.Equal(t, len(Collect(numbers).All()), 10)
	assert.Equal(t, Collect(foo).All()[1], map[string]interface{}{"foo": 30})
}

func ExampleBaseCollection_All() {
	a := []string{"h", "e", "l", "l", "o"}
	fmt.Println(Collect(a).All())

	// Output: [h e l l o]
}

func TestCollection_Mode(t *testing.T) {
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

	assert.Equal(t, Collect(`["h", "e", "l", "l", "o", "w", "o", "l", "d"]`).Mode(), []interface{}{"l"})
	assert.Equal(t, Collect(foo2).Mode("foo"), []interface{}{40})
}

func ExampleBaseCollection_Mode() {

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

	fmt.Println(Collect(`["h", "e", "l", "l", "o", "w", "o", "l", "d"]`).Mode())
	fmt.Println(Collect(foo2).Mode("foo"))

	// Output: [l]
	// [40]
}

func TestCollection_Chunk(t *testing.T) {
	a := []string{"h", "e", "l", "l", "o"}

	assert.Equal(t, Collect(foo).Chunk(2).value[0][0], map[string]interface{}{"foo": 10})
	assert.Equal(t, len(Collect(numbers).Chunk(3).value), 4)
	assert.Equal(t, Collect(a).Chunk(3).value[0][2], "l")
}

func ExampleBaseCollection_Chunk() {
	a := []string{"h", "e", "l", "l", "o"}
	fmt.Println(Collect(a).Chunk(3).ToMultiDimensionalArray()[0])

	// Output: [h e l]
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

	c := Collect(numbers)
	assert.Equal(t, c.Contains(10), false)

	var callback CB = func(item, value interface{}) bool {
		return value.(decimal.Decimal).GreaterThan(nd(5))
	}

	assert.Equal(t, c.Contains(callback), true)

	a := []string{"2", "3", "4", "5", "6"}
	assert.Equal(t, Collect(a).Contains("5"), true)

}

func TestCollection_CountBy(t *testing.T) {
	a := []string{"h", "e", "l", "l", "o"}
	assert.Equal(t, Collect(a).CountBy()["l"], 2)
	assert.Equal(t, Collect(numbers).CountBy()[float64(8)], 2)

	c := Collect([]string{"alice@gmail.com", "bob@yahoo.com", "carlos@gmail.com"})

	var cb FilterFun = func(value interface{}) interface{} {
		return strings.Split(value.(string), "@")[1]
	}

	assert.Equal(t, c.CountBy(cb), map[interface{}]int{"gmail.com": 2, "yahoo.com": 1})
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

func ExampleBaseCollection_Diff() {
	a := []string{"h", "e", "l", "l", "o"}
	fmt.Println(Collect(a).Diff([]string{"e", "o"}).ToStringArray())

	// Output: [h l l]
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

func ExampleBaseCollection_Each() {
	a := []int{2, 3, 4, 5, 6, 7}
	fmt.Println(Collect(a).Each(func(item, value interface{}) (interface{}, bool) {
		return value.(decimal.Decimal).IntPart() + 2, false
	}).ToIntArray())

	// Output: [4 5 6 7 8 9]
}

func TestNumberArrayCollection_Every(t *testing.T) {

	a := []int{2, 3, 4, 5, 6, 7}

	assert.Equal(t, Collect(a).Every(func(item, value interface{}) bool {
		return value.(decimal.Decimal).GreaterThanOrEqual(nd(5))
	}), false)
}

func ExampleBaseCollection_Every() {
	a := []int{2, 3, 4, 5, 6, 7}
	fmt.Println(Collect(a).Every(func(item, value interface{}) bool {
		return value.(decimal.Decimal).GreaterThanOrEqual(nd(5))
	}))

	// Output: false
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

func ExampleBaseCollection_Except() {
	a := map[string]interface{}{
		"name": "mike",
		"sex":  1,
	}
	fmt.Println(Collect(a).Except([]string{"name"}).ToMap())

	// Output: map[sex:1]
}

func TestNumberArrayCollection_Filter(t *testing.T) {
	a := []int{2, 3, 4, 5, 6, 7}

	assert.Equal(t, Collect(a).Filter(func(item, value interface{}) bool {
		return value.(decimal.Decimal).IntPart() > 4
	}).ToIntArray(), []int{5, 6, 7})
}

func ExampleBaseCollection_Filter() {
	a := []int{2, 3, 4, 5, 6, 7}
	fmt.Println(Collect(a).Filter(func(item, value interface{}) bool {
		return value.(decimal.Decimal).IntPart() > 4
	}).ToIntArray())

	// Output: [5 6 7]
}

func TestNumberArrayCollection_First(t *testing.T) {
	a := []int{2, 3, 4, 5, 6, 7}

	assert.Equal(t, Collect(a).First(func(item, value interface{}) bool {
		return value.(decimal.Decimal).IntPart() > 4
	}), nd(5))
}

func ExampleBaseCollection_First() {
	a := []int{2, 3, 4, 5, 6, 7}
	fmt.Println(Collect(a).First(func(item, value interface{}) bool {
		return value.(decimal.Decimal).IntPart() > 4
	}))

	// Output: 5
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

func ExampleBaseCollection_FirstWhere() {
	a := []map[string]interface{}{
		{"name": "mike", "sex": 0},
		{"name": "Mary", "sex": 1},
		{"name": "Jane", "sex": 1},
	}
	fmt.Println(Collect(a).FirstWhere("sex", 1))

	// Output: map[name:Mary sex:1]
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

func ExampleBaseCollection_FlatMap() {
	a := map[string]interface{}{
		"name": "mike",
	}
	fmt.Println(Collect(a).FlatMap(func(value interface{}) interface{} {
		return "user_" + value.(string)
	}).ToMap())

	// Output: map[name:user_mike]
}

func TestMapCollection_Flip(t *testing.T) {
	a := map[string]interface{}{
		"name": "mike",
	}
	assert.Equal(t, reflect.DeepEqual(Collect(a).Flip().ToMap(), map[string]interface{}{
		"mike": "name",
	}), true)
}

func ExampleBaseCollection_Flip() {
	a := map[string]interface{}{
		"name": "mike",
	}
	fmt.Println(Collect(a).Flip().ToMap())

	// Output: map[mike:name]
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

func ExampleBaseCollection_Forget() {
	a := map[string]interface{}{
		"name": "mike",
		"sex":  1,
	}
	fmt.Println(Collect(a).Forget("name").ToMap())

	// Output: map[sex:1]
}

func TestNumberArrayCollection_ForPage(t *testing.T) {
	a := []int{2, 3, 4, 5, 6, 7}

	assert.Equal(t, Collect(a).ForPage(2, 3).ToIntArray(), []int{5, 6, 7})
}

func ExampleBaseCollection_ForPage() {
	a := []int{2, 3, 4, 5, 6, 7}

	fmt.Println(Collect(a).ForPage(2, 3).ToIntArray())

	// Output: [5 6 7]
}

func TestMapCollection_Get(t *testing.T) {
	a := map[string]interface{}{
		"name": "mike",
	}
	assert.Equal(t, Collect(a).Get("name"), "mike")
}

func ExampleBaseCollection_Get() {
	a := map[string]interface{}{
		"name": "mike",
	}

	fmt.Println(Collect(a).Get("name"))

	// Output: mike
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

func ExampleBaseCollection_GroupBy() {
	a := []map[string]interface{}{
		{"name": "mike", "sex": 0},
		{"name": "Mary", "sex": 1},
		{"name": "Jane", "sex": 1},
	}

	fmt.Println(Collect(a).GroupBy("sex").ToMap()["1"])

	// Output: [map[name:Mary sex:1] map[name:Jane sex:1]]
}

func TestMapCollection_Has(t *testing.T) {
	a := map[string]interface{}{
		"name": "mike",
	}
	assert.Equal(t, Collect(a).Has("name"), true)
}

func ExampleBaseCollection_Has() {
	a := map[string]interface{}{
		"name": "mike",
	}

	fmt.Println(Collect(a).Has("name"))

	// Output: true
}

func TestMapArrayCollection_Implode(t *testing.T) {
	a := []map[string]interface{}{
		{"name": "mike", "sex": 0},
		{"name": "Mary", "sex": 1},
		{"name": "Jane", "sex": 1},
	}
	assert.Equal(t, Collect(a).Implode("name", "|"), "mike|Mary|Jane")
}

func ExampleBaseCollection_Implode() {
	a := []map[string]interface{}{
		{"name": "mike", "sex": 0},
		{"name": "Mary", "sex": 1},
		{"name": "Jane", "sex": 1},
	}

	fmt.Println(Collect(a).Implode("name", "|"))

	// Output: mike|Mary|Jane
}

func TestStringArrayCollection_Intersect(t *testing.T) {
	a := []string{"h", "e", "l", "l", "o"}
	assert.Equal(t, Collect(a).Intersect([]string{"e", "l", "l", "f"}).ToStringArray(), []string{"e", "l", "l"})
}

func ExampleBaseCollection_Intersect() {
	a := []string{"h", "e", "l", "l", "o"}

	fmt.Println(Collect(a).Intersect([]string{"e", "l", "l", "f"}).ToStringArray())

	// Output: [e l l]
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

func ExampleBaseCollection_IntersectByKeys() {
	a := map[string]interface{}{
		"name": "mike",
		"sex":  0,
	}

	fmt.Println(Collect(a).IntersectByKeys(map[string]interface{}{
		"name": "mike",
		"city": "beijing",
	}).ToMap())

	// Output: map[name:mike]
}

func TestStringArrayCollection_IsEmpty(t *testing.T) {
	a := []string{"h", "e", "l", "l", "o"}
	assert.Equal(t, Collect(a).IsEmpty(), false)
}

func ExampleBaseCollection_IsEmpty() {
	a := []string{"h", "e", "l", "l", "o"}

	fmt.Println(Collect(a).IsEmpty())

	// Output: false
}

func TestStringArrayCollection_IsNotEmpty(t *testing.T) {
	a := []string{"h", "e", "l", "l", "o"}
	assert.Equal(t, Collect(a).IsNotEmpty(), true)
}

func ExampleBaseCollection_IsNotEmpty() {
	a := []string{"h", "e", "l", "l", "o"}

	fmt.Println(Collect(a).IsNotEmpty())

	// Output: true
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

func ExampleBaseCollection_KeyBy() {
	a := []map[string]interface{}{
		{"name": "mike", "sex": 0},
		{"name": "Mary", "sex": 1},
		{"name": "Jane", "sex": 1},
	}

	fmt.Println(Collect(a).KeyBy("sex").ToMap()["1"])

	// Output: [map[name:Jane sex:1]]
}

func TestMapCollection_Keys(t *testing.T) {
	a := map[string]interface{}{
		"name": "mike",
		"sex":  1,
	}
	assert.Equal(t, Collect(a).Keys().ToStringArray(), []string{"name", "sex"})
}

func ExampleBaseCollection_Keys() {
	a := map[string]interface{}{
		"name": "mike",
		"sex":  1,
	}

	fmt.Println(Collect(a).Keys().ToStringArray())
}

func TestBaseCollection_Last(t *testing.T) {
	a := []int{2, 3, 4, 5, 6, 7}

	assert.Equal(t, Collect(a).Last(func(item, value interface{}) bool {
		return value.(decimal.Decimal).IntPart() > 4
	}), nd(7))
}

func ExampleBaseCollection_Last() {
	a := []int{2, 3, 4, 5, 6, 7}

	fmt.Println(Collect(a).Last(func(item, value interface{}) bool {
		return value.(decimal.Decimal).IntPart() > 4
	}))

	// Output: 7
}

func TestBaseCollection_MapToGroups(t *testing.T) {

}

func TestBaseCollection_MapWithKeys(t *testing.T) {

}

func TestBaseCollection_Median(t *testing.T) {
	a := []int{2, 3, 4, 5, 6, 7}

	assert.Equal(t, Collect(a).Median().Equal(nd(4.5)), true)
}

func ExampleBaseCollection_Median() {
	a := []int{2, 3, 4, 5, 6, 7}

	fmt.Println(Collect(a).Median())

	// Output: 4.5
}

func TestBaseCollection_Merge(t *testing.T) {
	a := []int{2, 3, 4, 5, 6, 7}

	assert.Equal(t, Collect(a).Merge([]int{8, 9}).ToIntArray(), []int{2, 3, 4, 5, 6, 7, 8, 9})
}

func ExampleBaseCollection_Merge() {
	a := []int{2, 3, 4, 5, 6, 7}

	fmt.Println(Collect(a).Merge([]int{8, 9}).ToIntArray())

	// Output: [2 3 4 5 6 7 8 9]
}

func TestBaseCollection_Pad(t *testing.T) {
	a := []int{2}

	assert.Equal(t, Collect(a).Pad(3, 0).ToIntArray(), []int{2, 0, 0})
}

func ExampleBaseCollection_Pad() {
	a := []int{2}

	fmt.Println(Collect(a).Pad(3, 0).ToIntArray())

	// Output: [2 0 0]
}

func TestBaseCollection_Partition(t *testing.T) {
	a := []int{2, 3, 4, 5, 6, 7}

	p1, p2 := Collect(a).Partition(func(i int) bool {
		return i > 3
	})

	assert.Equal(t, p1.ToIntArray(), []int{6, 7})
	assert.Equal(t, p2.ToIntArray(), []int{2, 3, 4, 5})
}

func ExampleBaseCollection_Partition() {
	a := []int{2, 3, 4, 5, 6, 7}

	p1, p2 := Collect(a).Partition(func(i int) bool {
		return i > 3
	})

	fmt.Println(p1.ToIntArray())
	fmt.Println(p2.ToIntArray())

	// Output: [6 7]
	// [2 3 4 5]
}

func TestBaseCollection_Pop(t *testing.T) {
	a := []int{2, 3, 4, 5, 6, 7}

	assert.Equal(t, Collect(a).Pop(), nd(7))
}

func ExampleBaseCollection_Pop() {
	a := []int{2, 3, 4, 5, 6, 7}

	fmt.Println(Collect(a).Pop())

	// Output: 7
}

func TestBaseCollection_Push(t *testing.T) {
	a := []int{2, 3, 4, 5, 6, 7}

	assert.Equal(t, Collect(a).Push(8).ToIntArray(), []int{2, 3, 4, 5, 6, 7, 8})
}

func ExampleBaseCollection_Push() {
	a := []int{2, 3, 4, 5, 6, 7}

	fmt.Println(Collect(a).Push(8).ToIntArray())

	// Output: [2 3 4 5 6 7 8]
}

func TestBaseCollection_Random(t *testing.T) {
	a := []int{2, 3, 4, 5, 6, 7}

	assert.Equal(t, Collect(a).Random().Value() != nd(3), true)
}

func ExampleBaseCollection_Random() {
	a := []int{2, 3, 4, 5, 6, 7}
	fmt.Println(Collect(a).Random().Value())
}

func TestBaseCollection_Reduce(t *testing.T) {
	a := []int{2, 3, 4, 5, 6, 7}

	assert.Equal(t, Collect(a).Reduce(func(i interface{}, i2 interface{}) interface{} {
		if i == nil {
			return i2
		} else {
			return i.(decimal.Decimal).Add(i2.(decimal.Decimal))
		}
	}), nd(27))
}

func ExampleBaseCollection_Reduce() {
	a := []int{2, 3, 4, 5, 6, 7}

	fmt.Println(Collect(a).Reduce(func(i interface{}, i2 interface{}) interface{} {
		if i == nil {
			return i2
		} else {
			return i.(decimal.Decimal).Add(i2.(decimal.Decimal))
		}
	}))

	// Output: 27
}

func TestBaseCollection_Reject(t *testing.T) {
	a := []int{2, 3, 4, 5, 6, 7}

	assert.Equal(t, Collect(a).Reject(func(item, value interface{}) bool {
		return value.(decimal.Decimal).GreaterThanOrEqual(nd(3))
	}).ToIntArray(), []int{2})
}

func ExampleBaseCollection_Reject() {
	a := []int{2, 3, 4, 5, 6, 7}

	fmt.Println(Collect(a).Reject(func(item, value interface{}) bool {
		return value.(decimal.Decimal).GreaterThanOrEqual(nd(3))
	}).ToIntArray())

	// Output: [2]
}

func TestBaseCollection_Reverse(t *testing.T) {
	a := []int{2, 3, 4, 5, 6, 7}

	assert.Equal(t, Collect(a).Reverse().ToIntArray(), []int{7, 6, 5, 4, 3, 2})
}

func ExampleBaseCollection_Reverse() {
	a := []int{2, 3, 4, 5, 6, 7}

	fmt.Println(Collect(a).Reverse().ToIntArray())

	// Output: [7 6 5 4 3 2]
}

func TestBaseCollection_Search(t *testing.T) {
	a := []int{2, 3, 4, 5, 6, 7}

	assert.Equal(t, Collect(a).Search(3), 1)

	var callback CB = func(item, value interface{}) bool {
		return value.(decimal.Decimal).GreaterThan(nd(3))
	}

	assert.Equal(t, Collect(a).Search(callback), 2)
}

func ExampleBaseCollection_Search() {
	a := []int{2, 3, 4, 5, 6, 7}

	fmt.Println(Collect(a).Search(3))

	// Output: 1
}

func TestBaseCollection_Shift(t *testing.T) {
	a := []int{2, 3, 4, 5, 6, 7}

	assert.Equal(t, Collect(a).Shift().ToIntArray(), []int{3, 4, 5, 6, 7})
}

func ExampleBaseCollection_Shift() {
	a := []int{2, 3, 4, 5, 6, 7}

	fmt.Println(Collect(a).Shift().ToIntArray())

	// Output: [3 4 5 6 7]
}

func TestBaseCollection_Shuffle(t *testing.T) {
	a := []int{2, 3, 4, 5, 6, 7}

	assert.Equal(t, reflect.DeepEqual(Collect(a).Shuffle().ToIntArray(), a), false)
}

func ExampleBaseCollection_Shuffle() {
	a := []int{2, 3, 4, 5, 6, 7}

	fmt.Println(Collect(a).Shuffle().ToIntArray())
}

func TestBaseCollection_Slice(t *testing.T) {
	a := []int{2, 3, 4, 5, 6, 7}

	assert.Equal(t, Collect(a).Slice(1, 2).ToIntArray(), []int{3, 4})
}

func ExampleBaseCollection_Slice() {
	a := []int{2, 3, 4, 5, 6, 7}

	fmt.Println(Collect(a).Slice(1, 2).ToIntArray())

	// Output: [3 4]
}

func TestBaseCollection_Sort(t *testing.T) {
	a := []int{4, 5, 2, 3, 6, 7}

	assert.Equal(t, Collect(a).Sort().ToIntArray(), []int{2, 3, 4, 5, 6, 7})
}

func ExampleBaseCollection_Sort() {
	a := []int{4, 5, 2, 3, 6, 7}

	fmt.Println(Collect(a).Sort().ToIntArray())

	// Output: [2 3 4 5 6 7]
}

func TestBaseCollection_SortByDesc(t *testing.T) {
	a := []int{4, 5, 2, 3, 6, 7}

	assert.Equal(t, Collect(a).SortByDesc().ToIntArray(), []int{7, 6, 5, 4, 3, 2})
}

func ExampleBaseCollection_SortByDesc() {
	a := []int{4, 5, 2, 3, 6, 7}

	fmt.Println(Collect(a).SortByDesc().ToIntArray())

	// Output: [7 6 5 4 3 2]
}

func TestBaseCollection_Split(t *testing.T) {
	a := []int{2, 3, 4, 5, 6, 7}

	assert.Equal(t, Collect(a).Split(3).Value(), [][]interface{}{{nd(2), nd(3), nd(4)}, {nd(5), nd(6), nd(7)}})
}

func ExampleBaseCollection_Split() {
	a := []int{2, 3, 4, 5, 6, 7}

	fmt.Println(Collect(a).Split(3).Value())
}

func TestBaseCollection_Unique(t *testing.T) {
	a := []int{4, 5, 5, 2, 2, 3, 6, 6, 7}

	assert.Equal(t, Collect(a).Unique().ToIntArray(), []int{4, 5, 2, 3, 6, 7})
}

func ExampleBaseCollection_Unique() {
	a := []int{4, 5, 5, 2, 2, 3, 6, 6, 7}

	fmt.Println(Collect(a).Unique().ToIntArray())

	// Output: [4 5 2 3 6 7]
}

func TestBaseCollection_WhereIn(t *testing.T) {
	a := []map[string]interface{}{
		{"name": "mike", "sex": 0},
		{"name": "Mary", "sex": 1},
		{"name": "Jane", "sex": 2},
	}
	assert.Equal(t, Collect(a).WhereIn("sex", []interface{}{1, 2}).ToMapArray(), []map[string]interface{}{
		{"name": "Mary", "sex": 1},
		{"name": "Jane", "sex": 2},
	})
}

func ExampleBaseCollection_WhereIn() {
	a := []map[string]interface{}{
		{"name": "mike", "sex": 0},
		{"name": "Mary", "sex": 1},
		{"name": "Jane", "sex": 2},
	}

	fmt.Println(Collect(a).WhereIn("sex", []interface{}{1, 2}).ToMapArray())
}

func TestBaseCollection_WhereNotIn(t *testing.T) {
	a := []map[string]interface{}{
		{"name": "mike", "sex": 0},
		{"name": "Mary", "sex": 1},
		{"name": "Jane", "sex": 2},
	}
	assert.Equal(t, Collect(a).WhereNotIn("sex", []interface{}{1, 2}).ToMapArray(), []map[string]interface{}{
		{"name": "mike", "sex": 0},
	})
}

func ExampleBaseCollection_WhereNotIn() {
	a := []map[string]interface{}{
		{"name": "mike", "sex": 0},
		{"name": "Mary", "sex": 1},
		{"name": "Jane", "sex": 2},
	}

	fmt.Println(Collect(a).WhereNotIn("sex", []interface{}{1, 2}).ToMapArray())

	// Output: [map[name:mike sex:0]]
}

func TestBaseCollection_ToJson(t *testing.T) {
	a := []int{4, 5, 5, 2, 2, 3, 6, 6, 7}

	assert.Equal(t, Collect(a).Unique().ToJson(), `["4","5","2","3","6","7"]`)
}

func ExampleBaseCollection_ToJson() {
	a := []int{4, 5, 5, 2, 2, 3, 6, 6, 7}

	fmt.Println(Collect(a).Unique().ToJson())

	// Output: ["4","5","2","3","6","7"]
}

func TestBaseCollection_ToNumberArray(t *testing.T) {
	a := []int{4, 5, 2, 3, 6, 7}

	assert.Equal(t, Collect(a).ToNumberArray(), []decimal.Decimal{nd(4), nd(5), nd(2),
		nd(3), nd(6), nd(7)})
}

func ExampleBaseCollection_ToNumberArray() {
	a := []int{4, 5, 2, 3, 6, 7}

	fmt.Println(Collect(a).ToNumberArray())
}

func TestBaseCollection_ToIntArray(t *testing.T) {
	a := []int{4, 5, 2, 3, 6, 7}

	assert.Equal(t, Collect(a).ToIntArray(), []int{4, 5, 2, 3, 6, 7})
}

func ExampleBaseCollection_ToIntArray() {
	a := []int{4, 5, 2, 3, 6, 7}

	fmt.Println(Collect(a).ToIntArray())

	// Output: [4 5 2 3 6 7]
}

func TestBaseCollection_ToStringArray(t *testing.T) {
	a := []string{"h", "e", "l", "l", "o"}
	assert.Equal(t, Collect(a).ToStringArray(), []string{"h", "e", "l", "l", "o"})
}

func ExampleBaseCollection_ToStringArray() {
	a := []string{"h", "e", "l", "l", "o"}

	fmt.Println(Collect(a).ToStringArray())

	// Output: [h e l l o]
}

func TestBaseCollection_ToMap(t *testing.T) {
	a := map[string]interface{}{
		"name": "mike",
		"sex":  0,
	}
	assert.Equal(t, reflect.DeepEqual(Collect(a).ToMap(), map[string]interface{}{
		"name": "mike",
		"sex":  0,
	}), true)
}

func ExampleBaseCollection_ToMap() {
	a := []string{"h", "e", "l", "l", "o"}

	fmt.Println(Collect(a).ToStringArray())

	// Output: [h e l l o]
}

func TestBaseCollection_ToMapArray(t *testing.T) {
	a := []map[string]interface{}{
		{"name": "Mary", "sex": 1},
		{"name": "Jane", "sex": 2},
	}
	assert.Equal(t, Collect(a).ToMapArray(), []map[string]interface{}{
		{"name": "Mary", "sex": 1},
		{"name": "Jane", "sex": 2},
	})
}

func ExampleBaseCollection_ToMapArray() {
	a := []map[string]interface{}{
		{"name": "Mary", "sex": 1},
		{"name": "Jane", "sex": 2},
	}

	fmt.Println(Collect(a).ToMapArray())

	// Output: [map[name:Mary sex:1] map[name:Jane sex:2]]
}

func TestBaseCollection_Where(t *testing.T) {
	a := []map[string]interface{}{
		{"name": "mike", "sex": 0},
		{"name": "Mary", "sex": 1},
		{"name": "Jane", "sex": 2},
	}
	assert.Equal(t, Collect(a).Where("sex", 2).ToMapArray(), []map[string]interface{}{
		{"name": "Jane", "sex": 2},
	})
}

func ExampleBaseCollection_Where() {
	a := []map[string]interface{}{
		{"name": "mike", "sex": 0},
		{"name": "Mary", "sex": 1},
		{"name": "Jane", "sex": 2},
	}

	Collect(a).Where("sex", ">", 1).ToMapArray()
	Collect(a).Where("sex", "<", 1).ToMapArray()
	fmt.Println(Collect(a).Where("sex", 2).ToMapArray())

	// Output: [map[name:Jane sex:2]]
}

func TestBaseCollection_Length(t *testing.T) {
	a := []map[string]interface{}{
		{"name": "mike", "sex": 0},
		{"name": "Mary", "sex": 1},
		{"name": "Jane", "sex": 2},
	}
	assert.Equal(t, Collect(a).Length(), 3)
}

func ExampleBaseCollection_Length() {
	a := []map[string]interface{}{
		{"name": "mike", "sex": 0},
		{"name": "Mary", "sex": 1},
		{"name": "Jane", "sex": 2},
	}

	fmt.Println(Collect(a).Length())

	// Output: 3
}

func TestBaseCollection_Select(t *testing.T) {
	a := []map[string]interface{}{
		{"name": "mike", "sex": 0},
		{"name": "Mary", "sex": 1},
		{"name": "Jane", "sex": 2},
	}
	assert.Equal(t, Collect(a).Select("sex").ToMapArray(), []map[string]interface{}{
		{"sex": 0},
		{"sex": 1},
		{"sex": 2},
	})
}

func ExampleBaseCollection_Select() {
	a := []map[string]interface{}{
		{"name": "mike", "sex": 0},
		{"name": "Mary", "sex": 1},
		{"name": "Jane", "sex": 2},
	}

	fmt.Println(Collect(a).Select("sex").ToMapArray())

	// Output: [map[sex:0] map[sex:1] map[sex:2]]
}

func TestBaseCollection_ToStruct(t *testing.T) {
	a := []map[string]interface{}{
		{"name": "mike", "sex": 0},
		{"name": "Mary", "sex": 1},
		{"name": "Jane", "sex": 2},
	}

	type People struct {
		Name string
		Sex  int
	}

	var people = make([]People, 3)

	Collect(a).ToStruct(&people)

	assert.Equal(t, people[0].Name, "mike")
}

func ExampleBaseCollection_ToStruct() {
	a := []map[string]interface{}{
		{"name": "mike", "sex": 0},
		{"name": "Mary", "sex": 1},
		{"name": "Jane", "sex": 2},
	}

	type People struct {
		Name string
		Sex  int
	}

	var people = make([]People, 3)

	Collect(a).ToStruct(&people)

	fmt.Println(people)

	// Output: [{mike 0} {Mary 1} {Jane 2}]
}
