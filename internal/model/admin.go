package model

type Admin struct {
	Model
	Name     string `json:"name" gorm:"size:30;not null;comment:用户名称"`
	Mobile   string `json:"mobile" gorm:"size:20;not null;index;comment:用户手机号"`
	Password string `json:"-" gorm:"not null;default:'';comment:用户密码"`
}

// TableName 指定表名
func (Admin) TableName() string {
	return "admin"
}

func FindAdminById() Admin {
	var admin Admin
	Db.Where("id = ?", 1).First(&admin)
	return admin
}
