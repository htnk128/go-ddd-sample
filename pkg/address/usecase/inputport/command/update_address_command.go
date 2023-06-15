package command

type UpdateAddressCommand struct {
	AddressID     string
	FullName      *string
	ZipCode       *string
	StateOrRegion *string
	Line1         *string
	Line2         *string
	PhoneNumber   *string
}
