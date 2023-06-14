package resource

type AccountResponse struct {
	AccountID         string `json:"account_id"`
	Name              string `json:"name"`
	NamePronunciation string `json:"name_pronunciation"`
	Email             string `json:"email"`
	Password          string `json:"password"`
	CreatedAt         int64  `json:"created_at"`
	DeletedAt         *int64 `json:"deleted_at"`
	UpdatedAt         int64  `json:"updated_at"`
}
