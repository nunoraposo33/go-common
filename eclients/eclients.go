package eclients

// Entity represents a client
type Entity struct {
	ID           int64   `json:"id"`
	ExternalKey  *string `json:"external_key"`
	Name         *string `json:"name"`
	Vat          *string `json:"vat"`
	Address      *string `json:"address"`
	PostalCode   *string `json:"postal_code"`
	City         *string `json:"city"`
	Country      *string `json:"country"`
	Phone        *string `json:"phone"`
	MobilePhone  *string `json:"mobile_phone"`
	Email        *string `json:"email"`
	ContactName  *string `json:"contact_name"`
	ContactPhone *string `json:"contact_phone"`
	ContactEmail *string `json:"contact_email"`
	Homepage     *string `json:"homepage"`
	Inactive     *bool   `json:"inactive"`
	CreatedAt    string  `json:"created_at"`
	UpdatedAt    string  `json:"updated_at"`
	DeletedAt    *string `json:"deleted_at"`
}
