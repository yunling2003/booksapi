package orm

type Order struct {
	OrmModel
	UserID			uint64		`json:"userID"`
	User			User		`json:"user" gorm:"association_save_reference:false;association_autoupdate:false"`
	BookID			uint64		`json:"bookID"`
	Book			Book		`json:"book" gorm:"association_save_reference:false;association_autoupdate:false"`	
}

func init() {
	db.AutoMigrate(&Order{})
}