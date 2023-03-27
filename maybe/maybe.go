/*
Maybe Generic with a complete monad definition and various helper functions.
*/
package maybe

import (
    "github.com/frankiethekneeman/gonads/slice"
)

type Maybe[T any] interface {
    isSome() bool
    get() T //Not Safe.  Panics on None. Intended for internal use.
}

type Some[T any] struct {
    val T
}

func (_ Some[T]) isSome() bool {
    return true
}

func (s Some[T]) get() T {
    return s.val
}

type None[T any] struct {}

func (_ None[T]) isSome() bool {
    return false
}

func (_ None[T]) get() T {
    panic("There is no value in this Maybe.  Calling Get() is an unsafe operation.")
}

// Creators

// Get a Maybe with a value in it.
func FromValue[T any](value T) Maybe[T] {
    return Some[T]{value}
}

// Get a Maybe without a value.
func FromNothing[T any]() Maybe[T] {
    return None[T]{}
}

// Functor Implementation

func Fmap[In any, Out any](f func (In) Out, m Maybe[In]) Maybe[Out] {
    if !m.isSome() {
        return None[Out]{}
    }
    return FromValue(f(m.get()))
}

// Applicative Implementation

func Pure[T any](v T) Maybe[T] {
    return FromValue(v)
}

func Fapply[In any, Out any](maybeFunc Maybe[func (In) Out], m Maybe[In]) Maybe[Out] {
    if !maybeFunc.isSome() {
        return None[Out]{}
    }
    return Fmap(maybeFunc.get(), m)
}


// Monad Implementation

// I'm not implementing `return` because it's a reserve word and for all Monads return == pure

func FlatMap[In any, Out any](f func(In) Maybe[Out], m Maybe[In]) Maybe[Out] {
    if !m.isSome() {
        return None[Out]{}
    }
    return f(m.get())
}

// Convenience functions

func IsSome[T any](m Maybe[T]) bool {
    return m.isSome()
}

func IsNone[T any](m Maybe[T]) bool {
    return !IsSome(m)
}

func ToSlice[T any] (m Maybe[T]) []T {
    if !m.isSome() {
        return []T{}
    }
    return []T{m.get()}
}

// This equates to a `safeHead` operation.  I'd prefer to call it that and put
// it in the slice module, but then it wouldn't compile because of a circular
// Dependency
func FromSlice[T any](ts []T) Maybe[T] {
    if len(ts) == 0 {
        return FromNothing[T]()
    }
    return FromValue(ts[0])
}

func Cat[T any](ms []Maybe[T]) []T {
    return slice.Flatten(slice.Fmap(ToSlice[T], ms))
}

func MapMaybe[T any, U any](f func (T) Maybe[U], ts []T) []U {
    return Cat(slice.Fmap(f, ts))
}
