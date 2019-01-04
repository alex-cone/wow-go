package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const yourTurn = " **YOUR TURN**"
const placeOrb = " **PLACE ORB**"

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

// writeLines writes the lines to the given file.
func writeLines(lines []string, path string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	w := bufio.NewWriter(file)
	m := make(map[string][]string)
	count := 1
	currentSection := ""
	for _, line := range lines {
		if len(line) > 3 {
			if strings.Contains(line, "Left") {
				if strings.Contains(currentSection, "Right") {
					count = count + 1
				}
				currentSection = "Left" + strconv.Itoa(count)
			} else if strings.Contains(line, "Right") {
				currentSection = "Right" + strconv.Itoa(count)
			} else {
				m[currentSection] = append(m[currentSection], strings.TrimSpace(line))
			}
		}
	}
	writeOneMacros(w, m)
	writeTwoMacros(w, m)
	writeThreeMacros(w, m)
	writeFourMacros(w, m)
	writeFiveMacros(w, m)
	return w.Flush()
}

func writeOneMacros(w *bufio.Writer, m map[string][]string) {
	//Name of left person
	if len(m["Left1"]) >= 1 {
		fmt.Fprintln(w, m["Left1"][len(m["Left1"])-1])
	} else {
		panic(fmt.Sprintf("Size of left1 is empty. Check the format of your file."))
	}

	//Macros for left person
	fmt.Fprintln(w, "/w "+m["Right1"][len(m["Right1"])-1]+placeOrb)
	if len(m["Left2"]) == 1 {
		fmt.Fprintln(w, "/w "+m["Left2"][0]+yourTurn)
	} else if len(m["Left2"]) == 2 {
		fmt.Fprintln(w, "/w "+m["Left2"][0]+yourTurn)
		fmt.Fprintln(w, "/w "+m["Left2"][1]+yourTurn)
	} else {
		panic(fmt.Sprintf("Size of left2 is empty. Check the format of your file."))
	}
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "")
	//Name of right person
	if len(m["Right1"]) >= 1 {
		fmt.Fprintln(w, m["Right1"][len(m["Right1"])-1])
	} else {
		panic(fmt.Sprintf("Size of right1 is empty. Check the format of your file."))
	}

	//Macros for right person
	fmt.Fprintln(w, "/w "+m["Left1"][len(m["Left1"])-1]+placeOrb)
	if len(m["Right2"]) == 1 {
		fmt.Fprintln(w, "/w "+m["Right2"][0]+yourTurn)
	} else if len(m["Right2"]) == 2 {
		fmt.Fprintln(w, "/w "+m["Right2"][0]+yourTurn)
		fmt.Fprintln(w, "/w "+m["Right2"][1]+yourTurn)
	} else {
		panic(fmt.Sprintf("Size of right2 is empty. Check the format of your file."))
	}
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "")
}

func writeTwoMacros(w *bufio.Writer, m map[string][]string) {
	//Name of left person
	if len(m["Left2"]) >= 1 {
		fmt.Fprintln(w, m["Left2"][len(m["Left2"])-1])
	} else {
		panic(fmt.Sprintf("Size of Left2 is empty. Check the format of your file."))
	}

	//Macros for left person
	fmt.Fprintln(w, "/w "+m["Right2"][len(m["Right2"])-1]+placeOrb)
	if len(m["Left3"]) == 1 {
		fmt.Fprintln(w, "/w "+m["Left3"][0]+yourTurn)
	} else if len(m["Left3"]) == 2 {
		fmt.Fprintln(w, "/w "+m["Left3"][0]+yourTurn)
		fmt.Fprintln(w, "/w "+m["Left3"][1]+yourTurn)
	} else {
		panic(fmt.Sprintf("Size of left3 is empty. Check the format of your file."))
	}

	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "")

	//Name of right person
	if len(m["Right2"]) >= 1 {
		fmt.Fprintln(w, m["Right2"][len(m["Right2"])-1])
	} else {
		panic(fmt.Sprintf("Size of Right2 is empty. Check the format of your file."))
	}

	//Macros for right person
	fmt.Fprintln(w, "/w "+m["Left2"][len(m["Left2"])-1]+placeOrb)
	if len(m["Right3"]) == 1 {
		fmt.Fprintln(w, "/w "+m["Right3"][0]+yourTurn)
	} else if len(m["Right3"]) == 2 {
		fmt.Fprintln(w, "/w "+m["Right3"][0]+yourTurn)
		fmt.Fprintln(w, "/w "+m["Right3"][1]+yourTurn)
	} else {
		panic(fmt.Sprintf("Size of right2 is empty. Check the format of your file."))
	}
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "")
}

