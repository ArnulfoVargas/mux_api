package database

import (
	"os"
	"strings"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Database = func () *gorm.DB {
  dsn := generateDSN()
  db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

  if err != nil {
    panic(err.Error())
  }

  return db
}

func generateDSN() string {
  sb := strings.Builder{}

  sb.WriteString("host=")
  sb.WriteString(os.Getenv("DB_HOST"))
  sb.WriteByte(' ')

  sb.WriteString("user=")
  sb.WriteString(os.Getenv("DB_USER"))
  sb.WriteByte(' ')

  sb.WriteString("dbname=")
  sb.WriteString(os.Getenv("DB_NAME"))
  sb.WriteByte(' ')

  sb.WriteString("port=")
  sb.WriteString(os.Getenv("DB_PORT"))
  sb.WriteByte(' ')

  sb.WriteString("sslmode=")
  sb.WriteString(os.Getenv("DB_SSL"))

  return sb.String()
}
