package go_map_reduce

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
        t.Errorf("should contains jake")
    }
    result = Contains(slice, "taro")
    if result {
        t.Errorf("should not contains taro")
    }

}

func TestReverse(t *testing.T) {
	slice := []string{"kei", "yoko", "blues", "jake"}
	result := Reverse(slice)

	if !CompareStringSlices(result, []string{"jake", "blues", "yoko", "kei"}) {
		 t.Errorf("Reverse failed")
	}

}
