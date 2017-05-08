package main

import (
    //"fmt"
    "os"
    "bufio"
    "fmt"
    "encoding/json"
    //"log"
)


func main() {
	//fmt.Println("GO STARTED2")

	//arrsiz := bufio.NewReader(os.Stdin)
	bio := bufio.NewReader(os.Stdin)

	arrsiz, _, _ := bio.ReadLine()
	var first int
	lolerr := json.Unmarshal(arrsiz, &first)
	arrsiz2 := int(first)
	if lolerr != nil {
		panic(lolerr)
	}

	if arrsiz2 > 0 {

		bio = bufio.NewReader(os.Stdin)
		line, _, _ := bio.ReadLine()
		//		arrsiz := bio.ReadLine()

		var p = [][]float32{}

		err := json.Unmarshal(line, &p)
		if err != nil {
			panic(err)
		}

		bytes2, err2 := json.Marshal(p)
		if err2 != nil {
			panic(err2)
		}
		s := string(bytes2[:])

		fmt.Println(s)
	}

}
