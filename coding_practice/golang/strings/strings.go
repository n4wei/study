package stringz

import (
	"strings"
)

// 4:33-4:41
// https://golang.org/pkg/strings/#Index
func Index(str, substr string) int {
	return NextIndex(str, substr, 0)
}

func searchForSubstr(str, substr string, i int) bool {
	for j:=0; j<len(substr); j++ {
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
	for ; i<=len(str)-len(substr); i++ {
		if searchForSubstr(str, substr, i) {
			return i
		}
	}
	return -1
}

// 4:58-5:06
// https://golang.org/pkg/strings/#Compare
func Compare(a, b string) int {
	for i:=0; i<len(a) && i<len(b); i++ {
		if a[i] < b[i] {
			return -1
		}
		if a[i] > b[i] {
			return 1
		}
	}

	if len(a) == len(b) {
		return 0
	}
	if len(a) < len(b) {
		return -1
	}
	return 1
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
	return NextIndex(str, suffix, len(str)-len(suffix)) != -1
}

// 9:46-10:16
// https://golang.org/pkg/strings/#Split
func Split(str, sep string) []string {
	result := []string{}
	if len(str) == 0 && len(sep) == 0 {
		return result
	}

	var i int
	if len(sep) == 0 {
		result = make([]string, 0, len(str))
		for i:=0; i<len(str); i++ {
			result = append(result, str[i:i+1])
		}
		return result
	}

	rest := str
	var first string
	for {
		i = NextIndex(rest, sep, 0)
		if i == -1 {
			break
		}
		first, rest = rest[:i], rest[i+len(sep):]
		result = append(result, first)
	}
	result = append(result, rest)

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
			if (n > 0 || n < 0) {
				if len(new) > 0 {
					results = append(results, new)
				}
				n--
			} else {
				results = append(results, old)
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