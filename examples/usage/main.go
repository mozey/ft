package main

import (
	"encoding/json"
	"fmt"
	"os"
	"text/template"

	"github.com/mozey/ft"
)

type Data struct {
	Required ft.String  `json:"required"`
	Optional ft.NString `json:"optional"`
}

// Valid validates all the required properties in the Data type
func (d *Data) Valid() bool {
	return d.Required.String == "123"
}
func main() {
	b := []byte(`{"required": 123,"optional": null}`)
	d := Data{}
	err := json.Unmarshal(b, &d)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	// ft.String doesn't have the Valid property
	fmt.Println(d.Required.String) // "123"
	fmt.Println(d.Optional.Valid)  // false
	fmt.Println("Optional string is empty ", d.Optional.String)

	if d.Valid() {
		fmt.Println("The data is valid")
	}

	t1 := template.Must(template.New("t1").Parse(
		`{{if not .Optional.Valid }}Required = {{.Required.String}} {{end}}
`))
	t1.Execute(os.Stdout, d) // Required = 123

	os.Exit(0)
}
