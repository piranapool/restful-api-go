package models

//import (
//    "github.com/jinzhu/gorm"
//)

type Group struct {
    ID    uint    `gorm:"primary_key;auto_increment"`
    Name  string  `gorm:"size:20;not null;unique"`
    Users []*User `gorm:"many2many:groups_users"`
}
