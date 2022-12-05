/*
 * Copyright (c) 2020. Jinlong Chen.
 */

package sms

import (
    "fmt"
    "io/ioutil"
    "path"
    "runtime"
    "testing"

    "github.com/qiniu/api.v7/v7/auth"
)

func getManager() *Manager {
    ak := getBitsAppKey()
    sk := getBitsSecretKey()

    qiniuManager := NewManager(auth.New(ak, sk))

    return qiniuManager
}

func TestNewManager(t *testing.T) {
    _, filename, _, _ := runtime.Caller(0)
    phoneData, _ := ioutil.ReadFile(path.Join(path.Dir(filename), "test.txt"))

    tID1 := "1217103952852553728"

    qiniuManager := getManager()

    response, err := qiniuManager.SendMessage(MessagesRequest{
        TemplateID: tID1,
        Mobiles:    []string{string(phoneData)},
        Parameters: map[string]interface{}{
            "module": "邮件",
            "reason": "被退信",
        },
    })

    if err != nil {
        t.Fatal(err)
    }

    fmt.Printf("job id: %s\n", response.JobID)
    // 1219176816170766336
}

func TestNewManager2(t *testing.T) {
    jobID := "1235997636973043712"

    qiniuManager := getManager()

    response, err := qiniuManager.QueryMessage(QueryMessageRequest{
        JobID: jobID,
    })
    if err != nil {
        t.Fatal(err)
    }

    fmt.Printf("messages: %v\n", response)

}
func TestNewManager3(t *testing.T) {
    qiniuManager := getManager()
    smsTemplate, err := qiniuManager.QueryTemplateByID("1217103952571539456")
    if err != nil {
        return
    }
    fmt.Printf("%v", smsTemplate)
    // var paramRegex = regexp.MustCompile(`\$\{([^\}]+)\}`)
    // templateContent := smsTemplate.Template
    // matches := paramRegex.FindAllStringSubmatch(templateContent, -1)
    // for _, group := range matches {
    //	templateContent = strings.ReplaceAll(templateContent, group[0], helper.GetValueAsString(params, group[1], ""))
    // }
    // smsContent := fmt.Sprintf("【%s】%s", signName, templateContent)
    // return smsContent, int(math.Ceil(float64(utf8.RuneCountInString(smsContent)) / 70.0))

}
