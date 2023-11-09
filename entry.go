package invoiceservice

import (
	_ "embed"
	"fmt"
	"io"
	"strings"
	"time"

	"github.com/signintech/gopdf"
	"github.com/spf13/cobra"
)

//go:embed "Inter/Inter Variable/Inter.ttf"
var interFont []byte

//go:embed "Inter/Inter Hinted for Windows/Desktop/Inter-Bold.ttf"
var interBoldFont []byte

type Invoice struct {
	Id    string `json:"id" yaml:"id"`
	Title string `json:"title" yaml:"title"`

	Logo            string `json:"logo" yaml:"logo"`
	From            string `json:"from" yaml:"from"`
	To              string `json:"to" yaml:"to"`
	Date            string `json:"date" yaml:"date"`
	PerformanceDate string `json:"performance_date" yaml:"performance_date"`
	Due             string `json:"due" yaml:"due"`

	Items      []string  `json:"items" yaml:"items"`
	Quantities []int     `json:"quantities" yaml:"quantities"`
	Rates      []float64 `json:"rates" yaml:"rates"`

	Tax        float64 `json:"tax" yaml:"tax"`
	IncludeTax bool    `json:"include_tax" yaml:"include_tax"`
	Discount   float64 `json:"discount" yaml:"discount"`
	Currency   string  `json:"currency" yaml:"currency"`

	Note    string `json:"note" yaml:"note"`
	Comment string `json:"comment" yaml:"comment"`

	Language string `json:"language" yaml:"language"`
}

func DefaultInvoice() Invoice {
	return Invoice{
		Id:              time.Now().Format("20060102"),
		Title:           "INVOICE",
		Rates:           []float64{25},
		Quantities:      []int{2},
		Items:           []string{"Paper Cranes"},
		From:            "Project Folded, Inc.",
		To:              "Untitled Corporation, Inc.",
		Date:            time.Now().Format("Jan 02, 2006"),
		PerformanceDate: time.Now().Format("Jan 02, 2006"),
		Due:             time.Now().AddDate(0, 0, 14).Format("Jan 02, 2006"),
		Tax:             0,
		IncludeTax:      false,
		Discount:        0,
		Currency:        "USD",
		Language:        "en",
	}
}

var (
	importPath     string
	output         string
	file           = Invoice{}
	defaultInvoice = DefaultInvoice()
)

var rootCmd = &cobra.Command{
	Use:   "invoice",
	Short: "Invoice generates invoices from the command line.",
	Long:  `Invoice generates invoices from the command line.`,
}

// var generateCmd = &cobra.Command{
// 	Use:   "generate",
// 	Short: "Generate an invoice",
// 	Long:  `Generate an invoice`,
// 	RunE: func(cmd *cobra.Command, args []string) error {
// 		// Read CLI params / JSON config into 'file'
// 		if importPath != "" {
// 			err := importData(importPath, &file, cmd.Flags())
// 			if err != nil {
// 				return err
// 			}
// 		}

// 		err := GenerateInvoice(file)
// 		if err != nil {
// 			return err
// 		}

// 		return nil
// 	},
// }

func GenerateInvoice(file Invoice, outputWriter io.Writer) (err error) {
	pdf := gopdf.GoPdf{}
	pdf.Start(gopdf.Config{
		PageSize: *gopdf.PageSizeA4,
	})
	pdf.SetMargins(40, 40, 40, 40)
	pdf.AddPage()
	err = pdf.AddTTFFontData("Inter", interFont)
	if err != nil {
		return err
	}

	err = pdf.AddTTFFontData("Inter-Bold", interBoldFont)
	if err != nil {
		return err
	}

	// Init localization service for language translations

	localizationService := &LocalizationServiceImpl{
		language: file.Language,
	}

	writeLogo(&pdf, file.Logo, file.From)
	writeTitle(&pdf, file.Title, file.Id, file.Date, file.PerformanceDate, localizationService)
	writeBillTo(&pdf, file.To, localizationService)
	writeHeaderRow(&pdf, localizationService)
	subtotal := 0.0
	for i := range file.Items {
		q := 1
		if len(file.Quantities) > i {
			q = file.Quantities[i]
		}

		r := 0.0
		if len(file.Rates) > i {
			r = file.Rates[i]
		}

		writeRow(&pdf, file.Items[i], q, r, file.Currency)
		subtotal += float64(q) * r
	}
	if file.Comment != "" {
		writeComment(&pdf, file.Comment, localizationService)
	}
	if file.Note != "" {
		writeNotes(&pdf, file.Note, localizationService)
	}
	var tax float64
	if file.IncludeTax {
		tax = subtotal * (1 - (1 / (1 + file.Tax)))
	} else {
		tax = subtotal * file.Tax
	}
	writeTotals(&pdf, subtotal, tax, file.IncludeTax, subtotal*file.Discount, file.Currency, localizationService)
	if file.Due != "" {
		writeDueDate(&pdf, file.Due, localizationService)
	}
	writeFooter(&pdf, file.Id)

	// Write either to io.Writer or file system based on whether a writer was provided
	if outputWriter != nil {
		_, err = pdf.WriteTo(outputWriter)
		if err != nil {
			return err
		}
	} else {
		output = strings.TrimSuffix(output, ".pdf") + ".pdf"
		err = pdf.WritePdf(output)
		if err != nil {
			return err
		}
	}

	fmt.Printf("Generated %s\n", output)

	return nil
}

// func main() {
// 	rootCmd.AddCommand(generateCmd)
// 	err := rootCmd.Execute()
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// }
