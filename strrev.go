package piscine

func StrRev(s string) string {
	reverse := ""
	for _, i := range s {
		reverse = string(i) + reverse
	}
	return reverse
}
