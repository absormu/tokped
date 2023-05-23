package job

import (
	"net/http"

	"github.com/absormu/tokped/app/entity"
	md "github.com/absormu/tokped/app/middleware"
	repojob "github.com/absormu/tokped/app/repository/job"
	pg "github.com/absormu/tokped/pkg/pagination"
	lg "github.com/absormu/tokped/pkg/response"
	resp "github.com/absormu/tokped/pkg/response"
	sdk "github.com/absormu/tokped/pkg/sdk"
	"github.com/labstack/echo/v4"
)

func GetJobList(c echo.Context, extractToken entity.ExtractToken) (e error) {
	logger := md.GetLogger(c)
	logger.WithField("request", extractToken).Info("usecase: GetJobList")

	description := c.QueryParam("description")
	location := c.QueryParam("location")
	fullTime := c.QueryParam("full_time")

	var jobLists []entity.JobList
	jobLists, e = repojob.RequestJobList(c, description, location, fullTime)
	if e != nil {
		logger.WithField("error", e.Error()).Error("Catch error RequestJobList request")
		return
	}

	total := int64(len(jobLists))
	var responseData []entity.JobList
	var data entity.JobList
	for _, jobList := range jobLists {
		data.ID = jobList.ID
		data.Type = jobList.Type
		data.Url = jobList.Url
		data.CreatedAt = jobList.CreatedAt
		data.Company = jobList.Company
		data.CompanyUrl = jobList.CompanyUrl
		data.Location = jobList.Location
		data.Title = jobList.Title
		data.Description = jobList.Description
		data.HowToApply = jobList.HowToApply
		data.CompanyLogo = jobList.CompanyLogo
		responseData = append(responseData, data)
	}

	meta, e := pg.Pagination(c)
	if e != nil {
		logger.WithField("error", e.Error()).Error("Catch error Pagination")
		e = resp.CustomError(c, http.StatusBadRequest, sdk.ERR_PARAM_ILLEGAL,
			lg.Language{Bahasa: nil, English: "Bad Request"}, nil, nil)
		return
	}

	metaPagination := pg.GenerateMeta(c, total, meta.Limit, meta.Page, meta.Offset, meta.Pagination, nil)

	e = resp.CustomError(c, http.StatusOK, sdk.ERR_SUCCESS,
		lg.Language{Bahasa: "Sukses", English: "Success"}, metaPagination, responseData)
	return
}

func GetJobDetail(c echo.Context, extractToken entity.ExtractToken) (e error) {
	logger := md.GetLogger(c)
	logger.WithField("request", extractToken).Info("usecase: GetJobDetail")

	idStr := c.Param("id")

	if e != nil {
		logger.WithField("error", e.Error()).Error("Catch error Pagination")
		e = resp.CustomError(c, http.StatusBadRequest, sdk.ERR_PARAM_ILLEGAL,
			lg.Language{Bahasa: nil, English: "Bad Request"}, nil, nil)
		return
	}

	var jobDetail entity.JobList
	var jobDetailEmpty entity.JobList
	jobDetail, e = repojob.RequestJobDetail(c, idStr)
	if e != nil {
		logger.WithField("error", e.Error()).Error("Catch error RequestJobDetail request")
		return
	}

	if jobDetail == jobDetailEmpty {
		logger.Error("Catch error jobDetail id not found")
		e = resp.CustomError(c, http.StatusNotFound, sdk.ERR_USER_NOT_FOUND,
			lg.Language{Bahasa: "jobDetail id tidak tersedia", English: "jobDetail id not found"}, nil, nil)
		return
	}

	var data entity.JobList
	data.ID = jobDetail.ID
	data.Type = jobDetail.Type
	data.Url = jobDetail.Url
	data.CreatedAt = jobDetail.CreatedAt
	data.Company = jobDetail.Company
	data.CompanyUrl = jobDetail.CompanyUrl
	data.Location = jobDetail.Location
	data.Title = jobDetail.Title
	data.Description = jobDetail.Description
	data.HowToApply = jobDetail.HowToApply
	data.CompanyLogo = jobDetail.CompanyLogo

	e = resp.CustomError(c, http.StatusOK, sdk.ERR_SUCCESS,
		lg.Language{Bahasa: "Sukses", English: "Success"}, nil, data)
	return
}
