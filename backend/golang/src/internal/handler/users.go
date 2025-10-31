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
		log.Println("error: user data bind")
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "server internal error"})
		return
	}

	// service に依存
	err := h.svc.SignUpUser(ctx.Request.Context(), userCreate.Name, userCreate.AvatarUrl, token)
	if err != nil {
		log.Println("error: failed to signup ", err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "failed to signup"})
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
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "failed to login"})
		return
	}

	user := dto.UserType{
		Id:   userRaw.Id,
		Name: userRaw.Name,
	}
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
	if err != nil || id <= 0 {
		log.Println("error: invalid user id")
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "invalid user id"})
		return
	}
	profileRaw, err := h.svc.GetProfile(ctx.Request.Context(), id)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "server internal error"})
		return
	}

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

func (h *UserHandler) GetFollowing(ctx *gin.Context) {
	// ヘッダー取り出し
	authz := ctx.GetHeader("Authorization")
	if !strings.HasPrefix(authz, "Bearer ") {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "missing Bearer token"})
		return
	}
	idStr := ctx.Param("user_id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil || id <= 0 {
		log.Println("error: invalid user")
	}
	followingUsers, err := h.svc.GetFollowingById(ctx, id)
	ctx.JSON(http.StatusOK, gin.H{"following": followingUsers})
}

func (h *UserHandler) CheckFollow(ctx *gin.Context) {
	// ヘッダー取り出し
	authz := ctx.GetHeader("Authorization")
	if !strings.HasPrefix(authz, "Bearer ") {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "missing bearer token"})
		return
	}
	token := strings.TrimPrefix(authz, "Bearer ")

	idStr := ctx.Param("user_id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil || id <= 0 {
		log.Println("error: invalid user id")
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "invalid user id"})
		return
	}
	followed, err := h.svc.CheckFollowById(ctx.Request.Context(), id, token)
	if err != nil {
		log.Println("failed to check follows")
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "server internal error"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"followed": followed})

}

func (h *UserHandler) DeleteFollowUser(ctx *gin.Context) {
	// ヘッダー取り出し
	authz := ctx.GetHeader("Authorization")
	if !strings.HasPrefix(authz, "Bearer ") {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "missing bearer token"})
		return
	}

	//　requestBody の読み取り
	var userDeleteFollow dto.UserFollowType
	if err := ctx.ShouldBindJSON(&userDeleteFollow); err != nil {
		log.Println("error: user follow")
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "server internal error"})
		return
	}

	err := h.svc.DeleteFollowUser(ctx.Request.Context(), userDeleteFollow.FollowerId, userDeleteFollow.FollowedId)
	if err != nil {
		log.Println("error: user follow")
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "server internal error"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "user follow successfully"})
}
