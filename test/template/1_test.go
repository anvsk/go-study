package template

import (
	"bytes"
	"log"

	"text/template"
	// "github.com/alecthomas/template"
)

func ExampleTemplate() {
	// Define a template.
	const letter2 = `
Dear {{.Name}},
{{if .Attended}}
It was a pleasure to see you at the wedding.{{else}}
It is a shame you couldn't make it to the wedding.{{end}}
{{with .Gift}}Thank you for the lovely {{.}}.
{{end}}
Best wishes,
Josie
`

	const letter = `
{{if true}}
hello
{{end}}
{{if true}}
world
{{end}}
	`

	// Prepare some data to insert into the template.
	type Recipient struct {
		Name, Gift string
		Attended   bool
	}
	var recipients = []Recipient{
		{"Aunt Mildred", "bone china tea set", true},
		{"Uncle John", "moleskin pants", false},
		{"Cousin Rodney", "", false},
	}

	// Create a new template and parse the letter into it.
	t := template.Must(template.New("letter").Parse(letter))
	var b = new(bytes.Buffer)
	// buf := io.ByteReader
	// Execute the template for each recipient.
	for _, r := range recipients {
		err := t.Execute(b, r)
		if err != nil {
			log.Println("executing template:", err)
		}
	}

	log.Print(string(b.Bytes()))

	// Output:
	// Dear Aunt Mildred,
	//
	// It was a pleasure to see you at the wedding.
	// Thank you for the lovely bone china tea set.
	//
	// Best wishes,
	// Josie
	//
	// Dear Uncle John,
	//
	// It is a shame you couldn't make it to the wedding.
	// Thank you for the lovely moleskin pants.
	//
	// Best wishes,
	// Josie
	//
	// Dear Cousin Rodney,
	//
	// It is a shame you couldn't make it to the wedding.
	//
	// Best wishes,
	// Josie
}
