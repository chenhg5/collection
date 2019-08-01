# Collection

Collection provides a fluent, convenient wrapper for working with arrays of data.

You can easily convert a map or an array into a Collection with the method ```Collect()```.
And then you can use the powerful and graceful apis of Collection to process the data.

In general, Collection are immutable, meaning every Collection method returns an entirely new Collection instance

## doc

[here](https://godoc.org/github.com/chenhg5/collection#Collection)

## example

```golang
a := []int{2,3,4,5,6,7}

Collect(a).Each(func(item, value interface{} {
    return value.(decimal.Decimal).IntPart() + 2, false
}).ToNumberArray()

// []decimal.Decimal{4,5,6,7,8,9}

b := []map[string]interface{}{
    {"name": "Jack", "sex": 0},
    {"name": "Mary", "sex": 1},
    {"name": "Jane", "sex": 1},
}

Collect(b).Where("name", "Jack").ToMapArray()[0]

// map[string]interface{}{"name": "Jack", "sex": 0}

``` 

more saying is useless, start coding!

## Contribution

pr is very welcome. 