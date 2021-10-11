//go:build ignore

package main

// THIS HURTS TO WATCH, YES, I KNOW.

// TODO:
// - Convert to "text/template"

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"strconv"
)

func main() {
	file, err := os.Create("tuple/tuple.go")
	if err != nil {
		panic(err)
	}

	if _, err := io.WriteString(file, "// Package generated by tuple/main.go, DO NOT EDIT!\n\npackage tuple\n"); err != nil {
		panic(err)
	}

	for n := 0; n < 32; n++ {
		typeName := "Tuple" + strconv.Itoa(n+1)
		funcName := "New" + strconv.Itoa(n+1)

		var fields string
		var generics string
		var params string
		var args string
		var assign string
		var accesses string
		var returns string
		for index := 0; index < n+1; index++ {
			upTName := "T" + strconv.Itoa(index)
			upVName := "V" + strconv.Itoa(index)
			downVName := "v" + strconv.Itoa(index)

			fields += upVName + " " + upTName
			generics += upTName
			params += upTName
			args += downVName + " " + upTName
			assign += "\n" + upVName + ": " + downVName
			accesses += "t." + upVName
			returns += upTName

			if index != n {
				fields += ";"
				generics += ","
				params += ","
				args += ", "
				assign += ","
				accesses += ","
				returns += ","
			}

			if index == n {
				generics += " any"
			}
		}

		if _, err := fmt.Fprintf(file, `
type %[1]s[%[2]s] struct {
	%[3]s
}

func %[9]s[%[2]s](%[5]s) %[1]s[%[4]s] {
	return %[1]s[%[4]s]{%[6]s`+",\n"+`}
}

func (t %[1]s[%[4]s]) Unpack() (%[7]s) {
	return %[8]s
}
`, typeName, generics, fields, params, args, assign, returns, accesses, funcName); err != nil {
			panic(err)
		}
	}

	cmd := exec.Command("gotip", "fmt", "./tuple")
	cmd.Stdout = os.Stdout
	if err := cmd.Run(); err != nil {
		panic(err)
	}
}
