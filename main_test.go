package main

import "testing"

func Benchmark_Read(b *testing.B) {
	for i := 0; i < b.N; i++ {
		read()
	}
}

func Benchmark_ReadAll(b *testing.B) {
	for i := 0; i < b.N; i++ {
		readAll()
	}
}
