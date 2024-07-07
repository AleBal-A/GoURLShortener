package main

import (
	"GoURLShortener/internal/config"
	"fmt"
)

func main() {
	cfg := config.ConfLoad()
	fmt.Println(cfg)
}
