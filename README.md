# Go-nads: Monads In Go!

A collection of some monadic behaviours (and some miscellaneous FP convenience functions)
written for Go!

## Existing Monads

### Slice

The slice monad looks a lot like the List monad!  Enjoy mapping and folding.

### Function

The function monad takes some work in Go, but you can compose, curry, and uncurry to your hearts' content.

## New (to Go) Monads

### Maybe

Is there a value here?  Maybe, maybe not!  Great for calculating defaults.

### Try

Monadically switch on success vs failure.  Enumerate long chains of failable operations without
having to constantly stop and check `if err != nil`.
