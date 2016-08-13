linmath [![GoDoc](https://godoc.org/github.com/xlab/linmath?status.svg)](https://godoc.org/github.com/xlab/linmath)
-------

A small library for linear math as required for computer graphics.
This is a Go-lang port of [github.com/datenwolf/linmath.h](https://github.com/datenwolf/linmath.h).

```
linmath.h provides the most used types required programming computer graphice:

vec3 -- 3 element vector of floats
vec4 -- 4 element vector of floats (4th component used for homogenous computations)
mat4x4 -- 4 by 4 elements matrix, computations are done in column major order
quat -- quaternion

The types are deliberately named like the types in GLSL. In fact they are meant to
be used for the client side computations and passing to same typed GLSL uniforms.
```

## License

MIT (re-licensed from the WTFPL).
