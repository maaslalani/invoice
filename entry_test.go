package invoiceservice

import (
	"testing"
)

func TestAll(t *testing.T) {
	output = "invoice"

	// Kleinbetragsrechnung example AT

	defaultInvoice.Id = "12345"
	defaultInvoice.Title = "Rechnungsbeleg"
	defaultInvoice.From = "CourseStart LLC"
	defaultInvoice.To = "Josef Büttgen\nHarbigstraße 61\n1A\n71032 Böblingen, DE\nVATIN: 12345"
	defaultInvoice.Items = []string{
		"25.11. CourseStart Turnier",
	}
	defaultInvoice.Date = "6.11.2023"
	defaultInvoice.PerformanceDate = "8.11.2023"
	defaultInvoice.Due = "10.11.2023"
	defaultInvoice.Quantities = []int{
		1,
	}
	defaultInvoice.Discount = 0
	defaultInvoice.Currency = "EUR"
	defaultInvoice.Tax = 0.2
	defaultInvoice.IncludeTax = true
	defaultInvoice.Note = "CourseStart LLC\nSome Street 123\n12345 Some Town\nSome Country"
	defaultInvoice.Comment = "It is the responsibility of the customer to determine the correct local treatment in respect of the receipt of these services\nincluding any reverse charge considerations."
	defaultInvoice.Language = "de"

	GenerateInvoice(defaultInvoice, nil)
}
