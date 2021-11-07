package main

func main() {

}

func multiplyMatricesMultithreaded(a [][]float32, b [][]float32) [][]float32 {
	channel := make(chan []float32)
	lenAx := len(a)
	lenAy := len(a[0])
	lenBy := len(b[0])

	result := make([][]float32, lenAx)

	for i := 0; i < lenAx; i++ {
		result[i] = make([]float32, lenBy)
	}

	for i := 0; i < lenAx; i++ {

		go func(x int) {
			for j := 0; j < lenBy; j++ {

				for k := 0; k < lenAy; k++ {
					result[x][j] += a[x][k] * b[k][j]
				}

			}
			channel <- result[x]
		}(i)
	}

	for i := 0; i < lenAx; i++ {
		<-channel
	}

	return result
}

func multiplyMatrices(a [][]float32, b [][]float32) [][]float32 {
	lenAx := len(a)
	lenAy := len(a[0])
	lenBy := len(b[0])

	result := make([][]float32, lenAx)

	for i := 0; i < lenAx; i++ {
		result[i] = make([]float32, lenBy)
	}

	for i := 0; i < lenAx; i++ {

		func(x int) {
			for j := 0; j < lenBy; j++ {

				for k := 0; k < lenAy; k++ {
					result[x][j] += a[x][k] * b[k][j]
				}

			}
		}(i)
	}

	return result
}
