package slice

// Functor Implementation

func Fmap[T any, U any](f func (T) U, ts []T) []U {
    toReturn := make([]U, len(ts))
    for i, t := range ts {
        toReturn[i] = f(t)
    }
    return toReturn
}

// Applicative Implementation

func Pure[T any](x T) []T {
    return []T{x}
}

func Fapply[T any, U any](fs []func(T) U, ts[]T) []U {
    tLen := len(ts)
    toReturn := make([]U, tLen * len(fs))
    for i, f := range fs {
        for j, t := range ts {
            toReturn[tLen * i + j] = f(t)
        }
    }
    return toReturn
}


// Monad Implementation

// I'm not implementing `return` because it's a reserve word and for all Monads return == pure

func FlatMap[T any, U any](f func(T) []U, ts []T) []U {
    return Flatten(Fmap(f, ts))
}

// Convenience Functions

func Flatten[T any](doubled [][]T) []T {
    size := FoldL(func(l int, r int) int { return l + r }, 0, Fmap(Length[T], doubled))
    toReturn := make([]T, size)
    next := 0
    for _, ts := range doubled {
        for _, t := range ts {
            toReturn[next] = t
            next++
        }
    }
    return toReturn
}

func FoldL[T any, U any](f func(U, T) U, init U, ts []T) U {
    result := init
    for _, t := range ts {
        result = f(result, t)
    }
    return result
}

//... Go demands that `len` be invoked, so I cant have a function reference.
func Length[T any](ts []T) int {
    return len(ts)
}
