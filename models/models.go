// models.go
package models

// Struct untuk response API
type Venue struct {
	ID              int      `json:"ID"`
	ParentID        int      `json:"ParentID"`
	MaxClubMember   int      `json:"MaxClubMember"`
	Name            string   `json:"Name"`
	UnitNumber      string   `json:"UnitNumber"`
	Large           string   `json:"Large"`
	Capacity        string   `json:"Capacity"`
	CapacityVisitor int      `json:"CapacityVisitor"`
	PhoneVenue      string   `json:"PhoneVenue"`
	Description     string   `json:"Description"`
	PrimaryImage    string   `json:"PrimaryImage"`
	IsActive        int      `json:"IsActive"`
	Status          int      `json:"Status"`
	UnitName        string   `json:"UnitName"`
	UnitSimpleName  string   `json:"UnitSimpleName"`
	UnitSheetName   string   `json:"UnitSheetName"`
	CategoryName    string   `json:"CategoryName"`
	Gallery         []string `json:"Gallery"`
	Rating          *string  `json:"Rating"` // Bisa null, jadi pakai pointer
}

// Struct untuk response utama dari API
type ApiResponse struct {
	Code    string  `json:"Code"`
	Status  string  `json:"Status"`
	Message string  `json:"Message"`
	Data    []Venue `json:"Data"`
}

type TokenResponse struct {
	Code    string `json:"Code"`
	Status  string `json:"Status"`
	Message string `json:"Message"`
	Token   string `json:"Token"`
}
