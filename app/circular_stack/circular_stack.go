package circular_stack

import (
	"github.com/p0poff/mock/app/storage"
)

type CircularStack struct {
	data  []storage.Route // массив для хранения данных стека
	top   int             // индекс вершины стека
	size  int             // текущий размер стека
	limit int             // максимальное количество элементов в стеке
}

func NewCircularStack(limit int) *CircularStack {
	return &CircularStack{
		data:  make([]storage.Route, limit),
		top:   -1,
		size:  0,
		limit: limit,
	}
}

func (cs *CircularStack) getTop() int {
	return (cs.top + 1) % cs.limit
}

func (cs *CircularStack) Push(value storage.Route) error {
	cs.top = cs.getTop()
	cs.data[cs.top] = value
	if cs.size < cs.limit {
		cs.size++
	}
	return nil
}

func (cs *CircularStack) All() []storage.Route {
	line := cs.getTop()

	if cs.size >= cs.limit {
		newSlice := make([]storage.Route, len(cs.data)) // Type замените на тип данных в вашем срезе
		copy(newSlice, cs.data)
		return append(newSlice[line:], cs.data[:line]...)
	}

	return cs.data[:line]
}