func writeThreeMacros(w *bufio.Writer, m map[string][]string) {
	//Name of left person
	if len(m["Left3"]) >= 1 {
		fmt.Fprintln(w, m["Left3"][len(m["Left3"])-1])
	} else {
		panic(fmt.Sprintf("Size of Left3 is empty. Check the format of your file."))
	}

	//Macros for left person
	fmt.Fprintln(w, "/w "+m["Right3"][len(m["Right3"])-1]+placeOrb)
	if len(m["Left4"]) == 1 {
		fmt.Fprintln(w, "/w "+m["Left4"][0]+yourTurn)
	} else if len(m["Left4"]) == 2 {
		fmt.Fprintln(w, "/w "+m["Left4"][0]+yourTurn)
		fmt.Fprintln(w, "/w "+m["Left4"][1]+yourTurn)
	} else {
		panic(fmt.Sprintf("Size of left4 is empty. Check the format of your file."))
	}

	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "")

	//Name of right person
	if len(m["Right3"]) >= 1 {
		fmt.Fprintln(w, m["Right3"][len(m["Right3"])-1])
	} else {
		panic(fmt.Sprintf("Size of Right3 is empty. Check the format of your file."))
	}

	//Macros for right person
	fmt.Fprintln(w, "/w "+m["Left3"][len(m["Left3"])-1]+placeOrb)
	if len(m["Right4"]) == 1 {
		fmt.Fprintln(w, "/w "+m["Right4"][0]+yourTurn)
	} else if len(m["Right4"]) == 2 {
		fmt.Fprintln(w, "/w "+m["Right4"][0]+yourTurn)
		fmt.Fprintln(w, "/w "+m["Right4"][1]+yourTurn)
	} else {
		panic(fmt.Sprintf("Size of right4 is empty. Check the format of your file."))
	}
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "")
}

func writeFourMacros(w *bufio.Writer, m map[string][]string) {
	//Name of left person
	if len(m["Left4"]) >= 1 {
		fmt.Fprintln(w, m["Left4"][len(m["Left4"])-1])
	} else {
		panic(fmt.Sprintf("Size of Left4 is empty. Check the format of your file."))
	}

	//Macros for left person
	fmt.Fprintln(w, "/w "+m["Right4"][len(m["Right4"])-1]+placeOrb)
	if len(m["Left5"]) == 1 {
		fmt.Fprintln(w, "/w "+m["Left5"][0]+yourTurn)
	} else if len(m["Left5"]) == 2 {
		fmt.Fprintln(w, "/w "+m["Left5"][0]+yourTurn)
		fmt.Fprintln(w, "/w "+m["Left5"][1]+yourTurn)
	} else {
		panic(fmt.Sprintf("Size of left5 is empty. Check the format of your file."))
	}

	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "")

	//Name of right person
	if len(m["Right4"]) >= 1 {
		fmt.Fprintln(w, m["Right4"][len(m["Right4"])-1])
	} else {
		panic(fmt.Sprintf("Size of Right4 is empty. Check the format of your file."))
	}

	//Macros for right person
	fmt.Fprintln(w, "/w "+m["Left4"][len(m["Left4"])-1]+placeOrb)
	if len(m["Right5"]) == 1 {
		fmt.Fprintln(w, "/w "+m["Right5"][0]+yourTurn)
	} else if len(m["Right5"]) == 2 {
		fmt.Fprintln(w, "/w "+m["Right5"][0]+yourTurn)
		fmt.Fprintln(w, "/w "+m["Right5"][1]+yourTurn)
	} else {
		panic(fmt.Sprintf("Size of right5 is empty. Check the format of your file."))
	}

	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "")
}

func writeFiveMacros(w *bufio.Writer, m map[string][]string) {
	//Name of left person
	if len(m["Left5"]) >= 1 {
		fmt.Fprintln(w, m["Left5"][len(m["Left5"])-1])
	} else {
		panic(fmt.Sprintf("Size of Left5 is empty. Check the format of your file."))
	}

	//Macros for left person
	fmt.Fprintln(w, "/w "+m["Right5"][len(m["Right5"])-1]+placeOrb)
	if len(m["Left1"]) == 1 {
		fmt.Fprintln(w, "/w "+m["Left1"][0]+yourTurn)
	} else if len(m["Left1"]) == 2 {
		fmt.Fprintln(w, "/w "+m["Left1"][0]+yourTurn)
		fmt.Fprintln(w, "/w "+m["Left1"][1]+yourTurn)
	} else {
		panic(fmt.Sprintf("Size of left1 is empty. Check the format of your file."))
	}

	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "")

	//Name of right person
	if len(m["Right5"]) >= 1 {
		fmt.Fprintln(w, m["Right5"][len(m["Right5"])-1])
	} else {
		panic(fmt.Sprintf("Size of Right5 is empty. Check the format of your file."))
	}

	//Macros for right person
	fmt.Fprintln(w, "/w "+m["Left5"][len(m["Left5"])-1]+placeOrb)
	if len(m["Right1"]) == 1 {
		fmt.Fprintln(w, "/w "+m["Right1"][0]+yourTurn)
	} else if len(m["Right1"]) == 2 {
		fmt.Fprintln(w, "/w "+m["Right1"][0]+yourTurn)
		fmt.Fprintln(w, "/w "+m["Right1"][1]+yourTurn)
	} else {
		panic(fmt.Sprintf("Size of right1 is empty. Check the format of your file."))
	}
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "")
}

func main() {
	lines, err := readLines("roster.txt")
	if err != nil {
		lines = runDocs()
		log.Println("no roster.txt file found in current directory...")
	}

	if err := writeLines(lines, "macro.txt"); err != nil {
		log.Fatalf("writeLines: %s", err)
	} else {
		log.Println("Writing to macro.txt from Google Sheet")
	}
}
