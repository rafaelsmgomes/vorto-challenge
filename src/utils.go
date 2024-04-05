package main

func allTrue(values []bool) bool {
	for _, val := range values {
		if !val {
			return false
		}
	}
	return true
}
