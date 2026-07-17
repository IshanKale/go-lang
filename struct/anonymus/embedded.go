package main
import "fmt"
func main(){
	mycar:= struct{
		model string
		wheels int
	}{
		model:"suzuki",
		wheels:4,
	}
	fmt.Println(mycar)

}
