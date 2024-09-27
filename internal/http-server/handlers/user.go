package handlers

import (
	"encoding/json"
	"fmt"
	"kanban/internal/domain/models"
	resp "kanban/internal/lib/api/response"
	"kanban/internal/lib/api/validation"
	"kanban/internal/lib/logger/sl"
	"net/http"
)

// Login Авторизация пользователя
// @Summary Авторизировать пользователя
// @Description Авторизация пользователя
// @Tags User
// @Router /user/login [post]
// @Param email body string true "Почта"
// @Param password body string true "Пароль"
// @Success 200 {object} models.AuthTokens "OK"
// @Failure 400 {object} resp.Response "Bad request"
// @Failure 500 {object} resp.Response "Internal server error"
func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	var op string = "User.Login"
	log := h.log.With("op", op)

	var credentials models.Credentials

	err := json.NewDecoder(r.Body).Decode(&credentials)
	if err != nil {
		log.With("op", op).With("error", err).Error("error decoding body")
		resp.WriteJSONResponse(w, http.StatusBadRequest, "invalid input parameters", nil)
		return
	}

	authTokens, err := h.userService.Login(r.Context(), credentials)
	if err != nil {
		log.With("op", op).With("error", err).Error("error calling Login")
		resp.WriteJSONResponse(w, http.StatusInternalServerError, "internal server error", nil)
		return
	}

	resp.WriteJSON(w, http.StatusOK, authTokens)
}

// Register Зарегистрировать пользователя
// @Summary Регистрация пользователя
// @Description Регистрация пользователя
// @Tags User
// @Router /user/register [post]
// @Param name body string true "Имя"
// @Param surname body string true "Фамилия"
// @Param middle_name body string false "Отчество"
// @Param email body string true "Почта"
// @Param password body string true "Пароль"
// @Param repeat_password body string true "Повторите пароль"
// @Success 201 {object} models.AuthTokens "Created"
// @Success 400 {object} resp.Response "Bad request"
// @Success 500 {object} resp.Response "Internal server error"
func (h *Handler) Register(w http.ResponseWriter, r *http.Request) {
	var op string = "User.Register"
	log := h.log.With("op", op)

	var credentials models.Credentials

	err := json.NewDecoder(r.Body).Decode(&credentials)
	if err != nil {
		resp.WriteJSONResponse(w, http.StatusBadRequest, "invalid input parameters", nil)
		return
	}

	if validationErrors := validation.ValidateStruct(credentials); len(validationErrors) > 0 {
		log.With("credentials", credentials).Error("invalid input parameters", sl.Err(fmt.Errorf("validation errors")))
		resp.WriteJSONResponse(w, http.StatusBadRequest, "invalid input parameters", nil)
		return
	}

	authTokens, err := h.userService.Register(r.Context(), credentials)
	if err != nil {
		resp.WriteJSONResponse(w, http.StatusInternalServerError, "internal server error", nil)
		return
	}

	resp.WriteJSON(w, http.StatusOK, authTokens)
}

func (h *Handler) Logout(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) Me(w http.ResponseWriter, r *http.Request) {

}
