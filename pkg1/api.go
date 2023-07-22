package pkg1

// ApiOne returns a string telling you that this is the API One
//
// Example:
//
//	res := pkg1.ApiOne()
//
// See also: [github.com/willianpc/example.pkg2.ApiTwo]
func ApiOne() string {
	return "I am the API One"
}

// ApiOneAndHalf returns a string telling you that this is the API One
//
// Example:
//
//	res := pkg1.ApiOneAndHalf()
//
// See also: [ApiOne]
func ApiOneAndHalf() string {
	return "I am the API One and a half"
}
