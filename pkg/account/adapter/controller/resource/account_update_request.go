package resource

type AccountUpdateRequest struct {
	Name              *string `json:"name"`
	NamePronunciation *string `json:"name_pronunciation"`
	Email             *string `json:"email"`
	Password          *string `json:"password"`
}
