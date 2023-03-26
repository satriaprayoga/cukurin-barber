package controllers

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/satriaprayoga/cukurin-barber/form"
	"github.com/satriaprayoga/cukurin-barber/middlewares"
	"github.com/satriaprayoga/cukurin-barber/models"
	"github.com/satriaprayoga/cukurin-barber/pkg/logging"
	"github.com/satriaprayoga/cukurin-barber/pkg/response"
	"github.com/satriaprayoga/cukurin-barber/pkg/utils"
	"github.com/satriaprayoga/cukurin-barber/services"
)

type FileUploadController struct {
	fileUploadService services.IFileUploadService
}

func NewFileUploadController(e *echo.Echo, fileUploadService services.IFileUploadService) {
	controller := &FileUploadController{
		fileUploadService: fileUploadService,
	}
	e.Static("/wwwroot", "wwwroot")
	r := e.Group("/fileupload")
	r.Use(middlewares.JWT)
	r.POST("", controller.CreateImage)
}

func (f *FileUploadController) CreateImage(e echo.Context) (err error) {
	ctx := e.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	var (
		resp          = response.Resp{R: e}
		imageFormList []models.FileUpload
		logger        = logging.Logger{}
	)

	mform, err := e.MultipartForm()
	if err != nil {
		return err
	}
	images := mform.File["upload_file"]
	pt := mform.Value["path"]
	logger.Info(pt)

	dir, err := os.Getwd()
	if err != nil {
		return err
	}
	claims, err := form.GetClaims(e)
	if err != nil {
		return resp.ResponseError(http.StatusBadRequest, fmt.Sprintf("%v", err), nil)
	}
	for i, image := range images {
		//Source
		src, err := image.Open()
		if err != nil {
			return err
		}
		defer src.Close()

		var dir_file = dir + "/wwwroot/uploads"
		var path_file = "/wwwroot/uploads"
		err = utils.IsNotExistMkDir(dir_file)
		if err != nil {
			return err
		}

		if pt[i] != "" {
			dirx := dir_file
			for _, val := range strings.Split(pt[i], "/") {
				dirx = dirx + "/" + val
				fmt.Print(dirx)
				err = utils.IsNotExistMkDir(dirx)
				if err != nil {
					return err
				}
			}
			dir_file = fmt.Sprintf("%s/%s", dir_file, pt[i])

			path_file = fmt.Sprintf("%s/%s", path_file, pt[i])

		}
		fileNameAndUnix := fmt.Sprintf("%d_%s", utils.GetTimeNow().Unix(), image.Filename)

		// Destination
		dest := fmt.Sprintf("%s/%s", dir_file, fileNameAndUnix)
		dst, err := os.Create(dest)
		if err != nil {
			return err
		}
		defer dst.Close()

		// Copy
		if _, err = io.Copy(dst, src); err != nil {
			return err
		}

		var imageForm models.FileUpload
		imageForm.FileName = fileNameAndUnix
		imageForm.FilePath = fmt.Sprintf("%s/%s", path_file, fileNameAndUnix)
		imageForm.FileType = filepath.Ext(fileNameAndUnix)
		imageForm.UserInput = claims.Username
		imageForm.UserEdit = claims.Username

		err = f.fileUploadService.CreateFileUpload(ctx, &imageForm)
		if err != nil {
			return err
		}
		imageFormList = append(imageFormList, imageForm)
	}

	return resp.Response(http.StatusOK, "Ok", imageFormList)
}
