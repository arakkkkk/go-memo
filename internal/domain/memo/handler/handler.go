package handler

import (
	"encoding/json"
	"memo/internal/domain/memo"
	"memo/internal/domain/memo/usecase"
	"memo/internal/util/response"
	"net/http"
)

type Handler struct {
	usecase *usecase.Usecase
}

func New(usecase *usecase.Usecase) *Handler {
	return &Handler{
		usecase: usecase,
	}
}

// Get all memos
// @Summary Get all memos
// @Description Get all users
// @Success 200 {object} []entity.Memo
// @Failure 500 {string} response.Error
// @router /api/v1/user [get]
func (h *Handler) List(w http.ResponseWriter, r *http.Request) {
	memos, err := h.usecase.List(r.Context())
	if err != nil {
		response.Error(w, http.StatusBadRequest, err)
		return
	}
	response.Json(w, http.StatusCreated, memos)
}

// Create memo
// @Summary Create memo
// @Description Create memo
// @Accept json
// @Produce json
// @Param CreateRequest body memo.CreateRequest true ""
// @Success 201 {object} entity.Memo
// @Failure 400 {string} response.Error
// @Failure 500 {string} response.Error
// @router /api/v1/memos/create [post]
func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	var req memo.CreateRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.Error(w, http.StatusBadRequest, err)
		return
	}
	memo, err := h.usecase.Create(r.Context(), &req)
	if err != nil {
		response.Error(w, http.StatusBadRequest, err)
		return
	}
	response.Json(w, http.StatusCreated, memo)
}
