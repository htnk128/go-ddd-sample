package controller

import (
	"net/http"
	"strconv"

	"github.com/friendsofgo/errors"
	"github.com/labstack/echo/v4"

	"github.com/htnk128/go-ddd-sample/pkg/account/adapter/controller/resource"
	"github.com/htnk128/go-ddd-sample/pkg/account/adapter/gateway/db"
	"github.com/htnk128/go-ddd-sample/pkg/account/adapter/gateway/rest"
	"github.com/htnk128/go-ddd-sample/pkg/account/adapter/presenter"
	"github.com/htnk128/go-ddd-sample/pkg/account/usecase/inputport"
	"github.com/htnk128/go-ddd-sample/pkg/account/usecase/inputport/command"
	"github.com/htnk128/go-ddd-sample/pkg/account/usecase/outputport"
	"github.com/htnk128/go-ddd-sample/pkg/account/usecase/outputport/dto"
	sharedResource "github.com/htnk128/go-ddd-sample/pkg/shared/adapter/controller/resource"
	sharedUseCase "github.com/htnk128/go-ddd-sample/pkg/shared/usecase"
)

type AccountController struct {
	inAccountUseCase  inputport.AccountUseCase
	outAccountUseCase outputport.AccountUseCase
}

func NewAccountController() (*AccountController, error) {
	repositories, err := db.NewRepositories()
	if err != nil {
		return nil, errors.Wrap(err, "failed to create database repository.")
	}

	r := repositories.AccountRepository
	s := rest.NewAddressBookRestService()
	return &AccountController{
		inAccountUseCase:  inputport.NewAccountUseCase(r, s),
		outAccountUseCase: presenter.NewAccountUseCase(),
	}, nil
}

const (
	AccountIDParam = "account_id"
)

func (ac *AccountController) Find() echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()

		cmd := command.FindAccountCommand{
			AccountID: c.Param(AccountIDParam),
		}
		a, err := ac.inAccountUseCase.Find(ctx, cmd)
		if err != nil {
			res := errorResponseFromError(err)
			return c.JSON(res.Error.Status, res)
		}

		res := responseFromDTO(ac.outAccountUseCase.DTO(a))

		return c.JSON(http.StatusOK, res)
	}
}

func (ac *AccountController) FindAll() echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()

		l := c.QueryParam("limit")
		limit, err := strconv.Atoi(l)
		if err != nil {
			limit = 10
		}
		o := c.QueryParam("offset")
		offset, err := strconv.Atoi(o)
		if err != nil {
			offset = 0
		}

		cmd := command.FindAllAccountCommand{
			Limit:  limit,
			Offset: offset,
		}

		cnt, as, err := ac.inAccountUseCase.FindAll(ctx, cmd)
		if err != nil {
			res := errorResponseFromError(err)
			return c.JSON(res.Error.Status, res)
		}

		pDto := ac.outAccountUseCase.PaginationDTO(as, int(cnt), limit, offset)
		data := make([]*resource.AccountResponse, len(pDto.Data))
		for i, a := range pDto.Data {
			data[i] = responseFromDTO(a)
		}
		res := &resource.AccountResponses{
			Count:   pDto.Count,
			HasMore: pDto.HasMore(),
			Data:    data,
		}

		return c.JSON(http.StatusOK, res)
	}
}

func (ac *AccountController) Create() echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()

		var req resource.AccountCreateRequest
		if err := c.Bind(&req); err != nil {
			errType := sharedUseCase.InvalidRequestErrorType
			code := sharedUseCase.InvalidRequestErrorStatus
			message := errType + ": " + err.Error()
			res := sharedResource.NewErrorResponse(errType, code, message)
			return c.JSON(code, res)
		}

		cmd := command.CreateAccountCommand{
			Name:              req.Name,
			NamePronunciation: req.NamePronunciation,
			Email:             req.Email,
			Password:          req.Password,
		}

		a, err := ac.inAccountUseCase.Create(ctx, cmd)
		if err != nil {
			res := errorResponseFromError(err)
			return c.JSON(res.Error.Status, res)
		}

		res := responseFromDTO(ac.outAccountUseCase.DTO(a))

		return c.JSON(http.StatusOK, res)
	}
}

func (ac *AccountController) Update() echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()

		var req resource.AccountUpdateRequest
		if err := c.Bind(&req); err != nil {
			errType := sharedUseCase.InvalidRequestErrorType
			code := sharedUseCase.InvalidRequestErrorStatus
			message := errType + ": " + err.Error()
			res := sharedResource.NewErrorResponse(errType, code, message)
			return c.JSON(code, res)
		}

		cmd := command.UpdateAccountCommand{
			Name:              req.Name,
			NamePronunciation: req.NamePronunciation,
			Email:             req.Email,
			Password:          req.Password,
		}

		a, err := ac.inAccountUseCase.Update(ctx, cmd)
		if err != nil {
			res := errorResponseFromError(err)
			return c.JSON(res.Error.Status, res)
		}

		res := responseFromDTO(ac.outAccountUseCase.DTO(a))

		return c.JSON(http.StatusOK, res)
	}
}

func (ac *AccountController) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()

		cmd := command.DeleteAccountCommand{
			AccountID: c.Param(AccountIDParam),
		}
		a, err := ac.inAccountUseCase.Delete(ctx, cmd)
		if err != nil {
			res := errorResponseFromError(err)
			return c.JSON(res.Error.Status, res)
		}

		res := responseFromDTO(ac.outAccountUseCase.DTO(a))

		return c.JSON(http.StatusOK, res)
	}
}

func errorResponseFromError(err error) *sharedResource.ErrorResponse {
	errorType := sharedUseCase.ErrorType(err)
	code := sharedUseCase.ErrorStatus(err)
	message := sharedUseCase.ErrorMessage(err)
	return sharedResource.NewErrorResponse(errorType, code, message)
}

func responseFromDTO(dto *dto.AccountDTO) *resource.AccountResponse {
	return &resource.AccountResponse{
		AccountID:         dto.AccountID,
		Name:              dto.Name,
		NamePronunciation: dto.NamePronunciation,
		Email:             dto.Email,
		Password:          dto.Password,
		CreatedAt:         dto.CreatedAt,
		DeletedAt:         dto.DeletedAt,
		UpdatedAt:         dto.UpdatedAt,
	}
}
