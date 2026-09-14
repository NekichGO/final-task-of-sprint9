package main

import (
	"fmt"
	"math/rand/v2"
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
		return nil
	}
	nums := make([]int, 0, size)
	for i := 0; i < size; i++ {
		randomNum := rand.Int()
		nums = append(nums, randomNum)
	}
	return nums
}

// maximum returns the maximum number of elements.
func maximum(data []int) int {
	// ваш код здесь
	if len(data) == 0 {
		return 0
	}
	maxNum := data[0]
	for _, v := range data {
		if v > maxNum {
			maxNum = v
		}
	}
	return maxNum
}

// maxChunks returns the maximum number of elements in a chunks.
func maxChunks(data []int) int {
	// ваш код здесь
	var wg sync.WaitGroup

	if len(data) <= 1 {
		return 0
	}
	// Вычисляем размер среза
	lenSlice := len(data) / CHUNKS
	result := make([]int, CHUNKS)
	for i := 0; i < CHUNKS; i++ {
		start := i * lenSlice
		end := start + lenSlice
		if end > len(data) {
			end = len(data)
		}
		wg.Add(1)
		go func(data []int) {
			defer wg.Done()
			maxNum := maximum(data)
			result[i] = maxNum
		}(data[start:end])
	}
	wg.Wait()
	return maximum(result)

}

func main() {
	fmt.Printf("Генерируем %d целых чисел", SIZE)
	// ваш код здесь
	nums := generateRandomElements(SIZE)

	fmt.Println("Ищем максимальное значение в один поток")
	// ваш код здесь
	start := time.Now()
	maxNumOneStream := maximum(nums)
	elapsed := time.Since(start)

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", maxNumOneStream, elapsed)

	fmt.Printf("Ищем максимальное значение в %d потоков", CHUNKS)
	// ваш код здесь
	startChunks := time.Now()
	maxNumChunks := maxChunks(nums)
	elapsedChunks := time.Since(startChunks)

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", maxNumChunks, elapsedChunks)
}
