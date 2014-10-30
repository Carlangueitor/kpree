/*
 Copyright (C) 2014 Charly Román <charly@croman.mx>

 This program is free software: you can redistribute it and/or modify
 it under the terms of the GNU General Public License as published by
 the Free Software Foundation, either version 3 of the License, or
 (at your option) any later version.

 This program is distributed in the hope that it will be useful,
 but WITHOUT ANY WARRANTY; without even the implied warranty of
 MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 GNU General Public License for more details.

 You should have received a copy of the GNU General Public License
 along with this program.  If not, see <http://www.gnu.org/licenses/>.)
*/

package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

const usage = `
Usage:
	kpree <presentation> <outfile>
	kpree file:///home/charly/index.html presentation.pdf`

func main() {
	if len(os.Args) < 3 {
		fmt.Println(usage)
		os.Exit(1)
	}

	slidesName := strings.Replace(os.Args[2], ".pdf", "", 1)
	capturer := exec.Command("slimerjs", "capture.js", os.Args[1], slidesName)
	_, err := capturer.CombinedOutput()

	if err == nil {
		mergerCommand := (`ls ` + slidesName +
			`-*.png | sort -n | tr '\n' ' ' | sed 's/$/\ ` + os.Args[2] +
			`/' | xargs convert`)
		merger := exec.Command("/bin/bash", "-c", mergerCommand)
		_, _err := merger.CombinedOutput()

		if _err != nil {
			log.Fatal("Error generating PDF file.")
		}

		exec.Command("/bin/bash", "-c", `rm `+slidesName+`-*.png`).Run()
	}
}
