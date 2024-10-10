// save_data.go
package controller

import (
	"integrasi-gbk-online/config"
	"integrasi-gbk-online/models"
	"log"
	"strings"
)

func SaveVenuesToDB(venues []models.Venue) error {
	for _, venue := range venues {
		galleryString := strings.Join(venue.Gallery, ",")

		query := `INSERT INTO integrasi_venue_online (ID, ParentID, MaxClubMember, Name, UnitNumber, Large, Capacity, CapacityVisitor, PhoneVenue, Description, PrimaryImage, IsActive, Status, UnitName, UnitSimpleName, UnitSheetName, CategoryName, Gallery, Rating) 
                  VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
                  ON DUPLICATE KEY UPDATE
                  ParentID = VALUES(ParentID),
				  MaxClubMember = VALUES(MaxClubMember),
				  Name = VALUES(Name),
				  UnitNumber = VALUES(UnitNumber),
				  Large = VALUES(Large),
				  Capacity = VALUES(Capacity),
				  CapacityVisitor = VALUES(CapacityVisitor),
				  PhoneVenue = VALUES(PhoneVenue),
				  Description = VALUES(Description),
				  PrimaryImage = VALUES(PrimaryImage),
				  IsActive = VALUES(IsActive),
				  Status = VALUES(Status),
				  UnitName = VALUES(UnitName),
				  UnitSimpleName = VALUES(UnitSimpleName),
				  UnitSheetName = VALUES(UnitSheetName),
				  CategoryName = VALUES(CategoryName),
				  Gallery = VALUES(Gallery),
				  Rating = VALUES(Rating)`

		// Gunakan string gabungan sebagai argumen untuk gallery
		_, err := config.DB.Exec(query, venue.ID, venue.ParentID, venue.MaxClubMember, venue.Name, venue.UnitNumber, venue.Large, venue.Capacity, venue.CapacityVisitor, venue.PhoneVenue, venue.Description, venue.PrimaryImage, venue.IsActive, venue.Status, venue.UnitName, venue.UnitSimpleName, venue.UnitSheetName, venue.CategoryName, galleryString, venue.Rating)
		if err != nil {
			log.Printf("Error saat menyimpan venue ID %d: %v", venue.ID, err)
			return err
		}
	}

	log.Println("Data venue berhasil disimpan.")
	return nil
}
