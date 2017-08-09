package main

import (
	"fmt"
	"github.com/shomali11/maps"
)

func main() {
	concurrentMap := maps.NewShardedConcurrentMap(maps.WithNumberOfShards(100))
	concurrentMap.Set("name", "Raed Shomali")

	fmt.Println(concurrentMap.ContainsKey("name"))
	fmt.Println(concurrentMap.ContainsEntry("name", "Raed Shomali"))
	fmt.Println(concurrentMap.Get("name"))
	fmt.Println(concurrentMap.Size())
	fmt.Println(concurrentMap.IsEmpty())
	fmt.Println(concurrentMap.Keys())

	concurrentMap.Remove("name")

	fmt.Println(concurrentMap.ContainsKey("name"))
	fmt.Println(concurrentMap.ContainsEntry("name", "Raed Shomali"))
	fmt.Println(concurrentMap.Get("name"))
	fmt.Println(concurrentMap.Size())
	fmt.Println(concurrentMap.IsEmpty())
	fmt.Println(concurrentMap.Keys())

	concurrentMap.Set("name", "Raed Shomali")
	concurrentMap.Clear()

	fmt.Println(concurrentMap.ContainsKey("name"))
	fmt.Println(concurrentMap.ContainsEntry("name", "Raed Shomali"))
	fmt.Println(concurrentMap.Get("name"))
	fmt.Println(concurrentMap.Size())
	fmt.Println(concurrentMap.IsEmpty())
	fmt.Println(concurrentMap.Keys())
}
