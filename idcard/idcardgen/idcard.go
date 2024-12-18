package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/jinlongchen/golang-utilities/idcard"
	"github.com/jinlongchen/golang-utilities/rand"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Street struct {
	Code         string `json:"code" xml:"code"`
	Name         string `json:"name" xml:"name"`
	AreaCode     string `json:"areaCode" xml:"areaCode"`
	CityCode     string `json:"cityCode" xml:"cityCode"`
	ProvinceCode string `json:"provinceCode" xml:"provinceCode"`
}
type Province struct {
	Code string `json:"code" xml:"code"`
	Name string `json:"name" xml:"name"`
}

func generateIDCard(c echo.Context) error {
	streets := make([]*Street, 0)
	err := json.Unmarshal([]byte(idcard.StreetsJson), &streets)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Error unmarshalling streets")
	}
	provinces := make([]*Province, 0)
	err = json.Unmarshal([]byte(idcard.ProvincesJson), &provinces)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Error unmarshalling provinces")
	}

	areaIndex := rand.GetRandInt(0, len(streets))
	t := time.Now()
	sexStr := "M"
	if rand.GetRandInt(0, 2) == 1 {
		sexStr = "F"
	}
	number := idcard.GenResidentIdCard(streets[areaIndex].AreaCode, t, sexStr)
	provinceName := ""
	for _, province := range provinces {
		if province.Code == streets[areaIndex].ProvinceCode {
			provinceName = province.Name
		}
	}
	valid, _, _ := idcard.IsResidentIdCard(number)
	if !valid {
		return c.JSON(http.StatusInternalServerError, "Generated invalid ID: "+number)
	}
	return c.JSON(http.StatusOK, map[string]string{
		"id":        number,
		"birthdate": t.Format("2006-01-02"),
		"sex":       sexStr,
		"province":  provinceName,
	})
}

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/generate-id", generateIDCard)

	e.Logger.Fatal(e.Start(":8080"))
}
