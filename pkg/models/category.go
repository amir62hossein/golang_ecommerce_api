package models

type CATEGORY_ENUM string

const (
	CategoryElectronics CATEGORY_ENUM = "ELECTRONIC"
	CategoryClothing    CATEGORY_ENUM = "CLOTHING"
	CategoryBooks       CATEGORY_ENUM = "BOOK"
	CategoryFurniture   CATEGORY_ENUM = "FURNITURE"
	CategoryToys        CATEGORY_ENUM = "TOY"
	CategorySports      CATEGORY_ENUM = "SPORT"
)

type Category struct {
	DBModel
	Category CATEGORY_ENUM `json:"category" gorm:"not null;unique"  sql:"type:ENUM('ELECTRONIC', 'CLOTHING' , 'BOOK' , 'FURNITURE', 'TOY' , 'SPORT')"`
}
