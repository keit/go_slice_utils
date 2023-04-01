package main

import (
	"fmt"

	"github.com/keit/go_slice_utils"
)

func main() {
	slice := []string{"kei", "yoko", "blues", "jake"}
    fmt.Println(go_slice_utils.Contains(slice, "jake"))
}
