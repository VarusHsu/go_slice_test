package main

import (
	"fmt"
	"os"
	"testing"
	"time"
)

func printSlice(name string, s []int) {
	fmt.Printf("len(%s)=%d cap(%s)=%d\n", name, len(s), name, cap(s))
	if len(s) == 0 {
		fmt.Println("[]")
		return
	}
	fmt.Printf("%s=[", name)
	for idx, elem := range s {
		fmt.Printf("%d", elem)
		if idx != len(s)-1 {
			fmt.Print(", ")
		} else {
			fmt.Println("]")
		}
	}
}

func recoverWithPrint() {
	err := recover()
	if err != nil {
		fmt.Println("err occur")
		fmt.Println(err)
	} else {
		fmt.Println("no err occur")
	}
}

// Test_slice01
// * 初始化切片 s 长度和容量均为 10
// * 在 s 的基础上追加 append 一个元素
// * 切片 s 的内容、长度以及容量分别是什么？
func Test_slice01(t *testing.T) {
	s := make([]int, 10)
	s = append(s, 1)
	printSlice("s", s)
}

// Test_slice02
// * 初始化切片 s 长度为 0，容量为 10
// * 在 s 的基础上追加 append 一个元素
// * 切片 s 的内容、长度以及容量分别是什么？
func Test_slice02(t *testing.T) {
	s := make([]int, 0, 10)
	s = append(s, 1)
	printSlice("s", s)
}

// Test_slice03
// * 初始化切片 s 长度为 10，容量为 11
// * 在 s 的基础上追加 append 一个元素
// * 切片 s 的内容、长度以及容量分别是什么？
func Test_slice03(t *testing.T) {
	s := make([]int, 10, 11)
	s = append(s, 1)
	printSlice("s", s)
}

// Test_slice04
// * 初始化切片 s 长度为 10，容量为 12
// * 截取切片 s index = 8 往后的内容赋给 s1
// * s1 的内容、长度以及容量分别是什么？
func Test_slice04(t *testing.T) {
	s := make([]int, 10, 12)
	s1 := s[8:]
	printSlice("s", s)
	printSlice("s1", s1)
}

// Test_slice05
// * 初始化切片 s 长度为 10，容量为 12
// * 截取切片 s index 为 [8,9) 范围内的元素赋给切片 s1
// * s1 的内容、长度以及容量分别是什么？
func Test_slice05(t *testing.T) {
	s := make([]int, 10, 12)
	s1 := s[8:9]
	printSlice("s", s)
	printSlice("s1", s1)
}

// Test_slice06
// * 初始化切片 s 长度为 10，容量为 12
// * 截取切片 s index = 8 往后的内容赋给 s1
// * 修改 s1[0] 的值
// * 这个修改是否会影响到 s ？此时，s 的内容是什么？
func Test_slice06(t *testing.T) {
	s := make([]int, 10, 12)
	s1 := s[8:9]
	s1[0] = 1
	printSlice("s", s)
	printSlice("s1", s1)
}

// Test_slice07
// * 初始化切片 s 长度为 10，容量为 12
// * 访问 s[10] 是否会越界？
func Test_slice07(t *testing.T) {
	s := make([]int, 10, 12)
	defer recoverWithPrint()
	_ = s[10]
}

// Test_slice08
// * 初始化切片 s 长度为 10，容量为 12
// * 截取 s 中 index = 8 后面的内容赋给 s1
// * 在 s1 的基础上追加 []int{10,11,12} 3 个元素
// * 经过上述操作时候，访问 s[10] 是否会越界？
func Test_slice08(t *testing.T) {
	s := make([]int, 10, 12)
	s1 := s[8:]
	s1 = append(s1, 10, 11, 12)
	defer recoverWithPrint()
	_ = s[10]
}

// Test_slice09
// * 初始化切片 s 长度为 10，容量为 12
// * 截取切片 s index = 8 往后的内容赋给 s1
// * 在方法 changeSlice 中，对 s1[0] 进行修改
// * 经过上述操作之后，s 的内容是什么？
func Test_slice09(t *testing.T) {
	s := make([]int, 10, 12)
	s1 := s[8:]
	var changeSlice = func(s []int) {
		s[0] = 1
	}
	changeSlice(s1)
	printSlice("s", s)
	printSlice("s1", s1)
}

// Test_slice10
// * 初始化切片 s 长度为 10，容量为 12
// * 截取切片 s index = 8 往后的内容赋给 s1
// * 在方法 changeSlice 中，对 s1 进行 append 追加操作
// * 经过上述操作后，s 以及 s1 的内容、长度和容量分别是什么？
func Test_slice10(t *testing.T) {
	s := make([]int, 10, 12)
	s1 := s[8:]
	var changeSlice = func(s []int) {
		s = append(s, 1)
	}
	changeSlice(s1)
	printSlice("s", s)
	printSlice("s1", s1)
}

// Test_slice11
// * 初始化切片 s，内容为 []int{0,1,2,3,4}
// * 截取 s 中 index = 2 前面的内容（不含s[2]），并在此基础上追加 index = 3 后面的内容
// * s 的内容、长度和内容分别是什么？此时访问 s[4] 是否会越界？
func Test_slice11(t *testing.T) {
	s := []int{0, 1, 2, 3, 4}
	s1 := s[0:2]
	s1 = append(s1, s[3:]...)
	printSlice("s", s)
	defer recoverWithPrint()
	_ = s[4]
}

// Test_slice12
// * 初始化切片 s 长度和容量均为 512
// * 在 s 的基础上追加 append 一个元素
// * 经过上述操作后，切片s 的内容、长度以及容量分别是什么？
func Test_slice12(t *testing.T) {
	s := make([]int, 512)
	s = append(s, 1)
	printSlice("s", s)
}

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

// Test_slice13
// * slice 的长度的上限
// make 1 len slice cost 0 us
// make 2 len slice cost 0 us
// make 4 len slice cost 0 us
// make 8 len slice cost 1 us
// make 16 len slice cost 4 us
// make 32 len slice cost 0 us
// make 64 len slice cost 0 us
// make 128 len slice cost 0 us
// make 256 len slice cost 1 us
// make 512 len slice cost 1 us
// make 1024 len slice cost 0 us
// make 2048 len slice cost 0 us
// make 4096 len slice cost 1 us
// make 8192 len slice cost 3 us
// make 16384 len slice cost 1 us
// make 32768 len slice cost 2 us
// make 65536 len slice cost 4 us
// make 131072 len slice cost 6 us
// make 262144 len slice cost 337 us
// make 524288 len slice cost 127 us
// make 1048576 len slice cost 858 us
// make 2097152 len slice cost 903 us
// make 4194304 len slice cost 2249 us
// make 8388608 len slice cost 4148 us
// make 16777216 len slice cost 7057 us
// make 33554432 len slice cost 14250 us
// make 67108864 len slice cost 28218 us
// make 134217728 len slice cost 58944 us
// make 268435456 len slice cost 186106 us
// make 536870912 len slice cost 595593 us
// make 1073741824 len slice cost 1933853 us
// make 2147483648 len slice cost 101160 us
// make 4294967296 len slice cost 6442552 us
// make 8589934592 len slice cost 15700334 us
// process kill by signal 9 SIGKILL
func Test_slice13(t *testing.T) {
	makeSlices := func(l int) {
		now := time.Now().UnixMicro()
		_ = make([]int, l)
		now2 := time.Now().UnixMicro()
		fmt.Printf("make %d len slice cost %d us\n", l, now2-now)
	}
	i := 1
	for j := 0; j < 34; j++ {
		makeSlices(i)
		i <<= 1
	}
}
