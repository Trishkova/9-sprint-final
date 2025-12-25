package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	SIZE   = 100_000_000
	CHUNKS = 8
)

// generateRandomElements generates random elements.
func generateRandomElements(size int) []int {
	if size <= 0 {
		return []int{}
	}

	rand.Seed(time.Now().UnixNano())

	data := make([]int, size)

	for i := 0; i < size; i++ {
		data[i] = rand.Intn(SIZE)
	}

	return data
}

// maximum returns the maximum number of elements.
func maximum(data []int) int {
	if len(data) == 0 {
		return 0
	}

	maxVal := data[0]

	for i := 1; i < len(data); i++ {
		if data[i] > maxVal {
			maxVal = data[i]
		}
	}

	return maxVal
}

// maxChunks returns the maximum number of elements in a chunks.
func maxChunks(data []int) int {
	size := len(data)
	if size == 0 {
		return 0
	}

	chunkMaxes := make([]int, CHUNKS)

	var wg sync.WaitGroup

	chunkSize := size / CHUNKS

	for i := 0; i < CHUNKS; i++ {
		start := i * chunkSize
		end := start + chunkSize

		if i == CHUNKS-1 {
			end = size
		}

		wg.Add(1)

		go func(index int, part []int) {
			defer wg.Done()

			if len(part) > 0 {
				chunkMaxes[index] = maximum(part)
			}
		}(i, data[start:end])
	}

	wg.Wait()

	return maximum(chunkMaxes)
}

func main() {
	fmt.Printf("Генерируем %d целых чисел\n", SIZE)
	array := generateRandomElements(SIZE)

	fmt.Println("Ищем максимальное значение в один поток")
	now := time.Now()
	max := maximum(array)
	elapsed := time.Since(now)

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)

	fmt.Printf("Ищем максимальное значение в %d потоков\n", CHUNKS)
	now = time.Now()
	max = maxChunks(array)
	elapsed = time.Since(now)

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)
}
