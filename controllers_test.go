package main

import (
	"encoding/json"
	"net/http"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateEntityController(t *testing.T) {
	cfg.Setup("test")
	db.Setup(cfg)
	defer db.Close()
	rtr.Setup()

	entity := createTestEntityStruct(1)

	response := entity.GenerateTestRequest("POST", "")
	json.Unmarshal(response.Body.Bytes(), &entity)
	assert.Equal(t, http.StatusOK, response.Code)
	assert.Equal(t, "Name 1", entity.Name)
}

func TestSearchEntityController(t *testing.T) {
	cfg.Setup("test")
	db.Setup(cfg)
	defer db.Close()
	rtr.Setup()

	entity, _ := createTestEntityStruct(1).Create()
	createTestEntityStruct(2).Create()

	params := &EntitiesSearchParameters{
		FilterKind: "1",
		FilterName: "",

		SortField:     "id",
		SortDirection: "DESC",

		Limit:  100,
		Offset: 0,
	}

	entities := make([]*Entity, 0)
	response := entity.GenerateTestRequest("GET", params.generateSearchURLString())
	json.Unmarshal(response.Body.Bytes(), &entities)
	assert.Equal(t, http.StatusOK, response.Code)
	assert.Equal(t, 1, len(entities))
	assert.Equal(t, "Name 1", entities[0].Name)
}

func TestGetEntityController(t *testing.T) {
	cfg.Setup("test")
	db.Setup(cfg)
	defer db.Close()
	rtr.Setup()

	entity, _ := createTestEntityStruct(1).Create()

	response := entity.GenerateTestRequest("GET", "/"+strconv.Itoa(entity.ID))
	json.Unmarshal(response.Body.Bytes(), &entity)
	assert.Equal(t, http.StatusOK, response.Code)
	assert.Equal(t, "Name 1", entity.Name)
}

func TestUpdateEntityController(t *testing.T) {
	cfg.Setup("test")
	db.Setup(cfg)
	defer db.Close()
	rtr.Setup()

	entity, _ := createTestEntityStruct(1).Create()

	entity.Name = "Name 2"

	response := entity.GenerateTestRequest("PUT", "/"+strconv.Itoa(entity.ID))
	json.Unmarshal(response.Body.Bytes(), &entity)
	assert.Equal(t, http.StatusOK, response.Code)
	assert.Equal(t, "Name 2", entity.Name)
}
