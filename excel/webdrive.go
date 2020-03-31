package main

import (
	"fmt"
	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/chrome"
	"os"
	"strings"
	"time"
)

const (
	// These paths will be different on your system.
	seleniumPath    = "selenium.jar"
	geckoDriverPath = "chromedriver.exe"
	port            = 8080
)
func getSeivice()(*selenium.Service, error){
	opts := []selenium.ServiceOption{
		//selenium.StartFrameBuffer(),           // Start an X frame buffer for the browser to run in.
		selenium.ChromeDriver(geckoDriverPath), // Specify the path to GeckoDriver in order to use Firefox.
		selenium.Output(os.Stderr),             // Output debug information to STDERR.
	}
	selenium.SetDebug(false)
	service, err := selenium.NewSeleniumService(seleniumPath, port, opts...)
	//defer service.Stop() 关闭浏览器不存在的
	return service,err
}
func getChrome()(selenium.WebDriver, error){
	caps := selenium.Capabilities{
		"browserName": "chrome",
	}
	chromeCaps := chrome.Capabilities{
		Path:  "",
		Args: []string{
			"User-Agent=Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_2) AppleWebKit/604.4.7 (KHTML, like Gecko) Version/11.0.2 Safari/604.4.7",
			"cookie=think_language=zh-CN; PHPSESSID=kavoa1nh443a4igrpp7q5i0hi5; Hm_lvt_31325a0b97c093393242fde24a1a4a1d=1562133050; Hm_lpvt_31325a0b97c093393242fde24a1a4a1d=1562133055; token=c070353dcfe596189f3c1fc82d08d61d; user_id=60317",
			"accept=text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3",
			"Host=rest.apizza.net",
			"Referer= https://apizza.net/pro/",
			// 模拟user-agent，防反爬
		},
	}
	caps.AddChrome(chromeCaps)
	wd, err := selenium.NewRemote(caps, fmt.Sprintf("http://localhost:%d/wd/hub", port))
	return wd,err

}
func Login(wd selenium.WebDriver)  {
	if err:=wd.Get("https://www.bugclose.com/login.html");err!=nil{
		panic(err)
	}
	btn,err:=wd.FindElement(selenium.ByCSSSelector,"#userLoginForm > div:nth-child(1) > div > input")
	if err!=nil{
		panic(err)
	}
	_=btn.Click()
	err=btn.SendKeys("1658616397@qq.com")
	if err!=nil {
		panic(err)
	}
	//#loginBtn
	pas,err:=wd.FindElement(selenium.ByCSSSelector,"#userLoginForm > div:nth-child(2) > div > input")
	if err!=nil{
		panic(err)
	}
	_=btn.Click()
	err=pas.SendKeys("liqizxc886")
	log,err:=wd.FindElement(selenium.ByCSSSelector,"#loginBtn")
	if err!=nil{
		panic(err)
	}
	_=log.Click()
}
func main() {
	selenium.SetDebug(true)
	service, err := getSeivice()
	if err != nil {
		panic(err) // panic is used only as an example and is not otherwise recommended.
	}
	fmt.Println(service)
	wd,err:=getChrome()
	// Connect to the WebDriver instance running locally.
	if err != nil {
		panic(err)
	}
	Login(wd)
	//defer wd.Quit()
	//Train(wd)
}
func test(wd selenium.WebDriver) (bool, error) {
	elem, err := wd.FindElement(selenium.ByCSSSelector, ".btn-upload-submit")
	fmt.Println(wd.Title())
	time.Sleep(20000)
	if err != nil {
		return false,err
	}
	if elem!=nil{
		return true,nil
	}
	return false,nil
}
func aiqi(wd selenium.WebDriver){
	if err := wd.Get("http://unionlifetest.h2hlife.com/Application/Home/View/Redirect/upload.html?openid%3DoCDjWjlQMvrTvcFYf-05gGlCytCE%26redirectCode%3D1010%26age%3D1%E5%B2%815%E4%B8%AA%E6%9C%880%E5%91%A8%26task_id%3D89%26second_task%3D%E6%9C%AC%E6%9C%88%E9%A5%AE%E9%A3%9F%E6%8C%87%E5%8D%97%26"); err != nil {
		panic(err)
	}
	time.Sleep(500)
	// Get a reference to the text box containing code.
	for{
		result:=wd.Wait(test)
		if result==nil{
			break
		}
	}
	elem, err := wd.FindElement(selenium.ByCSSSelector, ".upload-rdcontent textarea")
	if err != nil {
		panic(err)
	}
	// Remove the boilerplate code already in the text box.
	if err := elem.Clear(); err != nil {
		panic(err)
	}

	// Enter some new code in text box.
	err = elem.SendKeys(`
		package main
		import "fmt"
		func main() {
			fmt.Println("Hello WebDriver!\n")
		}
	`)
	if err != nil {
		panic(err)
	}

	// Click the run button.
	btn, err := wd.FindElement(selenium.ByCSSSelector, "#upload")
	if err != nil {
		panic(err)
	}
	for i:=0;i<8;i++{
		err=btn.SendKeys("C:/excel/QQ图片20190531101135.png")
	}
	if err != nil {
		panic(err)
	}
	time.Sleep(2000)
	cli, err := wd.FindElement(selenium.ByCSSSelector, ".btn-upload-submit")
	if err != nil {
		panic(err)
	}
	err=cli.Click()
	if err != nil {
		panic(err)
	}
}
func Train(wd selenium.WebDriver){
	err:=wd.Get("https://exservice.12306.cn/excater/index.html")
	if err!=nil{
		panic(err)
	}
	//cli,err:=wd.FindElement(selenium.ByXPATH,"/html/body/div[2]/div[2]/ul/li[2]/a")
	//if err!=nil{
	//	panic(err)
	//}
	//err=cli.Click()
	//if err!=nil{
	//	panic(err)
	//}
	for{
		result:=wd.Wait(waitLogin)
		if result==nil{
			break
		}
		fmt.Println("等待用户登录并进入选票窗口")
		time.Sleep(2*time.Second)
	}
	setout,err:=wd.FindElements(selenium.ByCSSSelector,".bgc")
	if err!=nil{
		panic(err)
	}
	for _,z:=range setout  {
		stringResult,err:=z.Text()
		if err!=nil{
			panic(err)
		}
		s:=strings.Split(stringResult,"\n")
		train:=false
		canbuy:=false
		for i:=0;i<len(s);i++{
			if s[i]=="G152"{
				fmt.Println("11111111111")
				train=true
			}
			if strings.Contains(s[i],"有"){
				fmt.Println("2222222222222222")
				canbuy=true
			}
			//document.querySelector("#ticket_5l0000G10880 > td.no-br > a")
		}
		if train==true&&canbuy==true{
			fmt.Println("????????????")
			btn,err:=z.FindElement(selenium.ByCSSSelector,".no-br")
			if err!=nil{
				panic(err)
			}
			err=btn.Click()
			if err!=nil{
				panic(err)
			}
			break
		}
	}
	//_=setout.Click()
	//_=setout.Clear()
	//err=setout.SendKeys("广州")
	//if err!=nil{
	//	panic(err)
	//}
	//setoutclickresult,err:=wd.FindElement(selenium.ByLinkText,"广州")
	//err=setoutclickresult.Click()
	//if err!=nil{
	//	panic(err)
	//}
	//等待用户输入账号密码
	////输入账号密码 然后等待用户识别验证码并登录
	//user,err:=wd.FindElement(selenium.ByCSSSelector,"J-userName")
	//if err!=nil{
	//	panic(err)
	//}
	//_=user.Click()
	//err=user.SendKeys("13435110579")
	//if err!=nil{
	//	panic(err)
	//}
	//pass,err:=wd.FindElement(selenium.ByXPATH,"J-password")
	//if err!=nil{
	//	panic(err)
	//}
	//_=pass.Click()
	//err=pass.SendKeys("liqizxc886")
	//if err!=nil{
	//	panic(err)
	//}
}
func waitLogin(wd selenium.WebDriver) (bool, error) {
	elem, err := wd.FindElement(selenium.ByCSSSelector, ".bgc")
	fmt.Println(wd.Title())
	time.Sleep(20000)
	if err != nil {
		return false,err
	}
	if elem!=nil{
		return true,nil
	}
	return false,nil
}
