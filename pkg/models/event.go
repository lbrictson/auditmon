package models

import "time"

type Event struct {
	EventID           string                 `json:"event_id"`
	EventName         string                 `json:"event_name"`
	EventTime         time.Time              `json:"event_time"`
	Username          string                 `json:"username"`
	Resource          string                 `json:"resource"`
	EventSource       string                 `json:"event_source"`
	SourceIPAddress   string                 `json:"source_ip_address"`
	RequestID         string                 `json:"request_id"`
	ReadOnly          bool                   `json:"read_only"`
	EventData         map[string]interface{} `json:"event_data"`
	FrontendEventTime string                 `json:"-"`
}
