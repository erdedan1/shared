package proto_mapper

func ToBoolProto(value *bool) bool {
	if value == nil {
		return false
	}

	return *value
}

func FromBoolProto(value bool) *bool {
	return &value
}

func ToPointBoolProto(value *bool) *bool {
	return value
}

func FromPointBoolProto(value *bool) *bool {
	return value
}
