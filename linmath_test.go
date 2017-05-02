package linmath

import (
	"log"
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	native "github.com/xlab/linmath-go/linmath"
)

const seed = 777

func TestMat4x4FromQuat(t *testing.T) {
	assert := assert.New(t)

	for _, test := range []string{"0", "1", "N", "R"} {
		var expected string
		{
			m := n_m(test, seed)
			q := n_q(test, seed)
			native.Mat4x4FromQuat(m, q)
			expected = native.DumpMatrix(m, test)
		}

		var actual string
		{
			m := m(test, seed)
			q := q(test, seed)
			m.FromQuat(q)
			actual = DumpMatrix(m, test)
		}
		assert.Equal(expected, actual)
		log.Println(test, expected)
	}
}

func TestQuatFromMat4x4(t *testing.T) {
	assert := assert.New(t)

	for _, test := range []string{"0", "1", "N", "R"} {
		var expected string
		{
			m := n_m(test, seed)
			q := n_q(test, seed)
			native.QuatFromMat4x4(q, m)
			expected = native.DumpQuat(q, test)
		}

		var actual string
		{
			m := m(test, seed)
			q := q(test, seed)
			q.FromMat4x4(m)
			actual = DumpQuat(q, test)
		}
		assert.Equal(expected, actual)
		log.Println(test, expected)
	}
}

func TestMat4x4ScaleAniso(t *testing.T) {
	assert := assert.New(t)

	for _, test := range []string{"0", "1", "N", "R"} {
		var expected string
		{
			m1 := n_m(test, seed)
			m2 := n_m(test, seed)
			native.Mat4x4ScaleAniso(m1, m2, 1, 2, 3)
			expected = native.DumpMatrix(m1, test)
		}

		var actual string
		{
			m1 := m(test, seed)
			m2 := m(test, seed)
			m1.ScaleAniso(m2, 1, 2, 3)
			actual = DumpMatrix(m1, test)
		}
		assert.Equal(expected, actual)
		log.Println(test, expected)
	}
}

func TestMat4x4LookAt(t *testing.T) {
	assert := assert.New(t)

	for _, test := range []string{"0", "1", "N", "R"} {
		var expected string
		{
			m := n_m(test, seed)
			v1 := n_v3(test, seed)
			v2 := n_v3(test, seed+1)
			v3 := n_v3(test, seed+2)
			native.Mat4x4LookAt(m, v1, v2, v3)
			expected = native.DumpMatrix(m, test)
		}

		var actual string
		{
			m := m(test, seed)
			v1 := v3(test, seed)
			v2 := v3(test, seed+1)
			v3 := v3(test, seed+2)
			m.LookAt(v1, v2, v3)
			actual = DumpMatrix(m, test)
		}
		assert.Equal(expected, actual)
		log.Println(test, expected)
	}
}

func q(test string, seed ...int64) *Quat {
	rand := getRand(seed)

	switch test {
	case "0":
		q := new(Quat)
		for i := 0; i < 4; i++ {
			q[i] = 0
		}
		return q
	case "1":
		q := new(Quat)
		for i := 0; i < 4; i++ {
			q[i] = 1
		}
		return q
	case "N":
		q := new(Quat)
		n := 1
		for i := 0; i < 4; i++ {
			q[i] = float32(n)
			n++
		}
		return q
	case "R":
		q := new(Quat)
		for i := 0; i < 4; i++ {
			q[i] = float32(rand.Intn(100))
		}
		return q
	default:
		panic(test)
	}
}

func v3(test string, seed ...int64) *Vec3 {
	rand := getRand(seed)

	switch test {
	case "0":
		q := new(Vec3)
		for i := 0; i < 3; i++ {
			q[i] = 0
		}
		return q
	case "1":
		q := new(Vec3)
		for i := 0; i < 3; i++ {
			q[i] = 1
		}
		return q
	case "N":
		q := new(Vec3)
		n := 1
		for i := 0; i < 3; i++ {
			q[i] = float32(n)
			n++
		}
		return q
	case "R":
		q := new(Vec3)
		for i := 0; i < 3; i++ {
			q[i] = float32(rand.Intn(100))
		}
		return q
	default:
		panic(test)
	}
}

func v4(test string, seed ...int64) *Vec4 {
	rand := getRand(seed)

	switch test {
	case "0":
		q := new(Vec4)
		for i := 0; i < 4; i++ {
			q[i] = 0
		}
		return q
	case "1":
		q := new(Vec4)
		for i := 0; i < 4; i++ {
			q[i] = 1
		}
		return q
	case "N":
		q := new(Vec4)
		n := 1
		for i := 0; i < 4; i++ {
			q[i] = float32(n)
			n++
		}
		return q
	case "R":
		q := new(Vec4)
		for i := 0; i < 4; i++ {
			q[i] = float32(rand.Intn(100))
		}
		return q
	default:
		panic(test)
	}
}

