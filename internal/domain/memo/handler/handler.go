package handler

import (
	"encoding/json"
	"fmt"
	"memo/internal/domain/memo"
	"memo/internal/domain/memo/usecase"
	"memo/internal/util/response"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
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
// @Tags         memos
// @Success 200 {object} []entity.Memo
// @Failure 400 {object} response.ErrorSchema
// @Failure 500 {object} response.ErrorSchema
// @router /api/v1/memo [get]
func (h *Handler) List(w http.ResponseWriter, r *http.Request) {
	memos, err := h.usecase.List(r.Context())
	if err != nil {
		response.Error(w, http.StatusBadRequest, err)
		return
	}
	response.Json(w, http.StatusCreated, memos)
}

// Get memo
// @Summary Get memo
// @Description Get user
// @Tags         memos
// @Param memo_id query uuid.UUID true "memo id"
// @Success 200 {object} entity.Memo
// @Failure 400 {object} response.ErrorSchema
// @Failure 500 {object} response.ErrorSchema
// @router /api/v1/memo/{memo_id} [get]
func (h *Handler) Get(w http.ResponseWriter, r *http.Request) {
	strMemoID := chi.URLParam(r, "memo_id")
	memoID, err := uuid.Parse(strMemoID)
	if err != nil {
		response.Error(w, http.StatusBadRequest, fmt.Errorf("Paremeter is fault."))
		return
	}

	memos, err := h.usecase.Get(r.Context(), memoID)
	if err != nil {
		response.Error(w, http.StatusBadRequest, err)
		return
	}
	response.Json(w, http.StatusCreated, memos)
}

// Create memo
// @Summary Create memo
// @Description Create memo
// @Tags         memos
// @Accept json
// @Produce json
// @Param CreateRequest body memo.CreateRequest true "create request"
// @Success 201 {object} entity.Memo
// @Failure 400 {object} response.ErrorSchema
// @Failure 500 {object} response.ErrorSchema
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
