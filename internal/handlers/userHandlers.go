package handlers

import (
	"context"
	userservice "gomeWork/internal/userService"
	"gomeWork/internal/web/api"

	"github.com/labstack/echo/v4"
	"github.com/oapi-codegen/runtime/types"
)

type UserHandler struct {
	service userservice.UserService
}

func NewUserHandler(s userservice.UserService) *UserHandler {
	return &UserHandler{service: s}
}

func (h *UserHandler) GetUsers(ctx context.Context, request api.GetUsersRequestObject) (api.GetUsersResponseObject, error) {
	allUsers, err := h.service.GetAllUser()
	if err != nil {
		return nil, err
	}
	response := api.GetUsers200JSONResponse{}

	for _, usr := range allUsers {
		createdAtStr := usr.CreatedAt.Format("2006-01-02T15:04:05Z")
		updatedAtStr := usr.UpdatedAt.Format("2006-01-02T15:04:05Z")
		usrResp := api.User{
			Id:        usr.ID,
			Email:     types.Email(usr.Email),
			CreatedAt: createdAtStr,
			UpdatedAt: updatedAtStr,
		}
		response = append(response, usrResp)
	}
	return response, nil
}

func (h *UserHandler) PostUsers(ctx context.Context, request api.PostUsersRequestObject) (api.PostUsersResponseObject, error) {
	if request.Body == nil {
		return nil, echo.NewHTTPError(400, "request body is required")
	}
	if request.Body.Email == "" {
		return nil, echo.NewHTTPError(400, "email is required")
	}
	if request.Body.Password == "" {
		return nil, echo.NewHTTPError(400, "password is required")
	}
	user, err := h.service.CreateUser(string(request.Body.Email), request.Body.Password)
	if err != nil {
		return nil, echo.NewHTTPError(400, err.Error())
	}
	createdAtStr := user.CreatedAt.Format("2006-01-02T15:04:05Z")
	updatedAtStr := user.UpdatedAt.Format("2006-01-02T15:04:05Z")

	return api.PostUsers201JSONResponse{
		Id:        user.ID,
		Email:     types.Email(user.Email),
		CreatedAt: createdAtStr,
		UpdatedAt: updatedAtStr,
	}, nil
}
func (h *UserHandler) PatchUsersId(ctx context.Context, request api.PatchUsersIdRequestObject) (api.PatchUsersIdResponseObject, error) {
	if request.Body == nil || (request.Body.Email == nil && request.Body.Password == nil) {
		return nil, echo.NewHTTPError(400, "at least email or password must be provided for update")
	}

	var email, password string

	if request.Body.Email != nil {
		email = string(*request.Body.Email)
	}

	if request.Body.Password != nil {
		password = *request.Body.Password
	}

	user, err := h.service.UpdateUser(request.Id, email, password)
	if err != nil {
		return nil, echo.NewHTTPError(400, err.Error())
	}

	createdAtStr := user.CreatedAt.Format("2006-01-02T15:04:05Z")
	updatedAtStr := user.UpdatedAt.Format("2006-01-02T15:04:05Z")

	return api.PatchUsersId200JSONResponse{
		Id:        user.ID,
		Email:     types.Email(user.Email),
		CreatedAt: createdAtStr,
		UpdatedAt: updatedAtStr,
	}, nil
}

func (h *UserHandler) DeleteUsersId(ctx context.Context, request api.DeleteUsersIdRequestObject) (api.DeleteUsersIdResponseObject, error) {
	err := h.service.DeleteUser(request.Id)
	if err != nil {
		return api.DeleteUsersId404Response{}, nil
	}
	return api.DeleteUsersId204Response{}, nil
}
