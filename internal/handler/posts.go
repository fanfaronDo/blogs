package handler

import (
	"github.com/fanfaronDo/blogs/internal/domain"
	"github.com/gin-gonic/gin"
	"html/template"
	"math"
	"net/http"
	"strconv"
)

type Pagination struct {
	Next          int
	Previous      int
	CurrentPage   int
	TotalPage     int
	RecordPerPage []domain.Post
}

const (
	Limit = 3
)

func (h *Handler) getPosts(c *gin.Context) {
	paramID, ok := c.GetQuery("path")

	if !ok {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": paramID + " is required"})
		return
	}
	if paramID == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Param ID not found"})
		return
	}

	id, err := strconv.Atoi(paramID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Param ID must be an integer"})
		return
	}

	total, err := h.service.Posts.GetTotal()
	offset := id
	offset -= 1
	if offset != 0 || offset != total {
		offset *= Limit
	}

	posts, err := h.service.Posts.GetPosts(Limit, offset)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	pagination := Pagination{}
	pagination.Next = id + 1
	pagination.Previous = id - 1
	pagination.CurrentPage = id
	pagination.TotalPage = int(math.Ceil(float64(total) / float64(Limit)))
	pagination.RecordPerPage = posts

	tmpl, err := template.ParseFiles("web/templates/main.html")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to create templates"})
		return
	}

	c.Header("Content-Type", "text/html")
	err = tmpl.Execute(c.Writer, pagination)
	if err != nil {
		if _, err := c.Writer.WriteString("Page not found 404"); err != nil {
			return
		}
		return
	}
}

func (h *Handler) getPost(c *gin.Context) {
	postId := c.Param("id")
	id, err := strconv.Atoi(postId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Param ID must be an integer"})
		return
	}

	post, err := h.service.GetById(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, post)
}

func (h *Handler) createPost(c *gin.Context) {
	var post domain.Post
	if err := c.Bind(&post); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.service.Create(post)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, post)
}

func (h *Handler) updatePost(c *gin.Context) {
	postId := c.Param("id")
	id, err := strconv.Atoi(postId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Param ID must be an integer"})
		return
	}

	var post domain.Post
	if err = c.Bind(&post); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = h.service.Update(id, post)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, post)
}
