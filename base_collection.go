package collection

import (
	"encoding/json"
	"github.com/shopspring/decimal"
)

type BaseCollection struct {
	value  interface{}
	length int
}

func (c BaseCollection) Value() interface{} {
	return c.value
}

func (c BaseCollection) All() []interface{} {
	panic("not implement")
}

func (c BaseCollection) Avg(key ...string) decimal.Decimal {
	return c.Sum(key...).Div(decimal.New(int64(c.length), 0))
}

func (c BaseCollection) Sum(key ...string) decimal.Decimal {
	panic("not implement")
}

func (c BaseCollection) Min(key ...string) decimal.Decimal {
	panic("not implement")
}

func (c BaseCollection) Max(key ...string) decimal.Decimal {
	panic("not implement")
}

func (c BaseCollection) Join(delimiter string) string {
	panic("not implement")
}

func (c BaseCollection) Combine(value []interface{}) Collection {
	panic("not implement")
}

func (c BaseCollection) Pluck(key string) Collection {
	panic("not implement")
}

func (c BaseCollection) Mode(key ...string) []interface{} {
	panic("not implement")
}

func (c BaseCollection) Only(keys []string) Collection {
	panic("not implement")
}

func (c BaseCollection) Prepend(values ...interface{}) Collection {
	panic("not implement")
}

func (c BaseCollection) Pull(key interface{}) Collection {
	panic("not implement")
}

func (c BaseCollection) Put(key string, value interface{}) Collection {
	panic("not implement")
}

func (c BaseCollection) SortBy(key string) Collection {
	panic("not implement")
}

func (c BaseCollection) Take(num int) Collection {
	panic("not implement")
}

func (c BaseCollection) Average() {
	panic("not implement")
}

func (c BaseCollection) Chunk(num int) MultiDimensionalArrayCollection {
	panic("not implement")
}

func (c BaseCollection) Collapse() Collection {
	panic("not implement")
}

func (c BaseCollection) Concat(value interface{}) Collection {
	panic("not implement")
}

func (c BaseCollection) Contains(value interface{}, callback ...interface{}) bool {
	panic("not implement")
}

func (c BaseCollection) ContainsStrict(value interface{}, callback ...interface{}) bool {
	panic("not implement")
}

func (c BaseCollection) CountBy(callback ...interface{}) map[interface{}]int {
	panic("not implement")
}

func (c BaseCollection) CrossJoin(array ...[]interface{}) MultiDimensionalArrayCollection {
	panic("not implement")
}

func (c BaseCollection) Dd() {
	panic("not implement")
}

func (c BaseCollection) Diff(interface{}) Collection {
	panic("not implement")
}

func (c BaseCollection) DiffAssoc(map[string]interface{}) Collection {
	panic("not implement")
}

func (c BaseCollection) DiffKeys(map[string]interface{}) Collection {
	panic("not implement")
}

func (c BaseCollection) Dump() {
	panic("not implement")
}

func (c BaseCollection) Each(func(item, value interface{}) (interface{}, bool)) Collection {
	panic("not implement")
}

func (c BaseCollection) Every(CB) bool {
	panic("not implement")
}

func (c BaseCollection) Except([]string) Collection {
	panic("not implement")
}

func (c BaseCollection) Filter(CB) Collection {
	panic("not implement")
}

func (c BaseCollection) First(...CB) interface{} {
	panic("not implement")
}

func (c BaseCollection) FirstWhere(key string, values ...interface{}) map[string]interface{} {
	panic("not implement")
}

func (c BaseCollection) FlatMap(func(value interface{}) interface{}) Collection {
	panic("not implement")
}

func (c BaseCollection) Flip() Collection {
	panic("not implement")
}

func (c BaseCollection) Forget(string) Collection {
	panic("not implement")
}

func (c BaseCollection) ForPage(int, int) Collection {
	panic("not implement")
}

func (c BaseCollection) Get(string, ...interface{}) interface{} {
	panic("not implement")
}

func (c BaseCollection) GroupBy(string) Collection {
	panic("not implement")
}

func (c BaseCollection) Has(...string) bool {
	panic("not implement")
}

func (c BaseCollection) Implode(string, string) string {
	panic("not implement")
}

func (c BaseCollection) Intersect([]string) Collection {
	panic("not implement")
}

func (c BaseCollection) IntersectByKeys(map[string]interface{}) Collection {
	panic("not implement")
}

func (c BaseCollection) IsEmpty() bool {
	panic("not implement")
}

func (c BaseCollection) IsNotEmpty() bool {
	panic("not implement")
}

func (c BaseCollection) KeyBy(interface{}) Collection {
	panic("not implement")
}

func (c BaseCollection) Keys() Collection {
	panic("not implement")
}

func (c BaseCollection) Last(...CB) interface{} {
	panic("not implement")
}

func (c BaseCollection) MapToGroups(MapCB) Collection {
	panic("not implement")
}

func (c BaseCollection) MapWithKeys(MapCB) Collection {
	panic("not implement")
}

func (c BaseCollection) Median(key ...string) decimal.Decimal {
	panic("not implement")
}

func (c BaseCollection) Merge(interface{}) Collection {
	panic("not implement")
}

func (c BaseCollection) Nth(...int) Collection {
	panic("not implement")
}

func (c BaseCollection) Pad(int, interface{}) Collection {
	panic("not implement")
}

func (c BaseCollection) Partition(PartCB) (Collection, Collection) {
	panic("not implement")
}

func (c BaseCollection) Pop() interface{} {
	panic("not implement")
}

func (c BaseCollection) Push(interface{}) Collection {
	panic("not implement")
}

func (c BaseCollection) Random(...int) Collection {
	panic("not implement")
}

func (c BaseCollection) Reduce(ReduceCB) interface{} {
	panic("not implement")
}

func (c BaseCollection) Reject(CB) Collection {
	panic("not implement")
}

func (c BaseCollection) Reverse() Collection {
	panic("not implement")
}

func (c BaseCollection) Search(interface{}) int {
	panic("not implement")
}

func (c BaseCollection) Shift() Collection {
	panic("not implement")
}

func (c BaseCollection) Shuffle() Collection {
	panic("not implement")
}

func (c BaseCollection) Slice(...int) Collection {
	panic("not implement")
}

func (c BaseCollection) Sort() Collection {
	panic("not implement")
}

func (c BaseCollection) SortByDesc() Collection {
	panic("not implement")
}

func (c BaseCollection) Split(int) Collection {
	panic("not implement")
}

func (c BaseCollection) Splice(index ...int) Collection {
	panic("not implement")
}

func (c BaseCollection) Unique() Collection {
	panic("not implement")
}

func (c BaseCollection) WhereIn(string, []interface{}) Collection {
	panic("not implement")
}

func (c BaseCollection) WhereNotIn(string, []interface{}) Collection {
	panic("not implement")
}

func (c BaseCollection) ToJson() string {
	s, err := json.Marshal(c.value)
	if err != nil {
		panic(err)
	}
	return string(s)
}

func (c BaseCollection) ToNumberArray() []decimal.Decimal {
	panic("not implement")
}

func (c BaseCollection) ToStringArray() []string {
	panic("not implement")
}

func (c BaseCollection) ToMap() map[string]interface{} {
	panic("not implement")
}

func (c BaseCollection) ToMapArray() []map[string]interface{} {
	panic("not implement")
}

func (c BaseCollection) Where(key string, values ...interface{}) Collection {
	panic("not implement")
}

func (c BaseCollection) Count() int {
	return c.length
}
