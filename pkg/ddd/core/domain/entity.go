package domain

// Entity DDDにおけるエンティティの概念
type Entity interface {
	SameIdentityAs(other Entity) bool
}
