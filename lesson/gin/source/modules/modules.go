package modules

import (
  "database/sql"
  "fmt"
  // "time"

  _ "github.com/go-sql-driver/mysql"
  "golang.org/x/crypto/bcrypt"
)

const users_table_name = "users"

type User struct {
  Id int          `json: id`
  Name string     `json: name`
  Passwd string   `json: passwd`
}

func DBConnect() (db *sql.DB) {
    dbDriver := "mysql"
    // dbUser := "curtain_app_user"
    dbUser := "root"
    // dbPass := "curtain_app_PW#0"
    dbPass := "root00"
    dbName := "lesson_db"
    dbProtocol := "tcp(mariadb:3306)"
    dbOption := "?parseTime=true"
    db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@"+dbProtocol+"/"+dbName+dbOption)
    if err != nil {
      panic(err.Error())
    }
    return db
}

func UserAdd(name string, passwd string) {
  db := DBConnect()
  // date := time.Now()
  hash := Pw2Hash(passwd)
  insert_sql := fmt.Sprintf("insert into %s (name, passwd) values (?, ?);", users_table_name)
  _, err := db.Exec(insert_sql, name, hash)
  if err != nil {
    panic(err.Error())
  }
  // fmt.Printf("insert: %s(%v), %s(%v).\n", date, date, msg, msg)
}

func GetUsers() []User {
  db := DBConnect()
  select_sql := fmt.Sprintf("select * from %s;", users_table_name)
  res, err := db.Query(select_sql)
  if err != nil {
    panic(err.Error())
  }

  users := []User{}
  for res.Next() {
    var id int
    var name string
    var passwd string
    res.Scan(&id, &name, &passwd)
    users = append(users, User{id, name, passwd})
  }
  return users
}

func Pw2Hash(passwd string) string {
  hash, err := bcrypt.GenerateFromPassword([]byte(passwd), bcrypt.DefaultCost)
  if err != nil {
    panic(err.Error())
  }
  return string(hash)
}

func CompPw(hash string, passwd string) bool {
  err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(passwd))
  if err == nil {
    return true
  } else {
    return false
  }
}

/*
func main() {
  var str string
  fmt.Scan(&str)

  if str == "insert" {
    DBInsert("hoge")
  }

  logs := DBSelect()
  for i, log := range logs {
    fmt.Printf("log[%d]: %v\n", i, log)
  }
}
*/
