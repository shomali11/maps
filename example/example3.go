package main

import (
	"fmt"
	"github.com/shomali11/cmap"
)

func main() {
	concurrentMap := cmap.NewShardedConcurrentMap(cmap.WithNumberOfShards(100))
	concurrentMap.Set("name", "Raed Shomali")

	fmt.Println(concurrentMap.Get("name"))

	concurrentMap.Remove("name")

	fmt.Println(concurrentMap.Get("name"))
}
