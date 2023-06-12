package domain

// ValueObject DDDにおける値オブジェクトの概念インターフェース
type ValueObject interface {
	SameValueAs(other ValueObject) bool
}
