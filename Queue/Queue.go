package Queue

import "GTL/Container"

type Queue interface {
	Push(value interface{}) error

	Front() (interface{}, error)

	Pop() (interface{}, error)

	Container.Container
}
