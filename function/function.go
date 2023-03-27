package function

// Functor Implementation

func Fmap[I any, O any, N any](f func (O) N, g func (I) O) func (I) N {
    return func(x I) N {
        return f(g(x))
    }
}

// Applicative Implementation

func Pure[I any, O any](x O) func(I) O {
    return func(_ I) O {
        return x
    }
}

func Fapply[I any, O any, N any](f func(I) func (O) N, g func(I) O) func(I) N {
    return func(x I) N {
        return f(x)(g(x))
    }
}


// Monad Implementation

// I'm not implementing `return` because it's a reserve word and for all Monads return == pure

func FlatMap[I any, O any, N any](f func(O) func (I) N, g func(I) O) func(I) N {
    return func(x I) N {
        return f(g(x))(x)
    }
}

// Helpful Function functions

func Compose[A any, B any, C any](f func(B) C, g func (A) B) func (A) C {
    return Fmap(f, g)
}

func Curry[A any, B any, C any](f func(A, B) C) func (A) func (B) C {
    return func (a A) func (B) C {
        return func (b B) C {
            return f(a, b)
        }
    }
}

func Curry3[A any, B any, C any, D any](f func(A, B, C) D ) func (A) func (B) func (C) D {
    return func (a A) func (B) func (C) D {
        return func (b B) func (C) D {
            return func(c C) D {
                return f(a, b, c)
            }
        }
    }
}

//TODO Code gen a _bunch_ of Curry functions.

func Uncurry[A any, B any, C any](f func(A) func (B) C) func (A, B) C {
    return func (a A, b B) C {
        return f(a)(b)
    }
}


func Uncurry3[A any, B any, C any, D any](f func(A) func (B) func(C) D) func (A, B, C) D {
    return func (a A, b B, c C) D {
        return f(a)(b)(c)
    }
}

//TODO Code gen some uncurry functions.

func Swap[A any, B any, C any](f func(A) func(B) C) func(B) func(A) C {
    return func (b B) func (A) C {
        return func (a A) C {
            return f(a)(b)
        }
    }
}
