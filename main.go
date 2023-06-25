package main

import (
	_ "embed"
	"flag"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/signintech/gopdf"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

//go:embed Inter.ttf
var interFont []byte

//go:embed Inter-Bold.ttf
var interBoldFont []byte

type invoiceData struct {
	Id    string `json:"id" yaml:"id"`
	Title string `json:"title" yaml:"title"`

	Logo string `json:"logo" yaml:"logo"`
	From string `json:"from" yaml:"from"`
	To   string `json:"to" yaml:"to"`
	Date string `json:"date" yaml:"date"`
	Due  string `json:"due" yaml:"due"`

	Items      []string  `json:"items" yaml:"items"`
	Quantities []int     `json:"quantities" yaml:"quantities"`
	Rates      []float64 `json:"rates" yaml:"rates"`

	Tax      float64 `json:"tax" yaml:"tax"`
	Discount float64 `json:"discount" yaml:"discount"`
	Currency string  `json:"currency" yaml:"currency"`

	Note string `json:"note" yaml:"note"`
}

var (
	importPath string
	output     string
	file       invoiceData
)

func init() {
	viper.AutomaticEnv()

	generateCmd.Flags().StringVar(&file.Id, "id", time.Now().Format("20060102"), "ID")
	generateCmd.Flags().StringVar(&file.Title, "title", "INVOICE", "Title")

	generateCmd.Flags().Float64SliceVarP(&file.Rates, "rate", "r", []float64{25}, "Rates")
	generateCmd.Flags().IntSliceVarP(&file.Quantities, "quantity", "q", []int{2}, "Quantities")
	generateCmd.Flags().StringSliceVarP(&file.Items, "item", "i", []string{"Paper Cranes"}, "Items")

	generateCmd.Flags().StringVarP(&file.Logo, "logo", "l", "", "Company logo")
	generateCmd.Flags().StringVarP(&file.From, "from", "f", "Project Folded, Inc.", "Issuing company")
	generateCmd.Flags().StringVarP(&file.To, "to", "t", "Untitled Corporation, Inc.", "Recipient company")
	generateCmd.Flags().StringVar(&file.Date, "date", time.Now().Format("Jan 02, 2006"), "Date")
	generateCmd.Flags().StringVar(&file.Due, "due", time.Now().AddDate(0, 0, 14).Format("Jan 02, 2006"), "Payment due date")

	generateCmd.Flags().Float64Var(&file.Tax, "tax", 0, "Tax")
	generateCmd.Flags().Float64VarP(&file.Discount, "discount", "d", 0.0, "Discount")
	generateCmd.Flags().StringVarP(&file.Currency, "currency", "c", "USD", "Currency")

	generateCmd.Flags().StringVarP(&file.Note, "note", "n", "", "Note")
	generateCmd.Flags().StringVarP(&output, "output", "o", "invoice.pdf", "Output file (.pdf)")
	generateCmd.Flags().StringVar(&importPath, "import", "", "Imported file (.json/.yaml)")

	flag.Parse()
}

var rootCmd = &cobra.Command{
	Use:   "invoice",
	Short: "Invoice generates invoices from the command line.",
	Long:  `Invoice generates invoices from the command line.`,
}

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate an invoice",
	Long:  `Generate an invoice`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if importPath != "" {
			err := importData(importPath, &file)
			if err != nil {
				return err
			}
		}
		return createPdf()
	},
}

func createPdf() error {
	pdf := gopdf.GoPdf{}
	pdf.Start(gopdf.Config{
		PageSize: *gopdf.PageSizeA4,
	})
	pdf.SetMargins(40, 40, 40, 40)
	pdf.AddPage()
	err := pdf.AddTTFFontData("Inter", interFont)
	if err != nil {
		return err
	}

	err = pdf.AddTTFFontData("Inter-Bold", interBoldFont)
	if err != nil {
		return err
	}

	writeLogo(&pdf, file.Logo, file.From)
	writeTitle(&pdf, file.Title, file.Id)
	writeBillTo(&pdf, file.To)
	writeHeaderRow(&pdf)
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

		writeRow(&pdf, file.Items[i], q, r)
		subtotal += float64(q) * r
	}
	if file.Note != "" {
		writeNotes(&pdf, file.Note)
	}
	writeTotals(&pdf, subtotal, subtotal*file.Tax, subtotal*file.Discount)
	writeFooter(&pdf, file.Id)
	output = strings.TrimSuffix(output, ".pdf") + ".pdf"
	err = pdf.WritePdf(output)
	if err != nil {
		return err
	}

	fmt.Printf("Generated %s\n", output)

	return nil

}

func main() {
	rootCmd.AddCommand(generateCmd)
	err := rootCmd.Execute()
	if err != nil {
		log.Fatal(err)
	}
}
