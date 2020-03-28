package stringz

import (
	"strings"
)

// 4:33-4:41
// https://golang.org/pkg/strings/#Index
func Index(str, substr string) int {
	if len(str) == 0 || len(substr) == 0 {
		return -1
	}
	for i:=0; i<len(str); i++ {
		if str[i] == substr[0] {
			if searchForSubstr(str, substr, i) {
				return i
			}
		}
	}
	return -1
}

func searchForSubstr(str, substr string, i int) bool {
	for j:=1; j<len(substr); j++ {
		if substr[j] != str[i+j] {
			return false
		}
	}
	return true
}

// 4:42-4:47
func NextIndex(str, substr string, i int) int {
	if len(str) == 0 || len(substr) == 0 || i < 0 || i >= len(str) {
		return -1
	}
	for i<len(str) {
		if str[i] == substr[0] {
			if searchForSubstr(str, substr, i) {
				return i
			}
		}
		i++
	}
	return -1
}

// 4:58-5:06
// https://golang.org/pkg/strings/#Compare
func Compare(a, b string) int {
	if a == "" && b == "" {
		return 0
	}

	i:=0
	for {
		if i == len(a) && i == len(b) {
			return 0
		}
		if i == len(a) {
			return -1
		}
		if i == len(b) {
			return 1
		}

		if a[i] < b[i] {
			return -1
		}
		if a[i] > b[i] {
			return 1
		}

		i++
	}

	// deadcode
	return -1
}

// 5:06-5:07
// https://golang.org/pkg/strings/#Contains
func Contains(str, substr string) bool {
	return Index(str, substr) != -1
}

// 5:08-5:09
// https://golang.org/pkg/strings/#HasPrefix
func HasPrefix(str, prefix string) bool {
	return Index(str, prefix) == 0
}

// 5:10-5:12
// https://golang.org/pkg/strings/#HasSuffix
func HasSuffix(str, suffix string) bool {
	i := Index(str, suffix)
	return i != -1 && i == len(str)-len(suffix)
}

// 9:46-10:16
// https://golang.org/pkg/strings/#Split
func Split(str, sep string) []string {
	if len(str) == 0 && len(sep) == 0 {
		return []string{}
	}

	var result []string
	var i int

	if len(sep) == 0 {
		result = make([]string, 0, len(str))
		for i=0; i<len(str); i++ {
			result = append(result, str[i:i+1])
		}
		return result
	}

	sepIndices := []int{}

	i = -1
	for {
		i = NextIndex(str, sep, i+1)
		if i == -1 {
			break
		}
		sepIndices = append(sepIndices, i)
	}

	if len(sepIndices) == 0 && len(sep) > 0 {
		return []string{str}
	}

	result = make([]string, 0, len(sepIndices)+1)

	for i=0; i<len(sepIndices); i++ {
		if i == 0 {
			result = append(result, str[:sepIndices[0]])
		} else {
			result = append(result, str[sepIndices[i-1]+len(sep):sepIndices[i]])
		}
		if i == len(sepIndices)-1 {
			result = append(result, str[sepIndices[i]+len(sep):])
		}
	}

	return result
}

// 3:30-4:02
// https://golang.org/pkg/strings/#Replace
func Replace(str, old, new string, n int) string {
	if len(str) == 0 || n == 0 {
		return str
	}

	parts := Split(str, old)
	results := []string{}
	var i int

	if len(old) == 0 {
		for i=0; i<len(parts)-1; i++ {
			if len(new) > 0 {
				results = append(results, new)
			}
			results = append(results, parts[i])
		}
		results = append(results, new)
	} else {
		for i=0; i<len(parts)-1; i++ {
			results = append(results, parts[i])
			if len(new) > 0 && (n > 0 || n < 0) {
				results = append(results, new)
				n--
			}
		}
		results = append(results, parts[len(parts)-1])
	}

	return strings.Join(results, "")
}

// 4:05-4:13
// https://golang.org/pkg/strings/#Join
func Join(elems []string, sep string) string {
	if len(elems) == 0 {
		return ""
	}

	builder := new(strings.Builder)
	var i int

	for i=0; i<len(elems)-1; i++ {
		builder.WriteString(elems[i])
		builder.WriteString(sep)
	}
	builder.WriteString(elems[i])
	
	return builder.String()
}