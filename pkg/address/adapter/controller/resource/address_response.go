package resource

type AddressResponse struct {
	AddressID     string  `json:"address_id"`
	OwnerID       string  `json:"owner_id"`
	FullName      string  `json:"full_name"`
	ZipCode       string  `json:"zip_code"`
	StateOrRegion string  `json:"state_or_Region"`
	Line1         string  `json:"line1"`
	Line2         *string `json:"line2"`
	PhoneNumber   string  `json:"phone_number"`
	CreatedAt     int64   `json:"created_at"`
	DeletedAt     *int64  `json:"deleted_at"`
	UpdatedAt     int64   `json:"updated_at"`
}
