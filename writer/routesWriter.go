package writer

import (
	"fmt"
	"github.com/stoewer/go-strcase"
	"strings"
)

type routesWriter struct {
	RoutePath string
	StartFlag,
	EndFlag string
}

func (routesWriter *routesWriter) FilePath() string {
	return routesWriter.RoutePath
}
func (routesWriter *routesWriter) Between() *Between {
	return &Between{
		StartFlag: routesWriter.StartFlag,
		EndFlag:   routesWriter.EndFlag,
	}
}
func (routesWriter *routesWriter) Content(routes []Route) []string {
	return routesWriter.Group(routes).String()
}
func NewRoutesWriter(RoutePath, startFlag, endFlag string) *routesWriter {
	object := new(routesWriter)
	object.RoutePath = RoutePath
	object.StartFlag = startFlag
	object.EndFlag = endFlag
	return object
}

type actionGroup struct {
	Indent int
	Name,
	Title string
	Actions []*actBody
}

//路由组转化为将插入到路由文件的值
func (actGroupMap actGroupMap) String() []string {
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
		))
	}
	return routes
}

type actBody struct {
	Indent int
	Path,
	ModuleName,
	ModuleTitle,
	ActionName,
	ActionTitle string
	Method
	SetId bool
}

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

type actGroupMap map[string]*actionGroup

func (routesWriter *routesWriter) Group(routes []Route) actGroupMap {
	actGroupMap := make(actGroupMap)
	for _, route := range routes {
		body := new(actBody)
		//是否存在主键
		body.SetId = route.SetId
		//路由地址
		body.Path = route.Path
		//method
		body.Method = route.Method

		body.ModuleName = route.ModuleName
		body.ModuleTitle = route.ModuleTitle

		body.ActionName = route.ActionName
		body.ActionTitle = route.ActionTitle

		moduleName := ""
		if strings.Contains(body.ModuleName, "_") {
			moduleName = strings.Split(body.ModuleName, "_")[0]
		} else {
			moduleName = body.ModuleName
		}
		if _, exists := actGroupMap[moduleName]; exists {
			actGroupMap[moduleName].Actions = append(actGroupMap[moduleName].Actions, body)
		} else {
			actGroupMap[moduleName] = &actionGroup{
				Name:    body.ModuleName,
				Title:   body.ModuleTitle,
				Actions: []*actBody{body},
			}
		}
	}
	return actGroupMap
}
