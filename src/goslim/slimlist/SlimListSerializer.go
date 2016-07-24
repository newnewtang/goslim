package slimlist

import (
	"fmt"
)

const (
	LIST_OVERHEAD    = 9
	ELEMENT_OVERHEAD = 8
)

func (s *SlimList) SerializedLength() int {
	length := LIST_OVERHEAD

	for iterator := s.head; iterator != nil; iterator = iterator.next {
		length += len(iterator.GetStringWithNullAsANormalString()) + ELEMENT_OVERHEAD
	}

	return length
}

func (s *SlimList) Serialize() string {
	buf := fmt.Sprintf("[%06d:", s.GetLength())

	for iterator := s.head; iterator != nil; iterator = iterator.next {
		nodeString := iterator.GetStringWithNullAsANormalString()
		buf += fmt.Sprintf("%06d:%s:", len(nodeString), nodeString)
	}
	buf += "]"
	return buf
}
