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
	exercises, err := h.svc.GetRecordsById(ctx, id)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "server internal error"})
		return
	}

	log.Println("Get Exercise User ", id, " exercises: ", exercises)
	ctx.JSON(http.StatusOK, exercises)
}

func (h *ExerciseHandler) GetRecords(ctx *gin.Context) {
	authz := ctx.GetHeader("Authorization")
	if !strings.HasPrefix(authz, "Bearer ") {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "missing bearer token"})
		return
	}
	exercises, err := h.svc.GetRecords(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "missing bearer token"})
		return
	}
	ctx.JSON(http.StatusOK, exercises)
}

func (h *ExerciseHandler) CreateLike(ctx *gin.Context) {
	authz := ctx.GetHeader("Authorization")
	if !strings.HasPrefix(authz, "Bearer ") {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "missing bearer token"})
		return
	}
	token := strings.TrimPrefix(authz, "Bearer ")

	var recordLike dto.RecordLikeType
	if err := ctx.ShouldBindJSON(&recordLike); err != nil {
		log.Println("error: exercise like ", err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "server internal error"})
		return
	}

	err := h.svc.CreateLike(ctx.Request.Context(), recordLike.RecordId, token)
	if err != nil {
		log.Println("error: exercise like ", err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "server internal error"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "exercise like created successfully"})
}

func (h *ExerciseHandler) CheckLike(ctx *gin.Context) {
	authz := ctx.GetHeader("Authorization")
	if !strings.HasPrefix(authz, "Bearer ") {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "missing bearer token"})
		return
	}
	token := strings.TrimPrefix(authz, "Bearer ")

	exerciseRecordIdStr := ctx.Param("record_id")
	log.Println("record ->", exerciseRecordIdStr)
	exerciseRecordId, err := strconv.ParseInt(exerciseRecordIdStr, 10, 64)
	if err != nil {
		log.Println("failed to parse record_id")
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "bad request"})
		return
	}
	liked, err := h.svc.CheckLikeById(ctx.Request.Context(), exerciseRecordId, token)
	if err != nil {
		log.Println("error check like ", err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "server internal error"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"liked": liked})

}

func (h *ExerciseHandler) DeleteLike(ctx *gin.Context) {
	authz := ctx.GetHeader("Authorization")
	if !strings.HasPrefix(authz, "Bearer ") {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "missing bearer token"})
		return
	}
	token := strings.TrimPrefix(authz, "Bearer ")
	exerciseRecordIdStr := ctx.Param("record_id")
	log.Println("record str -> ", exerciseRecordIdStr)
	exerciseRecordId, err := strconv.ParseInt(exerciseRecordIdStr, 10, 64)
	if err != nil {
		log.Println("failed to parse record_id")
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "bad request"})
		return
	}

	err = h.svc.DeleteLikeById(ctx.Request.Context(), exerciseRecordId, token)
	if err != nil {
		log.Println("error delete like ", err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "server internal error"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "exercise like deleted successfully"})

}

func (h *ExerciseHandler) GenerateIllustration(ctx *gin.Context) {
	// ヘッダー取り出し
	authz := ctx.GetHeader("Authorization")
	if !strings.HasPrefix(authz, "Bearer ") {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "missing bearer token"})
		return
	}
	token := strings.TrimPrefix(authz, "Bearer ")

	s3KeyRaw := ctx.PostForm("s3_key")
	if s3KeyRaw == "" {
		log.Println("error")
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "server internal error"})
		return
	}

	cleanUpTimeRaw := ctx.PostForm("clean_up_time")
	if cleanUpTimeRaw == "" {
		log.Println("error clean up time is not set")
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "server internal error"})
		return
	}

	cleanUpDateRaw := ctx.PostForm("clean_up_date")
	if cleanUpDateRaw == "" {
		log.Println("error clean date time is not set")
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "server internal error"})
		return
	}

	comment := ctx.PostForm("comment")
	if comment == "" {
		log.Println("error comment is not set")
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "server internal error"})
		return
	}

	image, err := ctx.FormFile("file")
	if err != nil {
		log.Println("error", err)
		return
	}
	recordCreate := dto.RecordCreateType{
		ObjectKey:      s3KeyRaw,
		CleanUpTimeRaw: cleanUpTimeRaw,
		CleanUpDateRaw: cleanUpDateRaw,
		Comment:        comment,
	}

	s3Key, err := h.svc.GenerateImgAndUpload(ctx.Request.Context(), image, s3KeyRaw)
	if err != nil {
		log.Println("image dir error", err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "server internal error"})
		return
	}

	recordCreate.ObjectKey = s3Key

	err = h.svc.CreateRecord(ctx.Request.Context(), recordCreate.ObjectKey, recordCreate.CleanUpTimeRaw, recordCreate.CleanUpDateRaw, recordCreate.Comment, token)
	if err != nil {
		log.Println("error: create record", err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "server internal error"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "record is created successfully"})

}
