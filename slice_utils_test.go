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
