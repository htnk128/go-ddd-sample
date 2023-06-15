package resource

type AddressCreateRequest struct {
	OwnerID       string  `json:"owner_id"`
	FullName      string  `json:"full_name"`
	ZipCode       string  `json:"zip_code"`
	StateOrRegion string  `json:"state_or_region"`
	Line1         string  `json:"line1"`
	Line2         *string `json:"line2"`
	PhoneNumber   string  `json:"phone_number"`
}
