package goniq

// Ordered type constraint
// Inspired by: https://cs.opensource.google/go/x/exp/+/master:constraints/constraints.go (MIT License)
type Ordered interface {
	Float | Int | ~string
}

type Float interface {
	~float32 | ~float64
}

type Int interface {
	Signed | Unsigned
}

type Signed interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

type Unsigned interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
}
