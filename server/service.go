package server

import (
	"encoding/json"
	logger2 "github.com/mamoe/mirai-public-api/logger"
	"github.com/robfig/cron"
	"strconv"
	"time"
)

type Service struct {
	PluginPathList      []string
	PluginInfoList      []Plugin
	PluginBasicJsonInfo []BasicPluginInfo
	PluginInfoMap       map[string]Plugin
	lastUpdate          time.Time
	logger              *logger2.MiraiLogger
}

type Plugin struct {
	BasicInfo BasicPluginInfo
	JsonFile  string
}

func (this *Service) runSchedule() {
	this.logger.Log("==================\nStarting schedule....")
	this.scanPaths()
	this.scanDetailData()
	this.cacheJsonData()
	this.cacheMapData()
	this.logger.Log("Done.\n==================")
}

func (this *Service) cacheJsonData() { //缓存
	this.logger.Log("Starting cache JSON data for list query....")
	infoStrs := []BasicPluginInfo{}
	for _, value := range this.PluginInfoList {
		infoStrs = append(infoStrs, value.BasicInfo)
	}
	this.PluginBasicJsonInfo = infoStrs
}

func (this *Service) cacheMapData() { //缓存
	this.logger.Log("Starting cache map data for detailed query....")
	infoMap := make(map[string]Plugin)
	for _, value := range this.PluginInfoList {
		infoMap[value.BasicInfo.Name] = value
	}
	this.PluginInfoMap = infoMap
}

func (this *Service) scanPaths() {
	//扫描文件夹列表
	this.logger.Log("Starting scan paths....")
	basePath := GetMiraiApiServer().ConstructPluginPath()
	pathResult, err := walkDir(basePath)
	if err != nil {
		this.logger.Log("Error when scanning path!" + err.Error())
	}
	this.logger.Log("Done, found " + strconv.Itoa(len(pathResult)) + " plugins")
	this.PluginPathList = pathResult
}

func (this *Service) scanDetailData() {
	basePath := GetMiraiApiServer().ConstructPluginPath()
	//开始获取每个插件的信息
	pluginInfoList := []Plugin{}
	this.logger.Log("Starting scan each plugin....")
	for _, value := range this.PluginPathList {
		plugin := Plugin{}
		filePath := basePath + "/" + value + "/"
		jsonStr, err := ReadFile(filePath + "plugin.json")
		if err != nil {
			this.logger.Log("Error when reading " + filePath + " :" + err.Error())
			continue
		}
		plugin.JsonFile = jsonStr
		info := BasicPluginInfo{}
		err = json.Unmarshal([]byte(jsonStr), &info)
		if err != nil {
			this.logger.Log("Error when decoding " + filePath + " :" + err.Error())
			continue
		}
		plugin.BasicInfo = info
		pluginInfoList = append(pluginInfoList, plugin)
	}
	this.PluginInfoList = pluginInfoList
	this.lastUpdate = time.Now()
}

func (this *Service) initScan() {
	this.logger = logger2.NewLogger("MiraiService")
	this.logger.ColorPrefix = "@y"
	this.runSchedule()
	c := cron.New()                   // 新建一个定时任务对象
	c.AddFunc("0 */1 * * *", func() { //每十五分钟更新一次
		this.runSchedule()
	})
	c.Start()
}
