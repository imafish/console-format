package main

import (
	"fmt"
	"os"
	"time"

	cf "github.com/imafish/console-format"
)

func main() {
	err := cf.Init()
	if err != nil {
		fmt.Printf("%s\n", err.Error())
		os.Exit(1)
	}
	defer cf.Close()

	config := cf.Config{
		TextSuffixLength: 7,
		Padding:          '.',
	}
	cf.SetConfig(config)

	for i := 0; i < 11; i++ {
		cf.PrintInCurrentLine(cf.Line{Prefix: "abc", Suffix: fmt.Sprintf("[%d%%]", i*10)})
		time.Sleep(time.Second)
	}

}
