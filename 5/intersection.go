package main

// intersection проверяет, есть ли пересечение значений first и second,
// и возвращает уникальные пересечённые значения в порядке second.
func intersection(first, second []int) (bool, []int) {
	firstValues := make(map[int]struct{}, len(first))
	for _, value := range first {
		firstValues[value] = struct{}{}
	}

	result := make([]int, 0, min(len(first), len(second)))
	for _, value := range second {
		if _, exists := firstValues[value]; exists {
			result = append(result, value)
			// Удаляем значение, чтобы оно не повторилось в результате.
			delete(firstValues, value)
		}
	}

	return len(result) > 0, result
}
