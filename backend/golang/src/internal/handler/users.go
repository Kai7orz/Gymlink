package handler

import (
	"gymlink/internal/dto"
	"gymlink/internal/service"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

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
	var userCreate dto.UserCreateType
	if err := ctx.ShouldBindJSON(&userCreate); err != nil {
		log.Println("error: user create")
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "server internal error"})
		return
	}

	// service に依存
	err := h.svc.SignUpUser(ctx.Request.Context(), userCreate.Name, userCreate.AvatarUrl, token)
	if err != nil {
		log.Println("error: ", err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "server internal error"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Finish Setting Endpoints Successfully✅"})
}

func (h *UserHandler) LoginUser(ctx *gin.Context) {
	// ヘッダー取り出し
	authz := ctx.GetHeader("Authorization")
	if !strings.HasPrefix(authz, "Bearer ") {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "missing bearer token"})
		return
	}
	token := strings.TrimPrefix(authz, "Bearer ")

	userRaw, err := h.svc.LoginUser(ctx.Request.Context(), token)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "server internal error"})
		return
	}

	user := dto.UserType{
		Id:   userRaw.Id,
		Name: userRaw.Name,
	}
	log.Println("Get user successfully ✅：", user)
	ctx.JSON(http.StatusOK, user)
}

func (h *UserHandler) GetProfilebyId(ctx *gin.Context) {
	// ヘッダー取り出し
	authz := ctx.GetHeader("Authorization")
	if !strings.HasPrefix(authz, "Bearer ") {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "missing bearer token"})
		return
	}
	idStr := ctx.Param("user_id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		log.Println("failed to parse user_id")
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "server internal error"})
		return
	}
	profileRaw, err := h.svc.GetProfile(ctx.Request.Context(), id)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "server internal error"})
		return
	}
	log.Println("Get profile successfully ✅：", profileRaw)

	profile := dto.ProfileType{
		Id:            profileRaw.Id,
		Name:          profileRaw.Name,
		ProfileImage:  profileRaw.ProfileImage,
		FollowCount:   profileRaw.FollowCount,
		FollowerCount: profileRaw.FollowerCount,
	}

	ctx.JSON(http.StatusOK, profile)
}

func (h *UserHandler) FollowUser(ctx *gin.Context) {
	// ヘッダー取り出し
	authz := ctx.GetHeader("Authorization")
	if !strings.HasPrefix(authz, "Bearer ") {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "missing bearer token"})
		return
	}

	//　requestBody の読み取り
	var userFollow dto.UserFollowType
	if err := ctx.ShouldBindJSON(&userFollow); err != nil {
		log.Println("error: user follow")
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "server internal error"})
		return
	}

	err := h.svc.FollowUser(ctx.Request.Context(), userFollow.FollowerId, userFollow.FollowedId)
	if err != nil {
		log.Println("error: user follow")
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "server internal error"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "user follow successfully"})
}
