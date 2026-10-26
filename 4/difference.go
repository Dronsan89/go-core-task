package main

// difference возвращает элементы first, которых нет в second.
// Порядок и повторения элементов first сохраняются.
func difference(first, second []string) []string {
	result := make([]string, 0, len(first))
	m := make(map[string]struct{})

	for _, value := range second {
		m[value] = struct{}{}
	}

	for _, value := range first {
		if _, ok := m[value]; !ok {
			result = append(result, value)
		}
	}

	return result
}
