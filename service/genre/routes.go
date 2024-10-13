package genre

import "github.com/gin-gonic/gin"

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) RegisterRoutes(router *gin.RouterGroup) {
	router.POST("/list", h.searchGenres)
	router.POST("/", h.createGenre)
	router.GET("/:code", h.findGenreByCode)
	router.PUT("/:code", h.updateGenre)
	router.DELETE("/", h.deleteGenres)
}

func (h *Handler) createGenre(c *gin.Context) {

}

func (h *Handler) updateGenre(c *gin.Context) {

}

func (h *Handler) findGenreByCode(c *gin.Context) {

}

func (h *Handler) searchGenres(c *gin.Context) {

}

func (h *Handler) deleteGenres(c *gin.Context) {

}
