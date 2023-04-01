package main

import (
	"fmt"

	"github.com/keit/slice_utils"
)

func main() {
	slice := []string{"kei", "yoko", "blues", "jake"}
    fmt.Println(slice_utils.Contains(slice, "jake"))
}
