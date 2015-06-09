/****************
*****************
***** Notes *****
*****************
*****************

-All numbers are two's complement signed (Due to Amiga specifications).
-'M.K' identifier allows for 31 instruments, four channels. Located address 1080d
-Motorola byte order (Big endian)
-Pattern data is 1024 bytes
--Must find byte order for pattern number
-JSON is a picky file format!!!!!!!!

*/


package main

import ("fmt"
	"io/ioutil"
	)

type sigData struct{
	Title string
	Data []byte
}

func main(){
	fileDir := ""
	fmt.Print("File Directory: ")
	fmt.Scanf("%s", &fileDir)
	mod, _ := ioutil.ReadFile(fileDir)
	
	fmt.Println("Title: ", string(mod[0:20]))
	fmt.Println("File format: ", string(mod[1080:1084]))
	fmt.Println("Number of patterns in song: ", byte(mod[950]))
	fmt.Println("Pattern Table: ", mod[952: 1080])
	//fmt.Println("Sample 1 Name: ", string(mod[20: 44]))
	//for z := 0; z < 42; z++{
	//	fmt.Println(mod[(1084 + (1024* z)):(1088 + (1024* z))])
	//}
	temp := byte(0)
	for _, val := range mod[952: 1080]{
		if val > temp{
			temp = val
		}
	}
	fmt.Println("Highest value in pattern table: ", temp)
	m := new(sigData)

	m.Title = "Test"
	m.Data = mod[1084 + (1024 * 41): 8764 + (1024 * 41)]
	for _, x := range m.Data{
		fmt.Print(int8(x), ",")
	}
/*	b, _ := json.Marshal(m)
	ioutil.WriteFile("test.json", b, 0644)*/
}
