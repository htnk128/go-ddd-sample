package model

type AddressBookService interface {
	Find(accountID AccountID) (*AddressBook, error)
	Remove(accountAddressID AccountAddressID) error
}
