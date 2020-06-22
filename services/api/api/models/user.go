package models

//import (
//    "github.com/jinzhu/gorm"
//)

type User struct {
    ID        uint     `gorm:"primary_key;auto_increment"`
    FirstName string   `gorm:"size:20;not null"`
    LastName  string   `gorm:"size:20;not null"`
    UserID    string   `gorm:"size:12;not null;unique"`
    Groups    []*Group `gorm:"many2many:groups_users"`
}
