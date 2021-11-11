package pkgo

// NullableString is a nullable string for use in post/patch endpoints
type NullableString = *string

// NewNullableString creates a new nullable string with the given value
func NewNullableString(s string) NullableString {
	return &s
}

// NullableBool is a nullable bool for use in post/patch endpoints
type NullableBool = *bool

// NewNullableBool creates a new nullable bool with the given value
func NewNullableBool(b bool) NullableBool {
	return &b
}
