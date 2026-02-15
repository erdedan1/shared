package proto_mapper

import (
	"github.com/google/uuid"
)

func FromUUIDProto(s string) *uuid.UUID {
	parsed, err := uuid.Parse(s)
	if err != nil {
		return nil
	}

	return &parsed
}

func FromUUIDsProto(s []string) []uuid.UUID {
	result := make([]uuid.UUID, 0, len(s))
	for i, v := range s {
		parsed, err := uuid.Parse(v)
		if err != nil {
			return nil
		}
		result[i] = parsed
	}

	return result
}

func ToUUIDProto(u uuid.UUID) string {
	return u.String()
}

func ToUUIDsProto(u []uuid.UUID) []string {
	result := make([]string, 0, len(u))
	for i, v := range u {
		result[i] = ToUUIDProto(v)
	}
	return result
}
