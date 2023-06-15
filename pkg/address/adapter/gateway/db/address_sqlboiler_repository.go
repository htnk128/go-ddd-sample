package db

import (
	"context"
	"database/sql"
	"time"

	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"

	gatewayModel "github.com/htnk128/go-ddd-sample/pkg/address/adapter/gateway/db/model"
	"github.com/htnk128/go-ddd-sample/pkg/address/domain/model"
	"github.com/htnk128/go-ddd-sample/pkg/address/domain/repository"
)

type addressSQLBoilerRepository struct {
	db *sql.DB
}

func newUserRepository(db *sql.DB) repository.AddressRepository {
	return &addressSQLBoilerRepository{db: db}
}

func (asr *addressSQLBoilerRepository) Find(ctx context.Context, addressID model.AddressID, lock bool) (*model.Address, error) {
	q := []qm.QueryMod{
		qm.Where("address_id = ?", addressID.ID()),
	}
	if lock {
		q = append(q, qm.For("UPDATE"))
	}
	ac, err := gatewayModel.Addresses(q...).One(ctx, asr.db)
	if err != nil {
		return nil, err
	}
	m, err := addressFromGatewayModel(ac)
	if err != nil {
		return nil, err
	}

	return m, nil
}

func (asr *addressSQLBoilerRepository) FindAll(ctx context.Context, ownerID model.OwnerID) ([]*model.Address, error) {
	acs, err := gatewayModel.Addresses(
		qm.Where("owner_id = ?", ownerID.ID()),
	).All(ctx, asr.db)
	if err != nil {
		return nil, err
	}

	all := make([]*model.Address, len(acs))
	for i, ac := range acs {
		m, err := addressFromGatewayModel(ac)
		if err != nil {
			return nil, err
		}
		all[i] = m
	}

	return all, nil
}

func (asr *addressSQLBoilerRepository) Add(ctx context.Context, address *model.Address) error {
	return addressFromDomainModel(address).Insert(ctx, asr.db, boil.Infer())
}

func (asr *addressSQLBoilerRepository) Set(ctx context.Context, address *model.Address) (int64, error) {
	return addressFromDomainModel(address).Update(ctx, asr.db, boil.Infer())
}

func (asr *addressSQLBoilerRepository) NextAddressID() *model.AddressID {
	return model.GenerateAddressID()
}

func addressFromGatewayModel(address *gatewayModel.Address) (*model.Address, error) {
	id, err := model.NewAddressID(address.AddressID)
	if err != nil {
		return nil, err
	}
	ownerID, err := model.NewOwnerID(address.OwnerID)
	if err != nil {
		return nil, err
	}
	fullName, err := model.NewFullName(address.FullName)
	if err != nil {
		return nil, err
	}
	zipCode, err := model.NewZipCode(address.ZipCode)
	if err != nil {
		return nil, err
	}
	stateOrRegion, err := model.NewStateOrRegion(address.StateOrRegion)
	if err != nil {
		return nil, err
	}
	line1, err := model.NewLine1(address.Line1)
	if err != nil {
		return nil, err
	}
	var line2 *model.Line2
	if address.Line2.Valid {
		line2, err = model.NewLine2(address.Line2.String)
		if err != nil {
			return nil, err
		}
	}
	phoneNumber, err := model.NewPhoneNumber(address.PhoneNumber)
	if err != nil {
		return nil, err
	}
	var deletedAt *time.Time
	if address.DeletedAt.Valid {
		deletedAt = &address.DeletedAt.Time
	}

	return &model.Address{
		ID:            *id,
		OwnerID:       *ownerID,
		FullName:      *fullName,
		ZipCode:       *zipCode,
		StateOrRegion: *stateOrRegion,
		Line1:         *line1,
		Line2:         line2,
		PhoneNumber:   *phoneNumber,
		CreatedAt:     address.CreatedAt,
		DeletedAt:     deletedAt,
		UpdatedAt:     address.UpdatedAt,
	}, nil
}

func addressFromDomainModel(address *model.Address) *gatewayModel.Address {
	var line2 null.String
	if address.Line2 != nil {
		line2 = null.StringFrom(address.Line2.Value())
	}

	return &gatewayModel.Address{
		AddressID:     address.ID.ID(),
		OwnerID:       address.OwnerID.ID(),
		FullName:      address.FullName.Value(),
		ZipCode:       address.ZipCode.Value(),
		StateOrRegion: address.StateOrRegion.Value(),
		Line1:         address.Line1.Value(),
		Line2:         line2,
		PhoneNumber:   address.PhoneNumber.Value(),
		CreatedAt:     address.CreatedAt,
		DeletedAt:     null.TimeFromPtr(address.DeletedAt),
		UpdatedAt:     address.UpdatedAt,
	}
}
