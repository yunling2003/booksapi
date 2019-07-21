package orm

import (	
  "time"
  "fmt"
  "strconv"
  "reflect"
	"booksapi/src/config"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type EmptyOrmModel struct {	
}

type OrmModel struct {
	EmptyOrmModel
	ID 			uint64 		`gorm:"primary_key" json:"id"`
	CreatedAt 	time.Time 	`json:"createdAt"`
	UpdatedAt 	time.Time 	`json:"updatedAt"`
	DeletedAt 	*time.Time 	`sql:"index" json:"-"`	
}

var (
	db *gorm.DB
)

func init() {
	db = OpenDB()
}

func OpenDB() *gorm.DB {
	logSQL, _ := strconv.ParseBool(config.All["logsql"])
	driver := config.All["dbdriver"]
	connectionString := config.All["connectionstring"]

	db,err := gorm.Open(driver, connectionString)
	db.LogMode(logSQL)

	if err != nil {
		panic("Failed to connect database")
	}

	return db
}

func (*EmptyOrmModel) GetAll(objs interface{}) error {
	return db.Find(objs).Error
}

func (*EmptyOrmModel) Create(obj interface{}) error {
	if rowsAffected := db.Create(obj).RowsAffected; rowsAffected == 0 {
    typeName := reflect.TypeOf(obj)
    return fmt.Errorf("Could not Create the obj[%s]", typeName.String())
  }
  
  return nil
}

func (*EmptyOrmModel) Update(obj interface{}) error {
  if rowsAffected := db.Save(obj).RowsAffected; rowsAffected == 0 {
    typeName := reflect.TypeOf(obj)
    return fmt.Errorf("Could not Update the obj[%s]", typeName.String())
  }

  return nil
}

func (*EmptyOrmModel) Get(obj interface{}, id uint64, preloads ...string) error {
	tmpDb := db
	for _, preLoad := range preloads {
		tmpDb = tmpDb.Preload(preLoad)
	}
	if rowsAffected := tmpDb.First(obj, id).RowsAffected; rowsAffected == 0 {
		typeName := reflect.TypeOf(obj)
		return fmt.Errorf("Could not Get the obj[%s] by ID [%v]", typeName.String(), id)
	}

	return nil
}