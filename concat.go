package piscine

func Concat(str1 string,str2 string)string  {

	resultat:=make([]byte,len(str1)+len(str2))
	j:=0
	for i,_:=range result{
		if(i<len(str1)){
			result[i]=str1[i]
		}else{
			result[i]=str2[j]
			j++
		}

	}
	return string(resultat) //Retourner le resultat du programme
}
