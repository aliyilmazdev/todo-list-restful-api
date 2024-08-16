package user

type User struct {
	ID       uint   `json:"id" gorm:"primaryKey;autoIncrement:true;not null;unique"`
	Name	 string `json:"name"`
	Surname  string `json:"surname"`
}