package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"github.com/shopspring/decimal"
)

func Collect(src interface{}) Collection {
	switch src.(type) {
	case []string:
		var c StringArrayCollection
		c.value = src.([]string)
		c.length = len(src.([]string))
		return c
	case []map[string]interface{}:
		var c MapArrayCollection
		c.value = src.([]map[string]interface{})
		c.length = len(src.([]map[string]interface{}))
		return c
	case map[string]interface{}:
		var c MapCollection
		c.value = src.(map[string]interface{})
		c.length = len(src.(map[string]interface{}))
		return c
	case []int:
		var c NumberArrayCollection
		var d = make([]decimal.Decimal, len(src.([]int)))
		for k, v := range src.([]int) {
			d[k] = decimal.New(int64(v), 0)
		}
		c.value = d
		c.length = len(src.([]int))
		return c
	case []int8:
		var c NumberArrayCollection
		var d = make([]decimal.Decimal, len(src.([]int8)))
		for k, v := range src.([]int8) {
			d[k] = decimal.New(int64(v), 0)
		}
		c.value = d
		c.length = len(src.([]int8))
		return c
	case []int16:
		var c NumberArrayCollection
		var d = make([]decimal.Decimal, len(src.([]int16)))
		for k, v := range src.([]int16) {
			d[k] = decimal.New(int64(v), 0)
		}
		c.value = d
		c.length = len(src.([]int16))
		return c
	case []int32:
		var c NumberArrayCollection
		var d = make([]decimal.Decimal, len(src.([]int32)))
		for k, v := range src.([]int32) {
			d[k] = decimal.New(int64(v), 0)
		}
		c.value = d
		c.length = len(src.([]int32))
		return c
	case []int64:
		var c NumberArrayCollection
		var d = make([]decimal.Decimal, len(src.([]int64)))
		for k, v := range src.([]int64) {
			d[k] = decimal.New(v, 0)
		}
		c.value = d
		c.length = len(src.([]int64))
		return c
	case []float32:
		var c NumberArrayCollection
		var f = make([]decimal.Decimal, len(src.([]float32)))
		for k, v := range src.([]float32) {
			f[k] = decimal.NewFromFloat32(v)
		}
		c.value = f
		c.length = len(src.([]float32))
		return c
	case []float64:
		var c NumberArrayCollection
		var f = make([]decimal.Decimal, len(src.([]float64)))
		for k, v := range src.([]float64) {
			f[k] = decimal.NewFromFloat(v)
		}
		c.value = f
		c.length = len(src.([]float64))
		return c
	case []interface{}:
		if len(src.([]interface{})) == 0 {
			panic("wrong value")
		}
		switch src.([]interface{})[0].(type) {
		case map[string]interface{}:
			var c MapArrayCollection
			var f = make([]map[string]interface{}, len(src.([]interface{})))
			for k, v := range src.([]interface{}) {
				f[k] = v.(map[string]interface{})
			}
			c.value = f
			c.length = len(src.([]interface{}))
			return c
		case decimal.Decimal:
			var c NumberArrayCollection
			var f = make([]decimal.Decimal, len(src.([]interface{})))
			for k, v := range src.([]interface{}) {
				f[k] = v.(decimal.Decimal)
			}
			c.value = f
			c.length = len(src.([]interface{}))
			return c
		case string:
			var c StringArrayCollection
			var f = make([]string, len(src.([]interface{})))
			for k, v := range src.([]interface{}) {
				f[k] = v.(string)
			}
			c.value = f
			c.length = len(src.([]interface{}))
			return c
		case []uint8:
			var c StringArrayCollection
			var f = make([]string, len(src.([]interface{})))
			for k, v := range src.([]interface{}) {
				f[k] = string(v.([]uint8))
			}
			c.value = f
			c.length = len(src.([]interface{}))
			return c
		default:
			panic("wrong type")
		}
	default:
		panic("wrong type")
	}
}

