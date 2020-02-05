package main

import (
	"net/http"

	"github.com/gilperopiola/frutils"
	"github.com/gin-gonic/gin"
)

func CreateEntity(c *gin.Context) {
	var entity *Entity
	c.BindJSON(&entity)

	if entity.Name == "" || entity.Kind == "" {
		c.JSON(http.StatusBadRequest, "name and kind required")
		return
	}

	entity, err := entity.Create()
	if err != nil {
		c.JSON(http.StatusBadRequest, db.BeautifyError(err))
		return
	}

	c.JSON(http.StatusOK, entity)
}

func SearchEntities(c *gin.Context) {
	entity := &Entity{}

	params := &EntitiesSearchParameters{
		FilterKind: c.Query("kind"),
		FilterName: c.Query("name"),

		SortField:     c.Query("sortField"),
		SortDirection: c.Query("sortDirection"),

		Limit:  frutils.ToInt(c.Query("limit")),
		Offset: frutils.ToInt(c.Query("offset")),
	}

	if params.Limit == 0 {
		params.Limit = 99999
	}

	entities, err := entity.Search(params)
	if err != nil {
		c.JSON(http.StatusBadRequest, db.BeautifyError(err))
		return
	}

	c.JSON(http.StatusOK, entities)
}

func GetEntity(c *gin.Context) {
	entity := &Entity{ID: frutils.ToInt(c.Param("id_entity"))}

	entity, err := entity.Get()
	if err != nil {
		c.JSON(http.StatusBadRequest, db.BeautifyError(err))
		return
	}

	c.JSON(http.StatusOK, entity)
}

func UpdateEntity(c *gin.Context) {
	var entity *Entity
	c.BindJSON(&entity)

	if entity.Name == "" || entity.Kind == "" {
		c.JSON(http.StatusBadRequest, "name and kind required")
		return
	}

	entity.ID = frutils.ToInt(c.Param("id_entity"))

	entity, err := entity.Update()
	if err != nil {
		c.JSON(http.StatusBadRequest, db.BeautifyError(err))
		return
	}

	c.JSON(http.StatusOK, entity)
}

func UpdateEntityImportance(c *gin.Context) {
	entity := &Entity{ID: frutils.ToInt(c.Param("id_entity")), Importance: frutils.ToInt(c.Param("importance"))}

	entity, err := entity.UpdateImportance()
	if err != nil {
		c.JSON(http.StatusBadRequest, db.BeautifyError(err))
		return
	}

	c.JSON(http.StatusOK, entity)
}

func DisableEntity(c *gin.Context) {
	entity := &Entity{ID: frutils.ToInt(c.Param("id_entity"))}

	entity, err := entity.Disable()
	if err != nil {
		c.JSON(http.StatusBadRequest, db.BeautifyError(err))
		return
	}

	c.JSON(http.StatusOK, entity)
}
