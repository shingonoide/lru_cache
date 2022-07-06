package main

import (
	"fmt"

	"github.com/shingonoide/lru_cache/internal/cache_lru"
)

func main() {
	cache := cache_lru.NewCacheLRU(10)
	cache.Write("a", "bcd")
	cache.Write("b", "123dfg")
	fmt.Println(cache.Read("a"))
	fmt.Println(cache.Read("b"))
	fmt.Println(cache.Length())
}
