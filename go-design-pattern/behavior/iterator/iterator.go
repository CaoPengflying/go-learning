package iterator

type Iterator interface {
	HasNext() bool
	Next()
	CurrentEnum() interface{}
}

type ArrayIterator struct {
	arraySlice []interface{}
	cursor     int
}

func (i *ArrayIterator) HasNext() bool {
	return len(i.arraySlice) > i.cursor
}

func (i *ArrayIterator) Next() {
	i.cursor++
}

func (i *ArrayIterator) CurrentEnum() interface{} {
	return i.arraySlice[i.cursor]
}

func NewArrayIterator(arraySlice []interface{}) *ArrayIterator {
	return &ArrayIterator{
		arraySlice: arraySlice,
		cursor:     0,
	}
}
