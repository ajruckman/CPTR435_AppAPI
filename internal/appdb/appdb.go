package appdb

import (
    . "github.com/ajruckman/xlib"
    "github.com/jmoiron/sqlx"
    _ "github.com/lib/pq"

    "api/internal/schema"
)

var (
    DB *sqlx.DB
)

const (
    dbstring = "user=temp_readings_mgr password=temp_readings_mgr host=localhost dbname=temp_readings sslmode=disable"
)

func init() {
    var err error
    DB, err = sqlx.Open("postgres", dbstring)
    Err(err)
}

func InsertReading(reading schema.Readings) {
    _, err := DB.NamedExec(`INSERT INTO readings (time, reading) VALUES (:Time, :Reading)`, reading)
    Err(err)
}

func GetReadings() (readings []schema.Readings) {
    err := DB.Select(&readings, "SELECT * FROM readings ORDER BY id DESC")
    Err(err)

    return
}
