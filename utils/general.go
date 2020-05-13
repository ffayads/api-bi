package utils

import (
    "io"
    "os"
	"fmt"
	"time"
	"path/filepath"
	"mime/multipart"
	"github.com/globalsign/mgo/bson"
)

func StringToDate(dateString *string) (*time.Time, error) {
    layout := "2006-01-02 15:04:05"
    date, err := time.Parse(layout, *dateString)
    if err != nil {
	    return nil,err
    }
	return &date,nil
}

func UploadFile(file multipart.File, fileHeader *multipart.FileHeader) (string, error) {
    defer file.Close()
	path := os.Getenv("CDN")
	tempFileName := bson.NewObjectId().Hex() + filepath.Ext(fileHeader.Filename)
    fileCre, err := os.Create(path+tempFileName)
    if err != nil {
        fmt.Println(err)
    }
    defer fileCre.Close()
    _, err = io.Copy(fileCre, file)
    if err != nil {
        fmt.Println(err)
    }
	return tempFileName, nil
}
