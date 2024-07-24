package common

func CastPointerToInt(i *int64) *int {
	if i == nil {
		return nil
	}
	v := int(*i)
	return &v
}

func From[T any](val T) *T {
	return &val
}

func EmptyStringIfNil(s *string) *string {
	if s == nil {
		return From("")
	}
	return s
}

func ValueOrDefault[T any](v *T, def T) T {
	if v == nil {
		return def
	}
	return *v
}
