package main
import "fmt"
type car struct{
	model string
	wheels int
}
type vhecles struct{
	weight int 
	car
}
//now vheichels contain all members of car
func main(){
	temp := vhecles{
		weight: 5,
		car:car{model:"suzuki",
		wheels:5,},
	}
	fmt.Println(temp)
}
