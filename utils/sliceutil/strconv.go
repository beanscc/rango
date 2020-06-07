package sliceutil

import "strconv"

// Itoa []int转换成 []string
func Itoa(xi []int) []string {
	if len(xi) == 0 {
		return nil
	}

	xs := make([]string, 0, len(xi))
	for i := 0; i < len(xi); i++ {
		xs = append(xs, strconv.Itoa(xi[i]))
	}

	return xs
}

// Atoi []string 转换成 []int
func Atoi(xs []string) ([]int, error) {
	if len(xs) == 0 {
		return nil, nil
	}

	xi := make([]int, 0, len(xs))
	for i := 0; i < len(xs); i++ {
		ti, err := strconv.Atoi(xs[i])
		if err != nil {
			return xi, err
		}
		xi = append(xi, ti)
	}

	return xi, nil
}
