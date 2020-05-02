package utils

func SliceHasString(needle string, haystack []string) bool {
	for _, v := range haystack {
		if v == needle {
			return true
		}
	}

	return false
}

func SliceUniqueStrings(strings []string) []string {
	result := make([]string, 0)

	for _, v := range strings {
		if SliceHasString(v, result) {
			continue
		}

		result = append(result, v)
	}

	return result
}

func SliceUnionStrings(stringSlices ...[]string) []string {
	result := make([]string, 0)

	for _, stringSlice := range stringSlices {
		stringSlice = SliceUniqueStrings(stringSlice)

		for _, v := range stringSlice {
			if SliceHasString(v, result) {
				continue
			}

			result = append(result, v)
		}
	}

	return result
}
