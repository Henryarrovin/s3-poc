package handlers

import (
	"net/http"
	"s3-poc/data"

	"github.com/gin-gonic/gin"
)

type S3Handler struct {
	repo *data.S3Repository
}

func NewS3Handler(repo *data.S3Repository) *S3Handler {
	return &S3Handler{repo: repo}
}

func (h *S3Handler) CreateBucket(c *gin.Context) {
	bucketName := c.Param("name")
	if err := h.repo.CreateBucket(bucketName); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "bucket created"})
}

func (h *S3Handler) ListObjects(c *gin.Context) {
	bucketName := c.Param("name")
	objects, err := h.repo.ListObjects(bucketName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, objects)
}
