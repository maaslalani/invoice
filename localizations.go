package invoiceservice

type LocalizationServiceImpl struct {
	language string
}

type LocalizationService interface {
	serviceDate() string
	billTo() string
	note() string
	from() string
	item() string
	qty() string
	rate() string
	amount() string
	subtotal() string
	discount() string
	taxLabelATInclVat() string
	total() string
	dueDate() string
}

func (l *LocalizationServiceImpl) serviceDate() string {
	switch l.language {
	case "de":
		return "Leistungsdatum"
	}

	return "Service Date"
}

func (l *LocalizationServiceImpl) billTo() string {
	switch l.language {
	case "de":
		return "AUSGESTELLT AN"
	}

	return "BILL TO"
}

func (l *LocalizationServiceImpl) note() string {
	switch l.language {
	case "de":
		return "Hinweis"
	}

	return "Note"
}

func (l *LocalizationServiceImpl) from() string {
	switch l.language {
	case "de":
		return "Ausgestellt von"
	}

	return "From"
}

func (l *LocalizationServiceImpl) item() string {
	switch l.language {
	case "de":
		return "BESCHREIBUNG"
	}

	return "ITEM"
}

func (l *LocalizationServiceImpl) qty() string {
	switch l.language {
	case "de":
		return "MENGE"
	}

	return "QTY"
}

func (l *LocalizationServiceImpl) rate() string {
	switch l.language {
	case "de":
		return "PREIS"
	}

	return "RATE"
}

func (l *LocalizationServiceImpl) amount() string {
	switch l.language {
	case "de":
		return "BETRAG"
	}

	return "AMOUNT"
}

func (l *LocalizationServiceImpl) subtotal() string {
	switch l.language {
	case "de":
		return "Zwischensumme"
	}

	return "Subtotal"
}

func (l *LocalizationServiceImpl) discount() string {
	switch l.language {
	case "de":
		return "Ermäßigung"
	}

	return "Discount"
}

func (l *LocalizationServiceImpl) taxLabelATInclVat() string {
	switch l.language {
	case "de":
		return "Inkl. 20% Ust."
	}

	return "Incl. 20% VAT"
}

func (l *LocalizationServiceImpl) total() string {
	switch l.language {
	case "de":
		return "Gesamtbetrag"
	}

	return "Total"
}

func (l *LocalizationServiceImpl) dueDate() string {
	switch l.language {
	case "de":
		return "Fällig am"
	}

	return "Due Date"
}
