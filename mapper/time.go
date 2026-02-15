package proto_mapper

import (
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func ToDtoTime(value *timestamppb.Timestamp) *time.Time {
	if value == nil || !value.IsValid() {
		return nil
	}
	t := value.AsTime()
	return &t
}

func ToTimestampProto(value *time.Time) *timestamppb.Timestamp {
	if value == nil || value.IsZero() {
		return nil
	}

	return timestamppb.New(*value)
}
