package piscine

func Any(f func(string) bool, arr []string) bool {

	for _,i:=range arr{
		if f(i){
			return true //Retourne TRUE pour le tableau de string
		}
	}
	return false //Retourne false pour le tableau de string
}