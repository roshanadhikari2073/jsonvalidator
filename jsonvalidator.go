package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strings"
)

var rgx = regexp.MustCompile(`\((.*?)\)`)

func main() {
	input, err := ioutil.ReadFile("test.json")
	if err != nil {
		log.Fatalln(err)
	}
	lines := strings.Split(string(input), "\n")

	for i, line := range lines {
		if strings.Contains(line, "ObjectId") {
			rs := rgx.FindStringSubmatch(lines[i])
			fixedLine := fmt.Sprintf("\"_id\" : %s,", rs[1])
			lines[i] = strings.ReplaceAll(line, line, fixedLine)
		}

	}
	output := strings.Join(lines, "\n")
	fmt.Println(string(output))
}
