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
	array_str := strings.Fields(str)
	if len(array_str) == 0 {
		return nil
	}

	for _, v := range array_str {
		_, ok := dictionary[v]
		if ok {
			dictionary[v] = dictionary[v] + 1
		} else {
			dictionary[v] = 1
		}
	}

	items := make([]Item, len(array_str))
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

	used_words := make([]string, 10)

	if len(items) >= 10 {
		for i := 0; i < 10; i++ {
			used_words[i] = items[i].Key
		}
	} else {
		for i, v := range items {
			used_words[i] = v.Key
		}
	}

	return used_words
}
