package hw03frequencyanalysis

import (
	"sort"
	"strings"
)

func Top10(text string) []string {
	words := strings.Fields(text)
	cache := make(map[string]int)
	for _, v := range words {
		if v != "" {
			if _, ok := cache[v]; !ok {
				cache[v] = 1
			} else {
				cache[v]++
			}
		}
	}
	result := make([]string, 0)
	subSlice := make([]string, 0)
	maxKeys := make([]string, 0)
	if len(cache) < 10 {
		return result
	}
	for len(result) < 10 {
		maxVal := -1
		maxKey := ""
		for key, val := range cache {
			if val > maxVal {
				subSlice = make([]string, 0)
				maxKeys = make([]string, 0)
				maxVal = val
				maxKey = key
				subSlice = append(subSlice, maxKey)
				maxKeys = append(maxKeys, maxKey)
			} else if val == maxVal {
				maxKeys = append(maxKeys, key)
				subSlice = append(subSlice, key)
			}
		}
		sort.Strings(subSlice)
		result = append(result, subSlice...)
		for _, k := range maxKeys {
			delete(cache, k)
		}
		subSlice = make([]string, 0)
		maxKeys = append(maxKeys, maxKey)
	}
	return result[:10]
}