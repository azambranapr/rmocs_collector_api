package batch

import (
	"errors"
	"github.com/azambranapr/rmocs_collector_api/src/data/models/batch"
	"github.com/azambranapr/rmocs_collector_api/src/data/models/generalResponse"
	"github.com/azambranapr/rmocs_collector_api/src/domain/repositories/batchService"
	"github.com/dustin/go-humanize"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

var (
	ErrNotFound  = errors.New("data not found")
	ErrEmpty     = errors.New("data table is empty")
	ErrNoChanges = errors.New("there are no changes or data not found")
)

type DbRepository struct {
	repository batchService.IdbService
}

func New(repo batchService.IdbService) *DbRepository {
	return &DbRepository{
		repository: repo,
	}
}

func (r *DbRepository) GetAll(ctx *gin.Context) {
	batchModels, err := r.repository.GetAll()

	if err != nil {
		resp := generalResponse.NewResponse(generalResponse.Error, err.Error(), "0", err)
		ctx.IndentedJSON(http.StatusInternalServerError, resp)
		return
	}

	if len(batchModels) <= 0 {
		response := generalResponse.NewResponse(generalResponse.Error, ErrEmpty.Error(), "0", nil)
		ctx.IndentedJSON(http.StatusBadRequest, response)
		return
	}
	var it = int64(len(batchModels))
	response := generalResponse.NewResponse(generalResponse.Message, "ok", humanize.Comma(it), batchModels)

	ctx.IndentedJSON(http.StatusOK, response)
}
func (r *DbRepository) GetById(c *gin.Context) {
	snum := c.Param("id")

	num, err := strconv.Atoi(snum)
	if err != nil {
		response := generalResponse.NewResponse(generalResponse.Error, err.Error(), "0", err)
		c.IndentedJSON(http.StatusBadRequest, response)
		return
	}
	ms, err := r.repository.GetById(num)
	if err != nil {
		response := generalResponse.NewResponse(generalResponse.Error, err.Error(), "0", err)
		c.IndentedJSON(http.StatusBadRequest, response)
		return
	}

	if len(ms) <= 0 {
		response := generalResponse.NewResponse(generalResponse.Error, ErrNotFound.Error(), "0", nil)
		c.IndentedJSON(http.StatusBadRequest, response)
		return
	}

	it := int64(len(ms))
	response := generalResponse.NewResponse(generalResponse.Message, "ok", humanize.Comma(it), ms)
	c.IndentedJSON(http.StatusOK, response)
}
func (r *DbRepository) Create(ctx *gin.Context) {
	var batchModel *batch.Model

	if err := ctx.BindJSON(&batchModel); err != nil {
		resp := generalResponse.NewResponse(generalResponse.Error, err.Error(), "0", err)
		ctx.IndentedJSON(http.StatusInternalServerError, resp)
		return
	}

	if err := r.repository.Create(batchModel); err != nil {
		resp := generalResponse.NewResponse(generalResponse.Error, err.Error(), "0", err)
		ctx.IndentedJSON(http.StatusInternalServerError, resp)
		return
	}

	response := generalResponse.NewResponse("Insert-"+generalResponse.Message, "ok", "1", batchModel)
	ctx.IndentedJSON(http.StatusCreated, response)

}
