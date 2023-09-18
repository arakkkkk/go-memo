package handler

import (
	"encoding/json"
	"memo/internal/domain/user"
	"memo/internal/domain/user/usecase"
	"memo/internal/util/response"
	"net/http"
	"time"
)

type Handler struct {
	usecase *usecase.Usecase
}

func New(usecase *usecase.Usecase) *Handler {
	return &Handler{
		usecase: usecase,
	}
}

// Register
// @Summary Register
// @Description Register
// @Tags         users
// @Accept json
// @Produce json
// @Param RegisterRequest body user.RegisterRequest true "register request"
// @Success 201 {string} ok
// @Failure 400 {object} response.ErrorSchema
// @Failure 500 {object} response.ErrorSchema
// @router /api/v1/user/register [post]
func (h *Handler) Register(w http.ResponseWriter, r *http.Request) {
	var req user.RegisterRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.Error(w, http.StatusBadRequest, err)
		return
	}
	token, err := h.usecase.Register(r.Context(), &req)
	if err != nil {
		response.Error(w, http.StatusBadRequest, err)
		return
	}

	cookie := http.Cookie{
		Name:     "jwt",
		Path:     "/",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HttpOnly: true,
	}
	http.SetCookie(w, &cookie)

	response.Json(w, http.StatusOK, "ok")
}

// Login
// @Summary Login
// @Description Login
// @Tags         users
// @Accept json
// @Produce json
// @Param LoginRequest body user.LoginRequest true "login request"
// @Failure 400 {object} response.ErrorSchema
// @Failure 500 {object} response.ErrorSchema
// @router /api/v1/user/login [post]
func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	var req user.LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.Error(w, http.StatusBadRequest, err)
		return
	}
	token, err := h.usecase.Login(r.Context(), &req)
	if err != nil {
		response.Error(w, http.StatusBadRequest, err)
		return
	}

	cookie := http.Cookie{
		Name:     "jwt",
		Path:     "/",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HttpOnly: true,
	}
	http.SetCookie(w, &cookie)

	response.Json(w, http.StatusOK, "ok")
}

// Logout
// @Summary Logout
// @Description Logout
// @Tags         users
// @Failure 400 {object} response.ErrorSchema
// @Failure 500 {object} response.ErrorSchema
// @router /api/v1/user/logout [get]
func (h *Handler) Logout(w http.ResponseWriter, r *http.Request) {
	cookie := http.Cookie{
		Name:     "jwt",
		Path:     "/",
		Value:    "",
		Expires:  time.Now().Add(time.Hour * 24),
		HttpOnly: true,
	}
	http.SetCookie(w, &cookie)

	response.Json(w, http.StatusOK, "ok")
}

// Get all users
// @Summary Get all users
// @Description Get all users
// @Tags         users
// @Success 200 {object} []entity.Users
// @Failure 400 {object} response.ErrorSchema
// @Failure 500 {object} response.ErrorSchema
// @router /api/v1/user/list [get]
func (h *Handler) List(w http.ResponseWriter, r *http.Request) {
	users, err := h.usecase.List(r.Context())
	if err != nil {
		response.Error(w, http.StatusBadRequest, err)
	}
	response.Json(w, http.StatusOK, users)
}
