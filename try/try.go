package try

type Try[T any] interface {
    isSuccess() bool
    get() T //Not Safe.  Panics on Failure. Intended for internal use.
    getErr() error //Not Safe.  Panics on Sucess.  Intended for internal use.
}

type Success[T any] struct {
    result T
}

func (_ Success[T]) isSuccess() bool {
    return true
}

func (s Success[T]) get() T {
    return s.result
}

func (_ Success[T]) getErr() error {
    panic("This was successful, there is no Error to surface.  Calling getError() is an unsafe operation")
}

type Failure[T any] struct {
    err error
}

func (_ Failure[T]) isSuccess() bool {
    return false
}

func (_ Failure[T]) get() T {
    panic("There is no result in this Try.  Calling get() is an unsafe operation.")
}

func (f Failure[T]) getErr() error {
    return f.err
}

// Creators

// Get a Try with a value in it.
func FromResult[T any](result T) Try[T] {
    return Success[T]{result}
}

// Get a Try without a value.
func FromError[T any](e error) Try[T] {
    return Failure[T]{e}
}

func FromReturn[T any](result T, err error) Try[T] {
    if err != nil {
        return FromError[T](err)
    }
    return FromResult(result)
}

// Functor Implementation

func Fmap[In any, Out any](f func (In) Out, t Try[In]) Try[Out] {
    if !t.isSuccess() {
        return Failure[Out]{t.getErr()}
    }
    return FromResult(f(t.get()))
}

// Applicative Implementation

func Pure[T any](result T) Try[T] {
    return FromResult(result)
}

func Fapply[In any, Out any](tryFunc Try[func (In) Out], t Try[In]) Try[Out] {
    if !tryFunc.isSuccess() {
        return Failure[Out]{t.getErr()}
    }
    return Fmap(tryFunc.get(), t)
}


// Monad Implementation

// I'm not implementing `return` because it's a reserve word and for all Monads return == pure

func FlatMap[In any, Out any](f func(In) Try[Out], t Try[In]) Try[Out] {
    if !t.isSuccess() {
        return Failure[Out]{t.getErr()}
    }
    return f(t.get())
}

// Convenience functions

func Wrap[In any, Out any](f func(In) (Out, error)) func(In) Try[Out] {
    // Ideally, this would work like:
    // return function.Compose(FromReturn[Out], f)
    // But golang does some special magic around the dual returns that make it not work.
    return func(i In) Try[Out] {
        return FromReturn(f(i))
    }
}

func Extract[T any](t Try[T]) (T, error) {
    if t.isSuccess() {
        return t.get(), nil
    }
    var noop T
    return noop, t.getErr()
}
