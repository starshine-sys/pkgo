package pkgo

// NullableString is a nullable string for use in post/patch endpoints
type NullableString = *string

// NewNullableString creates a new nullable string with the given value
func NewNullableString(s string) NullableString {
	return &s
}
