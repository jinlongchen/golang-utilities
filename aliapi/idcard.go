package aliapi

import (
	"encoding/base64"
	"encoding/json"
	gohttp "net/http"
	"github.com/jinlongchen/golang-utilities/http"
)

const (
	OCR_IDCARD_API_URL = `https://dm-51.data.aliyun.com/rest/160601/ocr/ocr_idcard.json`
)

func (api *AliApiHelper) OCRIDCardFace(data []byte) (*IDCardFace, error) {
	respData, err := api.ocrIDCard(data, "face")
	if err != nil {
		return nil, err
	}
	ret := &IDCardFace{}
	err = json.Unmarshal(respData, ret)
	if err != nil {
		return nil, err
	}
	return ret, nil
}

func (api *AliApiHelper) OCRIDCardBack(data []byte) (*IDCardBack, error) {
	respData, err := api.ocrIDCard(data, "back")
	if err != nil {
		return nil, err
	}
	ret := &IDCardBack{}
	err = json.Unmarshal(respData, ret)
	if err != nil {
		return nil, err
	}
	return ret, nil
}

func (api *AliApiHelper) ocrIDCard(data []byte, side string) (ret []byte, err error) {
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
	_, respData, err := http.PostDataWithHeaders(
		OCR_IDCARD_API_URL,
		gohttp.Header{
			"Authorization": []string{"APPCODE " + api.cfg.GetString("aliapi.ocr.idcard.appcode")},
		},
		"application/json",
		jsonBody,
	)

	//fmt.Printf("%v\n", respHeader)
	//fmt.Println(string(respData))

	ret = respData

	return
}

/*
{
    "address": "陕西省甘泉县****5号",
    "birth": "19890218",
    "config_str": "{\"side\":\"face\"}",
    "face_rect": {
        "angle": -89.867210388183594,
        "center": {
            "x": 675.3818359375,
            "y": 272.56576538085938
        },
        "size": {
            "height": 156.00003051757812,
            "width": 156.00001525878906
        }
    },
    "face_rect_vertices": [
        {
            "x": 753.20086669921875,
            "y": 350.746337890625
        },
        {
            "x": 597.20123291015625,
            "y": 350.384765625
        },
        {
            "x": 597.56280517578125,
            "y": 194.38519287109375
        },
        {
            "x": 753.56243896484375,
            "y": 194.74676513671875
        }
    ],
    "name": "高静",
    "nationality": "汉",
    "num": "6106****0067",
    "request_id": "20181116130758_6b45831f0ea48f24e4992f6e4489ef5d",
    "sex": "女",
    "success": true
}
*/
type IDCardFace struct {
	Address     string `json:"address"`
	Birth       string `json:"birth"`
	ConfigStr   string `json:"config_str"`
	Name        string `json:"name"`
	Nationality string `json:"nationality"`
	Num         string `json:"num"`
	RequestID   string `json:"request_id"`
	Sex         string `json:"sex"`
	Success     bool   `json:"success"`
	FaceRect    struct {
		Angle  float64 `json:"angle"`
		Center struct {
			X float64 `json:"x"`
			Y float64 `json:"y"`
		} `json:"center"`
		Size struct {
			Width  float64 `json:"width"`
			Height float64 `json:"height"`
		} `json:"size"`
	} `json:"face_rect"`
	FaceRectVertices []struct {
		X float64 `json:"x"`
		Y float64 `json:"y"`
	} `json:"face_rect_vertices"`
}

/*
{
    "config_str": "{\"side\":\"back\"}",
    "end_date": "20251008",
    "issue": "上海市公安局徐汇分局",
    "request_id": "20181116132317_6800f0186340de14bf29ce6c60c6c5b9",
    "start_date": "20051008",
    "success": true
}
*/
type IDCardBack struct {
	ConfigStr string `json:"config_str"`
	EndDate   string `json:"end_date"`
	Issue     string `json:"issue"`
	RequestID string `json:"request_id"`
	StartDate string `json:"start_date"`
	Success   bool   `json:"success"`
}
