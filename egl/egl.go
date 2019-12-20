package egl

func ErrString(errno uintptr) string {
	switch errno {
	case SUCCESS:
		return "SUCCESS"
	case NOT_INITIALIZED:
		return "NOT_INITIALIZED"
	case BAD_ACCESS:
		return "BAD_ACCESS"
	case BAD_ALLOC:
		return "BAD_ALLOC"
	case BAD_ATTRIBUTE:
		return "BAD_ATTRIBUTE"
	case BAD_CONFIG:
		return "BAD_CONFIG"
	case BAD_CONTEXT:
		return "BAD_CONTEXT"
	case BAD_CURRENT_SURFACE:
		return "BAD_CURRENT_SURFACE"
	case BAD_DISPLAY:
		return "BAD_DISPLAY"
	case BAD_MATCH:
		return "BAD_MATCH"
	case BAD_NATIVE_PIXMAP:
		return "BAD_NATIVE_PIXMAP"
	case BAD_NATIVE_WINDOW:
		return "BAD_NATIVE_WINDOW"
	case BAD_PARAMETER:
		return "BAD_PARAMETER"
	case BAD_SURFACE:
		return "BAD_SURFACE"
	case CONTEXT_LOST:
		return "CONTEXT_LOST"
	}
	return "EGL: unknown error"
}
