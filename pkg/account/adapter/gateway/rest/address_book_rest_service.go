package rest

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"github.com/friendsofgo/errors"

	"github.com/htnk128/go-ddd-sample/pkg/account/domain/model"
)

type addressClientConfig struct {
	url string
}

func newAddressClientConfig() *addressClientConfig {
	url, b := os.LookupEnv("ADDRESS_URL")
	if !b {
		url = "http://localhost:8081/addresses"
	}

	return &addressClientConfig{
		url: url,
	}
}

type addressClient struct {
	config *addressClientConfig
}

func newAddressClient() *addressClient {
	c := newAddressClientConfig()
	return &addressClient{
		config: c,
	}
}

type addressBookRestService struct {
	client *addressClient
}

func NewAddressBookRestService() model.AddressBookService {
	return &addressBookRestService{newAddressClient()}
}

func (abs *addressBookRestService) Find(accountID model.AccountID) (*model.AddressBook, error) {
	res, err := abs.client.findAll(accountID.ID())
	if err != nil {
		return nil, err
	}

	addresses := make([]*model.AccountAddress, len(res.Data))
	for i, ar := range res.Data {
		id, err := model.NewAccountAddressID(ar.AddressID)
		if err != nil {
			return nil, err
		}

		var deletedAt *time.Time
		if ar.DeletedAt != nil {
			t := time.Unix(0, *ar.DeletedAt*int64(time.Millisecond))
			deletedAt = &t
		}

		addresses[i] = model.NewAccountAddress(*id, deletedAt)
	}

	return model.NewAddressBook(addresses), nil
}

func (abs *addressBookRestService) Remove(accountAddressID model.AccountAddressID) error {
	_, err := abs.client.delete(accountAddressID.ID())
	if err != nil {
		return err
	}

	return nil
}

type (
	addressResponse struct {
		AddressID string `json:"address_id"`
		DeletedAt *int64 `json:"deleted_at"`
	}
	addressResponses struct {
		Data []addressResponse `json:"data"`
	}
)

func (ac *addressClient) findAll(accountID string) (*addressResponses, error) {
	res, err := http.Get(fmt.Sprintf("%s?ownerId=%s", ac.config.url, accountID))
	if err != nil {
		return nil, errors.Wrap(err, "failed to fetch address.")
	}

	defer func() {
		_ = res.Body.Close()
	}()

	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, errors.Wrap(err, "failed to fetch address.")
	}

	if res.StatusCode < http.StatusOK || res.StatusCode >= http.StatusBadRequest {
		return nil, errors.Wrap(err, "failed to fetch address.")
	}

	var addressRes addressResponses
	if err := json.Unmarshal(b, &addressRes); err != nil {
		return nil, errors.Wrap(err, "failed to parse response.")
	}
	return &addressRes, nil
}

func (ac *addressClient) delete(accountAddressID string) (*addressResponse, error) {
	req, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s", ac.config.url, accountAddressID), nil)
	if err != nil {
		panic(err)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "failed to delete address.")
	}

	defer func() {
		_ = res.Body.Close()
	}()

	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, errors.Wrap(err, "failed to delete address.")
	}

	if res.StatusCode < http.StatusOK || res.StatusCode >= http.StatusBadRequest {
		return nil, errors.Wrap(err, "failed to delete address.")
	}

	var addressRes addressResponse
	if err := json.Unmarshal(b, &addressRes); err != nil {
		return nil, errors.Wrap(err, "failed to parse response.")
	}
	return &addressRes, nil
}
