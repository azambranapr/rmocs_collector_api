package routers

import (
	"github.com/azambranapr/rmocs_collector_api/src/core/models/version"
	"github.com/azambranapr/rmocs_collector_api/src/data/datasources"
	"github.com/azambranapr/rmocs_collector_api/src/domain/usecases/batch"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func getVersion(c *gin.Context) {
	var v version.Version
	v.Title = "Oil collector System 2023"
	v.Version = "v0.0.1"
	v.Build = "2023.4.24"

	v.EntryPoint = make(version.Entrys, 0)
	m := version.Entry{}

	var us []version.Uris
	us = append(us, version.Uris{Uri: "/v1/total"})
	us = append(us, version.Uris{Uri: "/v1/get/:Id"})
	us = append(us, version.Uris{Uri: "/v1/get/all"})
	m.Get = us

	var ps []version.Uris
	ps = append(ps, version.Uris{Uri: "/v1/insert/"})
	ps = append(ps, version.Uris{Uri: "/v1/update/"})
	m.Post = ps

	var dl []version.Uris
	dl = append(dl, version.Uris{Uri: "/v1/delete/:Id"})
	dl = append(dl, version.Uris{Uri: "/v1/delete/all"})
	m.Delete = dl

	v.EntryPoint = append(v.EntryPoint, m)

	c.IndentedJSON(http.StatusOK, v)
}

func NewRouter() *gin.Engine {

	datasources.New()
	myStorage, err := datasources.DAO()
	if err != nil {
		log.Fatalf("DAO: %v", err)
	}

	route := gin.Default()

	route.GET("/v1/version", getVersion)

	// Data CRUD
	grData := route.Group("/v1/")
	service := batch.New(myStorage)
	grData.GET("/", getVersion)
	grData.GET("/batch/all", service.GetAll)
	grData.GET("/batch/:id", service.GetById)

	return route
}
