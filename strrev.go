package piscine

func StrRev(s string) string {
	reverse := ""         // var string
	for _, i := range s { // loop searches all string characters (var i = char)
		reverse = string(i) + reverse //
	}
	return reverse
}

/*
s(string)=musi
reverse=string(i)+reverse:
"m"+" "="m"
"u"+"m"="um"
"s"+"um"="sum"
"i"+"sum"="isum"*/
