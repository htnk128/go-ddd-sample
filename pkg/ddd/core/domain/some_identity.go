package domain

// SomeIdentity 何らかのドメインを識別するIDを表現した値オブジェクト
type SomeIdentity struct {
	Identity

	id string
}

func (si *SomeIdentity) ID() string {
	return si.id
}

func (si *SomeIdentity) Equals(other *SomeIdentity) bool {
	return si.SameValueAs(other)
}

func (si *SomeIdentity) SameValueAs(other *SomeIdentity) bool {
	return si.id == other.id
}

func (si *SomeIdentity) String() string {
	return si.id
}

const (
	SomeIdentityMinLength = 1
	SomeIdentityMaxLength = 64
	SomeIdentityPattern   = "[\\p{Alnum}-_]*"
)

func NewSomeIdentity(id string) *SomeIdentity {
	return &SomeIdentity{id: id}
}
