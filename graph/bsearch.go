package graph

func BSearch(vertices []*Vertex, ID int) *Vertex {
	begin := 0
	end := len(vertices) - 1

	for {
		if begin > end {
			break
		}

		mid := (begin + end) / 2
		currentValue := vertices[mid]

		if currentValue.ID == ID {
			return currentValue
		}

		if currentValue.ID > ID {
			end = mid + 1
			continue
		}

		begin = mid - 1
	}

	return nil
}
