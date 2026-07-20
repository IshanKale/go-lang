package main
import "fmt"
type shape interface {
	area() float64
	perimeter() float64
}
type circle struct {
	radius float64
}
type rectangle struct{
	h float64
	w float64
}
func (c circle) area() float64{
	return 3.14*c.radius
}
func (c circle) perimeter() float64{
	return 2*3.14*c.radius
}
func (r rectangle) area() float64{
	return r.h*r.w
}
func (r rectangle) perimeter() float64{
	return 2*(r.h+r.w)
}

func main(){
	var s shape
	fmt.Println(s)
	s=circle{
		radius:3.0,
	}
	fmt.Println(s.area())
	fmt.Println(s.perimeter())
	fmt.Println(s)
	s= rectangle{
		h:2,
		w:3,
	}
	fmt.Println(s.area())
	fmt.Println(s.perimeter())
	fmt.Println(s)
}
