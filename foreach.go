
package piscine

func ForEach(maf func(int), arr []int) {
	for _,i:= range arr{
		maf(i)
	}
}