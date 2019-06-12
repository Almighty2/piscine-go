package main
var nb int
func IsNegative(nb int) {

  if nb<0 {
  	print("T\n")
  }else {
  	print("F\n")
  }
}
func main() {
	IsNegative(1)
	IsNegative(0)
	IsNegative(-1)
}
