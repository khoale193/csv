package handler

import (
	"bufio"
	"errors"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"github.com/khoale193/csv/models"
	"github.com/khoale193/csv/modules/sync/service"
	"github.com/khoale193/csv/pkg/app"
	"github.com/khoale193/csv/pkg/e"
	"github.com/khoale193/csv/pkg/setting"
)

func SyncOhlcData(c *gin.Context) {
	var appG = app.Gin{C: c}
	form := SyncOhlcDataFormValidator{}
	if httpCode, err := form.BindAndValid(c); err != nil {
		appG.ResponseError(httpCode, app.NewValidatorError(err), nil)
		return
	}
	if err := form.service.SyncData(); err != nil {
		appG.Response(c, http.StatusOK, e.Msg[e.ERROR], e.ERROR, nil)
		return
	}
	// return that we have successfully uploaded our file!
	fmt.Fprintf(c.Writer, "Successfully Uploaded File")
}

type SyncOhlcDataFormValidator struct {
	service service.SyncOhlcService
}

func (v *SyncOhlcDataFormValidator) BindAndValid(c *gin.Context) (int, error) {
	file, handler, err := c.Request.FormFile("file")
	if err != nil {
		fmt.Println("Error Retrieving the File")
		fmt.Println(err)
		return http.StatusInternalServerError, errors.New(e.Msg[e.ERROR_IN_UPDATE_SUBJECT])
	}
	defer file.Close()
	// Create a temporary file within our temp-csv directory that follows
	_ = os.MkdirAll(setting.PathSetting.ResourceCsv, os.ModePerm)
	// a particular naming pattern
	if err := c.SaveUploadedFile(handler, setting.PathSetting.ResourceCsv+uuid.New().String()+filepath.Ext(handler.Filename)); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "Unable to save the file",
		})
		return http.StatusInternalServerError, errors.New(e.Msg[e.ERROR_CAN_NOT_SAVE_FILE])
	}
	var fileErr error
	scanner := bufio.NewScanner(file)
	// First line should be Header
	if scanner.Scan() {
		if line := strings.Split(scanner.Text(), ","); len(line) != e.NumberOfColumn {
			return http.StatusInternalServerError, errors.New(e.Msg[e.ERROR_INVALID_NUMBER_OF_COLUMN])
		}
	}
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		if line := strings.Split(scanner.Text(), ","); len(line) != e.NumberOfColumn {
			fileErr = errors.New(e.Msg[e.ERROR_INVALID_NUMBER_OF_COLUMN])
			break
		} else {
			// millisecond to int64
			millisecond, err := strconv.Atoi(line[0])
			if err != nil {
				fileErr = errors.New(e.Msg[e.ERROR_INVALID_UNIX_TIMESTAMP])
				break
			}
			open, err := strconv.ParseFloat(line[2], 64)
			if err != nil {
				fileErr = errors.New(e.Msg[e.ERROR_INVALID_UNIX_TIMESTAMP])
				break
			}
			high, err := strconv.ParseFloat(line[3], 64)
			if err != nil {
				fileErr = errors.New(e.Msg[e.ERROR_INVALID_UNIX_TIMESTAMP])
				break
			}
			low, err := strconv.ParseFloat(line[4], 64)
			if err != nil {
				fileErr = errors.New(e.Msg[e.ERROR_INVALID_UNIX_TIMESTAMP])
				break
			}
			close, err := strconv.ParseFloat(line[5], 64)
			if err != nil {
				fileErr = errors.New(e.Msg[e.ERROR_INVALID_UNIX_TIMESTAMP])
				break
			}
			v.service.Data = append(v.service.Data, models.Ohlc{Unix: int64(millisecond), Symbol: line[1], Open: open, High: high, Low: low, Close: close})
		}
	}
	if fileErr != nil {
		return http.StatusInternalServerError, errors.New(e.Msg[e.ERROR_IN_UPDATE_SUBJECT])
	}
	return http.StatusOK, nil
}
