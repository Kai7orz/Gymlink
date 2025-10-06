package handler

import (
	"gymlink/internal/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type userTypeDTO struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

// handler が何に依存するかを明示
type UserHandler struct {
	svc *service.UserService
}

func NewUserHandler(svc *service.UserService) *UserHandler {
	return &UserHandler{svc: svc}
}

func (h *UserHandler) GetUserByParam(ctx *gin.Context) {
	userIdStr := ctx.Param("id")
	if userIdStr == "" {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "user_id is required"})
		return
	}
	// validation
	userId, err := strconv.ParseInt(userIdStr, 10, 64)
	if err != nil || userId <= 0 {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "user_id is required"})
		return
	}

	// service に依存
	user, err := h.svc.GetUser(ctx.Request.Context(), userId)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "server internal error"})
		return
	}

	userRes := userTypeDTO{
		Id:   user.Id,
		Name: user.Name,
	}

	ctx.JSON(http.StatusOK, userRes)
}
