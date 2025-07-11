package main

import "testing"

func TestGenerateRandomElements(t *testing.T) {
	//empty slice
	arr := generateRandomElements(0)
	if arr == nil {
		t.Error("nil array")
	}
	if len(arr) != 0 {
		t.Errorf("expected length 0, got %d", len(arr))
	}

	//negative size
	arr = generateRandomElements(-111)
	if len(arr) != 0 {
		t.Errorf("expected length 0 for negative size, got %d", len(arr))
	}

	//casual size
	size := 100
	arr = generateRandomElements(size)
	if len(arr) != size {
		t.Errorf("expected length %d, got %d", size, len(arr))
	}
	for i, v := range arr {
		if v < 1 || v > 1_000_000 {
			t.Errorf("element %d out of bounds: got %d", i, v)
		}
	}
}

func TestMaximum(t *testing.T) {
	//empty slice
	_, err := maximum([]int{})
	if err == nil {
		t.Error("expected error for empty data")
	}

	//one element
	v, err := maximum([]int{123})
	if err != nil {
		t.Errorf("did not expect error for one element: %v", err)
	}
	if v != 123 {
		t.Errorf("expected 42, got %d", v)
	}

	//few elements
	v, err = maximum([]int{1, 2, 13, 4, 5})
	if err != nil {
		t.Errorf("did not expect error: %v", err)
	}
	if v != 13 {
		t.Errorf("expected 13, got %d", v)
	}

	//max is first element
	v, err = maximum([]int{1000, 2, 3, 4, 5})
	if v != 1000 {
		t.Errorf("expected 1000, got %d", v)
	}

	//max is in the end
	v, err = maximum([]int{1, 2, 3, 4, 500})
	if v != 500 {
		t.Errorf("expected 500, got %d", v)
	}
}
