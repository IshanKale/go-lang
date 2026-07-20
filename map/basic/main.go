package main
import "fmt"
func main(){
	mp:=make(map[string][]int)
	fmt.Println(mp)
	mp["abc"] = append(mp["abc"],6)
	fmt.Println(mp)

}