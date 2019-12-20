package gles20

import "errors"

var (
	InvalidEnum                 = errors.New("Invalid Enum")
	InvalidValue                = errors.New("Invalid Value")
	InvalidOperation            = errors.New("Invalid Operation")
	InvalidFrameBufferOperation = errors.New("Invalid Frame Buffer Operation")
	OutOfMemory                 = errors.New("Out Of Memory")
)

func int2error(err int) error {
	switch err {
	case NO_ERROR:
		return nil
	case INVALID_ENUM:
		return InvalidEnum
	case INVALID_VALUE:
		return InvalidValue
	case INVALID_OPERATION:
		return InvalidOperation
	}
	return errors.New("Unknown error")
}
