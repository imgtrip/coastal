package errors

import (
	"coastal/internal/env"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func Unknown(message ...string) error {
	return adapter(codes.Unknown, message...)
}

func InvalidArgument(message ...string) error {
	return adapter(codes.InvalidArgument, message...)
}

func NotFound(message ...string) error {
	return adapter(codes.NotFound, message...)
}
func AlreadyExists(message ...string) error {
	return adapter(codes.AlreadyExists, message...)
}
func PermissionDenied(message ...string) error {
	return adapter(codes.PermissionDenied, message...)
}
func ResourceExhausted(message ...string) error {
	return adapter(codes.ResourceExhausted, message...)
}
func FailedPrecondition(message ...string) error {
	return adapter(codes.FailedPrecondition, message...)
}
func Aborted(message ...string) error {
	return adapter(codes.Aborted, message...)
}
func OutOfRange(message ...string) error {
	return adapter(codes.OutOfRange, message...)
}
func Unimplemented(message ...string) error {
	return adapter(codes.Unimplemented, message...)
}
func Internal(message ...string) error {
	return adapter(codes.Internal, message...)
}

func DataLoss(message ...string) error {
	return adapter(codes.DataLoss, message...)
}

func New(c codes.Code, format string, a ...interface{}) error {
	return status.Errorf(c, format, a...)
}

func adapter(code codes.Code, message ...string) error {
	msg := ""
	if env.Process.Debug && len(message) > 0 {
		msg = message[0]
	}

	return New(code, msg)
}