type Collection interface {
	Value() interface{}

	// reference: https://laravel.com/docs/5.8/collections#method-all
	All() []interface{}

	// reference: https://laravel.com/docs/5.8/collections#method-avg
	Avg(key ...string) decimal.Decimal

	// reference: https://laravel.com/docs/5.8/collections#method-sum
	Sum(key ...string) decimal.Decimal

	// reference: https://laravel.com/docs/5.8/collections#method-min
	Min(key ...string) decimal.Decimal

	// reference: https://laravel.com/docs/5.8/collections#method-max
	Max(key ...string) decimal.Decimal

	// reference: https://laravel.com/docs/5.8/collections#method-join
	Join(delimiter string) string

	// reference: https://laravel.com/docs/5.8/collections#method-combine
	Combine(value []interface{}) Collection

	// reference: https://laravel.com/docs/5.8/collections#method-count
	Count() int

	// reference: https://laravel.com/docs/5.8/collections#method-pluck
	Pluck(key string) Collection

	// reference: https://laravel.com/docs/5.8/collections#method-mode
	Mode(key ...string) []interface{}

	// reference: https://laravel.com/docs/5.8/collections#method-only
	Only(keys []string) Collection

	// reference: https://laravel.com/docs/5.8/collections#method-prepend
	Prepend(values ...interface{}) Collection

	// reference: https://laravel.com/docs/5.8/collections#method-pull
	Pull(key interface{}) Collection

	// reference: https://laravel.com/docs/5.8/collections#method-put
	Put(key string, value interface{}) Collection

	// reference: https://laravel.com/docs/5.8/collections#method-sortby
	SortBy(key string) Collection

	// reference: https://laravel.com/docs/5.8/collections#method-take
	Take(num int) Collection

	// reference: https://laravel.com/docs/5.8/collections#method-average
	Average()

	// reference: https://laravel.com/docs/5.8/collections#method-chunk
	Chunk(num int) MultiDimensionalArrayCollection

	// reference: https://laravel.com/docs/5.8/collections#method-collapse
	Collapse() Collection

	// reference: https://laravel.com/docs/5.8/collections#method-concat
	Concat(value interface{}) Collection

	// reference: https://laravel.com/docs/5.8/collections#method-contains
	Contains(value interface{}, callback ...interface{}) bool

	// reference: https://laravel.com/docs/5.8/collections#method-containsStrict
	ContainsStrict(value interface{}, callback ...interface{}) bool

	// reference: https://laravel.com/docs/5.8/collections#method-countBy
	CountBy(callback ...interface{}) map[interface{}]int

	// reference: https://laravel.com/docs/5.8/collections#method-crossJoin
	CrossJoin(array ...[]interface{}) MultiDimensionalArrayCollection

	// reference: https://laravel.com/docs/5.8/collections#method-dd
	Dd()

	// reference: https://laravel.com/docs/5.8/collections#method-diff
	Diff(interface{}) Collection

	// reference: https://laravel.com/docs/5.8/collections#method-diffassoc
	DiffAssoc(map[string]interface{}) Collection

	// reference: https://laravel.com/docs/5.8/collections#method-diffkeys
	DiffKeys(map[string]interface{}) Collection

	// reference: https://laravel.com/docs/5.8/collections#method-dump
	Dump()

	// reference: https://laravel.com/docs/5.8/collections#method-duplicates

	// reference: https://laravel.com/docs/5.8/collections#method-each
	Each(func(item, value interface{}) (interface{}, bool)) Collection

	// reference: https://laravel.com/docs/5.8/collections#method-eachSpread
	EachSpread()

	// reference: https://laravel.com/docs/5.8/collections#method-every
	Every(CB) bool

	// reference: https://laravel.com/docs/5.8/collections#method-except
	Except([]string) Collection

	// reference: https://laravel.com/docs/5.8/collections#method-filter
	Filter(CB) Collection

	// reference: https://laravel.com/docs/5.8/collections#method-first
	First(...CB) interface{}

	// reference: https://laravel.com/docs/5.8/collections#method-firstWhere
	FirstWhere(key string, values ...interface{}) map[string]interface{}

	// reference: https://laravel.com/docs/5.8/collections#method-flatMap
	FlatMap(func(value interface{}) interface{}) Collection

	// reference: https://laravel.com/docs/5.8/collections#method-flatten
	Flatten() Collection

	// reference: https://laravel.com/docs/5.8/collections#method-flip
	Flip() Collection

	// reference: https://laravel.com/docs/5.8/collections#method-forget
	Forget(string) Collection

	// reference: https://laravel.com/docs/5.8/collections#method-forPage
	ForPage(int, int) Collection

	// reference: https://laravel.com/docs/5.8/collections#method-get
	Get(string, ...interface{}) interface{}

	// reference: https://laravel.com/docs/5.8/collections#method-groupBy
	GroupBy(string) Collection

	// reference: https://laravel.com/docs/5.8/collections#method-has
	Has(...string) bool

	// reference: https://laravel.com/docs/5.8/collections#method-implode
	Implode(string, string) string

	// reference: https://laravel.com/docs/5.8/collections#method-intersect
	Intersect([]string) Collection

	// reference: https://laravel.com/docs/5.8/collections#method-intersectByKeys
	IntersectByKeys(map[string]interface{}) Collection

	// reference: https://laravel.com/docs/5.8/collections#method-isEmpty
	IsEmpty() bool

	// reference: https://laravel.com/docs/5.8/collections#method-isNotEmpty
	IsNotEmpty() bool

	// reference: https://laravel.com/docs/5.8/collections#method-join

	// reference: https://laravel.com/docs/5.8/collections#method-keyBy
	KeyBy()

	// reference: https://laravel.com/docs/5.8/collections#method-keys
	Keys()

	// reference: https://laravel.com/docs/5.8/collections#method-last
	Last()

	// reference: https://laravel.com/docs/5.8/collections#method-macro
	Macro()

	// reference: https://laravel.com/docs/5.8/collections#method-make
	Make()

	// reference: https://laravel.com/docs/5.8/collections#method-map
	Map()

	// reference: https://laravel.com/docs/5.8/collections#method-mapInto
	MapInto()

	// reference: https://laravel.com/docs/5.8/collections#method-mapSpread
	MapSpread()

	// reference: https://laravel.com/docs/5.8/collections#method-mapToGroups
	MapToGroups()

	// reference: https://laravel.com/docs/5.8/collections#method-mapWithKeys
	MapWithKeys()

	// reference: https://laravel.com/docs/5.8/collections#method-median
	Median()

	// reference: https://laravel.com/docs/5.8/collections#method-merge
	Merge()

	// reference: https://laravel.com/docs/5.8/collections#method-nth
	Nth()

	// reference: https://laravel.com/docs/5.8/collections#method-pad
	Pad()

	// reference: https://laravel.com/docs/5.8/collections#method-partition
	Partition()

	// reference: https://laravel.com/docs/5.8/collections#method-pipe
	Pipe()

	// reference: https://laravel.com/docs/5.8/collections#method-pop
	Pop()

	// reference: https://laravel.com/docs/5.8/collections#method-push
	Push()

	// reference: https://laravel.com/docs/5.8/collections#method-random
	Random()

	// reference: https://laravel.com/docs/5.8/collections#method-reduce
	Reduce()

	// reference: https://laravel.com/docs/5.8/collections#method-reject
	Reject()

	// reference: https://laravel.com/docs/5.8/collections#method-reverse
	Reverse()

	// reference: https://laravel.com/docs/5.8/collections#method-search
	Search()

	// reference: https://laravel.com/docs/5.8/collections#method-shift
	Shift()

	// reference: https://laravel.com/docs/5.8/collections#method-shuffle
	Shuffle()

	// reference: https://laravel.com/docs/5.8/collections#method-slice
	Slice()

	// reference: https://laravel.com/docs/5.8/collections#method-some
	Some()

	// reference: https://laravel.com/docs/5.8/collections#method-sort
	Sort()

	// reference: https://laravel.com/docs/5.8/collections#method-sortByDesc
	SortByDesc()

	// reference: https://laravel.com/docs/5.8/collections#method-sortKeys
	SortKeys()

	// reference: https://laravel.com/docs/5.8/collections#method-sortKeysDesc
	SortKeysDesc()

	// reference: https://laravel.com/docs/5.8/collections#method-splice
	Splice(index ...int) Collection

	// reference: https://laravel.com/docs/5.8/collections#method-split
	Split()

	// reference: https://laravel.com/docs/5.8/collections#method-tap
	Tap()

	// reference: https://laravel.com/docs/5.8/collections#method-times
	Times()

	// reference: https://laravel.com/docs/5.8/collections#method-transform
	Transform()

	// reference: https://laravel.com/docs/5.8/collections#method-union
	Union()

	// reference: https://laravel.com/docs/5.8/collections#method-unique
	Unique()

	// reference: https://laravel.com/docs/5.8/collections#method-uniqueStrict
	UniqueStrict()

	// reference: https://laravel.com/docs/5.8/collections#method-unless
	Unless()

	// reference: https://laravel.com/docs/5.8/collections#method-unlessEmpty
	UnlessEmpty()

	// reference: https://laravel.com/docs/5.8/collections#method-unlessNotEmpty
	UnlessNotEmpty()

	// reference: https://laravel.com/docs/5.8/collections#method-unwrap
	Unwrap()

	// reference: https://laravel.com/docs/5.8/collections#method-values
	Values()

	// reference: https://laravel.com/docs/5.8/collections#method-when
	When()

	// reference: https://laravel.com/docs/5.8/collections#method-whenEmpty
	WhenEmpty()

	// reference: https://laravel.com/docs/5.8/collections#method-whenNotEmpty
	WhenNotEmpty()

	// reference: https://laravel.com/docs/5.8/collections#method-whereStrict
	WhereStrict()

	// reference: https://laravel.com/docs/5.8/collections#method-whereBetween
	WhereBetween()

	// reference: https://laravel.com/docs/5.8/collections#method-whereIn
	WhereIn()

	// reference: https://laravel.com/docs/5.8/collections#method-whereInStrict
	WhereInStrict()

	// reference: https://laravel.com/docs/5.8/collections#method-whereInstanceOf
	WhereInstanceOf()

	// reference: https://laravel.com/docs/5.8/collections#method-whereNotBetween
	WhereNotBetween()

	// reference: https://laravel.com/docs/5.8/collections#method-whereNotIn
	WhereNotIn()

	// reference: https://laravel.com/docs/5.8/collections#method-whereNotInStrict
	WhereNotInStrict()

	// reference: https://laravel.com/docs/5.8/collections#method-wrap
	Wrap()

	// reference: https://laravel.com/docs/5.8/collections#method-zip
	Zip()

	ToJson() string

	ToNumberArray() []decimal.Decimal

	ToStringArray() []string

	ToMap() map[string]interface{}

	ToMapArray() []map[string]interface{}

	Where(key string, value interface{}) Collection
}

