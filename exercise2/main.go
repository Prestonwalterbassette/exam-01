package main

func CompterLettres(s string) int {
	compteur := 0
	for i := 0; i < len(s); i++ {
		c := s[i]
        if (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') {
            compteur++
        }
    }
    return compteur
}
