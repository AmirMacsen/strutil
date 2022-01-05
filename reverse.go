package strutil

func ReverseStr(s string)string {

	var a [] rune
	for _, k:= range []rune(s){
		defer func(v rune) {
			a = append(a,v)
		}(k)
	}
	return string(a)
}