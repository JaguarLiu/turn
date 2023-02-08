package api

import (
	"math/rand"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/JaguarLiu/turn/service"

	"github.com/gin-gonic/gin"
)

var fileSrv *service.FileSrv

func NewFileRouter(r *gin.Engine, fs *service.FileSrv) {
	group := r.Group("/files")
	group.POST("/upload", upload)
	fileSrv = fs
}
func upload(ctx *gin.Context) {
	t := time.Now()
	rand.Seed(t.UnixNano())
	rootPath, _ := os.Getwd()
	form, err := ctx.MultipartForm()
	if err != nil {
		ctx.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}
	files := form.File["files"]
	if len(files) == 0 {
		ctx.JSON(400, gin.H{
			"message": "file not found",
		})
	}
	for _, file := range files {
		fixNum := rand.Intn(20)
		fileName := t.Format("20060102") + strconv.Itoa(fixNum) + ".xlsx"
		tempPath := filepath.Join(rootPath, "temp", fileName)
		ctx.SaveUploadedFile(file, tempPath)
		fileSrv.Save(fileName, tempPath)
	}
	ctx.JSON(200, gin.H{
		"message": "ok",
	})
}
