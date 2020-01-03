package main

import (
	"github.com/crusj/laravelCmd/cmd"
	_ "github.com/crusj/laravelCmd/init"
	"github.com/crusj/laravelCmd/writer"
	"gopkg.in/alecthomas/kingpin.v2"
	"os"
)

var (
	app                 = kingpin.New("laravelCmd", "laravel cmd tools")
	admin               = app.Command("admin", "laravel admin tool")
	adminConfig         = admin.Flag("config", "config file path").Default("admin.json").String()
	adminRoutePath      = admin.Flag("path", "route file path").Default("app/Admin/routes.php").String()
	adminControllerPath = admin.Flag("c", "controller dir").Default("app/Admin/Controllers/").String()
	adminRouteStart     = admin.Flag("s", "route start flag").Default("//route_start").String()
	adminRouteEnd       = admin.Flag("e", "route end flag").Default("//route_end").String()
	adminNumber         = admin.Command("number", "write route of number")
	adminNu             = adminNumber.Arg("number", "route table number").Required().Int()
	adminAll            = admin.Command("all", " write all routes")
	adminPrint          = admin.Command("list", "list route table")

	apiRoute      = app.Command("route", "analysis api document and write routes to api route file")
	apiRoutePath  = apiRoute.Flag("path", "api route file path").Default("routes/api.php").String()
	apiRouteStart = apiRoute.Flag("s", "route start flag").Default("//route_start").String()
	apiRouteEnd   = apiRoute.Flag("e", "route end flag").Default("//route_end").String()
	apiDocument   = apiRoute.Flag("config", "api document path").Default("api.json").String()
	apiParser     = apiRoute.Flag("parser", "api parser").Default("yapi").String()

	//服务
	service         = app.Command("service", "analysis api document and write routes to every service file")
	serviceDocument = service.Flag("config", "service document path").Default("api.json").String()
	servicePath     = service.Flag("path", "service path").Default("app/Http/Services/").String()
	serviceStart    = service.Flag("s", "service start flag").Default("//service_start").String()
	serviceEnd      = service.Flag("e", "service end flag").Default("//service_end").String()
	serviceParser   = service.Flag("parser", "api parser").Default("yapi").String()

	//控制器
	controller         = app.Command("controller", "analysis api document and write routes to every controller file")
	controllerDocument = controller.Flag("config", "controller document path").Default("api.json").String()
	controllerPath     = controller.Flag("path", "controller path").Default("app/Http/Controllers/").String()
	controllerStart    = controller.Flag("s", "controller start flag").Default("//controller_start").String()
	controllerEnd      = controller.Flag("e", "controller end flag").Default("//controller_end").String()
	controllerParser   = controller.Flag("parser", "api parser").Default("yapi").String()

	//request
	request          = app.Command("request", "analysis api document and write routes to every request file")
	requestDocument  = request.Flag("config", "request document path").Default("api.json").String()
	requestPath      = request.Flag("path", "request path").Default("app/Http/Requests/").String()
	requestParser    = request.Flag("parser", "api parser").Default("yapi").String()
	requestRule      = request.Command("rule", "write rules to request file")
	requestRuleStart = requestRule.Flag("s", "request rule start flag").Default("//rule_start").String()
	requestRuleEnd   = requestRule.Flag("e", "request rule end flag").Default("//rule_end").String()

	requestAttribute      = request.Command("attribute", "write attribtue to request file")
	requestAttributeStart = requestAttribute.Flag("s", "request attribute start flag").Default("//attribute_start").String()
	requestAttributeEnd   = requestAttribute.Flag("e", "request attribute end flag").Default("//attribute_end").String()
)

func main() {
	switch kingpin.MustParse(app.Parse(os.Args[1:])) {
	//写入规则属性
	case requestAttribute.FullCommand():
		if *requestParser == "yapi" {
			requestAttributeWriter := writer.NewRequestAttributeWriter(*requestPath, *requestAttributeStart, *requestAttributeEnd)
			routeParser := writer.NewRoutesParser(*requestDocument)
			err := writer.Writes(routeParser, requestAttributeWriter)
			if err != nil {
				kingpin.Errorf("%s\n", err)
			}
		} else {
			kingpin.Fatalf("不支持的文档解析器%s,当前仅支持yapi文档\n", *apiParser)
		}
	//写入规则
	case requestRule.FullCommand():
		if *requestParser == "yapi" {
			requestRuleWriter := writer.NewRequestRuleWriter(*requestPath, *requestRuleStart, *requestRuleEnd)
			routeParser := writer.NewRoutesParser(*requestDocument)
			err := writer.Writes(routeParser, requestRuleWriter)
			if err != nil {
				kingpin.Errorf("%s\n", err)
			}
		} else {
			kingpin.Fatalf("不支持的文档解析器%s,当前仅支持yapi文档\n", *apiParser)
		}
	//向controller写入方法
	case controller.FullCommand():
		if *controllerParser == "yapi" {
			routeParser := writer.NewRoutesParser(*controllerDocument)
			controllerWriter := writer.NewControllerWriter(*controllerPath, *controllerStart, *controllerEnd)
			err := writer.Writes(routeParser, controllerWriter)
			if err != nil {
				kingpin.Errorf("%s\n", err)
			}
		} else {
			kingpin.Fatalf("不支持的文档解析器%s,当前仅支持yapi文档\n", *apiParser)
		}
	//向service文件中写入方法
	case service.FullCommand():
		if *serviceParser == "yapi" {
			routeParser := writer.NewRoutesParser(*serviceDocument)
			serviceWriter := writer.NewServiceWriter(*servicePath, *serviceStart, *serviceEnd)
			err := writer.Writes(routeParser, serviceWriter)
			if err != nil {
				kingpin.Errorf("%s\n", err)
			}
		} else {
			kingpin.Fatalf("不支持的文档解析器%s,当前仅支持yapi文档\n", *apiParser)
		}
	//打印所有路由
	case adminPrint.FullCommand():
		laravelAdmin, err := cmd.NewLaravelAdmin(
			*adminConfig,
			*adminRoutePath,
			*adminControllerPath,
			*adminRouteStart,
			*adminRouteEnd,
		)
		if err != nil {
			kingpin.Errorf("%s\n", err)
		} else {
			laravelAdmin.List()
		}
	//生成指定序号路由
	case adminNumber.FullCommand():
		laravelAdmin, err := cmd.NewLaravelAdmin(
			*adminConfig,
			*adminRoutePath,
			*adminControllerPath,
			*adminRouteStart,
			*adminRouteEnd,
		)
		if err != nil {
			kingpin.Errorf("%s\n", err)
		} else {
			laravelAdmin.Make(*adminNu)
		}
	//生成指定序号路由
	case adminAll.FullCommand():
		laravelAdmin, err := cmd.NewLaravelAdmin(
			*adminConfig,
			*adminRoutePath,
			*adminControllerPath,
			*adminRouteStart,
			*adminRouteEnd,
		)
		if err != nil {
			kingpin.Errorf("%s\n", err)
		} else {
			laravelAdmin.MakeAll()
		}

	//api路由文件生成工具
	case apiRoute.FullCommand():
		if *apiParser == "yapi" {
			routeParser := writer.NewRoutesParser(*apiDocument)
			routeWriter := writer.NewRoutesWriter(*apiRoutePath, *apiRouteStart, *apiRouteEnd)
			err := writer.Write(routeParser, routeWriter)
			if err != nil {
				kingpin.Errorf("%s\n", err)
			}
		} else {
			kingpin.Fatalf("不支持的文档解析器%s,当前仅支持yapi文档\n", *apiParser)
		}
	}
}
