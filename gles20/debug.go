// +build !release

package gles20

func Check() {
	err := GetError()
	if err != nil {
		panic(err)
	}
}
