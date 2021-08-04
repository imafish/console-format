package main

import (
	"time"

	"github.com/fatih/color"
	cf "github.com/imafish/console-format"
)

func main() {
	defer cf.Close()

	config := cf.Config{
		TextSuffixLength: 7,
		Padding:          '.',
	}
	cf.SetConfig(config)

	for i := 0; i < 11; i++ {
		cf.Print("abc", color.New(color.FgGreen).Sprintf("[%d%%]", i*10))
		time.Sleep(time.Millisecond * time.Duration(210))
	}
	cf.PrintlnNoSuffix(color.New(color.FgBlue).Sprintf("DONE"))

	for i := 0; i < 11; i++ {
		cf.Print("abc", color.New(color.FgGreen).Sprintf("[%d%%]", i*10))
		time.Sleep(time.Millisecond * time.Duration(210))
	}
	cf.NextLine()
	cf.PrintlnNoSuffix(color.New(color.FgBlue).Sprintf("DONE"))
}
