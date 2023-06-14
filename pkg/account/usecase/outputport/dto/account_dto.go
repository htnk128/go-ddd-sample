package dto

type AccountDTO struct {
	AccountID,
	Name,
	NamePronunciation,
	Email,
	Password string
	CreatedAt int64
	DeletedAt *int64
	UpdatedAt int64
}
