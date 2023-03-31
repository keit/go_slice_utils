package go_map_reduce

func Contains[E comparable](s []E, v E) bool {
    for _, vs := range s {
        if v == vs {
            return true
        }
    }
    return false
}

func Reverse[E any](s []E) []E {
    result := make([]E, 0, len(s))
    for i := len(s) - 1; i >= 0; i-- {
        result = append(result, s[i])
    }
    return result
}

func CompareStringSlices(a, b []string) bool {
        if len(a) != len(b) {
                return false
        }
        for i, v := range a {
                if v != b[i] {
                        return false
                }
        }
        return true
}
