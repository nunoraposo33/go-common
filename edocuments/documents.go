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
	Ref            string          `json:"ref"`
	Stock          float32         `json:"stock"`
	Description    string          `json:"description"`
	Barcode        string          `json:"barcode"`
	AlternativeRef string          `json:"alternativeRef"`
	ControlStock   bool            `json:"controlStock"`
	Attributes     []ItemAttribute `json:"attributes"`
}

type ItemAttribute struct {
	AtrributeId      int    `json:"atrributeId"`
	AttributeValueId int    `json:"attributeValueId"`
	Name             string `json:"name"`
	AttributeName    string `json:"attributeName"`
}
