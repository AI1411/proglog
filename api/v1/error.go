package log_v1

import (
	"fmt"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/status"
	"net/http"
)

type ErrOffsetOutOfRange struct {
	Offset uint64
}

func (e ErrOffsetOutOfRange) GRPCStatus() *status.Status {
	st := status.New(
		http.StatusNotFound,
		fmt.Sprintf("offset %d out of range", e.Offset),
	)
	msg := fmt.Sprintf(" this requested offse is outside the log's raange; %d", e.Offset)

	d := &errdetails.LocalizedMessage{Locale: "ja", Message: msg}
	std, err := st.WithDetails(d)
	if err != nil {
		return st
	}
	return std
}

func (e ErrOffsetOutOfRange) Error() string {
	return e.GRPCStatus().Err().Error()
}
