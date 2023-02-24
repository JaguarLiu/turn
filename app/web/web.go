package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewWebRouter(r *gin.Engine) {
	r.LoadHTMLGlob("./app/template/*/*.html")
	r.GET("/navbar", navbar)
	r.GET("/sidebar", sidebar)
	group := r.Group("/web")
	{
		group.GET("/", index)
		group.GET("/upload", upload)

	}

}
func index(ctx *gin.Context) {

	ctx.HTML(http.StatusOK, "index.html", gin.H{
		"content": "index",
	})
}
func upload(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "upload.html", gin.H{
		"label": "Select files",
	})

}
func navbar(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "_navbar.html", nil)
}
func sidebar(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "_sidebar.html", nil)
}
