package phone

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"os"
	"regexp"
	"strings"
)

type PhoneNumbers struct {
	Id    string `json:"id" gorm:"column:id;primarykey;"`
	Phone string `json:"phone" gorm:"column:phone;"`
}

func (PhoneNumbers) TableName() string {
	return "phone_numbers"
}

func Start() (err error) {
	db, err := connectDB()
	if err != nil {
		return fmt.Errorf("failed to connect database: %s", err)
	}

	var phoneList []PhoneNumbers
	if err = db.Table(PhoneNumbers{}.TableName()).Find(&phoneList).Error; err != nil {
		return fmt.Errorf("failed to get all records in database: %s", err)
	}

	tmp := map[string]bool{}
	for _, item := range phoneList {
		fl := normalize(item.Phone)

		if val, _ := tmp[fl]; val {
			if err := db.Table(PhoneNumbers{}.TableName()).Delete(&PhoneNumbers{Id: item.Id}).Error; err != nil {
				fmt.Printf("error when delete id %s: %s", item.Id, err)
				continue
			}
		}

		tmp[fl] = true
		if err := db.Table(PhoneNumbers{}.TableName()).Where("id = ?", item.Id).Update("phone", fl).Error; err != nil {
			fmt.Printf("error when delete id %s: %s", item.Id, err)
		}
	}

	fmt.Printf("Normalize table %q successfully.\n", PhoneNumbers{}.TableName())
	return
}

func connectDB() (db *gorm.DB, err error) {
	db, err = gorm.Open(postgres.New(postgres.Config{
		DSN:                  os.Getenv("DSN"),
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: false,
		}})
	return
}

func normalize[T string | int](in T) (out string) {
	sampleRegexp := regexp.MustCompile(`\d+`)
	s := fmt.Sprintf("%v", in)
	match := sampleRegexp.FindAllString(s, -1)
	if len(match) >= 1 {
		out = strings.Join(match, "")
	}
	return
}
