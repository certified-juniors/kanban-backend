package apiutils

import (
	"fmt"
	resp "kanban/internal/lib/api/response"
	"kanban/internal/lib/logger/sl"
	"log/slog"

	"net/http"
	"strconv"
	"strings"
	"time"
)

const (
	DefaultPage  = 1
	DefaultLimit = 20
)

func ParseTimeParam(value string, target **time.Time, fieldName string, log *slog.Logger, w http.ResponseWriter) {
	if value == "" {
		return
	}

	value = strings.TrimSpace(value)

	parsed, err := time.Parse("2006-01-02 15:04:05", value)
	if err != nil {
		log.With(fieldName, value).Error("Error parsing "+fieldName, sl.Err(err))
		resp.WriteJSONResponse(w, http.StatusBadRequest, "invalid "+fieldName+" format", nil)
		return
	}

	*target = &parsed
}

func ParseIntParam(value string, target **int, fieldName string, log *slog.Logger, w http.ResponseWriter) {
	if value == "" {
		return
	}
	parsed, err := strconv.Atoi(value)
	if err != nil {
		log.With(fieldName, value).Error("Error parsing "+fieldName, sl.Err(err))
		resp.WriteJSONResponse(w, http.StatusBadRequest, fieldName+" parameter must be an integer", nil)
		return
	}
	*target = &parsed
}

func ParseBoolParam(value string, target **bool, fieldName string, log *slog.Logger, w http.ResponseWriter) {
	if value == "" {
		return
	}
	parsed, err := strconv.ParseBool(value)
	if err != nil {
		log.With(fieldName, value).Error("Error parsing "+fieldName, sl.Err(err))
		resp.WriteJSONResponse(w, http.StatusBadRequest, fieldName+" parameter must be a boolean", nil)
		return
	}
	*target = &parsed
}

func ParseIntArrayParam(value string, target *[]int, fieldName string, log *slog.Logger, w http.ResponseWriter) {
	if value == "" {
		return
	}

	values := strings.Split(value, ",")
	var result []int

	for _, v := range values {
		v = strings.TrimSpace(v)
		parsed, err := strconv.Atoi(v)
		if err != nil {
			log.With(fieldName, v).Error("error parsing "+fieldName, sl.Err(err))
			resp.WriteJSONResponse(w, http.StatusBadRequest, fieldName+" parameter must be an array of integers", nil)
			return
		}
		result = append(result, parsed)
	}

	*target = result
}

func ParseStringArrayParam(value string, target *[]string, fieldName string, log *slog.Logger, w http.ResponseWriter) {
	if value == "" {
		return
	}

	values := strings.Split(value, ",")
	var result []string

	for _, v := range values {
		v = strings.TrimSpace(v)
		if len(v) == 0 {
			log.With(fieldName, v).Error("error parsing "+fieldName, sl.Err(fmt.Errorf("empty string value")))
			resp.WriteJSONResponse(w, http.StatusBadRequest, fieldName+" parameter must be an array of non-empty strings", nil)
			return
		}
		result = append(result, v)
	}

	*target = result
}

func ParsePagination(pageStr, limitStr string) (page, limit int) {
	page, limit = DefaultPage, DefaultLimit

	if p, err := strconv.Atoi(pageStr); err == nil && p > 0 {
		page = p
	}
	if l, err := strconv.Atoi(limitStr); err == nil && l > 0 {
		limit = l
	}

	return
}

func GetLastDayOfMonth(t time.Time) *time.Time {
	year, month, _ := t.Date()
	lastDay := time.Date(year, month+1, 0, 0, 0, 0, 0, t.Location())
	return &lastDay
}

func GetFirstDayOfMonth(t time.Time) *time.Time {
	year, month, _ := t.Date()
	firstDay := time.Date(year, month, 1, 0, 0, 0, 0, t.Location())
	return &firstDay
}
