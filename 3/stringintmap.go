package stringintmap

type StringIntMap struct {
	data map[string]int
}

func New() *StringIntMap {
	return &StringIntMap{data: make(map[string]int)}
}

// Add добавляет пару key-value или заменяет значение существующего ключа.
func (m *StringIntMap) Add(key string, value int) {
	if m.data == nil {
		m.data = make(map[string]int)
	}
	m.data[key] = value
}

// Remove удаляет элемент по ключу. Если ключа нет, метод ничего не делает.
func (m *StringIntMap) Remove(key string) {
	delete(m.data, key)
}

// Copy возвращает новую независимую карту со всеми элементами.
func (m *StringIntMap) Copy() map[string]int {
	result := make(map[string]int, len(m.data))
	for key, value := range m.data {
		result[key] = value
	}
	return result
}

// Exists сообщает, существует ли ключ в карте.
func (m *StringIntMap) Exists(key string) bool {
	_, exists := m.data[key]
	return exists
}

// Get возвращает значение по ключу и признак успешности операции.
func (m *StringIntMap) Get(key string) (int, bool) {
	value, exists := m.data[key]
	return value, exists
}
