package main
import "fmt"
func main(){
	mp:=make(map[int]map[int]string)
	if mp[8]==nil {
		mp[8]=make(map[int]string)
	}
	mp[8][8]="ji"
	fmt.Println(mp)
}