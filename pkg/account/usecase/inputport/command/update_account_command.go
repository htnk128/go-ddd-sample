package command

type UpdateAccountCommand struct {
	AccountID         string
	Name              *string
	NamePronunciation *string
	Email             *string
	Password          *string
}
