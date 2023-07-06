package models




type (

	BaseTable struct {
		Id int `gorm:"primarykey"`
		Name string `gorm:"not null"`
	}

	Categories struct {
		BaseModel
		BaseTable
	}
	Brands struct {
		BaseModel
		BaseTable
	}

	Sizes struct {
		BaseModel
		BaseTable
	}

	Colors struct {
		BaseModel
		BaseTable
	}

	Materials struct {
		BaseModel
		BaseTable
	}
	Products struct {
		BaseModel
		BaseTable
		CategoryID  int
		BrandID     int
		SizeID      int
		ColorID     int
		MaterialID  int
		Price       float64

		Category  Categories `gorm:"foreignKey:CategoryID"`
		Brand     Brands     `gorm:"foreignKey:BrandID"`
		Size      Sizes      `gorm:"foreignKey:SizeID"`
		Color     Colors     `gorm:"foreignKey:ColorID"`
		Material  Materials  `gorm:"foreignKey:MaterialID"`
	}
)