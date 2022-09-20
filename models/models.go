package models

type Field struct {
	ID   int
	Name string
}

type RuleRoute struct {
	RuleID        int
	OriginID      int
	DestinationID int
}

type RuleAirline struct {
	RuleID    int
	AirlineID int
}

type RuleAgency struct {
	RuleID   int
	AgencyID int
}

type RuleSupplier struct {
	RuleID     int
	SupplierID int
}

type Rule struct {
	ID            int
	RuleRoutes    []RuleRoute
	RuleAirlines  []RuleAirline
	RuleAgencies  []RuleAgency
	RuleSuppliers []RuleSupplier
	IsPercent     bool
	Amount        int
}

type Ticket struct {
	OriginID      string
	DestinationID string
	AirlineID     string
	AgencyID      string
	SupplierID    string
}

type RawRule struct {
	Routes []struct {
		Origin      string `json:"origin"`
		Destination string `json:"destination"`
	} `json:"routes"`
	Airlines    []string `json:"airlines"`
	Agencies    []string `json:"agencies"`
	Suppliers   []string `json:"suppliers"`
	AmountType  string   `json:"amountType"`
	AmountValue int      `json:"amountValue"`
}

type RawTicket struct {
	RuleID       int    `json:"ruleId"`
	Origin       string `json:"origin"`
	Destination  string `json:"destination"`
	Airline      string `json:"airline"`
	Agency       string `json:"agency"`
	Supplier     string `json:"supplier"`
	BasePrice    int    `json:"basePrice"`
	Markup       int    `json:"markup"`
	PayablePrice int    `json:"payablePrice"`
}

type Response struct {
	OK      bool   `json:"ok"`
	Message string `json:"message"`
}
