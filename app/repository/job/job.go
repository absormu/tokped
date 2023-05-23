package job

import (
	"encoding/json"

	"github.com/absormu/tokped/app/entity"
	md "github.com/absormu/tokped/app/middleware"
	cm "github.com/absormu/tokped/pkg/configuration"
	sdk "github.com/absormu/tokped/pkg/sdk"
	"github.com/labstack/echo/v4"
)

func RequestJobList(c echo.Context, description, location, fullTime string) (res []entity.JobList, e error) {

	logger := md.GetLogger(c)
	logger.WithField("request", "").Info("repository: RequestJobList")

	queryParams := map[string]string{}
	if description != "" {
		queryParams["description"] = description
	}
	if location != "" {
		queryParams["location"] = location
	}
	if fullTime == "true" {
		queryParams["type"] = "Full Time"
	}

	rawResponse, e := sdk.RawGetRequest(logger, cm.Config.JobListUrl, cm.Config.Timeout, queryParams)
	if e != nil {
		return
	}

	if e = json.Unmarshal(rawResponse, &res); e != nil {
		return
	}

	return
}

func RequestJobDetail(c echo.Context, id string) (res entity.JobList, e error) {

	logger := md.GetLogger(c)
	logger.WithField("request", "").Info("repository: RequestJobDetail")

	queryParams := map[string]string{}
	if id != "" {
		queryParams["id"] = id
	}

	rawResponse, e := sdk.RawGetRequestV2(logger, cm.Config.JobDetaillUrl, cm.Config.Timeout, id)
	if e != nil {
		return
	}

	if e = json.Unmarshal(rawResponse, &res); e != nil {
		return
	}

	return
}
