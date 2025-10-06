package handler

import (
	"gymlink/internal/service"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type userTypeDTO struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

// handler が何に依存するかを明示
type UserHandler struct {
	svc service.UserService
}

func NewUserHandler(svc service.UserService) *UserHandler {
	return &UserHandler{svc: svc}
}

func (h *UserHandler) GetUserByParam(ctx *gin.Context) {
	authz := ctx.GetHeader("Authorization")
	if !strings.HasPrefix(authz, "Bearer ") {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "missing bearer token"})
		return
	}
	token := strings.TrimPrefix(authz, "Bearer ")
	// service に依存
	user, err := h.svc.GetUser(ctx.Request.Context(), token)
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
