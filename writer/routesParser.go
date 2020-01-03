package writer

import (
	"encoding/json"
	"github.com/crusj/logger"
	"io/ioutil"
	"os"
	"strings"
)

type routesParser struct {
	Route []Route
}

func (routesParser routesParser) Parse() []Route {
	return routesParser.Route
}
func NewRoutesParser(path string) *routesParser {
	parser := new(routesParser)
	f, err := os.Open(path)
	if err != nil {
		logger.Painc("打开文件%v,失败,%s", path, err)
	}
	defer f.Close()

	if j, err := ioutil.ReadAll(f); err != nil {
		logger.Painc("读取文件%v,失败,%s", path, err)
	} else {
		var groups []*YaPiGroup
		if err := json.Unmarshal(j, &groups); err != nil {
			logger.Painc("分析文件%v,失败,%s", path, err)
		}
		//分析文件
		for _, group := range groups {
			for i, list := range group.List {
				if list.ReqBodyOtherJson != "" {
					list.ReqBodyOther = new(YaPiListReqBodyOther)
					if err := json.Unmarshal([]byte(list.ReqBodyOtherJson), &list.ReqBodyOther); err != nil {
						logger.Painc("分析%req_body_other失败", list.Title)
					} else {
						group.List[i].ReqBodyOther = list.ReqBodyOther
					}
				}
			}
			for _, item := range group.decompose() {
				parser.Route = append(parser.Route, item)
			}
		}
	}
	return parser
}

type YaPiGroup struct {
	Name string      `json:"name"`
	List []*YaPiList `json:"list"`
}

func (it *YaPiGroup) decompose() (routes []Route) {
	for _, list := range it.List {
		route := new(Route)
		route.Name = list.Title
		route.Method = Method(list.Method)
		route.Path = list.Path

		reqParams := make([]*QueryParam, 0)
		for _, reqParam := range list.ReqQuery {
			queryParam := new(QueryParam)
			if reqParam.Required == "1" {
				queryParam.Required = true
			}
			queryParam.Name = reqParam.Name
			queryParam.Desc = reqParam.Desc
			desEach := strings.Split(queryParam.Desc, "|")
			switch len(desEach) {
			case 0, 1:
				queryParam.Type = "string"
			case 2:
				queryParam.Type = desEach[0]
			case 3:
				queryParam.Type = desEach[0]
				queryParam.Desc = reqParam.Desc
			default:
				queryParam.Type = "string"
				queryParam.Desc = ""
			}
			reqParams = append(reqParams, queryParam)
		}
		route.QueryParams = reqParams
		route.ModuleTitle = it.Name
		route.ActionTitle = list.Title
		route.SetId = strings.Contains(route.Path, "{id}")
		requestParam := make([]*RequestParam, 0)
		if list.ReqBodyOther != nil {
			//需要存储的参数
			for name, property := range list.ReqBodyOther.Properties {
				t := new(RequestParam)
				t.Name = name
				t.Desc = property.Description
				t.Type = property.Type
				for _, requireName := range list.ReqBodyOther.Required {
					if t.Name == requireName {
						t.Required = true
						break
					}
				}
				requestParam = append(requestParam, t)
			}
		}

		route.RequestParam = requestParam
		paths := strings.Split(strings.Trim(route.Path, "/"), "/")
		switch len(paths) {
		case 3:
			route.ActionName = paths[2]
		case 2:
			if strings.Contains(route.Path, "{id}") {
				route.ActionName = ""
			} else {
				route.ActionName = paths[1]
			}
		default:
			route.ActionName = ""
		}
		route.ModuleName = paths[0]
		routes = append(routes, *route)
	}

	return routes
}

type YaPiList struct {
	Path   string `json:"path"`
	Method string `json:"method"`
	Title  string `json:"title"`
	//查询参数
	ReqQuery         []*YaPiListReqQuery `json:"req_query"`
	ReqBodyOtherJson string              `json:"req_body_other"`
	ReqBodyOther     *YaPiListReqBodyOther
}

//查询参数
type YaPiListReqQuery struct {
	Required string `json:"required"`
	Name     string `json:"name"`
	Desc     string `json:"desc"`
}

//存储参数
type YaPiListReqBodyOther struct {
	Properties map[string]YapiProperties `json:"properties"`
	Required   []string                  `json:"required"`
}
type YapiProperties struct {
	Type        string `json:"type"`
	Description string `json:"description"`
}
