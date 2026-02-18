package helper

func PtrString(s string) *string {
	return &s
}

func PtrUint16(v uint16) *uint16 {
	return &v
}

func DerefUint16(v *uint16) uint16 {
	if v == nil {
		return 0
	}
	return *v
}

func DerefUint32(v *uint32) uint32 {
	if v == nil {
		return 0
	}
	return *v
}

func DerefString(v *string) string {
	if v == nil {
		return ""
	}
	return *v
}
