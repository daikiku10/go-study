package utils

// ToMap スライスをKeyValueのマップに変換する
// キーは任意の要素、値は要素自体を格納する
// valuesがnilの場合はからのマップを返す
func ToMap[V any, K comparable](values *[]V, keyExtractor func(V) K) map[K]V {
	if values == nil {
		return map[K]V{}
	}

	result := make(map[K]V, len(*values))
	for _, value := range *values {
		result[keyExtractor(value)] = value
	}
	return result
}
