package app

import (
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/khoale193/csv/pkg/e"
	"github.com/khoale193/csv/pkg/pagination"
)

type Gin struct {
	C *gin.Context
}

type Response struct {
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
	Error   string      `json:"error"`
	Message string      `json:"message"`
	Status  string      `json:"status"`
}

// Response setting gin.JSON
func (g *Gin) Response(c *gin.Context, httpCode int, status string, errCode int, data interface{}) {
	origin := c.Request.Header.Get("Origin")
	c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
	c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
	g.C.JSON(httpCode, Response{
		Code:    httpCode,
		Data:    data,
		Error:   e.Msg[errCode],
		Message: e.Msg[errCode],
		Status:  status,
	})
	return
}

type ResponseError struct {
	Message string `json:"message"`
}

func (g *Gin) ResponseError(httpCode int, message string, data interface{}) {
	g.C.JSON(httpCode, ResponseError{
		Message: message,
	})
	return
}

func (g *Gin) WritePaginatorHeaderResponse(paginator pagination.Paginator) {
	if g.C.Request.Header.Get("X-Paging-Page") != "" {
		g.C.Writer.Header().Set("Access-Control-Expose-Headers", strings.Join([]string{"X-Paging-Page", "X-Paging-Size", "X-Paging-Pages", "X-Paging-Count"}, ","))
		g.C.Writer.Header().Set("X-Paging-Page", strconv.Itoa(paginator.Page()))
		g.C.Writer.Header().Set("X-Paging-Size", strconv.Itoa(paginator.PerPageNums()))
		g.C.Writer.Header().Set("X-Paging-Pages", strconv.Itoa(paginator.PageNums()))
		g.C.Writer.Header().Set("X-Paging-Count", strconv.FormatInt(paginator.Nums(), 10))
	}
}

func (g *Gin) SimpleResponse(c *gin.Context, httpCode int, data interface{}) {
	origin := c.Request.Header.Get("Origin")
	c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
	c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
	g.C.JSON(httpCode, struct {
		Data interface{} `json:"data"`
	}{Data: data})
	return
}
