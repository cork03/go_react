package domain

type Company struct {
	Name       string `json:"name"`
	PostalCode string `json:"postal_code"`
	Prefecture string `json:"prefecture"`
	Town       string `json:"town"`
	Area       string `json:"area"`
	Tel        string `json:"tel"`
}
