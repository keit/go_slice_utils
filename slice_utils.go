package slice_utils

func Contains[T comparable](s []T, v T) bool {
    for _, vs := range s {
        if v == vs {
            return true
        }
    }
    return false
}

func Reverse[T any](s []T) []T {
    result := make([]T, 0, len(s))
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

func CompareIntSlices(a, b []int) bool {
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


func Map[T any, M any](a []T, f func(T) M) []M {
    n := make([]M, len(a))
    for i, e := range a {
        n[i] = f(e)
    }
    return n
}

func Filter[T any](s []T, f func(T) bool) []T {
    result := []T{}

    for _, v := range s {
        if f(v) {
            result = append(result, v)
        }
    }
    return result
}
