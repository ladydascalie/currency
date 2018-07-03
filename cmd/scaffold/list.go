package scaffold

type Entry struct {
	Code        string `xml:" Ccy,omitempty" json:"AlphanumericCode,omitempty"`
	MinorUnits  string `xml:" CcyMnrUnts,omitempty" json:"MinorUnits,omitempty"`
	Country     string `xml:" CtryNm,omitempty" json:"CountryName,omitempty"`
	Description string `xml:" CcyNm,omitempty" json:"CurrencyName,omitempty"`
}

type Table struct {
	Entries []*Entry `xml:" CcyNtry,omitempty" json:"Entries,omitempty"`
}

type ISO4217 struct {
	AttrPublished string `xml:"Pblshd,attr"  json:",omitempty"` // maxLength=10
	Table         *Table `xml:" CcyTbl,omitempty" json:"Table,omitempty"`
}
