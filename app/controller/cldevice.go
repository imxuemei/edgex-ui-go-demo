package controller

import (
	"bytes"
	"github.com/edgexfoundry/device-sdk-go/pkg/startup"
	yxmtest "github.com/edgexfoundry/edgex-ui-go/vds/simpledevice"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/edgexfoundry/edgex-ui-go/app/common"
)

func GetCldeviceNameTest(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	w.Header().Set(common.ContentTypeKey, common.JsonContentType)
	w.Write([]byte("yxmtestName"))
}

func StartSimpledevice(w http.ResponseWriter, r *http.Request){
	//使用go语言的 goroutine 机制创建子进程来运行simple device示例
	go startSimpleDeviceThread()

	defer r.Body.Close()
	w.Header().Set(common.ContentTypeKey, common.JsonContentType)
	w.Write([]byte("starting"))
}

func StopSimpledevice(w http.ResponseWriter, r *http.Request)  {
	//todo 未找到关闭方法，暂时不做处理

	defer r.Body.Close()
	w.Header().Set(common.ContentTypeKey, common.JsonContentType)
	w.Write([]byte("stoping"))
}

func startSimpleDeviceThread(){
	sd := yxmtest.SimpleDriver{}
	localConfProfile := ""
	localConfDir := "E:\\workspace-edgex\\src\\github.com\\edgexfoundry\\edgex-ui-go\\vds\\simpledevice\\res"
	log.Println(localConfDir)
	log.Println(localConfProfile)
	startup.Bootstrap("Simple-Device-yxm", "1.0.0", &sd, localConfProfile, localConfDir)
}




func putStop(){
	url := "http://10.109.24.136:49990/api/v1/device/Simple-Device-yxm/SwitchButton"
	log.Println("URL:>", url)
	item := "off"
	updateParams := `{"SwitchButton":"` + item + `"}`
	var jsonStr = []byte(updateParams)
	req, _ := http.NewRequest(http.MethodPut, url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, _ := client.Do(req)
	log.Println("response Status:", resp.Status)
	log.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	log.Println("response Body:", string(body))
}
