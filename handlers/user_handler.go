package handlers

import (
	"net/http"
	"s3-poc/data"
	"s3-poc/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userRepo *data.UserRepository
	s3Repo   *data.S3Repository
}

func NewUserHandler(userRepo *data.UserRepository, s3Repo *data.S3Repository) *UserHandler {
	return &UserHandler{userRepo: userRepo, s3Repo: s3Repo}
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	var req struct {
		Name       string `json:"name"`
		BucketName string `json:"bucket_name"`
	}
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if req.Name == "" || req.BucketName == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "name and bucket_name are required"})
		return
	}

	user := &models.User{
		Name:       req.Name,
		BucketName: req.BucketName,
	}

	if err := h.userRepo.CreateUser(user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err := h.s3Repo.CreateBucket(req.BucketName); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create bucket"})
		return
	}

	c.JSON(http.StatusOK, user)
}

func (h *UserHandler) GetUser(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	user, err := h.userRepo.GetUser(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}
