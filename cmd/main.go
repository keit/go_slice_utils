package main

import (
	"fmt"

	"github.com/keit/go_map_reduce"
)

func main() {
	slice := []string{"kei", "yoko", "blues", "jake"}
    fmt.Println(go_map_reduce.Contains(slice, "jake"))
}
