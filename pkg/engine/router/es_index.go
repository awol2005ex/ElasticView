package router

import (
	. "github.com/1340691923/ElasticView/api"
	"github.com/1340691923/ElasticView/pkg/api_config"
	. "github.com/gofiber/fiber/v2"
)

// ES索引 路由
func runEsIndex(app *App) {
	apiRouterConfig := api_config.NewApiRouterConfig()
	const AbsolutePath = "/api/es_index"
	esIndex := app.Group(AbsolutePath)
	{
		esIndexController := EsIndexController{}
		apiRouterConfig.MountApi(api_config.MountApiBasePramas{
			Remark:       "删除ES索引",
			Method:       api_config.MethodPost,
			AbsolutePath: AbsolutePath,
			RelativePath: "DeleteAction",
		}, esIndex.(*Group), true, esIndexController.DeleteAction)

		apiRouterConfig.MountApi(api_config.MountApiBasePramas{
			Remark:       "创建ES索引",
			Method:       api_config.MethodPost,
			AbsolutePath: AbsolutePath,
			RelativePath: "CreateAction",
		}, esIndex.(*Group), true, esIndexController.CreateAction)

		apiRouterConfig.MountApi(api_config.MountApiBasePramas{
			Remark:       "获取ES索引配置",
			Method:       api_config.MethodPost,
			AbsolutePath: AbsolutePath,
			RelativePath: "GetSettingsAction",
		}, esIndex.(*Group), true, esIndexController.GetSettingsAction)

		apiRouterConfig.MountApi(api_config.MountApiBasePramas{
			Remark:       "得到所有的ES索引名",
			Method:       api_config.MethodPost,
			AbsolutePath: AbsolutePath,
			RelativePath: "IndexNamesAction",
		}, esIndex.(*Group), true, esIndexController.IndexNamesAction)

		apiRouterConfig.MountApi(api_config.MountApiBasePramas{
			Remark:       "重建索引",
			Method:       api_config.MethodPost,
			AbsolutePath: AbsolutePath,
			RelativePath: "ReindexAction",
		}, esIndex.(*Group), true, esIndexController.ReindexAction)

		apiRouterConfig.MountApi(api_config.MountApiBasePramas{
			Remark:       "获取别名",
			Method:       api_config.MethodPost,
			AbsolutePath: AbsolutePath,
			RelativePath: "GetAliasAction",
		}, esIndex.(*Group), true, esIndexController.GetAliasAction)

		apiRouterConfig.MountApi(api_config.MountApiBasePramas{
			Remark:       "迁移别名到其他索引",
			Method:       api_config.MethodPost,
			AbsolutePath: AbsolutePath,
			RelativePath: "MoveAliasToIndex",
		}, esIndex.(*Group), true, esIndexController.MoveAliasToIndex)

		apiRouterConfig.MountApi(api_config.MountApiBasePramas{
			Remark:       "添加别名",
			Method:       api_config.MethodPost,
			AbsolutePath: AbsolutePath,
			RelativePath: "AddAliasToIndex",
		}, esIndex.(*Group), true, esIndexController.AddAliasToIndex)

		apiRouterConfig.MountApi(api_config.MountApiBasePramas{
			Remark:       "批量添加别名",
			Method:       api_config.MethodPost,
			AbsolutePath: AbsolutePath,
			RelativePath: "BatchAddAliasToIndex",
		}, esIndex.(*Group), true, esIndexController.BatchAddAliasToIndex)

		apiRouterConfig.MountApi(api_config.MountApiBasePramas{
			Remark:       "移除别名",
			Method:       api_config.MethodPost,
			AbsolutePath: AbsolutePath,
			RelativePath: "RemoveAlias",
		}, esIndex.(*Group), true, esIndexController.RemoveAlias)

		apiRouterConfig.MountApi(api_config.MountApiBasePramas{
			Remark:       "获取所有的索引配置信息",
			Method:       api_config.MethodPost,
			AbsolutePath: AbsolutePath,
			RelativePath: "GetSettingsInfoAction",
		}, esIndex.(*Group), true, esIndexController.GetSettingsInfoAction)

		apiRouterConfig.MountApi(api_config.MountApiBasePramas{
			Remark:       "获取索引的Stats",
			Method:       api_config.MethodPost,
			AbsolutePath: AbsolutePath,
			RelativePath: "StatsAction",
		}, esIndex.(*Group), true, esIndexController.StatsAction)

		apiRouterConfig.MountApi(api_config.MountApiBasePramas{
			Method:       api_config.MethodPost,
			AbsolutePath: AbsolutePath,
			RelativePath: "IndexsCountAction",
		}, esIndex.(*Group), false, esIndexController.IndexsCountAction)
	}
}
