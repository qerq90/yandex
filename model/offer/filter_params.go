package offer

type Params struct {
	DateFrom   string   `json:"dateFrom,omitempty"`
	DateTo     string   `json:"dateTo,omitempty"`
	UpdateFrom string   `json:"updateFrom,omitempty"`
	UpdateTo   string   `json:"updateTo,omitempty"`
	Orders     []int    `json:"orders,omitempty"`
	Statuses   []string `json:"statuses,omitempty"`
	HasCIS     bool     `json:"hasCis,omitempty"`
}
