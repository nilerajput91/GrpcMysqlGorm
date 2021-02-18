package model

type UserDB struct {
	Id       string `gorm:"primary_key; auto_increment; not_null"`
	Name     string
	Email    string
	Address  string
	Password string
}

func (e *UserDB) TableName() string {
	return "user"
}
