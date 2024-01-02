package file_service

import (
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/gin-gonic/gin"
	"github.com/shop_management/dto/file_dto"
	"github.com/shop_management/sm_error"
	"github.com/shop_management/sm_error/error_code"
	"github.com/shop_management/util"
	"github.com/xuri/excelize/v2"
	"io"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"time"
)

type fileService struct {
}

func NewFileService() *fileService {
	return &fileService{}
}

func (f *fileService) UploadFile(ctx *gin.Context, req *file_dto.UploadReq) (string, error) {
	finalFileName := util.FormatTime(time.Now()) + "_" + strconv.FormatInt(int64(rand.Intn(100)), 10) + "_" + req.File.Filename

	if req.MaxSize > 0 && req.File.Size > req.MaxSize {
		return "", sm_error.NewHttpError(error_code.FileSizeOutOfMax)
	}
	uploadedFile, err := req.File.Open()
	if err != nil {
		return "", err
	}
	defer uploadedFile.Close()

	tempFile, err := os.Create(finalFileName)
	if err != nil {
		return "", err
	}
	_, err = io.Copy(tempFile, uploadedFile)
	if err != nil {
		return "", err
	}
	err = tempFile.Sync()
	if err != nil {
		return "", err
	}
	tempFile.Close()

	tempLocalFile, err := os.Open(finalFileName)
	if err != nil {
		return "", err
	}
	provider, err := oss.NewEnvironmentVariableCredentialsProvider()
	if err != nil {
		return "", err
	}
	defer func() {
		tempLocalFile.Close()
		os.Remove(finalFileName)
	}()
	client, err := oss.New("oss-cn-shenzhen.aliyuncs.com", os.Getenv("OSS_ACCESS_KEY_ID"), os.Getenv("OSS_ACCESS_KEY_SECRET"), oss.SetCredentialsProvider(&provider))
	if err != nil {
		return "", err
	}
	bucket, err := client.Bucket("szwkoss")
	if err != nil {
		return "", err
	}
	err = bucket.PutObject(finalFileName, tempLocalFile)
	if err != nil {
		return "", err
	}

	return "https://szwkoss.oss-cn-shenzhen.aliyuncs.com/" + finalFileName, nil
}

func (f *fileService) UploadRemoteFile(ctx *gin.Context, url string, suffix string) (string, error) {
	finalFileName := util.FormatTime(time.Now()) + "_" + strconv.FormatInt(int64(rand.Intn(100)), 10) + "." + suffix

	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	provider, err := oss.NewEnvironmentVariableCredentialsProvider()
	if err != nil {
		return "", err
	}
	client, err := oss.New("oss-cn-shenzhen.aliyuncs.com", os.Getenv("OSS_ACCESS_KEY_ID"), os.Getenv("OSS_ACCESS_KEY_SECRET"), oss.SetCredentialsProvider(&provider))
	if err != nil {
		return "", err
	}
	bucket, err := client.Bucket("szwkoss")
	if err != nil {
		return "", err
	}
	err = bucket.PutObject(finalFileName, resp.Body)
	if err != nil {
		return "", err
	}
	return "https://szwkoss.oss-cn-shenzhen.aliyuncs.com/" + finalFileName, nil
}

func (f *fileService) UploadLocalFile(ctx *gin.Context, path string, suffix string) (string, error) {
	finalFileName := util.FormatTime(time.Now()) + "_" + strconv.FormatInt(int64(rand.Intn(100)), 10) + "." + suffix
	open, err := os.Open(path)
	if err != nil {
		return "", err
	}

	defer open.Close()

	provider, err := oss.NewEnvironmentVariableCredentialsProvider()
	if err != nil {
		return "", err
	}
	client, err := oss.New("oss-cn-shenzhen.aliyuncs.com", os.Getenv("OSS_ACCESS_KEY_ID"), os.Getenv("OSS_ACCESS_KEY_SECRET"), oss.SetCredentialsProvider(&provider))
	if err != nil {
		return "", err
	}
	bucket, err := client.Bucket("szwkoss")
	if err != nil {
		return "", err
	}
	err = bucket.PutObject(finalFileName, open)
	if err != nil {
		return "", err
	}
	return "https://szwkoss.oss-cn-shenzhen.aliyuncs.com/" + finalFileName, nil
}

func (f *fileService) ParseExcel(ctx *gin.Context, url string, dealData func(*gin.Context, [][]string) error) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer func() {
		resp.Body.Close()
	}()
	fileReader, err := excelize.OpenReader(resp.Body)
	if err != nil {
		return err
	}
	defer func() {
		fileReader.Close()
	}()
	rows, err := fileReader.Rows("sheet1")
	if err != nil {
		return err
	}
	batchSize := 50
	results := make([][]string, 0, 64)
	for rows.Next() {
		row, err := rows.Columns()
		if err != nil {
			break
		}
		results = append(results, row)
		if len(results) >= batchSize {
			err := dealData(ctx, results)
			if err != nil {
				return err
			}
			results = make([][]string, 0, 64)
		}
	}
	if len(results) != 0 {
		err := dealData(ctx, results)
		if err != nil {
			return err
		}
	}
	return nil
}
