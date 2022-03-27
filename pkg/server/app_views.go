package server

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/lbrictson/auditmon/pkg/storage"
	log "github.com/sirupsen/logrus"

	"github.com/labstack/echo/v4"
)

func (s *Server) viewIndexPage(c echo.Context) error {
	limit := 10000
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
	return c.Render(http.StatusOK, "index", map[string]any{
		"Filters": map[string]any{
			"filterType":   qFilterType,
			"filterString": qFilterString,
			"readOnly":     qReadOnly,
			"startTime":    qStart,
			"endTime":      qEnd,
		},
		"Events": events,
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