func newDecimalFromInterface(a interface{}) decimal.Decimal {
	var d decimal.Decimal

	switch a.(type) {
	case uint:
		d = decimal.New(int64(a.(uint)), 0)
	case uint8:
		d = decimal.New(int64(a.(uint8)), 0)
	case uint16:
		d = decimal.New(int64(a.(uint16)), 0)
	case uint32:
		d = decimal.New(int64(a.(uint32)), 0)
	case uint64:
		d = decimal.New(int64(a.(uint64)), 0)
	case int:
		d = decimal.New(int64(a.(int)), 0)
	case int8:
		d = decimal.New(int64(a.(int8)), 0)
	case int16:
		d = decimal.New(int64(a.(int16)), 0)
	case int32:
		d = decimal.New(int64(a.(int32)), 0)
	case int64:
		d = decimal.New(a.(int64), 0)
	case float32:
		d = decimal.NewFromFloat32(a.(float32))
	case float64:
		d = decimal.NewFromFloat(a.(float64))
	default:
		panic("wrong type")
	}

	return d
}

func isTrue(a interface{}) bool {
	switch a.(type) {
	case uint:
		return a.(uint) != uint(0)
	case uint8:
		return a.(uint8) != uint8(0)
	case uint16:
		return a.(uint16) != uint16(0)
	case uint32:
		return a.(uint32) != uint32(0)
	case uint64:
		return a.(uint64) != uint64(0)
	case int:
		return a.(int) != int(0)
	case int8:
		return a.(int8) != int8(0)
	case int16:
		return a.(int16) != int16(0)
	case int32:
		return a.(int32) != int32(0)
	case int64:
		return a.(int64) != int64(0)
	case float32:
		return a.(float32) != float32(0)
	case float64:
		return a.(float64) != float64(0)
	case string:
		return a.(string) != ""
	case bool:
		return a.(bool)
	default:
		return false
	}
}

