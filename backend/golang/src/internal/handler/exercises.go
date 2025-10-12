package handler

import (
	"gymlink/internal/service"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

type exerciseLikeTypeDTO struct {
	ExerciseRecordId int64 `json:"exercise_record_id" db:"exercise_record_id"`
}

type exerciseCreateTypeDTO struct {
	Image        string    `json:"exercise_image" db:"exercise_image"`
	ExerciseTime int64     `json:"exercise_time" db:"exercise_time"`
	Date         time.Time `json:"exercise_date" db:"exercise_date"`
	Comment      string    `json:"comment" db:"comment"`
}

type ExerciseHandler struct {
	svc service.ExerciseService
}

func NewExerciseHandler(svc service.ExerciseService) *ExerciseHandler {
	return &ExerciseHandler{svc: svc}
}

func (h *ExerciseHandler) GetExercisesById(ctx *gin.Context) {
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
	exercises, err := h.svc.GetExercisesById(ctx, id)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "server internal error"})
		return
	}

	log.Println("Get Exercise User ", id, " exercises: ", exercises)
	ctx.JSON(http.StatusOK, exercises)
}

func (h *ExerciseHandler) GetExercises(ctx *gin.Context) {
	authz := ctx.GetHeader("Authorization")
	if !strings.HasPrefix(authz, "Bearer ") {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "missing bearer token"})
		return
	}
	exercises, err := h.svc.GetExercises(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "missing bearer token"})
		return
	}
	ctx.JSON(http.StatusOK, exercises)
}

func (h *ExerciseHandler) CreateExercise(ctx *gin.Context) {
	// ヘッダー取り出し
	authz := ctx.GetHeader("Authorization")
	if !strings.HasPrefix(authz, "Bearer ") {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "missing bearer token"})
		return
	}
	token := strings.TrimPrefix(authz, "Bearer ")

	var exerciseCreate exerciseCreateTypeDTO
	if err := ctx.ShouldBindJSON(&exerciseCreate); err != nil {
		log.Println("error: exercise read body ", err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "server internal error"})
		return
	}

	err := h.svc.CreateExercise(ctx.Request.Context(), exerciseCreate.Image, exerciseCreate.ExerciseTime, exerciseCreate.Date, exerciseCreate.Comment, token)
	if err != nil {
		log.Println("error: exercise create ", err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "server internal error"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "exercise record created successfully"})
}

func (h *ExerciseHandler) CreateLike(ctx *gin.Context) {
	authz := ctx.GetHeader("Authorization")
	if !strings.HasPrefix(authz, "Bearer ") {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "missing bearer token"})
		return
	}
	token := strings.TrimPrefix(authz, "Bearer ")

	var exerciseLike exerciseLikeTypeDTO
	if err := ctx.ShouldBindJSON(&exerciseLike); err != nil {
		log.Println("error: exercise like ", err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "server internal error"})
		return
	}

	err := h.svc.CreateLike(ctx.Request.Context(), exerciseLike.ExerciseRecordId, token)
	if err != nil {
		log.Println("error: exercise like ", err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "server internal error"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "exercise like created successfully"})
}
