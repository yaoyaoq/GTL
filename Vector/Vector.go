package Vector

import "GTL/Container"

type Vector interface {
	Container.Container

	PushBack(value interface{}) error

	PopBack() (interface{}, error)

	Set(index, value int) error

	At(index int) (interface{}, error)

	Remove(start, end int) error

	Find(
		value interface{},
		less func(interface{}, interface{}) bool,
	) int
}
