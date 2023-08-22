package cluser_settings_service

import (
	"context"
	"github.com/1340691923/ElasticView/es_sdk/pkg"
	"github.com/tidwall/gjson"
	"net/http"
)

type ClusterSettingsService struct {
	esClient pkg.EsI
}

func NewClusterSettingsService(esClient pkg.EsI) *ClusterSettingsService {
	return &ClusterSettingsService{esClient: esClient}
}

func (this *ClusterSettingsService) GetSettings(ctx context.Context) (resByte []byte, err error) {

	req, err := http.NewRequest("GET", "/_cluster/settings?include_defaults=true", nil)

	if err != nil {
		return nil, err
	}

	res, err := this.esClient.PerformRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	if res.StatusErr() != nil {
		return nil, res.StatusErr()
	}

	return res.ResByte(), nil
}

func (this *ClusterSettingsService) GetPathRepo(settingsBytes []byte) (pathRepo []interface{}) {

	gjson.GetBytes(settingsBytes, "defaults.path.repo").ForEach(func(key, value gjson.Result) bool {
		pathRepo = append(pathRepo, value.Value())
		return true
	})

	return
}

func (this *ClusterSettingsService) GetAllowedUrls(settingsBytes []byte) (allowedUrls []interface{}) {

	gjson.GetBytes(settingsBytes, "defaults.repositories.url.allowed_urls").ForEach(func(key, value gjson.Result) bool {
		allowedUrls = append(allowedUrls, value.Value())
		return true
	})

	return
}
