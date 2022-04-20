// Author:- Roshan Adhikari
// Simple program to fix GIS Json error
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strings"
)

// regex to capture datas inside the small bracket paranthesis
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
		if strings.Contains(line, "NumberInt") {
			rs := rgx.FindStringSubmatch(lines[i])
			fixedLiner := fmt.Sprintf("\"ogc_fid\" : %s,", rs[1])
			lines[i] = strings.ReplaceAll(line, line, fixedLiner)
		}

	}
	output := strings.Join(lines, "\n")
	// writing the datas into new file
	e := ioutil.WriteFile("filesuccess.json", []byte(output), 0644)
	if e != nil {
		panic(e)
	}

}
