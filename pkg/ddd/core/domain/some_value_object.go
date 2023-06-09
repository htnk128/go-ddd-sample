package domain

import (
	"fmt"
)

// SomeValueObject 何らかの型の値を1つ持つ値オブジェクト
type SomeValueObject[V comparable] struct {
	ValueObject

	Value V
}

func (svo *SomeValueObject[V]) Equals(other *SomeValueObject[V]) bool {
	return svo.SameValueAs(other)
}

func (svo *SomeValueObject[V]) SameValueAs(other *SomeValueObject[V]) bool {
	return svo.Value == other.Value
}

func (svo *SomeValueObject[V]) String() string {
	return fmt.Sprintf("%v", svo.Value)
}
