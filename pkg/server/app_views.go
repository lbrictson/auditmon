package server

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/lbrictson/auditmon/pkg/storage"
	log "github.com/sirupsen/logrus"

	"github.com/labstack/echo/v4"
)

func (s *Server) viewIndexPage(c echo.Context) error {
	tz := c.Get("userData").(SessionData).Timezone
	loc, _ := time.LoadLocation(tz)
	if loc == nil {
		loc, _ = time.LoadLocation("UTC")
	}
	limit := s.maxQueryResults
	offSet := 0
	qFilterType := c.QueryParam("filterBy")
	if qFilterType == "" {
		qFilterType = "Username"
	}
	qFilterString := c.QueryParam("filterString")
	qReadOnly := c.QueryParam("readOnly")
	qStart := c.QueryParam("start")
	// Time format like 2022-03-21T01:18
	timeLayout := "2006-01-02T15:04"
	if qStart == "" {
		t := time.Now().AddDate(0, 0, -1)
		qStart = fmt.Sprint(t.Format(timeLayout))
	}
	qEnd := c.QueryParam("end")
	if qEnd == "" {
		t := time.Now().Add(5 * time.Minute)
		qEnd = fmt.Sprint(t.Format(timeLayout))
	}
	startTime, err := time.Parse(timeLayout, qStart)
	if err != nil {
		log.Error(err)
		startTime = time.Now().AddDate(0, 0, -1)
	}
	endTime, err := time.Parse(timeLayout, qEnd)
	if err != nil {
		endTime = time.Now().Add(5 * time.Minute)
		log.Error(err)
	}
	t := time.Now().In(loc)
	by := -1
	_, offset := t.Zone()
	if offset >= 0 {
		by = 1
	}
	endTime = endTime.Add(time.Duration(offset*by) * time.Second).In(loc)
	startTime = startTime.Add(time.Duration(offset*by) * time.Second).In(loc)
	query := storage.EventStoreQueryBuilder{
		Username:    nil,
		Resource:    nil,
		StartTime:   startTime,
		EndTime:     endTime,
		EventName:   nil,
		RequestID:   nil,
		EventSource: nil,
		EventIP:     nil,
		ReadOnly:    nil,
		Limit:       limit,
		Page:        offSet,
	}
	if qReadOnly != "" {
		b, ok := strconv.ParseBool(qReadOnly)
		if ok == nil {
			query.ReadOnly = &b
		}
	}
	if qFilterType != "" {
		if qFilterString != "" {
			switch strings.ToLower(qFilterType) {
			case "username":
				query.Username = &qFilterString
			case "requestid":
				query.RequestID = &qFilterString
			case "ipaddress":
				query.EventIP = &qFilterString
			case "resource":
				query.Resource = &qFilterString
			case "eventname":
				query.EventName = &qFilterString
			case "eventsource":
				query.EventSource = &qFilterString
			}
		}
	}
	events, err := s.eventStorage.Query(context.TODO(), query)
	if err != nil {
		log.Error(err)
	}
	limited := false
	if len(events) == limit {
		limited = true
	}
	for i := range events {
		events[i].FrontendEventTime = fmt.Sprintf(events[i].EventTime.In(loc).Format("2006-01-02 03:04pm MST"))
	}
	return c.Render(http.StatusOK, "index", map[string]any{
		"Filters": map[string]any{
			"filterType":   qFilterType,
			"filterString": qFilterString,
			"readOnly":     qReadOnly,
			"startTime":    qStart,
			"endTime":      qEnd,
		},
		"Events":   events,
		"Limited":  limited,
		"Limit":    limit,
		"Timezone": tz,
	})
}

func (s *Server) formIndexPage(c echo.Context) error {
	type FormData struct {
		FilterType   string `form:"filterBy"`
		FilterString string `form:"filterInput"`
		Start        string `form:"startDate"`
		End          string `form:"endDate"`
		ReadOnly     string `form:"readOnly"`
	}
	data := FormData{}
	c.Bind(&data)
	params := url.Values{
		"filterBy":     []string{data.FilterType},
		"filterString": []string{data.FilterString},
		"readOnly":     []string{data.ReadOnly},
		"start":        []string{data.Start},
		"end":          []string{data.End},
	}
	redirectUrl := "/?" + params.Encode()
	return c.Redirect(http.StatusSeeOther, redirectUrl)
}

func (s *Server) hookGetEventByID(c echo.Context) error {
	id := c.Param("id")
	e, err := s.eventStorage.GetByID(context.TODO(), id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, e)
	}
	return c.JSONPretty(http.StatusOK, e, " ")
}
