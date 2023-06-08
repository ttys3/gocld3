package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/ttys3/gocld3/cld3"
)

func main() {
	fmt.Println("tests begin")
	ts := time.Now()
	defer func() {
		fmt.Println("tests end. time elapsed: ", time.Since(ts))
	}()

	cld, _ := cld3.New(5, 500)
	defer cld.Free()

	var wg sync.WaitGroup
	for j := 0; j < 8; j++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := 0; i < 9999999; i++ {
				res := cld.FindLanguage("Hey, this is an english sentence")
				fmt.Println(res)
				res = cld.FindLanguage("Muy bien, gracias.")
				fmt.Println(res)

				topLangs := cld.FindTopNMostFreqLangs("Hey, this is an english sentence 这是一段中文 ja こんいちは", 3)
				if topLangs == nil {
					panic("topLangs is nil")
				}
				for _, r := range topLangs {
					fmt.Println(r)
				}
			}
		}()
	}

	wg.Wait()
	fmt.Print("sleeping for 15 minutes")
	time.Sleep(time.Second * 900)
}
