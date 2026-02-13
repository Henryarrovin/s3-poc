package handlers

import (
	"net/http"
	"s3-poc/services"

	"github.com/gin-gonic/gin"
)

type S3Handler struct {
	service *services.S3Service
}

func NewS3Handler(service *services.S3Service) *S3Handler {
	return &S3Handler{service: service}
}

func (h *S3Handler) CreateBucket(c *gin.Context) {
	name := c.Param("name")

	if err := h.service.CreateBucket(name); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "bucket created"})
}

func (h *S3Handler) ListObjects(c *gin.Context) {
	name := c.Param("name")

	objects, err := h.service.ListObjects(name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, objects)
}
