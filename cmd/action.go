package cmd

import (
	"fmt"
	myfile "github.com/crusj/file"
	"github.com/crusj/logger"
	"github.com/stoewer/go-strcase"
	"strconv"
	"strings"
)

//请求方法
type Method int

const (
	InvalidMethod Method = iota
	Get
	POST
	PUT
	DELETE
)

var MethodNames = []string{
	"INVALID",
	"GET",
	"POST",
	"PUT",
	"DELETE",
}

func (method Method) String() string {
	if int(method) < len(MethodNames) {
		return MethodNames[method]
	}
	return ""
}
func NewMethod(method string) Method {
	method = strings.ToUpper(method)
	for i, m := range MethodNames[1:] {
		if m == method {
			return Method(i + 1)
		}
	}
	return InvalidMethod
}

//请求内容
type Request struct {
	DataFormat interface{}
}

//返回内容
type Response struct {
	//返回内容数据类型
	Type ResponseType
	//返回内容格式
	DataFormat interface{}
}

//response类型
type ResponseType int

const (
	InvalidResponseType ResponseType = iota
	Array
	Object
	Null
)

var ResponseTypeNames = []string{
	"INVALID",
	"ARRAY",
	"OBJECT",
	"NULL",
}

func (res ResponseType) String() string {
	if int(res) < len(ResponseTypeNames) {
		return ResponseTypeNames[res]
	}
	return "ResponseType" + strconv.Itoa(int(res))
}

type Module struct {
	Name string
}

type ActionInterface interface {
	RoutePath() (string, error)
	MethodName() (Method, error)
	ModuleName() (string, error)
	ActionName() (string, error)
	ModuleTitle() string
	ActionTitle() string
	SetId() bool
	ResponseType() (ResponseType, error)
	Params() []*Param
}

type Param struct {
	Name string
	Type string
}

type actBody struct {
	Indent int
	Path,
	ModuleName,
	ModuleTitle,
	ActionName,
	ActionTitle string
	Method
	SetId  bool
	Params []*Param
	ResponseType
}

//普通路由
func (actBody actBody) String() string {
	routeTpl := "%s//%s\n%sRoute::%s('%s','%s');"

	if actBody.ActionName == "" {
		switch strings.ToUpper(actBody.Method.String()) {
		case "GET":
			if actBody.SetId == true {
				actBody.ActionName = "show"
			} else {
				actBody.ActionName = "index"
			}
		case "POST":
			actBody.ActionName = "store"
		case "PUT":
			if actBody.SetId == true {
				actBody.ActionName = "update"
			}
		case "DELETE":
			if actBody.SetId == true {
				actBody.ActionName = "destroy"
			}
		}
	}
	tmp := fmt.Sprintf(routeTpl,
		strings.Repeat(" ", actBody.Indent|8),
		actBody.ActionTitle,
		strings.Repeat(" ", actBody.Indent|8),
		strings.ToLower(actBody.Method.String()),
		actBody.Path,
		fmt.Sprintf("%s@%s", strcase.UpperCamelCase(actBody.ModuleName), actBody.ActionName))
	return tmp
}

type RoutePath struct {
	Path     string
	Start    string
	StartTag string
	End      string
	EndTag   string
}

func (routePath *RoutePath) addStartTag(line int, content string) string {
	if content == routePath.Start {
		if routePath.Start == "" {
			routePath.StartTag = "route_start"
		}
		return routePath.StartTag
	}
	return ""
}
func (routePath *RoutePath) addEndTag(line int, content string) string {
	if content == routePath.End {
		if routePath.End == "" {
			routePath.EndTag = "route_start"
		}
		return routePath.EndTag
	}
	return ""
}

type actionGroup struct {
	Indent int
	Name,
	Title string
	Actions []*actBody
}
type ActGroupMap map[string]*actionGroup

//路由组转化为将插入到路由文件的值
func (actGroupMap ActGroupMap) String() []string {
	routes := make([]string, 0)
	for _, actGroup := range actGroupMap {
		groupTpl := "%s//%s\n%sRoute::prefix('%s')->group(function(){\n%s\n%s})"
		abStr := make([]string, 0)
		for _, ab := range actGroup.Actions {
			abStr = append(abStr, ab.String())
		}
		routes = append(routes, fmt.Sprintf(groupTpl,
			strings.Repeat(" ", actGroup.Indent|4),
			actGroup.Title,
			strings.Repeat(" ", actGroup.Indent|4),
			actGroup.Name, strings.Join(abStr, "\n"),
			strings.Repeat(" ", actGroup.Indent|4),
		) )
	}
	return routes
}

//core
type actions struct {
	//分组缩进空格数
	groupIndent int
	//方法缩进空格数
	actionIndent int
	groups       ActGroupMap
	//路由路径
	routePath *RoutePath
	//路由
	routes []string
}

func (actions *actions) PrintRoutes() {
	for index, routeStr := range actions.routes {
		fmt.Printf("%d %s\n", index, routeStr)
	}
}

//创建实例
func newActions(acts []ActionInterface) (*actions, error) {
	actions := new(actions)
	actions.groups = make(ActGroupMap)
	for _, act := range acts {
		body := new(actBody)
		var err error

		//是否存在主键
		body.SetId = act.SetId()
		//路由地址
		body.Path, err = act.RoutePath()
		if err != nil {
			return nil, err
		}
		//method
		body.Method, err = act.MethodName()
		if err != nil {
			return nil, err
		}

		body.ModuleName, err = act.ModuleName()
		if err != nil {
			return nil, err
		}
		body.ModuleTitle = act.ModuleTitle()

		body.ActionName, err = act.ActionName()
		if err != nil {
			return nil, err
		}
		body.ActionTitle = act.ActionTitle()

		body.Params = act.Params()
		body.ResponseType, err = act.ResponseType()
		if err != nil {
			return nil, err
		}

		moduleName := ""
		if strings.Contains(body.ModuleName, "_") {
			moduleName = strings.Split(body.ModuleName, "_")[0]
		} else {
			moduleName = body.ModuleName
		}
		if _, exists := actions.groups[moduleName]; exists {
			actions.groups[moduleName].Actions = append(actions.groups[moduleName].Actions, body)
		} else {
			actions.groups[moduleName] = &actionGroup{
				Name:    body.ModuleName,
				Title:   body.ModuleTitle,
				Actions: []*actBody{body},
			}
		}
	}
	return actions, nil
}

//分析组到路由文件
func (actions *actions) AnalyzeGroup() {
	actions.routes = actions.groups.String()
}

//路由文件插入路由
func (actions *actions) InsertRoute() {
	//TODO 重复校验
	file, err := myfile.NewFile(actions.routePath.Path)
	if err != nil {
		logger.Painc(err)
	}
	file.Scan([]myfile.AddTag{actions.routePath.addStartTag, actions.routePath.addEndTag}...)
	if err := file.InsertBetween(actions.routePath.StartTag, actions.routePath.EndTag, actions.routes); err != nil {
		logger.Painc(err)
	}
	logger.Info("写入路由成功")
}

type ParseActions interface {
	Parse(routeJsonPath string) []ActionInterface
	RoutePath() *RoutePath
}

func NewActions(routeJsonPath string, parseActions ParseActions) *actions {
	acts := parseActions.Parse(routeJsonPath)
	actions, err := newActions(acts)
	if err != nil {
		logger.Painc("创建初始路由信息失败", err)
	}
	actions.routePath = parseActions.RoutePath()
	return actions
}
