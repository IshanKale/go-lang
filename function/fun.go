package main
import "fmt"
func add(a int,b int)int{
	return a+b
}
func add1(a,b int)int{
	return a+b
}
func cord()(x,y int){
	x=3
	y=0
	return
}
func cord1()(x,y int){
	return 6,7
}
func main()  {
	fmt.Println(add(3,4))
	fmt.Println(add1(3,4))
	fmt.Println(cord())
	fmt.Println(cord1())
}