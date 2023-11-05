package invoiceservice

import (
	"testing"
)

func TestAll(t *testing.T) {
	output = "invoice"

	// Kleinbetragsrechnung example AT

	defaultInvoice.Id = ""
	defaultInvoice.From = "Josef Büttgen"
	defaultInvoice.To = "Josef Büttgen"
	defaultInvoice.Currency = "EUR"
	defaultInvoice.Tax = 0.2
	defaultInvoice.IncludeTax = true
	defaultInvoice.To = ""
	defaultInvoice.Note = "CourseStart LLC\nJosef Büttgen\nHarbigstraße 61\n71032 Böblingen\nDeutschland"

	GenerateInvoice(defaultInvoice)
}
