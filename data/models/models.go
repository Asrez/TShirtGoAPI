package models




type (

	BaseTable struct {
		id int
		name string
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
)