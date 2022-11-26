package utils

import "strings"

type CustomError struct {
	ErrorID string `json:"errorId"`
	Inner   error
}

func (e *CustomError) Error() string {
	return strings.Join([]string{e.ErrorID, e.Inner.Error()}, `, `)
}

func NewCustomError(errorId string, err error) error {
	return &CustomError{errorId, err}
}

// golang style enum
type CustomErrorType int64

const (
	MissingHeaderError CustomErrorType = iota
	RtcMissingExternalIp
)

func (s CustomErrorType) String() string {
	switch s {
	case MissingHeaderError:
		return "MissingHeaderError"
	case RtcMissingExternalIp:
		return "RtcMissingExternalIp"
	}
	
	return "UnknownError"
}
