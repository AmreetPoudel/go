package template

import (
	"os"
	"text/template"
)

type item struct {
	Name  string
	Qty   int
	Price float64
}

type invoice struct {
	Customer string
	Items    []item
	Total    float64
}

const InvoiceTemplate = `
Invoice For: {{.Customer}}


Items: 
{{range .Items }} {{.Name}} Qty: {{.Qty}} ${{.Price}} each
{{end}}
Total: {{.Total}}
`

func main() {

	//lets make all data to fed to template so we will fill the above invoice struct
	data := invoice{
		Customer: "Amrit Poudel",
		Items: []item{
			{"1TB SSD", 2, 150},
			{"Macbook m4", 1, 2000},
			{"Airpod", 1, 500},
			{"apple watch", 1, 450},
		},
		Total: 2*150 + 1*2000 + 1*500 + 1*450,
	}
	// parse the templates string
	tmp1, err := template.New("invoice").Parse(InvoiceTemplate)
	if err != nil {
		panic(err)
	}
	err = tmp1.Execute(os.Stdout, data)
	if err != nil {
		panic(err)
	}

}
