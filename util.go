package linmath

import (
	"bytes"
	"fmt"
	"unsafe"
)

func (m *Mat4x4) Data() []byte {
	const mm = 0x7fffffff
	return (*[mm]byte)(unsafe.Pointer(m))[:SizeofMat4x4]
}

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

const (
	SizeofMat4x4 = 4 * 4 * 4
	SizeofVec4   = 4 * 4
	SizeofVec3   = 3 * 4
	SizeofVec2   = 2 * 4
)

type ArrayVec4 []Vec4

func (a ArrayVec4) Sizeof() int {
	return len(a) * SizeofVec4
}

func (a ArrayVec4) Data() []byte {
	const m = 0x7fffffff
	return (*[m]byte)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&a)).Data))[:len(a)*SizeofVec4]
}

type ArrayVec3 []Vec3

func (a ArrayVec3) Sizeof() int {
	return len(a) * SizeofVec3
}

func (a ArrayVec3) Data() []byte {
	const m = 0x7fffffff
	return (*[m]byte)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&a)).Data))[:len(a)*SizeofVec3]
}

type ArrayVec2 []Vec2

func (a ArrayVec2) Sizeof() int {
	return len(a) * SizeofVec2
}

func (a ArrayVec2) Data() []byte {
	const m = 0x7fffffff
	return (*[m]byte)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&a)).Data))[:len(a)*SizeofVec2]
}

type ArrayUint16 []uint16

func (a ArrayUint16) Sizeof() int {
	return len(a) * 2
}

func (a ArrayUint16) Data() []byte {
	const m = 0x7fffffff
	return (*[m]byte)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&a)).Data))[:len(a)*2]
}

type ArrayFloat32 []float32

func (a ArrayFloat32) Sizeof() int {
	return len(a) * 4
}

func (a ArrayFloat32) Data() []byte {
	const m = 0x7fffffff
	return (*[m]byte)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&a)).Data))[:len(a)*4]
}

type sliceHeader struct {
	Data uintptr
	Len  int
	Cap  int
}