func nd(a interface{}) decimal.Decimal {
	return newDecimalFromInterface(a)
}

type CB func(item, value interface{}) bool

func copyMap(m map[string]interface{}) map[string]interface{} {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	dec := gob.NewDecoder(&buf)
	err := enc.Encode(m)
	if err != nil {
		panic(err)
	}
	var cm map[string]interface{}
	err = dec.Decode(&cm)
	if err != nil {
		panic(err)
	}
	return cm
}

func dd(c Collection) {
	fmt.Println(c)
}

func dump(c Collection) {
	fmt.Println(c)
}

func newDecimalArray(src interface{}) []decimal.Decimal {
	switch src.(type) {
	case []int:
		var d = make([]decimal.Decimal, len(src.([]int)))
		for k, v := range src.([]int) {
			d[k] = decimal.New(int64(v), 0)
		}
		return d
	case []int8:
		var d = make([]decimal.Decimal, len(src.([]int8)))
		for k, v := range src.([]int8) {
			d[k] = decimal.New(int64(v), 0)
		}
		return d
	case []int16:
		var d = make([]decimal.Decimal, len(src.([]int16)))
		for k, v := range src.([]int16) {
			d[k] = decimal.New(int64(v), 0)
		}
		return d
	case []int32:
		var d = make([]decimal.Decimal, len(src.([]int32)))
		for k, v := range src.([]int32) {
			d[k] = decimal.New(int64(v), 0)
		}
		return d
	case []int64:
		var d = make([]decimal.Decimal, len(src.([]int64)))
		for k, v := range src.([]int64) {
			d[k] = decimal.New(v, 0)
		}
		return d
	case []float32:
		var f = make([]decimal.Decimal, len(src.([]float32)))
		for k, v := range src.([]float32) {
			f[k] = decimal.NewFromFloat32(v)
		}
		return f
	case []float64:
		var f = make([]decimal.Decimal, len(src.([]float64)))
		for k, v := range src.([]float64) {
			f[k] = decimal.NewFromFloat(v)
		}
		return f
	default:
		return nil
	}
}
