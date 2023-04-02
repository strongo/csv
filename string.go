package csv

import "strings"

type String string

// Values returns the values of s.
func (s String) Values() []string {
	if s == "" {
		return nil
	}
	return strings.Split(string(s), ",")
}

// Add adds the values v to s.
func (s String) Add(v ...string) String {
	result := string(s)
	for _, v := range v {
		if result == "" || strings.HasSuffix(result, ",") {
			result += v
		} else {
			result += "," + v
		}
	}
	return String(result)
}

// Remove removes the values v from s.
func (s String) Remove(v ...string) String {
	values := s.Values()
	result := make([]string, 0, len(values))
	for _, value := range values {
		for _, val := range v {
			if val == value {
				goto next
			}
		}
		result = append(result, value)
	next:
	}
	return String(strings.Join(result, ","))
}

// Set sets the value at index i to v.
func (s String) Set(i int, v string) String {
	values := s.Values()
	values[i] = v
	return String(strings.Join(values, ","))
}

// Contains returns true if s contains v.
func (s String) Contains(v string) bool {
	return Contains(string(s), v, ",")
}

// Contains returns true if s contains v, delimited by delimiter.
func Contains(s string, v string, delimiter string) bool { // TODO: implement scanner search
	if s == "" {
		return false
	}
	return s == v ||
		strings.HasPrefix(s, v+delimiter) ||
		strings.HasSuffix(s, delimiter+v) ||
		strings.Contains(s, delimiter+v+delimiter)
}