func m(test string, seed ...int64) *Mat4x4 {
	rand := getRand(seed)

	switch test {
	case "0":
		m := new(Mat4x4)
		for i := 0; i < 4; i++ {
			for j := 0; j < 4; j++ {
				m[i][j] = 0
			}
		}
		return m
	case "1":
		m := new(Mat4x4)
		for i := 0; i < 4; i++ {
			for j := 0; j < 4; j++ {
				m[i][j] = 1
			}
		}
		return m
	case "N":
		m := new(Mat4x4)
		n := 1
		for i := 0; i < 4; i++ {
			for j := 0; j < 4; j++ {
				m[i][j] = float32(n)
				n++
			}
		}
		return m
	case "R":
		m := new(Mat4x4)
		for i := 0; i < 4; i++ {
			for j := 0; j < 4; j++ {
				m[i][j] = float32(rand.Intn(100))
			}
		}
		return m
	default:
		panic(test)
	}
}

func n_q(test string, seed ...int64) *native.Quat {
	rand := getRand(seed)

	switch test {
	case "0":
		q := new(native.Quat)
		for i := 0; i < 4; i++ {
			q[i] = 0
		}
		return q
	case "1":
		q := new(native.Quat)
		for i := 0; i < 4; i++ {
			q[i] = 1
		}
		return q
	case "N":
		q := new(native.Quat)
		n := 1
		for i := 0; i < 4; i++ {
			q[i] = float32(n)
			n++
		}
		return q
	case "R":
		q := new(native.Quat)
		for i := 0; i < 4; i++ {
			q[i] = float32(rand.Intn(100))
		}
		return q
	default:
		panic(test)
	}
}

func n_v3(test string, seed ...int64) *native.Vec3 {
	rand := getRand(seed)

	switch test {
	case "0":
		q := new(native.Vec3)
		for i := 0; i < 3; i++ {
			q[i] = 0
		}
		return q
	case "1":
		q := new(native.Vec3)
		for i := 0; i < 3; i++ {
			q[i] = 1
		}
		return q
	case "N":
		q := new(native.Vec3)
		n := 1
		for i := 0; i < 3; i++ {
			q[i] = float32(n)
			n++
		}
		return q
	case "R":
		q := new(native.Vec3)
		for i := 0; i < 3; i++ {
			q[i] = float32(rand.Intn(100))
		}
		return q
	default:
		panic(test)
	}
}

func n_v4(test string, seed ...int64) *native.Vec4 {
	rand := getRand(seed)

	switch test {
	case "0":
		q := new(native.Vec4)
		for i := 0; i < 4; i++ {
			q[i] = 0
		}
		return q
	case "1":
		q := new(native.Vec4)
		for i := 0; i < 4; i++ {
			q[i] = 1
		}
		return q
	case "N":
		q := new(native.Vec4)
		n := 1
		for i := 0; i < 4; i++ {
			q[i] = float32(n)
			n++
		}
		return q
	case "R":
		q := new(native.Vec4)
		for i := 0; i < 4; i++ {
			q[i] = float32(rand.Intn(100))
		}
		return q
	default:
		panic(test)
	}
}

func n_m(test string, seed ...int64) *native.Mat4x4 {
	rand := getRand(seed)

	switch test {
	case "0":
		m := new(native.Mat4x4)
		for i := 0; i < 4; i++ {
			for j := 0; j < 4; j++ {
				m[i][j] = 0
			}
		}
		return m
	case "1":
		m := new(native.Mat4x4)
		for i := 0; i < 4; i++ {
			for j := 0; j < 4; j++ {
				m[i][j] = 1
			}
		}
		return m
	case "N":
		m := new(native.Mat4x4)
		n := 1
		for i := 0; i < 4; i++ {
			for j := 0; j < 4; j++ {
				m[i][j] = float32(n)
				n++
			}
		}
		return m
	case "R":
		m := new(native.Mat4x4)
		for i := 0; i < 4; i++ {
			for j := 0; j < 4; j++ {
				m[i][j] = float32(rand.Intn(100))
			}
		}
		return m
	default:
		panic(test)
	}
}

func getRand(seed []int64) *rand.Rand {
	if len(seed) > 0 {
		return rand.New(rand.NewSource(seed[0]))
	}
	return rand.New(rand.NewSource(time.Now().UnixNano()))
}
