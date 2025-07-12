package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGenerateRandomElements(t *testing.T) {
	tests := []struct {
		name    string
		size    int
		wantLen int
	}{
		{"Zero size", 0, 0},
		{"Negative size", -5, 0},
		{"Normal size", 30, 30},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			arr := generateRandomElements(tc.size)
			assert.NotNil(t, arr)
			assert.Equal(t, tc.wantLen, len(arr))
			for _, v := range arr {
				assert.GreaterOrEqual(t, v, 1)
			}
		})
	}

}

func TestMaximum(t *testing.T) {
	tests := []struct {
		name  string
		slice []int
		want  int
	}{
		{"Empty slice", []int{}, 0},
		{"One element", []int{42}, 42},
		{"All equal", []int{5, 5, 5, 5}, 5},
		{"Max in middle", []int{2, 9, 7, 5}, 9},
		{"Max negative", []int{-5, -3, -8, -1}, -1},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := maximum(tc.slice)
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestMaxChunks(t *testing.T) {
	tests := []struct {
		name  string
		slice []int
		want  int
	}{
		{"Empty", []int{}, 0},
		{"Short", []int{1, 99, 3}, 99},
		{"One", []int{7}, 7},
		{"Chunked", []int{1, 5, 11, 4, 3, 99, 8, 33, 4, 2, 70, 50, 20, 19, 14, 88}, 99},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := maxChunks(tc.slice)
			assert.Equal(t, tc.want, got)
		})
	}
}
