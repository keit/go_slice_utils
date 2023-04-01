package slice_utils

import (
	"testing"
)

// func TestMain(m *testing.M) {
// 	// call flag.Parse() here if TestMain uses flags
// 	os.Exit(m.Run())
// }
func TestContains(t *testing.T) {
	slice := []string{"kei", "yoko", "blues", "jake"}
    result := Contains(slice, "jake")
    if !result {
        t.Errorf("should return true for jake with %s", slice)
    }
    result = Contains(slice, "taro")
    if result {
        t.Errorf("should return false for taro with %s", slice)
    }

}

func TestReverse(t *testing.T) {
	slice := []string{"kei", "yoko", "blues", "jake"}
	result := Reverse(slice)
    expected := []string{"jake", "blues", "yoko", "kei"}
	if !CompareStringSlices(result, expected) {
		 t.Errorf("should return %s but got %s", expected, result)
	}

}

func TestCompareStringSlices(t *testing.T) {
    slice1 := []string{"kei", "yoko", "blues", "jake"}
    slice2 := []string{"kei", "yoko", "blues", "jake"}
    slice3 := []string{"kei", "yoko", "blues"}

    if !CompareStringSlices(slice1, slice2) {
        t.Errorf("should return true for slice1: %s and slice2: %s", slice1, slice2)
    }

    if CompareStringSlices(slice1, slice3) {
        t.Errorf("should return false for slice1: %s and slice3: %s", slice1, slice3)
    }
}

func TestMap(t *testing.T) {
    slice := []string{"kei", "yoko", "blues", "jake"}
    expected := []int{3, 4, 5, 4}

    result := Map(slice, func(s string) int {
        return len(s)
    })


    if !CompareIntSlices(expected, result) {
		 t.Errorf("should return %v but got %v", expected, result)
    }
}

func TestFilter(t *testing.T) {
    slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
    expected := []int{1,3,5,7, 9}

    result := Filter(slice, func(i int) bool {
        return i % 2 == 1
    })


    if !CompareIntSlices(expected, result) {
		 t.Errorf("should return %v but got %v", expected, result)
    }
}
