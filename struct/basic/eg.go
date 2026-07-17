package main
import "fmt"
type wheel struct{
	num int
	rubur bool
}
type car struct{
	model string
	company string
	wheel wheel
}
func main(){
	suzuki:=car{}
	suzukiwheel:=wheel{}
	fmt.Println(suzuki)
	fmt.Println(suzukiwheel)
	suzuki.company="suzuki"
	suzuki.model="wagonr"
	suzukiwheel.num=4
	suzukiwheel.rubur=true
	suzuki.wheel=suzukiwheel
	fmt.Println(suzuki)
	suzukiwheel.num=6
	fmt.Println(suzuki)
}
