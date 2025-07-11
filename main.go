package main

import (
	"errors"
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
		data[i] = rand.Intn(1_000_000) + 1
	}
	return data
}

// maximum returns the maximum number of elements.
func maximum(data []int) (int, error) {
	// ваш код здесь
	if len(data) == 0 {
		return 0, errors.New("data is empty")
	}
	max := data[0]
	for _, v := range data[1:] {
		if v > max {
			max = v
		}
	}
	return max, nil
}

// maxChunks returns the maximum number of elements in a chunks.
func maxChunks(data []int) (int, error) {
	// ваш код здесь
	if len(data) == 0 {
		return 0, errors.New("data is empty")
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
		wg.Add(1)
		go func(idx, a, b int) {
			defer wg.Done()
			if a >= b {
				maxs[idx] = 0
				return
			}
			max := data[a]
			for _, v := range data[a+1 : b] {
				if v > max {
					max = v
				}
			}
			maxs[idx] = max
		}(i, start, end)
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
	max, err := maximum(data)
	elapsed := time.Since(t1).Microseconds()
	if err != nil {
		fmt.Println("Error: ", err.Error())
	}

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)

	fmt.Printf("Ищем максимальное значение в %d потоков\n", CHUNKS)
	// ваш код здесь
	t2 := time.Now()
	max, err = maxChunks(data)
	elapsed = time.Since(t2).Microseconds()
	if err != nil {
		fmt.Println("Error: ", err.Error())
	}

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)
}
