package cmd

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/crusj/logger"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
)

//yapi接口文档,请求实现ActInterface
type YaPi struct {
	Path,
	Method,
	Name,
	Module,
	//module title
	Mt,
	Action,
	//action title
	At string
	Request  Request
	Response Response
}

func (yapi YaPi) MethodName() (Method, error) {
	method := NewMethod(yapi.Method)
	if method == InvalidMethod {
		return InvalidMethod, errors.New(InvalidMethod.String())
	}
	return method, nil
}
func (yapi YaPi) ModuleName() (string, error) {
	if yapi.Module == "" {
		return "", errors.New(fmt.Sprintf("%s %s", yapi.Path, "module name为空"))
	}
	return yapi.Module, nil
}
func (yapi YaPi) ActionName() (string, error) {
	return yapi.Action, nil
}
func (yapi YaPi) RoutePath() (string, error) {
	if yapi.Path == "" {
		return "", errors.New(fmt.Sprintf("%s", "route  path为空"))
	}
	return yapi.Path, nil
}
func (yapi YaPi) SetId() bool {
	reg, _ := regexp.Compile(`.*\{id\}.*`)
	find := reg.FindString(yapi.Path)
	return find != ""
}
func (yapi YaPi) ResponseType() (ResponseType, error) {
	return Null, nil
}
func (yapi YaPi) Params() []*Param {
	return nil
}
func (yapi YaPi) ModuleTitle() string {
	return yapi.Mt
}
func (yapi YaPi) ActionTitle() string {
	return yapi.At
}

type YaPiGroup struct {
	Name string      `json:"name"`
	List []*YaPiList `json:"list"`
}

func (it *YaPiGroup) decompose() (yapis []YaPi) {
	for _, list := range it.List {
		yapi := new(YaPi)
		yapi.Name = list.Title
		yapi.Method = list.Method
		yapi.Path = list.Path

		yapi.Mt = it.Name
		yapi.At = list.Title

		paths := strings.Split(strings.Trim(yapi.Path, "/"), "/")
		if len(paths) == 3 {
			yapi.Action = paths[2]
		}
		yapi.Module = paths[0]
		yapis = append(yapis, *yapi)
	}

	return yapis
}

type YaPiList struct {
	Path   string `json:"path"`
	Method string `json:"method"`
	Title  string `json:"title"`
}
type yaPiParse struct {
	routePath *RoutePath
}

func NewYaPiPath(routeFilePath, routeStart, routeEnd, routeStartTag, routeEndTag string) *yaPiParse {
	rp := &RoutePath{
		Path:     routeFilePath,
		Start:    routeStart,
		StartTag: routeStartTag,
		End:      routeEnd,
		EndTag:   routeEndTag,
	}
	return &yaPiParse{routePath: rp}
}

func (it *yaPiParse) RoutePath() *RoutePath {
	return it.routePath
}
func (it *yaPiParse) Parse(routeJsonPath string) []ActionInterface {
	f, err := os.Open(routeJsonPath)
	if err != nil {
		logger.Painc("打开文件%v,失败,%s", routeJsonPath, err)
	}
	defer f.Close()

	if j, err := ioutil.ReadAll(f); err != nil {
		logger.Painc("读取文件%v,失败,%s", routeJsonPath, err)
	} else {
		var groups []*YaPiGroup
		var ret []ActionInterface
		if err := json.Unmarshal(j, &groups); err != nil {
			logger.Painc("分析文件%v,失败,%s", routeJsonPath, err)
		}
		//分析文件
		for _, group := range groups {
			for _, item := range group.decompose() {
				ret = append(ret, item)
			}
		}
		return ret
	}
	return nil
}
