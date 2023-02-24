package api

import (
	"math/rand"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/JaguarLiu/turn/model"
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
	resp := model.Payload{
		Message: "OK",
	}
	if err != nil {
		resp.Message = err.Error()
		ctx.JSON(http.StatusBadRequest, resp)
		return
	}
	files := form.File["files"]
	if len(files) == 0 {
		resp.Message = model.File_not_found
		ctx.JSON(http.StatusBadRequest, resp)
	}
	for _, file := range files {
		fixNum := rand.Intn(20)
		fileName := t.Format("20060102") + strconv.Itoa(fixNum) + ".xlsx"
		tempPath := filepath.Join(rootPath, "temp", fileName)
		ctx.SaveUploadedFile(file, tempPath)
		fileSrv.Save(fileName, tempPath)
	}
	ctx.JSON(http.StatusBadRequest, resp)
}
