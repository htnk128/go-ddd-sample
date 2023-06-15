package command

type CreateAddressCommand struct {
	OwnerID       string
	FullName      string
	ZipCode       string
	StateOrRegion string
	Line1         string
	Line2         *string
	PhoneNumber   string
}
