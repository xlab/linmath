package linmath

import (
	"bytes"
	"fmt"
	"unsafe"
)

func DumpMatrix(m *Mat4x4, note string) string {
	buf := new(bytes.Buffer)
	fmt.Fprintf(buf, "[mat4x4] %s: \n", note)
	for i := 0; i < 4; i++ {
		fmt.Fprintf(buf, "%.3f, %.3f, %.3f, %.3f\n", m[i][0], m[i][1], m[i][2], m[i][3])
	}
	return buf.String()
}

func DumpVec4(v *Vec4, note string) string {
	buf := new(bytes.Buffer)
	fmt.Fprintf(buf, "[vec4] %s: \n", note)
	fmt.Fprintf(buf, "%.3f, %.3f, %.3f, %.3f\n", v[0], v[1], v[2], v[3])
	return buf.String()
}

func (m *Mat4x4) Sizeof() int {
	return 4 * 4 * 4
}

func (m *Mat4x4) Slice() []float32 {
	hdr := &sliceHeader{
		Len: 4 * 4, Cap: 4 * 4,
		Data: uintptr(unsafe.Pointer(m)),
	}
	return *(*[]float32)(unsafe.Pointer(hdr))
}

func (m *Mat4x4) CopyTo(buf *[4][4]float32) {
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			buf[i][j] = m[i][j]
		}
	}
}

func (v *Vec4) Sizeof() int {
	return 4 * 4 * 4
}

func (v *Vec4) Slice() []float32 {
	hdr := &sliceHeader{
		Len: 4, Cap: 4,
		Data: uintptr(unsafe.Pointer(v)),
	}
	return *(*[]float32)(unsafe.Pointer(hdr))
}

func (v *Vec4) CopyTo(buf *[4]float32) {
	for i := 0; i < 4; i++ {
		buf[i] = v[i]
	}
}

type sliceHeader struct {
	Data uintptr
	Len  int
	Cap  int
}
