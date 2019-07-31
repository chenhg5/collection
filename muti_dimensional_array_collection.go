package collection

type MultiDimensionalArrayCollection struct {
	value [][]interface{}
	BaseCollection
}

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

func (c MultiDimensionalArrayCollection) Concat(value interface{}) Collection {
	return MultiDimensionalArrayCollection{
		value:          append(c.value, value.([][]interface{})...),
		BaseCollection: BaseCollection{length: c.length + len(value.([][]interface{}))},
	}
}

func (c MultiDimensionalArrayCollection) Dd() {
	dd(c)
}

func (c MultiDimensionalArrayCollection) Dump() {
	dump(c)
}
