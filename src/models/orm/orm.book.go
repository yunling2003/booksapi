package orm

type Book struct {
	OrmModel
	Name 					string  `json:"name"`       //名称
	Author					string  `json:"author"`     //作者
	Publisher				string  `json:"publisher"`  //出版社
	CoverUrl				string	`json:"coverUrl"`	//封面URL
	FileUrl					string  `json:"fileUrl"`    //文件URL
}

func init() {
	db.AutoMigrate(&Book{})
}