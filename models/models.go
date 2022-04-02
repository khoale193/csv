package models

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/khoale193/csv/pkg/s"
)

func LimitCondition(limit int) string {
	if limit == s.CancelLimit {
		return ""
	} else if limit > 0 {
		return fmt.Sprintf(" limit %d ", limit)
	} else {
		return ""
	}
}

func OffsetCondition(offset int) string {
	if offset == s.CancelOffset {
		return ""
	} else if offset > 0 {
		return fmt.Sprintf(" offset %d ", offset)
	} else {
		return ""
	}
}

func ColName(i interface{}, fieldByName string) string {
	field, ok := reflect.TypeOf(i).Elem().FieldByName(fieldByName)
	if !ok {
		return ""
	}
	return field.Tag.Get("db")
}

func ColumnsName(tableName string, i interface{}, columnName []string) string {
	var bar []string
	for _, item := range columnName {
		field, ok := reflect.TypeOf(i).Elem().FieldByName(item)
		if !ok {
			continue
		}
		bar = append(bar, tableName+"."+field.Tag.Get("db"))
	}
	if len(bar) > 0 {
		return strings.Join(bar, ",")
	}
	return ""
}

func ColumnsNameValueList(i interface{}, columnName []string) string {
	var bar []string
	for _, item := range columnName {
		field, ok := reflect.TypeOf(i).Elem().FieldByName(item)
		if !ok {
			continue
		}
		bar = append(bar, ":"+field.Tag.Get("db"))
	}
	if len(bar) > 0 {
		return strings.Join(bar, ",")
	}
	return ""
}

func SelectQueryBuilder(str ...string) string {
	return strings.Join(str, ",")
}
