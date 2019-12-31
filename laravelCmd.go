package main

import (
	"github.com/crusj/laravelCmd/cmd"
	_ "github.com/crusj/laravelCmd/init"
	"github.com/crusj/laravelCmd/writer"
	"gopkg.in/alecthomas/kingpin.v2"
	"os"
)

var (
	app                 = kingpin.New("laravelcmd", "laravel cmd tools")
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
)

func main() {
	switch kingpin.MustParse(app.Parse(os.Args[1:])) {
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
