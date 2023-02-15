package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewWebRouter(r *gin.Engine) {
	r.LoadHTMLGlob("./app/template/pages/*.html")
	group := r.Group("/web")
	group.GET("/upload", upload)
}

func upload(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "upload.html", gin.H{
		"label": "Select files",
	})
}
