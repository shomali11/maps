package main

import (
	"fmt"
	"github.com/shomali11/maps"
)

func main() {
	concurrentMap := maps.NewShardedConcurrentMap(maps.WithNumberOfShards(100))
	concurrentMap.Set("name", "Raed Shomali")

	fmt.Println(concurrentMap.Contains("name"))
	fmt.Println(concurrentMap.Get("name"))
	fmt.Println(concurrentMap.Size())

	concurrentMap.Remove("name")

	fmt.Println(concurrentMap.Contains("name"))
	fmt.Println(concurrentMap.Get("name"))
	fmt.Println(concurrentMap.Size())

	concurrentMap.Set("name", "Raed Shomali")
	concurrentMap.Clear()

	fmt.Println(concurrentMap.Contains("name"))
	fmt.Println(concurrentMap.Get("name"))
	fmt.Println(concurrentMap.Size())
}
