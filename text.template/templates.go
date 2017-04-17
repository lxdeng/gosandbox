package main

import (
	"fmt"
	"os"
	"text/template"
)

type User struct {
	User string
}

/*
 template is a string or file containing one or more portions enclosed in double braces, {{…}}, called actions. These actions are processed by the template engine
*/
func main() {
	t := template.Must(template.ParseFiles("./templates/sample1.txt"))
	t.Execute(os.Stdout, nil)

	p := User{"jack"}
	t.Execute(os.Stdout, p)

	t2 := t.Lookup("sample1.txt")
	fmt.Println(t2)
}
