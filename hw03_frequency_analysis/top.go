package hw03frequencyanalysis

import (
	"sort"
	"strings"
)

type Word struct {
	w       string
	repeats int
}

func Top10(str string) []string {
	words := strings.Fields(str)
	resMap := make(map[string]int)
	for _, word := range words {
		resMap[word]++
	}
	resList := make([]Word, 0)
	for key, value := range resMap {
		resList = append(resList, Word{key, value})
	}
	sort.Slice(resList, func(i, j int) bool {
		if resList[i].repeats == resList[j].repeats {
			return resList[i].w < resList[j].w
		}
		return resList[i].repeats > resList[j].repeats
	})
	result := make([]string, 0)
	for _, word := range resList {
		result = append(result, word.w)
	}
	if len(result) > 10 {
		return result[:10]
	}
	return result
}
