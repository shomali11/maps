package main

import (
	"fmt"
	"github.com/shomali11/maps"
)

func main() {
	concurrentMap := maps.NewConcurrentMultiMap()
	concurrentMap.Set("names", []interface{}{"Raed Shomali"})
	concurrentMap.Append("names", "Dwayne Johnson")

	fmt.Println(concurrentMap.ContainsKey("names"))
	fmt.Println(concurrentMap.ContainsEntry("names", "Raed Shomali"))
	fmt.Println(concurrentMap.Get("names"))
	fmt.Println(concurrentMap.Size())
	fmt.Println(concurrentMap.IsEmpty())
	fmt.Println(concurrentMap.Keys())

	concurrentMap.Remove("names")

	fmt.Println(concurrentMap.ContainsKey("names"))
	fmt.Println(concurrentMap.ContainsEntry("names", "Raed Shomali"))
	fmt.Println(concurrentMap.Get("names"))
	fmt.Println(concurrentMap.Size())
	fmt.Println(concurrentMap.IsEmpty())
	fmt.Println(concurrentMap.Keys())

	concurrentMap.Append("names", "Raed Shomali")
	concurrentMap.Clear()

	fmt.Println(concurrentMap.ContainsKey("names"))
	fmt.Println(concurrentMap.ContainsEntry("names", "Raed Shomali"))
	fmt.Println(concurrentMap.Get("names"))
	fmt.Println(concurrentMap.Size())
	fmt.Println(concurrentMap.IsEmpty())
	fmt.Println(concurrentMap.Keys())
}
