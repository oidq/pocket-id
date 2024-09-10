package model

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

type AuditLog struct {
	Base

	Event     AuditLogEvent
	IpAddress string
	UserAgent string
	UserID    string
	Data      AuditLogData
}

type AuditLogData map[string]string

type AuditLogEvent string

const (
	AuditLogEventSignIn                 AuditLogEvent = "SIGN_IN"
	AuditLogEventClientAuthorization    AuditLogEvent = "CLIENT_AUTHORIZATION"
	AuditLogEventNewClientAuthorization AuditLogEvent = "NEW_CLIENT_AUTHORIZATION"
	AuditLogEventAddPasskey             AuditLogEvent = "ADD_NEW_PASSKEY"
	AuditLogEventRemovePasskey          AuditLogEvent = "REMOVE_PASSKEY"
)

// Scan and Value methods for GORM to handle the custom type

func (e *AuditLogEvent) Scan(value interface{}) error {
	*e = AuditLogEvent(value.(string))
	return nil
}

func (e AuditLogEvent) Value() (driver.Value, error) {
	return string(e), nil
}

func (d *AuditLogData) Scan(value interface{}) error {
	if v, ok := value.([]byte); ok {
		return json.Unmarshal(v, d)
	} else {
		return errors.New("type assertion to []byte failed")
	}
}

func (d AuditLogData) Value() (driver.Value, error) {
	return json.Marshal(d)
}