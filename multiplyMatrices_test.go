package main

import (
	"math/rand"
	"testing"
	"time"
)

func makeRandomMatrix() [][]float32 {
	rand.Seed(time.Now().UnixNano())
	res := make([][]float32, 1024*2)

	for i := range res {
		res[i] = make([]float32, 1024*2)
		for j := range res[i] {
			res[i][j] = randFloat32(-1024, 1024)
		}
	}

	return res
}

func getTwoMatrices() ([][]float32, [][]float32) {
	return makeRandomMatrix(), makeRandomMatrix()
}

func randFloat32(min, max int) float32 {
	return rand.Float32() * float32(rand.Intn(max-min)+min)
}

func BenchmarkMultiThreaded(b *testing.B) {
	A, B := getTwoMatrices()

	multiplyMatricesMultithreaded(A, B)
}

func BenchmarkSingleThreaded(b *testing.B) {
	A, B := getTwoMatrices()

	multiplyMatrices(A, B)
}
