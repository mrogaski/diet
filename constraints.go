package diet

// Integer copies contraints.Integer from golang.org/x/exp.
type Integer interface {
	Signed | Unsigned
}

// Signed copies contraints.Signed from golang.org/x/exp.
type Signed interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

// Unsigned copies contraints.Unsigned from golang.org/x/exp.
type Unsigned interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}
