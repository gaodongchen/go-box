package gobox

func Chunk(list []interface{}, limit int) [][]interface{} {
	result := [][]interface{}{}
	for i := 0; i < len(list); i += limit {
		length := Min(i+limit, len(list))
		batch := list[i:length]
		result = append(result, batch)
	}

	return result
}
