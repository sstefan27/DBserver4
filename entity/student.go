package entity

type Student struct {
	StudentID string   `gorm:"type:varchar(100);not null;column:id"`
	LastName  string   `gorm:"type:varchar(100);not null"`
	FirstName string   `gorm:"type:varchar(100);not null"`
	Age       int      `gorm:"type:int;not null"`
	Class     []*Class `gorm:"many2many:class_students;"`
}
