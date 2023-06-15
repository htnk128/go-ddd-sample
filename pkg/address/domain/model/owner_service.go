package model

type OwnerService interface {
	Find(ownerID OwnerID) (*Owner, error)
}
