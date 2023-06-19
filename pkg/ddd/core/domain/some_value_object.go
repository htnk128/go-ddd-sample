package domain

import (
	"fmt"
)

// SomeValueObject 何らかの型の値を1つ持つ値オブジェクト
type SomeValueObject[V comparable] struct {
	ValueObject

	value V
}

func (svo *SomeValueObject[V]) Value() V {
	return svo.value
}

func (svo *SomeValueObject[V]) String() string {
	return fmt.Sprintf("%v", svo.value)
}

func NewSomeValueObject[V comparable](value V) *SomeValueObject[V] {
	return &SomeValueObject[V]{value: value}
}
