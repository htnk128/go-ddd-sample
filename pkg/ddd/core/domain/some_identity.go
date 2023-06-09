package domain

// SomeIdentity 何らかのドメインを識別するIDを表現した値オブジェクト
type SomeIdentity struct {
	Identity

	ID string
}

func (si *SomeIdentity) SameValueAs(other *SomeIdentity) bool {
	return si.ID == other.ID
}

func (si *SomeIdentity) String() string {
	return si.ID
}

const (
	SomeIdentityMinLength = 1
	SomeIdentityMaxLength = 64
	SomeIdentityPattern   = "[\\p{Alnum}-_]*"
)
