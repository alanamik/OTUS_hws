package hw03frequencyanalysis

import (
	"sort"
	"strings"
)

type Item struct {
	Key   string
	Value int
}

func Top10(str string) []string {
	if str == "" || str == " " {
		return nil
	}
	dictionary := make(map[string]int)
	arrayStr := strings.Fields(str)
	if len(arrayStr) == 0 {
		return nil
	}

	for _, v := range arrayStr {
		_, ok := dictionary[v]
		if ok {
			dictionary[v]++
		} else {
			dictionary[v] = 1
		}
	}

	items := make([]Item, len(arrayStr))
	i := 0
	for k, v := range dictionary {
		items[i] = Item{k, v}
		i++
	}

	sort.Slice(items, func(i, j int) bool {
		if items[i].Value != items[j].Value {
			return items[i].Value > items[j].Value
		}

		return items[i].Key < items[j].Key
	})

	usedWords := make([]string, 10)

	if len(items) >= 10 {
		for i := 0; i < 10; i++ {
			usedWords[i] = items[i].Key
		}
	} else {
		for i, v := range items {
			usedWords[i] = v.Key
		}
	}

	return usedWords
}
