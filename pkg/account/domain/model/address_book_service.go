package model

type AddressBookService interface {
	Find(accountID AccountID) *AddressBook
	Remove(accountAddressID AccountAddressID)
}
