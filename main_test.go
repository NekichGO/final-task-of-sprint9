package main

// Пишите тесты в этом файле
import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateRandomElements(t *testing.T) {
	var valuesList = []struct {
		size, length int
	}{
		{0, 0},
		{1, 1},
		{1000, 1000},
		{100000, 100000},
	}
	for _, v := range valuesList {
		result := generateRandomElements(v.size)
		assert.Equal(t, v.length, len(result))
	}

}

func TestMaxElement(t *testing.T) {
	list := []struct {
		data   []int
		result int
	}{
		{
			data:   []int{},
			result: 0,
		},
		{
			data:   []int{55},
			result: 55,
		},
		{
			data:   []int{34, 123, 234, 652},
			result: 652,
		},
		{
			data:   []int{959, 123, 754, 23},
			result: 959,
		},
		{
			data:   []int{-6, -3, -10},
			result: -3,
		},
		{
			data:   []int{959, 123, 754, 23, 959},
			result: 959,
		},
		{
			data:   nil,
			result: 0,
		},
		{
			data:   []int{55, 100, 200, 900, 500, 600, 700},
			result: 900,
		},
		{
			data:   []int{-5, -2, 0, -6},
			result: 0,
		},
	}
	for _, v := range list {
		maxNumber := maximum(v.data)
		assert.Equal(t, v.result, maxNumber)
	}
}
