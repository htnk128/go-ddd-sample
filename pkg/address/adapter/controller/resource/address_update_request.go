package resource

type AddressUpdateRequest struct {
	FullName      *string `json:"full_name"`
	ZipCode       *string `json:"zip_code"`
	StateOrRegion *string `json:"state_or_region"`
	Line1         *string `json:"line1"`
	Line2         *string `json:"line2"`
	PhoneNumber   *string `json:"phone_number"`
}
