package stringutil

// Reverse methods reverses the given string
// Returns reversed string
func Reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

/*
  Revese method is exported (Means that it's assible oudside the package).
  This is due to capitalisation, first letter capitalized means exported.
  This applies to methods, Structs, variables, etc...
*/
