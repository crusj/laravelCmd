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

		route.ModuleTitle = it.Name
		route.ActionTitle = list.Title
		route.SetId = strings.Contains(route.Path, "{id}")
		paths := strings.Split(strings.Trim(route.Path, "/"), "/")
		if len(paths) == 3 {
			route.ActionName = paths[2]
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
}
