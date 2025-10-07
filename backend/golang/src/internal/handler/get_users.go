package handler

import (
	"gymlink/internal/service"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type userTypeDTO struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

type userCreateTypeDTO struct {
	Name      string `json:"name"`
	AvatarUrl string `json:"avatar_url"`
}

// handler が何に依存するかを明示
type UserHandler struct {
	svc service.UserService
}

func NewUserHandler(svc service.UserService) *UserHandler {
	return &UserHandler{svc: svc}
}

func (h *UserHandler) SignUpUserById(ctx *gin.Context) {
	// ヘッダー取り出し
	authz := ctx.GetHeader("Authorization")
	if !strings.HasPrefix(authz, "Bearer ") {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "missing bearer token"})
		return
	}
	token := strings.TrimPrefix(authz, "Bearer ")

	//　requestBody の読み取り
	var userCreate userCreateTypeDTO
	if err := ctx.ShouldBindJSON(&userCreate); err != nil {
		log.Println("error: user create")
		return
	}

	// service に依存
	err := h.svc.SignUpUser(ctx.Request.Context(), userCreate.Name, userCreate.AvatarUrl, token)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "server internal error"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "User Created Successfully✅"})
}
