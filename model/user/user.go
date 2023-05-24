package user

type User struct {
	ID       int    `json:"id" gorm:"primary_key"`
	Username string `json:"username" gorm:"type:varchar(20);not null;unique"`
	Password string `json:"password" gorm:"type:char(60);not null"` // bcrypt hashed
	Role     string `json:"role" gorm:"type:varchar(10);not null"`
}
