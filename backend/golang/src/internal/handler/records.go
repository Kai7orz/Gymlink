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

type RecordHandler struct {
	svc service.RecordService
}

func NewRecordHandler(svc service.RecordService) *RecordHandler {
	return &RecordHandler{svc: svc}
}

func (h *RecordHandler) GetRecordsById(ctx *gin.Context) {
	authz := ctx.GetHeader("Authorization")
	if !strings.HasPrefix(authz, "Bearer ") {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "missing bearer token"})
		return
	}
	idStr := ctx.Param("user_id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		log.Println("failed to parse user_id")
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "server internal error"})
		return
	}
	records, err := h.svc.GetRecordsById(ctx, id)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "server internal error"})
		return
	}
	ctx.JSON(http.StatusOK, records)
}

func (h *RecordHandler) GetRecords(ctx *gin.Context) {
	authz := ctx.GetHeader("Authorization")
	if !strings.HasPrefix(authz, "Bearer ") {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "missing bearer token"})
		return
	}
	records, err := h.svc.GetRecords(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "missing bearer token"})
		return
	}
	ctx.JSON(http.StatusOK, records)
}

func (h *RecordHandler) DeleteRecord(ctx *gin.Context) {
	authz := ctx.GetHeader("Authorization")
	if !strings.HasPrefix(authz, "Bearer ") {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "missing bearer token"})
		return
	}
	token := strings.TrimPrefix(authz, "Bearer ")

	userIdStr := ctx.Param("user_id")
	recordIdStr := ctx.Param("record_id")
	if userIdStr == "" || recordIdStr == "" {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": " failed to receive paramas: user_id or record_id"})
		return
	}

	userId, err := strconv.ParseInt(userIdStr, 10, 64)
	if err != nil {
		log.Println("failed to parse user_id")
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "failed to parse user_id"})
	}

	recordId, err := strconv.ParseInt(recordIdStr, 10, 64)
	if err != nil {
		log.Println("failed to parse record_id")
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "bad request"})
		return
	}

	h.svc.DeleteRecordById(ctx.Request.Context(), userId, recordId, token)
}

func (h *RecordHandler) CreateLike(ctx *gin.Context) {
	authz := ctx.GetHeader("Authorization")
	if !strings.HasPrefix(authz, "Bearer ") {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "missing bearer token"})
		return
	}
	token := strings.TrimPrefix(authz, "Bearer ")

	var recordLike dto.RecordLikeType
	if err := ctx.ShouldBindJSON(&recordLike); err != nil {
		log.Println("error: record like ", err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "server internal error"})
		return
	}

	err := h.svc.CreateLike(ctx.Request.Context(), recordLike.RecordId, token)
	if err != nil {
		log.Println("error: record like ", err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "server internal error"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "record like created successfully"})
}

func (h *RecordHandler) CheckLike(ctx *gin.Context) {
	authz := ctx.GetHeader("Authorization")
	if !strings.HasPrefix(authz, "Bearer ") {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "missing bearer token"})
		return
	}
	token := strings.TrimPrefix(authz, "Bearer ")

	recordIdStr := ctx.Param("record_id")
	recordId, err := strconv.ParseInt(recordIdStr, 10, 64)
	if err != nil {
		log.Println("failed to parse record_id")
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "bad request"})
		return
	}
	liked, err := h.svc.CheckLikeById(ctx.Request.Context(), recordId, token)
	if err != nil {
		log.Println("error check like ", err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "server internal error"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"liked": liked})

}

func (h *RecordHandler) DeleteLike(ctx *gin.Context) {
	authz := ctx.GetHeader("Authorization")
	if !strings.HasPrefix(authz, "Bearer ") {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "missing bearer token"})
		return
	}
	token := strings.TrimPrefix(authz, "Bearer ")
	recordIdStr := ctx.Param("record_id")
	recordId, err := strconv.ParseInt(recordIdStr, 10, 64)
	if err != nil {
		log.Println("failed to parse record_id")
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "bad request"})
		return
	}

	err = h.svc.DeleteLikeById(ctx.Request.Context(), recordId, token)
	if err != nil {
		log.Println("error delete like ", err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "server internal error"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "record like deleted successfully"})

}

func (h *RecordHandler) GenerateIllustration(ctx *gin.Context) {
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
