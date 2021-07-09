package entity

type Class struct {
	ClassID          string     `gorm:"type:varchar(100);not null;column:id"`
	Title            string     `gorm:"type:varchar(100);not null"`
	ClassDescription string     `gorm:"type:varchar(100);not null"`
	Students         []*Student `gorm:"many2many:class_students;"`
}
