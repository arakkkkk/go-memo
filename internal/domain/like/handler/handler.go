package handler

import (
	"encoding/json"
	"fmt"
	"memo/internal/domain/like"
	"memo/internal/domain/like/usecase"
	"memo/internal/util/response"
	"net/http"
	"strconv"

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

// Get users likes
// @Summary Get users likes
// @Description Get users likes
// @Tags         like
// @Param user_id query int true "user id"
// @Success 200 {object} []entity.Like
// @Failure 400 {object} response.ErrorSchema
// @Failure 500 {object} response.ErrorSchema
// @router /api/v1/like [get]
func (h *Handler) List(w http.ResponseWriter, r *http.Request) {
	strUserID := r.URL.Query().Get("user_id")
	var userID int
	if strUserID == "" {
		userID = 0
	} else {
		var err error
		userID, err = strconv.Atoi(strUserID)
		if err != nil {
			response.Error(w, http.StatusBadRequest, fmt.Errorf("Paremeter is fault."))
			return
		}
	}

	likes, err := h.usecase.List(r.Context(), userID)
	if err != nil {
		response.Error(w, http.StatusBadRequest, err)
		return
	}
	response.Json(w, http.StatusCreated, likes)
}

// Count memo likes
// @Summary Count memo likes
// @Description Count memo likes
// @Tags         like
// @Param memo_id query uuid.UUID true "memo id"
// @Success 200 {object} []entity.Like
// @Failure 400 {object} response.ErrorSchema
// @Failure 500 {object} response.ErrorSchema
// @router /api/v1/like/count [get]
func (h *Handler) Count(w http.ResponseWriter, r *http.Request) {
	strMemoID := r.URL.Query().Get("memo_id")
	memoID, err := uuid.Parse(strMemoID)
	if err != nil {
		response.Error(w, http.StatusBadRequest, fmt.Errorf("Paremeter is fault."))
		return
	}

	countLiked, err := h.usecase.Count(r.Context(), memoID)
	if err != nil {
		response.Error(w, http.StatusBadRequest, err)
		return
	}
	response.Json(w, http.StatusCreated, map[string]int{"count": countLiked})
}

// Whether memo is liked or not
// @Summary Check the user liked specified memo
// @Description Check the user liked specified memo
// @Tags         like
// @Param memo_id query uuid.UUID true "memo id"
// @Param user_id query int true "user id"
// @Failure 400 {object} response.ErrorSchema
// @Failure 500 {object} response.ErrorSchema
// @router /api/v1/like/count [get]
func (h *Handler) IsLiked(w http.ResponseWriter, r *http.Request) {
	strMemoID := r.URL.Query().Get("memo_id")
	memoID, err := uuid.Parse(strMemoID)
	if err != nil {
		response.Error(w, http.StatusBadRequest, fmt.Errorf("Paremeter is fault."))
		return
	}
	strUserID := r.URL.Query().Get("user_id")
	userID, err := strconv.Atoi(strUserID)
	if err != nil {
		response.Error(w, http.StatusBadRequest, fmt.Errorf("Paremeter is fault."))
		return
	}

	isLiked, err := h.usecase.IsLiked(r.Context(), memoID, userID)
	if err != nil {
		response.Error(w, http.StatusBadRequest, err)
		return
	}
	response.Json(w, http.StatusCreated, map[string]bool{"is_liked": isLiked})
}

// Add like
// @Summary Add like
// @Description Add like
// @Tags         like
// @Accept json
// @Produce json
// @Param AddRequest body like.AddRequest true "add request"
// @Success 201 {object} entity.Like
// @Failure 400 {object} response.ErrorSchema
// @Failure 500 {object} response.ErrorSchema
// @router /api/v1/like/add [post]
func (h *Handler) Add(w http.ResponseWriter, r *http.Request) {
	var req like.AddRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.Error(w, http.StatusBadRequest, err)
		return
	}
	like, err := h.usecase.Add(r.Context(), &req)
	if err != nil {
		response.Error(w, http.StatusBadRequest, err)
		return
	}
	response.Json(w, http.StatusCreated, like)
}

// Delete like
// @Summary Delete like
// @Description Delete like
// @Tags         like
// @Accept json
// @Produce json
// @Param DeleteRequest body like.DeleteRequest true "delete request"
// @Success 201 {object} entity.Like
// @Failure 400 {object} response.ErrorSchema
// @Failure 500 {object} response.ErrorSchema
// @router /api/v1/like/delete [post]
func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {
	var req like.DeleteRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.Error(w, http.StatusBadRequest, err)
		return
	}
	_, err := h.usecase.Delete(r.Context(), &req)
	if err != nil {
		response.Error(w, http.StatusBadRequest, err)
		return
	}
	response.Json(w, http.StatusCreated, "ok")
}
