package mystrings

func Reverse(input string) string{
	s:=[]rune(input)
	for i,j:=0,len(s)-1;i<j;i,j=i+1,j-1{
		temp:=s[i]
		s[i]=s[j]
		s[j]=temp
	}
	return string(s)
}