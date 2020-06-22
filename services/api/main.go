package main

import (
    "fmt"
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/postgres"
    m "github.com/piranapool/restful-api-go/api/models"
)

func main() {
    dialect := "postgres"
    sslmode := "disable"
    host := "postgres"
    port := "5432"
    user := "dev_user"
    password := "dev_password"
    dbname := "dev_directory"

    dbUri := fmt.Sprintf("sslmode=%s host=%s port=%s user=%s password=%s dbname=%s", 
        sslmode, host, port, user, password, dbname)

    db, err := gorm.Open(dialect, dbUri)
    if err != nil {
        fmt.Print(err)
        return
    }

    fmt.Println(dbUri)
    db.AutoMigrate(&m.Group{}, &m.User{})
    db.Table("groups_users").AddForeignKey("group_id", "groups(id)", "CASCADE", "CASCADE")
    db.Table("groups_users").AddForeignKey("user_id", "users(id)", "CASCADE", "CASCADE")
}
