package main

import (
	"fmt"
	"github.com/shomali11/cmap"
)

func main() {
	concurrentMap := cmap.NewShardedConcurrentMap()
	concurrentMap.Set("name", "Raed Shomali")

	fmt.Println(concurrentMap.Get("name"))

	concurrentMap.Remove("name")

	fmt.Println(concurrentMap.Get("name"))
}
