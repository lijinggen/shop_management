package file_server

import (
	"github.com/gin-gonic/gin"
	"github.com/shop_management/po/file_po"
	"github.com/shop_management/server/assembly/file_assembly"
	"github.com/shop_management/service"
	"github.com/shop_management/service/file_service"
	"github.com/shop_management/sm_error"
	"github.com/shop_management/sm_error/error_code"
	"github.com/shop_management/vars"
	"net/http"
	"strconv"
)

type FileServer struct {
	fileService service.FileServiceInterface
}

func NewFileServer() *FileServer {
	return &FileServer{
		fileService: file_service.NewFileService(),
	}
}

func (f *FileServer) Upload(ctx *gin.Context) {
	file, err := ctx.FormFile("file")
	if err != nil {
		ctx.JSON(http.StatusOK, sm_error.NewHttpError(error_code.FileUploadError))
		return
	}
	maxSize, _ := strconv.ParseInt(ctx.PostForm("max_size"), 10, 64)
	uploadFileUrl, err := f.fileService.UploadFile(ctx, file_assembly.ConvertURPoToDto(&file_po.UploadReq{
		File:    file,
		MaxSize: maxSize,
	}))
	if err != nil {
		vars.Log.Errorf("FileServer.Upload err:%v", err)
		if _, ok := err.(*sm_error.Error); ok {
			ctx.JSON(http.StatusOK, err)
			return
		}
		ctx.JSON(http.StatusOK, sm_error.NewHttpError(error_code.FileUploadError))
		return
	}
	ctx.JSON(http.StatusOK, file_po.UploadResp{Url: uploadFileUrl})
}
