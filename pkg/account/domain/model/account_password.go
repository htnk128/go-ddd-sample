package model

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"golang.org/x/crypto/pbkdf2"

	"github.com/friendsofgo/errors"

	"github.com/htnk128/go-ddd-sample/pkg/ddd/core/domain"
)

type AccountPassword struct {
	*domain.SomeValueObject[string]
}

func (ap *AccountPassword) Format() string {
	return "*****"
}

func NewAccountPasswordWithHash(value string, id AccountID) (*AccountPassword, error) {
	const (
		minLength      = 6
		maxLength      = 100
		iterationCount = 100
		keyLength      = 256
	)

	if minLength <= len(value) && len(value) <= maxLength {
		secret := []byte(value)
		salt := sha256.New()
		salt.Write([]byte(id.ID()))
		saltBytes := salt.Sum(nil)

		derivedKey := pbkdf2.Key(secret, saltBytes, iterationCount, keyLength/8, sha256.New)
		encodedKey := hex.EncodeToString(derivedKey)

		return &AccountPassword{SomeValueObject: domain.NewSomeValueObject(encodedKey)}, nil
	}

	return nil, errors.New(fmt.Sprintf("AccountPassword must be between %d and %d characters.", minLength, maxLength))
}

func NewAccountPassword(value string) *AccountPassword {
	return &AccountPassword{SomeValueObject: domain.NewSomeValueObject(value)}
}
