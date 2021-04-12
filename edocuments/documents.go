package edocuments

type DocumentFetchQuery struct {
	Type   string `json:"type"`
	Series string `json:"series"`
	Number string `json:"number"`
}

type Document struct {
	Type             string  `json:"type"` // FT ; FR ; FS  ; DI  ;
	Serie            string  `json:"serie"`
	Number           string  `json:"number"`
	UserId           int     `json:"userId"`
	Observations     *string `json:"observations"`
	ClientId         string  `json:"clientId"`
	ClientName       string  `json:"clientName"`
	ClientAddress    string  `json:"clientAddress"`
	ClientVat        string  `json:"clientVat"`
	ClientPostalCode string  `json:"clientPostalCode"`
	ClientCity       string  `json:"clientCity"`
	UniqueID         string  `json:"uniqueID"`
	Total            float32 `json:"total"`
	SubTotal         float32 `json:"subTotal"`
	Rows             []struct {
		ItemKeyId       string  `json:"itemKeyId"`
		ItemDescription string  `json:"itemDescription"`
		Quantity        string  `json:"quantity"`
		UnitPrice       float32 `json:"unitPrice"`
		IncludesTaxes   bool    `json:"includesTaxes"`
		TaxType         string  `json:"TaxType"` // ISE ; NOR  ; RED  ; INT
		TaxValue        float32 `json:"taxValue"`
	}
}

type ItemFetchQuery struct {
	Barcode string `json:"barcode"`
}

type Item struct {
	Ref            string
	Stock          float32
	Description    string
	Barcode        string
	AlternativeRef string
	Attributes     []ItemAttribute
}

type ItemAttribute struct {
	AtrributeId      int
	AttributeValueId int
	Name             string
	AttributeName    string
}
