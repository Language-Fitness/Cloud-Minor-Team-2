package helper

func DereferenceArrayIfNeeded(value interface{}) []string {
	var newArray []string

	if IsNil(value) {
		return newArray
	}

	if myArray, ok := value.([]*string); ok {
		for _, pointer := range myArray {
			if pointer == nil || IsNil(pointer) {
				continue
			}

			value := *pointer
			newArray = append(newArray, value)
		}
	}

	return newArray
}
