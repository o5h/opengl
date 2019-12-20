package egl

import (
	"log"
)

var rgb888 = []int32{
	RENDERABLE_TYPE, OPENGL_ES2_BIT,
	SURFACE_TYPE, WINDOW_BIT,
	BLUE_SIZE, 8,
	GREEN_SIZE, 8,
	RED_SIZE, 8,
	DEPTH_SIZE, 16,
	STENCIL_SIZE, 8,
	NONE,
}

func CreateEGLSurface(nativeDisplay NativeDisplay, nativeWindow NativeWindow) (context Context, display Display, surface Surface, err error) {
	var displayAttribPlatforms = [][]int32{
		// Default
		[]int32{
			PLATFORM_ANGLE_TYPE_ANGLE,
			PLATFORM_ANGLE_TYPE_DEFAULT_ANGLE,
			PLATFORM_ANGLE_MAX_VERSION_MAJOR_ANGLE, DONT_CARE,
			PLATFORM_ANGLE_MAX_VERSION_MINOR_ANGLE, DONT_CARE,
			NONE,
		},
		// Direct3D 11
		[]int32{
			PLATFORM_ANGLE_TYPE_ANGLE,
			PLATFORM_ANGLE_TYPE_D3D11_ANGLE,
			PLATFORM_ANGLE_MAX_VERSION_MAJOR_ANGLE, DONT_CARE,
			PLATFORM_ANGLE_MAX_VERSION_MINOR_ANGLE, DONT_CARE,
			NONE,
		},
		// Direct3D 9
		[]int32{
			PLATFORM_ANGLE_TYPE_ANGLE,
			PLATFORM_ANGLE_TYPE_D3D9_ANGLE,
			PLATFORM_ANGLE_MAX_VERSION_MAJOR_ANGLE, DONT_CARE,
			PLATFORM_ANGLE_MAX_VERSION_MINOR_ANGLE, DONT_CARE,
			NONE,
		},
		// Direct3D 11 with WARP
		//   https://msdn.microsoft.com/en-us/library/windows/desktop/gg615082.aspx
		[]int32{
			PLATFORM_ANGLE_TYPE_ANGLE,
			PLATFORM_ANGLE_TYPE_D3D11_ANGLE,
			PLATFORM_ANGLE_DEVICE_TYPE_ANGLE,
			PLATFORM_ANGLE_DEVICE_TYPE_WARP_ANGLE,
			PLATFORM_ANGLE_MAX_VERSION_MAJOR_ANGLE, DONT_CARE,
			PLATFORM_ANGLE_MAX_VERSION_MINOR_ANGLE, DONT_CARE,
			NONE,
		},
	}

	display = NO_DISPLAY
	for i, displayAttrib := range displayAttribPlatforms {
		lastTry := i == len(displayAttribPlatforms)-1
		display, err = GetPlatformDisplayEXT(PLATFORM_ANGLE_ANGLE, nativeDisplay, displayAttrib)

		if display == NO_DISPLAY {
			if !lastTry {
				continue
			}
			log.Printf("eglGetPlatformDisplayEXT failed: %v", err)
			return NO_CONTEXT, NO_DISPLAY, NO_SURFACE, err
		}

		var major, minor int
		if major, minor, err = Initialize(display); err != nil {
			if !lastTry {
				continue
			}
			log.Printf("eglInitialize failed: %v", err)
			return NO_CONTEXT, NO_DISPLAY, NO_SURFACE, err
		}

		log.Printf("Version = %v.%v", major, minor)
		break
	}

	if err = BindAPI(OPENGL_ES_API); err != nil {
		return NO_CONTEXT, NO_DISPLAY, NO_SURFACE, err
	}

	var numConfigs int32
	var config Config
	config, numConfigs, err = ChooseConfig(display, rgb888)

	if numConfigs == 0 {
		log.Printf("eglChooseConfig failed: %v", GetError())
		return
	}
	if numConfigs <= 0 {
		log.Printf("eglChooseConfig found no valid config")
		return
	}

	surface, err = CreateWindowSurface(display, config, nativeWindow, nil)
	if err != nil {
		log.Printf("eglCreateWindowSurface failed: %v", err)
		return
	}

	contextAttribs := []int32{
		CONTEXT_CLIENT_VERSION, 2,
		NONE,
	}
	context, err = CreateContext(display, config, NO_CONTEXT, contextAttribs)
	if err != nil {
		log.Printf("eglCreateContext failed: %v", err)
		return
	}

	return
}
