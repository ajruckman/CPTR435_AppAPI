package endpoint

import (
    "fmt"
    "net/http"

    "github.com/gin-gonic/gin"

    "api/internal/appdb"
    "api/internal/schema"
)

func Reading_Create(c *gin.Context) {
    var reading schema.Readings

    if err := c.ShouldBindJSON(&reading); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    appdb.InsertReading(reading)

    fmt.Println(reading.Time, reading.Reading)
    c.JSON(http.StatusOK, gin.H{})
}

func Reading_List(c *gin.Context) {
    readings := appdb.GetReadings()

    c.JSON(200, readings)
}
