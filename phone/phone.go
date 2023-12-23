package phone

import (
	"fmt"
	"os"
	"regexp"
	"strings"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type Numbers struct {
	Id    string `json:"id" gorm:"column:id;primarykey;"`
	Phone string `json:"phone" gorm:"column:phone;"`
}

const TableName = "phone_numbers"

func (Numbers) TableName() string {
	return TableName
}

func Start() (err error) {
	db, err := ConnectDB(os.Getenv("DSN"))
	if err != nil {
		return fmt.Errorf("failed to connect database: %s", err)
	}

	var phoneList []Numbers
	if err = db.Table(Numbers{}.TableName()).Find(&phoneList).Error; err != nil {
		return fmt.Errorf("failed to get all records in database: %s", err)
	}

	tmp := map[string]bool{}
	for _, item := range phoneList {
		fl := Normalize(item.Phone)

		if val, _ := tmp[fl]; val {
			if err := db.Table(Numbers{}.TableName()).Delete(&Numbers{Id: item.Id}).Error; err != nil {
				fmt.Printf("error when delete id %s: %s", item.Id, err)
				continue
			}
		}

		tmp[fl] = true
		if err := db.Table(Numbers{}.TableName()).Where("id = ?", item.Id).Update("phone", fl).Error; err != nil {
			fmt.Printf("error when delete id %s: %s", item.Id, err)
		}
	}

	fmt.Printf("Normalize table %q successfully.\n", Numbers{}.TableName())
	return
}

func ConnectDB(dsn string) (db *gorm.DB, err error) {
	db, err = gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: false,
		}})
	return
}

func Normalize[T string | int](in T) (out string) {
	sampleRegexp := regexp.MustCompile(`\d+`)
	s := fmt.Sprintf("%v", in)
	match := sampleRegexp.FindAllString(s, -1)
	if len(match) >= 1 {
		out = strings.Join(match, "")
	}
	return
}
