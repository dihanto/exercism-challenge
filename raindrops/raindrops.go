package raindrops

import "strconv"

func Convert(number int) string {
	sounds := map[int]string{
		3: "Pling",
		5: "Plang",
		7: "Plong",
	}
	res := ""
	for divisor, sound := range sounds {
		if number%divisor == 0 {
			res += sound
		}
	}
	if res == "" {
		res = strconv.Itoa(number)
	}
	return res
}
