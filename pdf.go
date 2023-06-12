package main

import (
	"strconv"

	"github.com/signintech/gopdf"
)

const (
	quantityColumnOffset = 360
	rateColumnOffset     = 405
	amountColumnOffset   = 480
)

const (
	subtotalLabel = "Subtotal"
	discountLabel = "Discount"
	taxLabel      = "Tax"
	totalLabel    = "Total"
)

func writeLogo(pdf *gopdf.GoPdf, logo string, from string) {
	if logo != "" {
		_ = pdf.Image(logo, pdf.GetX(), pdf.GetY(), &gopdf.Rect{W: 44, H: 44})
		pdf.Br(64)
	}
	_ = pdf.SetFont("Inter", "", 12)
	pdf.SetTextColor(55, 55, 55)
	_ = pdf.Cell(nil, from)
	pdf.Br(36)
	pdf.SetStrokeColor(225, 225, 225)
	pdf.Line(pdf.GetX(), pdf.GetY(), 100, pdf.GetY())
	pdf.Br(36)
}

func writeTitle(pdf *gopdf.GoPdf, title, id string) {
	_ = pdf.SetFont("Inter-Bold", "", 24)
	pdf.SetTextColor(0, 0, 0)
	_ = pdf.Cell(nil, title)
	pdf.Br(36)
	_ = pdf.SetFont("Inter", "", 12)
	pdf.SetTextColor(100, 100, 100)
	_ = pdf.Cell(nil, "#")
	pdf.SetTextColor(100, 100, 100)
	_ = pdf.Cell(nil, id)
	pdf.Br(48)
}

func writeBillTo(pdf *gopdf.GoPdf, to string) {
	pdf.SetTextColor(75, 75, 75)
	_ = pdf.SetFont("Inter", "", 9)
	_ = pdf.Cell(nil, "BILL TO")
	pdf.Br(14)
	pdf.SetTextColor(75, 75, 75)
	_ = pdf.SetFont("Inter", "", 15)
	_ = pdf.Cell(nil, to)
	pdf.Br(64)
}

func writeHeaderRow(pdf *gopdf.GoPdf) {
	_ = pdf.SetFont("Inter", "", 9)
	pdf.SetTextColor(55, 55, 55)
	_ = pdf.Cell(nil, "ITEM")
	pdf.SetX(quantityColumnOffset)
	_ = pdf.Cell(nil, "QTY")
	pdf.SetX(rateColumnOffset)
	_ = pdf.Cell(nil, "RATE")
	pdf.SetX(amountColumnOffset)
	_ = pdf.Cell(nil, "AMOUNT")
	pdf.Br(24)
}

func writeNotes(pdf *gopdf.GoPdf, notes string) {
	pdf.SetY(650)

	_ = pdf.SetFont("Inter", "", 10)
	pdf.SetTextColor(55, 55, 55)
	_ = pdf.Cell(nil, "Notes")
	pdf.Br(18)
	pdf.SetTextColor(0, 0, 0)
	_ = pdf.Cell(nil, notes)
	pdf.Br(48)
}

func writeFooter(pdf *gopdf.GoPdf, id string) {
	pdf.SetY(800)

	_ = pdf.SetFont("Inter", "", 10)
	pdf.SetTextColor(55, 55, 55)
	_ = pdf.Cell(nil, id)
	pdf.SetStrokeColor(225, 225, 225)
	pdf.Line(pdf.GetX()+10, pdf.GetY()+6, 550, pdf.GetY()+6)
	pdf.Br(48)
}

func writeRow(pdf *gopdf.GoPdf, item string, quantity int, rate float64) {
	_ = pdf.SetFont("Inter", "", 11)
	pdf.SetTextColor(0, 0, 0)

	total := float64(quantity) * rate
	amount := strconv.FormatFloat(total, 'f', 2, 64)

	_ = pdf.Cell(nil, item)
	pdf.SetX(quantityColumnOffset)
	_ = pdf.Cell(nil, strconv.Itoa(quantity))
	pdf.SetX(rateColumnOffset)
	_ = pdf.Cell(nil, currencySymbols[currency]+strconv.FormatFloat(rate, 'f', 2, 64))
	pdf.SetX(amountColumnOffset)
	_ = pdf.Cell(nil, currencySymbols[currency]+amount)
	pdf.Br(24)
}

func writeTotals(pdf *gopdf.GoPdf, subtotal float64, tax float64, discount float64) {
	pdf.SetY(650)

	writeTotal(pdf, subtotalLabel, subtotal)
	if tax > 0 {
		writeTotal(pdf, taxLabel, tax)
	}
	if discount > 0 {
		writeTotal(pdf, discountLabel, discount)
	}
	writeTotal(pdf, totalLabel, subtotal+tax-discount)
}

func writeTotal(pdf *gopdf.GoPdf, label string, total float64) {
	_ = pdf.SetFont("Inter", "", 9)
	pdf.SetTextColor(75, 75, 75)
	pdf.SetX(rateColumnOffset)
	_ = pdf.Cell(nil, label)
	pdf.SetTextColor(0, 0, 0)
	_ = pdf.SetFontSize(12)
	pdf.SetX(amountColumnOffset - 15)
	if label == totalLabel {
		_ = pdf.SetFont("Inter-Bold", "", 11.5)
	}
	_ = pdf.Cell(nil, currencySymbols[currency]+strconv.FormatFloat(total, 'f', 2, 64))
	pdf.Br(24)
}
