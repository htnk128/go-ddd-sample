package controller

import (
	"net/http"

	"github.com/friendsofgo/errors"
	"github.com/labstack/echo/v4"

	"github.com/htnk128/go-ddd-sample/pkg/address/adapter/controller/resource"
	"github.com/htnk128/go-ddd-sample/pkg/address/adapter/gateway/db"
	"github.com/htnk128/go-ddd-sample/pkg/address/adapter/gateway/rest"
	"github.com/htnk128/go-ddd-sample/pkg/address/adapter/presenter"
	"github.com/htnk128/go-ddd-sample/pkg/address/usecase/inputport"
	"github.com/htnk128/go-ddd-sample/pkg/address/usecase/inputport/command"
	"github.com/htnk128/go-ddd-sample/pkg/address/usecase/outputport"
	"github.com/htnk128/go-ddd-sample/pkg/address/usecase/outputport/dto"
	sharedResource "github.com/htnk128/go-ddd-sample/pkg/shared/adapter/controller/resource"
	sharedUseCase "github.com/htnk128/go-ddd-sample/pkg/shared/usecase"
)

type AddressController struct {
	inAddressUseCase  inputport.AddressUseCase
	outAddressUseCase outputport.AddressUseCase
}

func NewAddressController() (*AddressController, error) {
	repositories, err := db.NewRepositories()
	if err != nil {
		return nil, errors.Wrap(err, "failed to create database repository.")
	}

	r := repositories.AddressRepository
	s := rest.NewOwnerRestService()
	return &AddressController{
		inAddressUseCase:  inputport.NewAddressUseCase(r, s),
		outAddressUseCase: presenter.NewAddressUseCase(),
	}, nil
}

const (
	AddressIDParam = "address_id"
)

func (ac *AddressController) Find() echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()

		cmd := command.FindAddressCommand{
			AddressID: c.Param(AddressIDParam),
		}

		a, err := ac.inAddressUseCase.Find(ctx, cmd)
		if err != nil {
			res := errorResponseFromCustomError(err)
			return c.JSON(res.Error.Status, res)
		}

		res := responseFromDTO(ac.outAddressUseCase.DTO(a))

		return c.JSON(http.StatusOK, res)
	}
}

func (ac *AddressController) FindAll() echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()

		cmd := command.FindAllAddressCommand{
			OwnerID: c.QueryParam("owner_id"),
		}

		as, err := ac.inAddressUseCase.FindAll(ctx, cmd)
		if err != nil {
			res := errorResponseFromCustomError(err)
			return c.JSON(res.Error.Status, res)
		}

		data := make([]*resource.AddressResponse, len(as))
		for i, a := range as {
			data[i] = responseFromDTO(ac.outAddressUseCase.DTO(a))
		}
		res := &resource.AddressResponses{
			Data: data,
		}

		return c.JSON(http.StatusOK, res)
	}
}

func (ac *AddressController) Create() echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()

		var req resource.AddressCreateRequest
		if err := c.Bind(&req); err != nil {
			res := errorResponseFromError(err)
			return c.JSON(res.Error.Status, res)
		}

		cmd := command.CreateAddressCommand{
			OwnerID:       req.OwnerID,
			FullName:      req.FullName,
			ZipCode:       req.ZipCode,
			StateOrRegion: req.StateOrRegion,
			Line1:         req.Line1,
			Line2:         req.Line2,
			PhoneNumber:   req.PhoneNumber,
		}

		a, err := ac.inAddressUseCase.Create(ctx, cmd)
		if err != nil {
			res := errorResponseFromCustomError(err)
			return c.JSON(res.Error.Status, res)
		}

		res := responseFromDTO(ac.outAddressUseCase.DTO(a))

		return c.JSON(http.StatusOK, res)
	}
}

func (ac *AddressController) Update() echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()

		var req resource.AddressUpdateRequest
		if err := c.Bind(&req); err != nil {
			res := errorResponseFromError(err)
			return c.JSON(res.Error.Status, res)
		}

		cmd := command.UpdateAddressCommand{
			AddressID:     c.Param(AddressIDParam),
			FullName:      req.FullName,
			ZipCode:       req.ZipCode,
			StateOrRegion: req.StateOrRegion,
			Line1:         req.Line1,
			Line2:         req.Line2,
			PhoneNumber:   req.PhoneNumber,
		}

		a, err := ac.inAddressUseCase.Update(ctx, cmd)
		if err != nil {
			res := errorResponseFromCustomError(err)
			return c.JSON(res.Error.Status, res)
		}

		res := responseFromDTO(ac.outAddressUseCase.DTO(a))

		return c.JSON(http.StatusOK, res)
	}
}

func (ac *AddressController) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()

		cmd := command.DeleteAddressCommand{
			AddressID: c.Param(AddressIDParam),
		}
		a, err := ac.inAddressUseCase.Delete(ctx, cmd)
		if err != nil {
			res := errorResponseFromCustomError(err)
			return c.JSON(res.Error.Status, res)
		}

		res := responseFromDTO(ac.outAddressUseCase.DTO(a))

		return c.JSON(http.StatusOK, res)
	}
}

func errorResponseFromCustomError(err error) *sharedResource.ErrorResponse {
	errorType := sharedUseCase.ErrorType(err)
	code := sharedUseCase.ErrorStatus(err)
	message := sharedUseCase.ErrorMessage(err)
	return sharedResource.NewErrorResponse(errorType, code, message)
}

func errorResponseFromError(err error) *sharedResource.ErrorResponse {
	errType := sharedUseCase.InvalidRequestErrorType
	code := sharedUseCase.InvalidRequestErrorStatus
	message := errType + ": " + err.Error()
	return sharedResource.NewErrorResponse(errType, code, message)
}

func responseFromDTO(dto *dto.AddressDTO) *resource.AddressResponse {
	return &resource.AddressResponse{
		AddressID:     dto.AddressID,
		FullName:      dto.FullName,
		ZipCode:       dto.ZipCode,
		StateOrRegion: dto.StateOrRegion,
		Line1:         dto.Line1,
		Line2:         dto.Line2,
		CreatedAt:     dto.CreatedAt,
		DeletedAt:     dto.DeletedAt,
		UpdatedAt:     dto.UpdatedAt,
	}
}
