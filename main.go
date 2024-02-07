package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

type CandyData struct {
	Name  string     `json:"name"`
	Candy string     `json:"candy"`
	Eaten int        `json:"eaten"`
	Date  CustomTime `json:"date"`
}

type CandyDataSlice []CandyData

type CustomTime struct {
	time.Time
}

const dateFormat = "2006-01-02"

func (ct *CustomTime) UnmarshalJSON(input []byte) error {
	if len(input) < 3 { 
		return fmt.Errorf("invalid date: %s", input)
	}
	strInput := string(input[1 : len(input)-1]) 
	newTime, err := time.Parse(dateFormat, strInput)
	if err != nil {
		return fmt.Errorf("parse time error: %v for input %s", err, strInput)
	}
	ct.Time = newTime
	return nil
}

func fetchCandyData() ([]byte, error) {
	resp, err := http.Get("https://candystore.zimpler.net/api/candy")
	if err != nil {
		log.Printf("Error fetching candy data from external API: %v", err)
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Error reading response body: %v", err)
		return nil, err
	}

	return body, nil
}

func main() {
	e := echo.New()

	e.GET("/api/candy", func(c echo.Context) error {
		data, err := fetchCandyData()
		if err != nil {
			e.Logger.Errorf("Failed to fetch candy data: %v", err)
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to fetch candy data"})
		}

		var candyData CandyDataSlice
		if err := json.Unmarshal(data, &candyData); err != nil {
			e.Logger.Errorf("Failed to parse candy data: %v", err)
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to parse candy data"})
		}

		return c.JSON(http.StatusOK, candyData)
	})

	e.Logger.Fatal(e.Start(":3000"))
}
