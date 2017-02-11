package stringutil

// Not assible outside the package.
var myName = "Dylan"

// MyLastname is assible outside the package.
var MyLastname = "van Gils"

// GetName is returining the not exported (not visable) 'myName' variable
func GetName() string {
	return myName
}
