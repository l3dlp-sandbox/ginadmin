package comment

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"html/template"
	"math"
	"math/rand"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"time"
)

/**
获取项目根目录
*/
func RootPath() (path string, err error) {
	sysType := runtime.GOOS

	if sysType == "linux" {
		path, err = filepath.Abs(filepath.Dir(os.Args[0]))
		if err != nil {
			fmt.Printf("path err %v", err)
		}
	}

	if sysType == "windows" {
		path, err = os.Getwd()
		if err != nil {
			fmt.Printf("path err %v", err)
		}
	}
	return
}

/**
生成随机字符串
*/
func RandString(len int) string {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	bytes := make([]byte, len)
	for i := 0; i < len; i++ {
		b := r.Intn(26) + 65
		bytes[i] = byte(b)
	}
	return string(bytes)
}

/**
密码加密
*/
func Encryption(password string, salt string) string {
	str := fmt.Sprintf("%s%s", password, salt)
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

/**
模板分页器
*/
type PageData struct {
	Count     int
	Data      interface{}
	Page      int
	PageHtml  template.HTML
	PageCount int
}

func PageOperation(c *gin.Context, db *gorm.DB, limit int, page int, data interface{}) PageData {
	var count float64
	db.Offset((page - 1) * limit).Limit(limit).Find(data)
	db.Count(&count)
	pageCount := int(math.Ceil(count / float64(limit)))
	url := c.FullPath()

	paramUrl := ""
	for k, v := range c.Request.URL.Query() {
		if k != "p" {
			paramUrl += "&" + k + "=" + v[0]
		}
	}
	fmt.Println(paramUrl)
	pageHtml := "<nav aria-label='Page navigation'><ul class='pagination'><li><a href='" + url + "?p=1' aria-label='Previous'><span aria-hidden='true'>&laquo;</span></a></li>"

	if pageCount < 5 {
		for i := 1; i <= pageCount; i++ {
			pageHtml += "<li><a href='" + url + "?p=" + strconv.Itoa(i) + paramUrl + "'>" + strconv.Itoa(i) + "</a></li>"
		}
	} else {
		if page > 2 && page < pageCount-2 {
			for i := page - 2; i <= page+2; i++ {
				pageHtml += "<li><a href='" + url + "?p=" + strconv.Itoa(i) + paramUrl + "'>" + strconv.Itoa(i) + "</a></li>"
			}
		}

		if page <= 2 {
			var maxPage int
			if pageCount > 5 {
				maxPage = 5
			} else {
				maxPage = pageCount
			}
			for i := 1; i <= maxPage; i++ {
				pageHtml += "<li><a href='" + url + "?p=" + strconv.Itoa(i) + paramUrl + "'>" + strconv.Itoa(i) + "</a></li>"
			}
		}

		if page >= pageCount-2 {
			for i := pageCount - 4; i <= pageCount; i++ {
				pageHtml += "<li><a href='" + url + "?p=" + strconv.Itoa(i) + paramUrl + "'>" + strconv.Itoa(i) + "</a></li>"
			}
		}
	}
	pageHtml += "<li><a href='" + url + "?p=" + strconv.Itoa(pageCount) + paramUrl + "' aria-label='Next'><span aria-hidden='true'>&raquo;</span></a></li><li></li></ul></nav>"
	return PageData{
		Count:     1,
		Data:      data,
		Page:      1,
		PageHtml:  template.HTML(pageHtml),
		PageCount: pageCount,
	}
}
