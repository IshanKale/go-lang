package main
import "fmt"
type rec struct{
	height int
	width int
}
func (r rec) area() int {
	return r.height*r.width
}
func main(){
	r1:=rec{
		height: 6,
		width: 5,
	}
	fmt.Print(r1.area())

}
