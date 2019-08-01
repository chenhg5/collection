package collection

type MultiDimensionalArrayCollection struct {
	value [][]interface{}
	BaseCollection
}

// Collapse collapses a collection of arrays into a single, flat collection.
func (c MultiDimensionalArrayCollection) Collapse() Collection {
	if len(c.value[0]) == 0 {
		return Collect([]interface{}{})
	}
	length := 0
	for _, v := range c.value {
		length += len(v)
	}

	d := make([]interface{}, length)
	index := 0
	for _, innerSlice := range c.value {
		for _, v := range innerSlice {
			d[index] = v
			index++
		}
	}

	return Collect(d)
}

// Concat appends the given array or collection values onto the end of the collection.
func (c MultiDimensionalArrayCollection) Concat(value interface{}) Collection {
	return MultiDimensionalArrayCollection{
		value:          append(c.value, value.([][]interface{})...),
		BaseCollection: BaseCollection{length: c.length + len(value.([][]interface{}))},
	}
}

// Dd dumps the collection's items and ends execution of the script.
func (c MultiDimensionalArrayCollection) Dd() {
	dd(c)
}

// Dump dumps the collection's items.
func (c MultiDimensionalArrayCollection) Dump() {
	dump(c)
}
