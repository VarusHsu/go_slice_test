package main

import "unsafe"

/*
 #include "malloc.h"
*/
import "C"

type Slice[T any] struct {
	len   int
	cap   int
	array unsafe.Pointer
}

func (s *Slice[T]) Append(elements ...T) Slice[T] {
	return Slice[T]{}
}

func Make[T any](params ...int) Slice[T] {
	if len(params) == 0 {
		panic("not enough params to make slice")
	}

	if len(params) >= 3 {
		panic("too many params to make slice")
	}

	if params[0] < 0 {
		panic("slice length can't less than zero")
	}
	if len(params) == 2 && params[1] < params[0] {
		panic("slice length can't less than capacity")
	}

	s := Slice[T]{
		len: params[0],
	}

	if len(params) == 2 {
		s.cap = params[1]
	} else {
		s.cap = params[0]
	}
	var t T
	// unsafe.Sizeof(t)
	s.array = C.c_malloc(C.int(int(unsafe.Sizeof(t)) * s.len))
	return s
}

func (s *Slice[T]) SubSlice() Slice[T] {
	return Slice[T]{}
}

func (s *Slice[T]) Get(index int) (T, error) {
	var t T
	return t, nil
}

func (s *Slice[T]) Set(index int, element T) error {

	return nil
}

func (s *Slice[T]) Range(func(index int, element T)) {
	if s.len == 0 {
		return
	}
	for i := 0; i < s.len; i++ {

	}
}
