package helper

func DereferenceArrayIfNeeded(value interface{}) []string {
	var newArray []string

	if myArray, ok := value.([]*string); ok {
		for _, pointer := range myArray {
			if pointer == nil {
				continue
			}

			value := *pointer
			newArray = append(newArray, value)
		}
	}

	return newArray
}
