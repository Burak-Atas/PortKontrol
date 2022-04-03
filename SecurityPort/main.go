package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
)

var Host = flag.String("h", "", "kontor etmek istediğiniz adres")
var NumberSlice = flag.Int("s", 1, "Kontrol etmek istediğniz aralık")

func main() {
	flag.Parse()
	if !NumberSliceCont() {
		*NumberSlice = 1
	}
	s := bufio.NewScanner(os.Stdin)

	for s.Scan() {
		Sİnt, _ := strconv.Atoi(s.Text())
		for i := 0; i < *NumberSlice; i++ {
			fmt.Println(Port("tcp", *Host, Sİnt))
			fmt.Println()
			Sİnt += 1
		}

	}
}

func NumberSliceCont() bool {
	if 0 >= *NumberSlice || *NumberSlice > 20 {
		return false
	}
	return true
}
