package language

import (
	"cmp"
	"errors"
	"fmt"
	"math"
)

func init() {
	strings := []string{"Tomato", "Potato", "Srotato", "Potato"}

	filtered := Filter(strings, func(s string) bool { return s != "Potato" })
	fmt.Println(filtered)

	lengths := Map(filtered, func(s string) int { return len(s) })
	fmt.Println(lengths)

	sum := Reduce(lengths, 0, func(acc int, val int) int { return acc + val })
	fmt.Println(sum)

	// Find closer
	pair2Da := Pair[Point2D]{Point2D{1, 1}, Point2D{5, 5}}
	pair2Db := Pair[Point2D]{Point2D{10, 10}, Point2D{15, 5}}
	closer := FindCloser(pair2Da, pair2Db)
	fmt.Println(closer)

	pair3Da := Pair[Point3D]{Point3D{1, 1, 10}, Point3D{5, 5, 0}}
	pair3Db := Pair[Point3D]{Point3D{10, 10, 10}, Point3D{11, 5, 0}}
	closer2 := FindCloser(pair3Da, pair3Db)
	fmt.Println(closer2)

	var a uint = 18_243_324_324_234
	var b uint = 9_223_342_213_123
	fmt.Println(divAndRemainder2(a, b))
	type MyInt int
	var c MyInt = 14324
	var d MyInt = 35
	fmt.Println(divAndRemainderUserTypeFriendly(c, d))

	convert := Convert[int, int64](10)
	fmt.Println(convert)

	t1 := NewTree(cmp.Compare[int])
	t1.Add(10)
	t1.Add(20)
	t1.Add(30)
	fmt.Println("Tree contains", 30, ":", t1.Contains(30))
}

func Filter[T1 any](s []T1, f func(T1) bool) []T1 {
	var result []T1
	for _, v := range s {
		if f(v) {
			result = append(result, v)
		}
	}
	return result
}

func Map[T1, T2 any](s []T1, f func(T1) T2) []T2 {
	result := make([]T2, len(s))
	for i, v := range s {
		result[i] = f(v)
	}
	return result
}

func Reduce[T1, T2 any](s []T1, initializer T2, f func(T2, T1) T2) T2 {
	result := initializer
	for _, v := range s {
		result = f(result, v)
	}
	return result
}

// Differ

type Pair[T fmt.Stringer] struct {
	Val1 T
	Val2 T
}

type Differ[T any] interface {
	fmt.Stringer
	Diff(T) float64
}

func FindCloser[T Differ[T]](pair1, pair2 Pair[T]) Pair[T] {
	d1 := pair1.Val1.Diff(pair2.Val1)
	d2 := pair2.Val2.Diff(pair1.Val2)
	if d1 < d2 {
		return pair1
	} else {
		return pair2
	}
}

type Point2D struct {
	X, Y int
}

func (p2 Point2D) String() string {
	return fmt.Sprintf("{%d,%d}", p2.X, p2.Y)
}

func (p2 Point2D) Diff(from Point2D) float64 {
	x := p2.X - from.X
	y := p2.Y - from.Y
	return math.Sqrt(float64(x*x) + float64(y*y))
}

type Point3D struct {
	X, Y, Z int
}

func (p3 Point3D) String() string {
	return fmt.Sprintf("{%d,%d,%d}", p3.X, p3.Y, p3.Z)
}

func (p3 Point3D) Diff(from Point3D) float64 {
	x := p3.X - from.X
	y := p3.Y - from.Y
	z := p3.Z - from.Z
	return math.Sqrt(float64(x*x) + float64(y*y) + float64(z*z))
}

type Integer interface {
	int | int8 | int16 | int32 | int64 |
		uint | uint8 | uint16 | uint32 | uint64 | uintptr
}

func divAndRemainder2[T Integer](num, denom T) (T, T, error) {
	if denom == 0 {
		return num, num, errors.New("cannot divide by zero")
	}

	return num / denom, num % denom, nil
}

type UserTypeFriendlyInteger interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
}

func divAndRemainderUserTypeFriendly[T UserTypeFriendlyInteger](num, denom T) (T, T, error) {
	if denom == 0 {
		return num, num, errors.New("cannot divide by zero")
	}

	return num / denom, num % denom, nil
}

func Convert[T1, T2 Integer](in T1) T2 {
	return T2(in)
}

////////////////////////////////////////////////////////////////////////////

type OrderableFunc[T any] func(t1, t2 T) int

type Tree[T any] struct {
	f    OrderableFunc[T]
	root *Node[T]
}

func NewTree[T any](f OrderableFunc[T]) *Tree[T] {
	return &Tree[T]{
		f: f,
	}
}

func (t *Tree[T]) Add(val T) {
	t.root = t.root.Add(t.f, val)
}

func (t *Tree[T]) Contains(val T) bool {
	return t.root.Contains(t.f, val)
}

type Node[T any] struct {
	val         T
	left, right *Node[T]
}

func (t *Node[T]) Add(f OrderableFunc[T], val T) *Node[T] {
	if t == nil {
		return &Node[T]{val: val}
	}

	switch r := f(t.val, val); {
	case r <= -1:
		t.left = t.left.Add(f, val)
	case r >= 1:
		t.right = t.right.Add(f, val)
	}

	return t
}

func (n *Node[T]) Contains(f OrderableFunc[T], val T) bool {
	if n == nil {
		return false
	}

	switch r := f(n.val, val); {
	case r <= -1:
		return n.left.Contains(f, val)
	case r >= 1:
		return n.right.Contains(f, val)
	}

	return true
}
