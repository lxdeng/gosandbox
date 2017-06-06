package main

import (
	"fmt"
	"log"
	"os"
	"text/template"
	"time"
)

type User struct {
	User string
}

/*
 template is a string or file containing one or more portions enclosed in double braces, {{…}}, called actions. These actions are processed by the template engine.
*/
func main() {
	processFields()
	processFields2()
	processMultipleFiles()
	processBuiltInFuncPrintf()
	processUserDefinedFunc()
}

func formatDate(layout string, info string, d time.Time) string {
	return d.Format(layout) + " " + info
}

func processUserDefinedFunc() {
	log.Println("processUserDefinedFunc")

	var ts1 = `
This is a test
date: {{.Date | dateFormat "Jan 2, 2006" "Moon"}}`

	funcMap := template.FuncMap{
		"dateFormat": formatDate,
	}

	d := struct{ Date time.Time }{
		Date: time.Now(),
	}

	t := template.New("date")
	t.Funcs(funcMap)
	t.Parse(ts1)
	t.Execute(os.Stdout, d)
	fmt.Printf("\n\n")
}

func processBuiltInFuncPrintf() {
	log.Println("processBuiltInFuncPrintf")
	t := template.New("test")

	// calls printf to format 8 as binary number
	t = template.Must(t.Parse("test number: {{8 | printf \"%b\"}} "))

	t.Execute(os.Stdout, nil)
	fmt.Printf("\n\n")
}

func processFields() {
	log.Println("processFields")
	t := template.Must(template.ParseFiles("./templates/sample1.txt"))
	fmt.Printf("t.Name=%v\n", t.Name())

	t.Execute(os.Stdout, nil)

	p := User{"jack"}
	t.Execute(os.Stdout, p)

	tt := t.Lookup("sample1.txt")
	fmt.Printf("t.Lookup=%v\n\n", tt)
}

func processFields2() {
	log.Println("processFields2")
	t := template.New("t2")
	fmt.Printf("t.Name=%v\n", t.Name())

	tt := t.Lookup("t2") //nil, no definition
	fmt.Printf("t.Lookup=%v\n", tt)

	t = template.Must(t.Parse("Hi {{.User}}"))
	tt = t.Lookup("t2") //nil
	fmt.Printf("t.Lookup=%v\n", tt)

	p := User{"Tom"}
	t.Execute(os.Stdout, p)
	fmt.Printf("\n\n")
}

func processMultipleFiles() {
	log.Println("processMultipleFiles")
	t := template.Must(template.ParseFiles("./templates2/sample1.txt", "./templates2/sample2.txt"))
	fmt.Printf("t.Name=%v\n", t.Name())

	p := User{"jack"}
	t.Execute(os.Stdout, p)

	tt := t.Lookup("sample1.txt")
	fmt.Printf("t.Lookup=%v\n\n", tt)
}
