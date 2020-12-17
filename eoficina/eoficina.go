package eoficina

import "github.com/nunoraposo33/go-eclients-common/eclients"

// Repair Struct - used for creating new repairs
type Repair struct {
	ID            int64            `json:"id"`
	Client        eclients.Entity  `json:"client"`
	InvoiceClient *eclients.Entity `json:"invoice_client"`
	Vehicle       Vehicle          `json:"vehicle"`
	Options       []RepairOptions  `json:"options"`
	StateID       *int64           `json:"state_id"`
	State         *string          `json:"state"`
	CategoryID    *int64           `json:"category_id"`
	Category      *string          `json:"category"`
	RepairTypeID  *int64           `json:"type_id"`
	RepairType    *string          `json:"type"`
	PriorityID    *int64           `json:"priority_id"`
	Priority      *string          `json:"priority"`
	EntryDate     string           `json:"entry_date"`
	EntryKM       *int64           `json:"entry_km"`
	ExitDate      *string          `json:"exit_date"`
	ExitKM        *int64           `json:"exit_km"`
	CloseDate     *string          `json:"close_date"`
	Closed        bool             `json:"closed"`
	Repaired      bool             `json:"repaired"`
	OutOfGarage   bool             `json:"out_of_garage"`
	Invoiced      bool             `json:"invoiced"`
	IPO           *bool            `json:"ipo"`
	FuelLevel     *string          `json:"fuel_level"`
	ClientAbsent  bool             `json:"client_absent"`
	Todo          string           `json:"todo"`
	Observations  string           `json:"observations"`
	Externalkey   *string          `json:"integration"`
	Signature     *string          `json:"signature"`
	Damage        *string          `json:"damage"`
	UpdatedAt     string           `json:"updated_at"`
	DeletedAt     *string          `json:"deleted_at"`
	CreatedAt     string           `json:"created_at"`
}

// RepairOptions is used to describe selected options (checkboxes) in a repair
type RepairOptions struct {
	ID                int64  `json:"id"`
	OptionID          int64  `json:"option_id"`
	OptionDescription string `json:"description"`
	Value             string `json:"value"`
}

// Vehicle represents a vehicle in a repair
type Vehicle struct {
	ID                        int64   `json:"id"`
	ExternalKey               *string `json:"external_key"`
	Plate                     *string `json:"plate"`
	BrandID                   *int64  `json:"brand_id"`
	Brand                     *string `json:"brand"`
	ModelID                   *int64  `json:"model_id"`
	Model                     *string `json:"model"`
	VIN                       *string `json:"vin"`
	PlateDate                 *string `json:"plate_date"`
	InspectionDate            *string `json:"next_periodic_inspection"`
	Version                   *string `json:"version"`
	NextRevision              *string `json:"next_revision"`
	Kms                       *int64  `json:"kms"`
	FuelTypeID                *int64  `json:"fuel_type"`
	Fuel                      *string `json:"fuel"`
	Displacement              *string `json:"displacement"`
	Cilinders                 *string `json:"cilinders"`
	HorsePower                *int64  `json:"horse_power"`
	PowerKW                   *string `json:"power_kw"`
	Doors                     *int64  `json:"n_doors"`
	Seats                     *int64  `json:"n_seats"`
	WarantyEnd                *string `json:"waranty_end"`
	Observations              *string `json:"observations"`
	Color                     *string `json:"color"`
	ColorCode                 *string `json:"color_code"`
	Co2                       *string `json:"co2"`
	GrossWeight               *string `json:"gross_weight"`
	KeyCode                   *string `json:"key_code"`
	RadioCode                 *string `json:"radio_code"`
	TiresFrontSize            *string `json:"tires_size_front"`
	TiresBackSize             *string `json:"tires_size_back"`
	EngineCode                *string `json:"engine_code"`
	VehicleTypeID             *int64  `json:"vehicle_type_id"`
	VehicleType               *string `json:"vehicle_type"`
	VehicleFrameid            *int64  `json:"frame_type_id"`
	VehicleFrame              *string `json:"frame_type"`
	VehicleTransmissionTypeID *int64  `json:"transmission_type_id"`
	VehicleTransmission       *string `json:"transmission_type"`
	ClientID                  *int64  `json:"client_id"`
	CreatedAt                 *string `json:"updated_at"`
	UpdatedAt                 *string `json:"deleted_at"`
	DeletedAt                 *string `json:"created_at"`
}
