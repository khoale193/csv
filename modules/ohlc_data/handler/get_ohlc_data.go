package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/khoale193/csv/modules/ohlc_data/service"
	"github.com/khoale193/csv/pkg/app"
	b "github.com/khoale193/csv/pkg/binding"
	"github.com/khoale193/csv/pkg/e"
	"github.com/khoale193/csv/pkg/pagination"
	"github.com/khoale193/csv/pkg/setting"
)

func GetOhlcData(c *gin.Context) {
	var appG = app.Gin{C: c}
	form := GetOhlcDataForm{}
	if httpCode, err := form.BindAndValid(c); err != nil {
		appG.ResponseError(httpCode, app.NewValidatorError(err), nil)
		return
	}
	data, err := form.service.GetOhlcData()
	if err != nil {
		appG.Response(c, http.StatusInternalServerError, e.Msg[e.ERROR], e.ERROR, nil)
		return
	} else if data == nil {
		appG.Response(c, http.StatusOK, e.Msg[e.SUCCESS], e.SUCCESS, []interface{}{})
		return
	}
	appG.WritePaginatorHeaderResponse(*form.service.Pagination)
	appG.Response(c, http.StatusOK, e.GetMsg(e.SUCCESS), e.SUCCESS, data)
}

type GetOhlcDataForm struct {
	service service.GetOhlcDataForm
}

func (v *GetOhlcDataForm) BindAndValid(c *gin.Context) (int, error) {
	v.service.Symbol = b.BindString(c.Request.URL.Query(), "symbol")
	v.service.Pagination = pagination.NewPaginator(c.Request, setting.AppSetting.PageSize, nil)
	return http.StatusOK, nil
}
