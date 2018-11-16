package aliapi

import (
	"encoding/base64"
	"encoding/json"
	gohttp "net/http"
	"github.com/jinlongchen/golang-utilities/http"
	"fmt"
)

const (
	OCR_VEHICLE_LICENSE_API_URL = `https://dm-53.data.aliyun.com/rest/160601/ocr/ocr_vehicle.json`
)

func (api *AliApiHelper) OCRVehicleLicenseFace(data []byte) (*VehicleLicenseFace, error) {
	respData, err := api.ocrVehicleLicense(data, "face")
	if err != nil {
		return nil, err
	}
	ret := &VehicleLicenseFace{}
	err = json.Unmarshal(respData, ret)
	if err != nil {
		return nil, err
	}
	return ret, nil
}
func (api *AliApiHelper) OCRVehicleLicenseBack(data []byte) (*VehicleLicenseBack, error) {
	respData, err := api.ocrVehicleLicense(data, "back")
	if err != nil {
		return nil, err
	}
	ret := &VehicleLicenseBack{}
	err = json.Unmarshal(respData, ret)
	if err != nil {
		return nil, err
	}
	return ret, nil
}

func (api *AliApiHelper) ocrVehicleLicense(data []byte, side string) (ret []byte, err error) {
	bt := base64.StdEncoding.EncodeToString(data)
	confBody, err := json.Marshal(&struct {
		Side string `json:"side"`
	}{
		Side: side,
	})
	jsonBody, err := json.Marshal(&struct {
		Image     string `json:"image"`
		Configure string `json:"configure"`
	}{
		Image:     bt,
		Configure: string(confBody),
	})
	println(string(jsonBody))
	respHeader, respData, err := http.PostDataWithHeaders(
		OCR_VEHICLE_LICENSE_API_URL,
		gohttp.Header{
			"Authorization": []string{"APPCODE " + api.cfg.GetString("aliapi.ocr.idcard.appcode")},
		},
		"application/json",
		jsonBody,
	)

	fmt.Printf("%v\n", respHeader)
	fmt.Println(string(respData))

	ret = respData

	return
}

/*
正面
{
    "addr": "内蒙古自****2单元4层2号",
    "config_str": "{\"side\":\"face\"}",
    "engine_num": "51***80",
    "issue_date": "20130722",
    "model": "解放牌CA4****EA80",
    "owner": "呼伦贝尔市晓明运输有限公司",
    "plate_num": "蒙E****8",
    "register_date": "20101207",
    "request_id": "20181116223213_f17142d3ee1205cc65e5d10619e65e12",
    "success": true,
    "use_character": "货运",
    "vehicle_type": "重型半挂牵引车",
    "vin": "LFWS****754"
}
反面
{
    "config_str":"{\"side\": \"back\" }", 
    "appproved_passenger_capacity":"5人",   
    "approved_load":"",                     
    "file_no":"530100001466",               
    "gross_mass":"2000kg",                  
    "inspection_record":"检验有效期至2014年09月云A(01)", 
    "overall_dimension":"4945x1845x1480mm",  
    "traction_mass":"",                      
    "unladen_mass":"1505kg"                  
    "plate_num":"云AD8V02",                  
    "success":true,             
    "request_id":"20180131144149_c440540b20a4dc079a10680ff60b2d2a"
}
*/
type VehicleLicenseFace struct {
	Addr string `json:"addr" xml:"addr"`
	ConfigStr string `json:"config_str" xml:"config_str"`
	EngineNum string `json:"engine_num" xml:"engine_num"`
	IssueDate string `json:"issue_date" xml:"issue_date"`
	Model string `json:"model" xml:"model"`
	Owner string `json:"owner" xml:"owner"`
	PlateNum string `json:"plate_num" xml:"plate_num"`
	RegisterDate string `json:"register_date" xml:"register_date"`
	RequestID string `json:"request_id" xml:"request_id"`
	Success bool `json:"success" xml:"success"`
	UseCharacter string `json:"use_character" xml:"use_character"`
	VehicleType string `json:"vehicle_type" xml:"vehicle_type"`
	Vin string `json:"vin" xml:"vin"`
}
type VehicleLicenseBack struct {
	AppprovedPassengerCapacity string `json:"appproved_passenger_capacity" xml:"appproved_passenger_capacity"`
	ApprovedLoad string `json:"approved_load" xml:"approved_load"`
	ConfigStr string `json:"config_str" xml:"config_str"`
	FileNo string `json:"file_no" xml:"file_no"`
	GrossMass string `json:"gross_mass" xml:"gross_mass"`
	InspectionRecord string `json:"inspection_record" xml:"inspection_record"`
	OverallDimension string `json:"overall_dimension" xml:"overall_dimension"`
	PlateNum string `json:"plate_num" xml:"plate_num"`
	RequestID string `json:"request_id" xml:"request_id"`
	Success bool `json:"success" xml:"success"`
	TractionMass string `json:"traction_mass" xml:"traction_mass"`
	UnladenMass string `json:"unladen_mass" xml:"unladen_mass"`
}