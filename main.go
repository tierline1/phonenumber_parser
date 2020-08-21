package main

import (
	"github.com/ttacon/libphonenumber"
	"github.com/euskadi31/go-tokenizer"
	"fmt"
	"os"
	"bufio"
	"time"
	
)


func main() {
	//var phone string
	fmt.Println("Введи текст:")
	in := bufio.NewReader(os.Stdin)
	line, err := in.ReadString('\n')
	if err != nil {
			fmt.Println(err)
			os.Exit(0)
	}
	start := time.Now()
	t := tokenizer.New()

    tokens := t.Tokenize(line)
	fmt.Println("")
	//os.Exit(0)
	for _, token := range tokens {
        
		num, err := libphonenumber.Parse(token, "RU")
		if err != nil {
			//fmt.Println(err)
			//fmt.Println(" - no number")
		} else {
			
			formattedNum := libphonenumber.Format(num, libphonenumber.INTERNATIONAL)
			fmt.Printf("%s - %s\n",token,formattedNum)
		}
		
    }
	elapsed := time.Since(start)
    fmt.Printf("\n\nПарсинг занял: %s", elapsed)
    
    
	
}
