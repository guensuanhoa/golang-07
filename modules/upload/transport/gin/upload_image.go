package ginupload

import (
	"net/http"

	"example.com/g07-food-delivery/component/appctx"
	uploadbiz "example.com/g07-food-delivery/modules/upload/biz"
	"github.com/gin-gonic/gin"
)

func UploadImage(appCtx appctx.AppContext) func(*gin.Context) {
	return func(c *gin.Context) {
		fileHeader, err := c.FormFile("file")

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		folder := c.DefaultPostForm("folder", "img")

		file, err := fileHeader.Open()

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Lo khuc duoi co van de thi ham defer se close file truoc khi ket thuc
		defer file.Close() // we can close here

		dataBytes := make([]byte, fileHeader.Size)
		if _, err := file.Read(dataBytes); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		//imgStore := uploadstorage.NewSQLStore(db)
		// Tang biz chi quan tam toi mang byte ma thoi
		biz := uploadbiz.NewUploadBiz(appCtx.UploadProvider())
		img, err := biz.Upload(c.Request.Context(), dataBytes, folder, fileHeader.Filename)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		img.Fulfill(appCtx.UploadProvider().GetDomain())

		c.JSON(http.StatusOK, gin.H{"data": img})
	}
}
