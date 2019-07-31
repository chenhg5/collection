package main

import "github.com/shopspring/decimal"

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

func (c BaseCollection) EachSpread() {
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

func (c BaseCollection) Flatten() Collection {
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

func (c BaseCollection) KeyBy() {
	panic("not implement")
}

func (c BaseCollection) Keys() {
	panic("not implement")
}

func (c BaseCollection) Last() {
	panic("not implement")
}

func (c BaseCollection) Macro() {
	panic("not implement")
}

func (c BaseCollection) Make() {
	panic("not implement")
}

func (c BaseCollection) Map() {
	panic("not implement")
}

func (c BaseCollection) MapInto() {
	panic("not implement")
}

func (c BaseCollection) MapSpread() {
	panic("not implement")
}

func (c BaseCollection) MapToGroups() {
	panic("not implement")
}

func (c BaseCollection) MapWithKeys() {
	panic("not implement")
}

func (c BaseCollection) Median() {
	panic("not implement")
}

func (c BaseCollection) Merge() {
	panic("not implement")
}

func (c BaseCollection) Nth() {
	panic("not implement")
}

func (c BaseCollection) Pad() {
	panic("not implement")
}

func (c BaseCollection) Partition() {
	panic("not implement")
}

func (c BaseCollection) Pipe() {
	panic("not implement")
}

func (c BaseCollection) Pop() {
	panic("not implement")
}

func (c BaseCollection) Push() {
	panic("not implement")
}

func (c BaseCollection) Random() {
	panic("not implement")
}

func (c BaseCollection) Reduce() {
	panic("not implement")
}

func (c BaseCollection) Reject() {
	panic("not implement")
}

func (c BaseCollection) Reverse() {
	panic("not implement")
}

func (c BaseCollection) Search() {
	panic("not implement")
}

func (c BaseCollection) Shift() {
	panic("not implement")
}

func (c BaseCollection) Shuffle() {
	panic("not implement")
}

func (c BaseCollection) Slice() {
	panic("not implement")
}

func (c BaseCollection) Some() {
	panic("not implement")
}

func (c BaseCollection) Sort() {
	panic("not implement")
}

func (c BaseCollection) SortByDesc() {
	panic("not implement")
}

func (c BaseCollection) SortKeys() {
	panic("not implement")
}

func (c BaseCollection) SortKeysDesc() {
	panic("not implement")
}

func (c BaseCollection) Split() {
	panic("not implement")
}

func (c BaseCollection) Splice(index ...int) Collection {
	panic("not implement")
}

func (c BaseCollection) Tap() {
	panic("not implement")
}

func (c BaseCollection) Times() {
	panic("not implement")
}

func (c BaseCollection) Transform() {
	panic("not implement")
}

func (c BaseCollection) Union() {
	panic("not implement")
}

func (c BaseCollection) Unique() {
	panic("not implement")
}

func (c BaseCollection) UniqueStrict() {
	panic("not implement")
}

func (c BaseCollection) Unless() {
	panic("not implement")
}

func (c BaseCollection) UnlessEmpty() {
	panic("not implement")
}

func (c BaseCollection) UnlessNotEmpty() {
	panic("not implement")
}

func (c BaseCollection) Unwrap() {
	panic("not implement")
}

func (c BaseCollection) Values() {
	panic("not implement")
}

func (c BaseCollection) When() {
	panic("not implement")
}

func (c BaseCollection) WhenEmpty() {
	panic("not implement")
}

func (c BaseCollection) WhenNotEmpty() {
	panic("not implement")
}

func (c BaseCollection) WhereStrict() {
	panic("not implement")
}

func (c BaseCollection) WhereBetween() {
	panic("not implement")
}

func (c BaseCollection) WhereIn() {
	panic("not implement")
}

func (c BaseCollection) WhereInStrict() {
	panic("not implement")
}

func (c BaseCollection) WhereInstanceOf() {
	panic("not implement")
}

func (c BaseCollection) WhereNotBetween() {
	panic("not implement")
}

func (c BaseCollection) WhereNotIn() {
	panic("not implement")
}

func (c BaseCollection) WhereNotInStrict() {
	panic("not implement")
}

func (c BaseCollection) Wrap() {
	panic("not implement")
}

func (c BaseCollection) Zip() {
	panic("not implement")
}

func (c BaseCollection) ToJson() string {
	panic("not implement")
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

func (c BaseCollection) Where(key string, value interface{}) Collection {
	panic("not implement")
}

func (c BaseCollection) Count() int {
	return c.length
}
