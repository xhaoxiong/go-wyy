package TEST

import (
	"github.com/iqysf/gorm"
	"time"
)

type StaffOperationLog struct {
	gorm.Model
	Path   string
	Method string
	Ip     string
	Input  string
}

type WorkFlow struct {
	gorm.Model
	Type    string
	Step    string
	StaffId int64
	Status  int
}

type Roles struct {
	gorm.Model
	Name       string
	Slug       string
	Permission []Permission `gorm:"many2many:role_permissions;"`
	Staffs     []Staffs     `gorm:"many2many:role_staffs;"`
	Menu       []Menu       `gorm:"many2many:role_menus;"`
}

type Permission struct {
	gorm.Model
	Name       string
	Slug       string
	HttpMethod string
	HttpPath   string
	Roles      []Roles  `gorm:"many2many:role_permissions;"`
	Staffs     []Staffs `gorm:"many2many:permission_staffs"`
}

type Staffs struct {
	gorm.Model
	UserName            string   `gorm:"unique"`
	Password            string
	Name                string
	Avatar              string
	IsOnline            int
	PassCount           int
	RefusedCount        int
	TransferCount       int
	AverageOperatedTime time.Time
	Permissions         []Staffs `gorm:"many2many:permission_staffs"`
	Roles               []Roles  `gorm:"many2many:role_staffs"`
}

type Menu struct {
	ParentId int
	Order    int
	Title    string
	Icon     string  `gorm:"type:longtext"`
	Uri      string  `gorm:"type:longtext"`
	Roles    []Roles `gorm:"many2many:role_menus;"`
}
