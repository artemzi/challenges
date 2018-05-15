package subdomain

/*
Subdomain Visit Count problem from leetcode contest 78

A website domain like "discuss.leetcode.com" consists of various subdomains. At the top level, we have "com", at the next level, we have "leetcode.com", and at the lowest level, "discuss.leetcode.com". When we visit a domain like "discuss.leetcode.com", we will also visit the parent domains "leetcode.com" and "com" implicitly.
Now, call a "count-paired domain" to be a count (representing the number of visits this domain received), followed by a space, followed by the address. An example of a count-paired domain might be "9001 discuss.leetcode.com".
We are given a list cpdomains of count-paired domains. We would like a list of count-paired domains, (in the same format as the input, and in any order), that explicitly counts the number of visits to each subdomain.
*/

import (
	"fmt"
	"strconv"
	"strings"
)

func subdomainVisits(cpdomains []string) []string {
	domainData := make([][]string, len(cpdomains))
	var (
		storage []string
		result  []string
	)

	for i := range cpdomains {
		s := strings.Split(cpdomains[i], " ")
		domainData[i] = s
	}

	for _, v := range domainData {
		tmp := strings.Split(v[1], ".")

		storage = append(storage, fmt.Sprintf("%v %v", v[0], strings.Join(tmp, ".")))
		storage = append(storage, fmt.Sprintf("%v %v", v[0], strings.Join(tmp[1:], ".")))
		storage = append(storage, fmt.Sprintf("%v %v", v[0], strings.Join(tmp[2:], ".")))

	}

	m := make(map[string]int, len(storage))
	for _, v := range storage {
		data := strings.Split(v, " ")
		i, _ := strconv.Atoi(data[0])
		if _, ok := m[data[1]]; ok {
			m[data[1]] += i
			continue
		}
		if data[1] == "" {
			continue
		}
		m[data[1]] = i
	}

	for k, v := range m {
		result = append(result, fmt.Sprintf("%v %v", v, k))
	}

	return result
}
