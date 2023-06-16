package rest

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/friendsofgo/errors"

	"github.com/htnk128/go-ddd-sample/pkg/address/domain/model"
)

type accountClientConfig struct {
	url string
}

func newAccountClientConfig() *accountClientConfig {
	url, b := os.LookupEnv("ACCOUNT_URL")
	if !b {
		url = "http://go-ddd-sample-account:8080/accounts"
	}

	return &accountClientConfig{
		url: url,
	}
}

type accountClient struct {
	config *accountClientConfig
}

func newAccountClient() *accountClient {
	c := newAccountClientConfig()
	return &accountClient{
		config: c,
	}
}

type ownerRestService struct {
	client *accountClient
}

func NewOwnerRestService() model.OwnerService {
	return &ownerRestService{newAccountClient()}
}

func (abs *ownerRestService) Find(ownerID model.OwnerID) (*model.Owner, error) {
	res, err := abs.client.find(ownerID.ID())
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}

	id, err := model.NewOwnerID(res.AccountID)
	if err != nil {
		return nil, err
	}
	var deletedAt *time.Time
	if res.DeletedAt != nil {
		t := time.Unix(0, *res.DeletedAt*int64(time.Millisecond))
		deletedAt = &t
	}

	return model.NewOwner(*id, deletedAt), nil
}

type (
	accountResponse struct {
		AccountID string `json:"account_id"`
		DeletedAt *int64 `json:"deleted_at"`
	}
)

func (ac *accountClient) find(accountID string) (*accountResponse, error) {
	res, err := http.Get(fmt.Sprintf("%s/%s", ac.config.url, accountID))
	if err != nil {
		return nil, errors.Wrap(err, "failed to fetch account.")
	}

	defer func() {
		_ = res.Body.Close()
	}()

	b, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, errors.Wrap(err, "failed to fetch account.")
	}

	if res.StatusCode == http.StatusNotFound {
		return nil, nil
	}

	if res.StatusCode < http.StatusOK || res.StatusCode >= http.StatusBadRequest {
		return nil, errors.Wrap(err, "failed to fetch account.")
	}

	var accountRes accountResponse
	if err := json.Unmarshal(b, &accountRes); err != nil {
		return nil, errors.Wrap(err, "failed to parse response.")
	}
	return &accountRes, nil
}
