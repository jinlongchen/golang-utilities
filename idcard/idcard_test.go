package idcard

import (
	"encoding/json"
	"testing"
	"time"
	"yijiu.com/common/rand"
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
func TestGenResidentIdCard(t *testing.T) {
	streets := make([]*Street, 0)
	err := json.Unmarshal([]byte(StreetsJson), &streets)
	if err != nil {
		return
	}
	provinces := make([]*Province, 0)
	err = json.Unmarshal([]byte(ProvincesJson), &provinces)
	if err != nil {
		return
	}

	for offset := 0; offset < 80; offset++ {
		areaIndex := rand.GetRandInt(0, len(streets))

		t := time.Now().Add(-time.Hour * 24 * time.Duration(rand.GetRandInt(10, 365))).AddDate(-offset, 0,0)

		var sexStr string
		if rand.GetRandInt(0, 2) == 1 {
			sexStr = "M"
		} else {
			sexStr = "F"
		}
		number := GenResidentIdCard(streets[areaIndex].AreaCode, t, sexStr)
		provinceName := ""
		for _, province := range provinces {
			if province.Code == streets[areaIndex].ProvinceCode {
				provinceName = province.Name
			}
		}
		valid, _, _ := IsResidentIdCard(number)
		if !valid {
			panic("invalid id:" + number)
		}
		println(number, t.Format("2006-01-02"), sexStr, provinceName)
	}

}

func TestIsResidentIdCard(t *testing.T) {
	valid, _, _ := IsResidentIdCard("510122199108247115")
	println(valid)
}