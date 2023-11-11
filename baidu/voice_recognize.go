package baidu

import (
    "net/url"

    "github.com/jinlongchen/golang-utilities/http"
    "github.com/jinlongchen/golang-utilities/json"
)

type VoiceRecognizeResponse struct {
    CorpusNo string   `json:"corpus_no"`
    ErrMsg   string   `json:"err_msg"`
    ErrNo    int      `json:"err_no"`
    Result   []string `json:"result"`
    Sn       string   `json:"sn"`
}

func (bd *Baidu) VoiceRecognizePro(
    voiceData []byte,
    appId, appSecret string) (*VoiceRecognizeResponse, error) {
    accessToken, err := bd.GetAccessTokenOpenApiByClient(appId, appSecret)
    if err != nil {
        bd.logf("cannot get access token(%v): %v", appId, err.Error())
        return nil, err
    }

    bdReqURL, _ := url.Parse(`https://vop.baidu.com/pro_api`)
    bdReqQuery := bdReqURL.Query()
    bdReqQuery.Set("token", accessToken.AccessToken)
    bdReqQuery.Set("dev_pid", "80001")
    bdReqQuery.Set("cuid", "fd342f603890941ea5416a9508c75f8cd437b54d")

    bdReqURL.RawQuery = bdReqQuery.Encode()

    bdRespData, err := http.PostData(bdReqURL.String(),
        `audio/pcm;rate=16000`,
        voiceData,
    )

    if err != nil {
        bd.logf("bd err: %v", err)
        return nil, err
    }
    bdResp := &VoiceRecognizeResponse{}
    err = json.Unmarshal(bdRespData, bdResp)
    if err != nil {
        bd.logf("bd err: %v", err)
        return nil, err
    }
    return bdResp, nil
}
