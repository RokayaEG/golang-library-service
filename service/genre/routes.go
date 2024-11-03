package genre

import (
	"math"
	"net/http"

	types "github.com/RokayaEG/golang-library-service/types/genre"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Handler struct {
	store types.GenreStore
}

func NewHandler(store types.GenreStore) *Handler {
	return &Handler{
		store: store,
	}
}

func (h *Handler) RegisterRoutes(router *gin.RouterGroup) {
	router.POST("/list", h.searchGenres)
	router.POST("/", h.createGenre)
	router.GET("/:code", h.findGenreByCode)
	router.PUT("/:code", h.updateGenre)
	router.DELETE("/", h.deleteGenres)
}

func (h *Handler) createGenre(c *gin.Context) {
	var genrePayload types.CreateGenrePayload

	if err := c.ShouldBindBodyWithJSON(&genrePayload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error})
		return
	}

	code := uuid.New().String()

	_genre, err := h.store.CreateGenre(types.Genre{
		Code:        code,
		Name:        genrePayload.Name,
		Slug:        genrePayload.Slug,
		Description: genrePayload.Description,
	})

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusCreated, _genre)

}

func (h *Handler) updateGenre(c *gin.Context) {
	code := c.Param("code")

	var genrePayload types.CreateGenrePayload

	if err := c.ShouldBindBodyWithJSON(&genrePayload); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}

	_genre, err := h.store.UpdateGenre(code, types.Genre{
		Code:        code,
		Name:        genrePayload.Name,
		Slug:        genrePayload.Slug,
		Description: genrePayload.Description,
	})

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, _genre)

}

func (h *Handler) findGenreByCode(c *gin.Context) {
	code := c.Param("code")

	_genre, err := h.store.FindGenreByCode(code)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, _genre)

}

func (h *Handler) searchGenres(c *gin.Context) {
	var genrePayload types.GenreSearch

	if err := c.ShouldBindBodyWithJSON(&genrePayload); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}

	limit := 10
	page := 1

	if genrePayload.Limit > 0 {
		limit = genrePayload.Limit
	}

	if genrePayload.Page > 0 {
		page = genrePayload.Page
	}

	offset := (page - 1) * limit

	total, err := h.store.CountGenres(genrePayload)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}

	if offset > total {
		offset = 0
	}

	totalPages := int(math.Ceil(float64(total) / float64(limit)))

	if page > totalPages {
		page = 1
	}

	_genre, err := h.store.SearchGenres(genrePayload, limit, offset)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}

	if len(_genre) == 0 {
		_genre = []types.Genre{}
	}

	result := types.GenreSearchResponse{
		Total: total,
		Data:  _genre,
		Page:  page,
		Limit: limit,
	}

	c.IndentedJSON(http.StatusOK, result)

}

func (h *Handler) deleteGenres(c *gin.Context) {
	var deleteGenrePayload types.DeleteGenrepayload

	if err := c.ShouldBindBodyWithJSON(&deleteGenrePayload); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}
	if len(deleteGenrePayload.Codes) > 0 {
		err := h.store.DeleteGenres(deleteGenrePayload.Codes)
		if err != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"err": err.Error()})
			return
		}
		c.IndentedJSON(http.StatusOK, gin.H{"message": "Deleted Successfully"})
		return
	}
	c.IndentedJSON(http.StatusBadRequest, gin.H{"err": "codes are required"})

}
