package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"math"
	"math/rand"
	"os"
	"path"
	"path/filepath"
	"time"

	"github.com/fatih/color"
	cf "github.com/imafish/console-format"
)

func main() {
	directory := flag.String("d", ".", "directory to parse")
	pattern := flag.String("p", "*", "file pattern to parse")
	round := flag.Int("r", 0, "rounds to go, 0 for infinite")
	flag.Parse()

	cf.SetConfig(cf.Config{
		Padding:          '.',
		TextSuffixLength: 7,
	})
	defer cf.Close()

	if *round > 0 {
		for i := 0; i < *round; i++ {
			readDirRecursive(*directory, *pattern)
		}
	} else {
		for {
			readDirRecursive(*directory, *pattern)
		}
	}
}

func readDirRecursive(directory string, pattern string) {
	red := color.New(color.FgRed).SprintFunc()

	files, err := ioutil.ReadDir(directory)
	if err != nil {
		cf.Println(fmt.Sprintf("Failed to open dir %s: %s", directory, err), red("ERROR"))
		return
	}
	cf.PrintlnNoSuffix(fmt.Sprintf("Entering dir %s", directory))

	for _, f := range files {
		name := f.Name()
		if f.IsDir() {
			readDirRecursive(path.Join(directory, name), pattern)
		}
		match, err := filepath.Match(pattern, name)
		if err != nil {
			cf.Println(fmt.Sprintf("Failed to parse filename %s: %s", name, err), red("ERROR"))
		}
		if match {
			doOneFile(f)
		}
	}
}

func doOneFile(f os.FileInfo) {
	filename := f.Name()
	red := color.New(color.FgRed).SprintFunc()
	green := color.New(color.FgGreen).SprintFunc()
	yellow := color.New(color.FgHiYellow).SprintFunc()
	warn := rand.Intn(35) == 0

	level := int(math.Log10(float64(f.Size())) - 3)
	if level < 1 {
		level = 1
	}

	for i := 0; i < 100; i += rand.Intn(70/level) + 7 {
		fail := rand.Intn(140) == 0
		time.Sleep(time.Millisecond * time.Duration(rand.Intn(70*level)))
		if fail {
			cf.Println(fmt.Sprintf("Failed to parse file %s", filename), red("ERROR"))
			return
		}

		cf.Print(fmt.Sprintf("Parsing %s", filename), fmt.Sprintf("[%d%%]", i))
	}

	var msg string
	if warn {
		msg = yellow("WARN")
	} else {
		msg = green("OK")
	}
	cf.Println(fmt.Sprintf("Parsing %s", filename), msg)
}
