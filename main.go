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
	// ваш код здесь
	if size <= 0 {
		return []int{}
	}
	data := make([]int, size)
	for i := range data {
		data[i] = rand.Int() + 1
	}
	return data
}

// maximum returns the maximum number of elements.
func maximum(data []int) int {
	// ваш код здесь
	if len(data) == 0 {
		return 0
	}
	max := data[0]
	for _, v := range data[1:] {
		if v > max {
			max = v
		}
	}
	return max
}

// maxChunks returns the maximum number of elements in a chunks.
func maxChunks(data []int) int {
	// ваш код здесь
	if len(data) == 0 {
		return 0
	}
	chunkSize := len(data) / CHUNKS
	if chunkSize == 0 {
		return maximum(data)
	}
	maxs := make([]int, CHUNKS)
	wg := sync.WaitGroup{}
	for i := 0; i < CHUNKS; i++ {
		start := i * chunkSize
		end := start + chunkSize
		if i == CHUNKS-1 {
			end = len(data)
		}
		chunk := data[start:end]
		wg.Add(1)
		go func(idx int, chunk []int) {
			defer wg.Done()
			maxs[idx] = maximum(chunk)
		}(i, chunk)
	}
	wg.Wait()
	return maximum(maxs)
}

func main() {
	fmt.Printf("Генерируем %d целых чисел\n", SIZE)
	// ваш код здесь
	rand.Seed(time.Now().UnixNano())
	data := generateRandomElements(SIZE)

	fmt.Println("Ищем максимальное значение в один поток")
	// ваш код здесь
	t1 := time.Now()
	max := maximum(data)
	elapsed := time.Since(t1).Microseconds()

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)

	fmt.Printf("Ищем максимальное значение в %d потоков\n", CHUNKS)
	// ваш код здесь
	t2 := time.Now()
	max = maxChunks(data)
	elapsed = time.Since(t2).Microseconds()
	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)
}
