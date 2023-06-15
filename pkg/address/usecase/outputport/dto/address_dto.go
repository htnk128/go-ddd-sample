package dto

type AddressDTO struct {
	AddressID,
	OwnerID,
	FullName,
	ZipCode,
	StateOrRegion,
	Line1 string
	Line2       *string
	PhoneNumber string
	CreatedAt   int64
	DeletedAt   *int64
	UpdatedAt   int64
}
