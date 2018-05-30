package main

import (
	"fmt"
	"sort"
	"strings"
)

/**
Поиск образца в тексте
Найти все вхождения строки Pattern в строку Text.
	Вход. Строки Pattern и Text.

	Выход. Все индексы i строки Text, начиная с которых стро-
		ка Pattern входит в Text:
		Text[i..i + |Pattern| − 1] = Pattern.
	Реализуйте алгоритм Карпа–Рабина.

	Формат входа. Образец Pattern и текст Text.

	Формат выхода. Индексы вхождений строки
		Pattern в строку Text в возрастающем порядке, используя индек-
		сацию с нуля.

	Ограничения. 1 ≤ |Pattern| ≤ |Text| ≤ 5 · 10 5 .
		Суммарная длина всех вхождений образ-
		ца в текста не превосходит 10 8 . Обе стро-
		ки содержат буквы латинского алфавита.
*/

func main() {
	var (
		pattern, text string
	)
	fmt.Scan(&pattern)
	fmt.Scan(&text)

	data := make(map[int]string, len(text))
	for i := range text {
		if len(text[i:]) >= len(pattern) {
			data[i] = text[i : i+len(pattern)]
		}
	}

	res := make([]int, 0, len(data))
	for k, val := range data {
		if pattern == val {
			res = append(res, k)
		}
	}
	sort.Ints(res)
	fmt.Println(fmt.Sprintf("%v", strings.Trim(strings.Join(strings.Split(fmt.Sprint(res), " "), " "), "[]")))
}

/*
Rabin-Karp algorithm not used in current task
*/

// const base = 16777619

// // Search searches given patterns in txt and returns the matched ones. Returns
// // empty string slice if there is no match.
// func Search(txt string, patterns []string) []string {
// 	in := indices(txt, patterns)
// 	matches := make([]string, len(in))
// 	i := 0
// 	for j, p := range patterns {
// 		if _, ok := in[j]; ok {
// 			matches[i] = p
// 			i++
// 		}
// 	}

// 	return matches
// }

// // indices returns the indices of the first occurence of each pattern in txt.
// func indices(txt string, patterns []string) map[int]int {
// 	n, m := len(txt), minLen(patterns)
// 	matches := make(map[int]int)

// 	if n < m || len(patterns) == 0 {
// 		return matches
// 	}

// 	var mult uint32 = 1 // mult = base^(m-1)
// 	for i := 0; i < m-1; i++ {
// 		mult = (mult * base)
// 	}

// 	hp := hashPatterns(patterns, m)
// 	h := hash(txt[:m])
// 	for i := 0; i < n-m+1 && len(hp) > 0; i++ {
// 		if i > 0 {
// 			h = h - mult*uint32(txt[i-1])
// 			h = h*base + uint32(txt[i+m-1])
// 		}

// 		if mps, ok := hp[h]; ok {
// 			for _, pi := range mps {
// 				pat := patterns[pi]
// 				e := i + len(pat)
// 				if _, ok := matches[pi]; !ok && e <= n && pat == txt[i:e] {
// 					matches[pi] = i
// 				}
// 			}
// 		}
// 	}
// 	return matches
// }

// func hash(s string) uint32 {
// 	var h uint32
// 	for i := 0; i < len(s); i++ {
// 		h = (h*base + uint32(s[i]))
// 	}
// 	return h
// }

// func hashPatterns(patterns []string, l int) map[uint32][]int {
// 	m := make(map[uint32][]int)
// 	for i, t := range patterns {
// 		h := hash(t[:l])
// 		if _, ok := m[h]; ok {
// 			m[h] = append(m[h], i)
// 		} else {
// 			m[h] = []int{i}
// 		}
// 	}

// 	return m
// }

// func minLen(patterns []string) int {
// 	if len(patterns) == 0 {
// 		return 0
// 	}

// 	m := len(patterns[0])
// 	for i := range patterns {
// 		if m > len(patterns[i]) {
// 			m = len(patterns[i])
// 		}
// 	}

// 	return m
// }
