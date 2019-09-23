package main

import (
	"testing"

	"github.com/gilperopiola/frutils"
	"github.com/stretchr/testify/assert"
)

func createTestEntityStruct(identifier int) *Entity {
	stringIdentifier := frutils.ToString(identifier)

	return &Entity{
		Name:        "Name " + stringIdentifier,
		Description: "Description " + stringIdentifier,
		Kind:        "Kind " + stringIdentifier,
		Importance:  identifier,
	}
}

func TestCreateEntity(t *testing.T) {
	cfg.Setup("test")
	db.Setup(cfg)
	defer db.Close()

	entity, err := createTestEntityStruct(1).Create()
	assert.NoError(t, err)
	assert.NotZero(t, entity.ID)
	assert.Equal(t, "Name 1", entity.Name)
	assert.Equal(t, "Description 1", entity.Description)
	assert.Equal(t, "Kind 1", entity.Kind)
	assert.Equal(t, 1, entity.Importance)
	assert.Equal(t, 1, entity.Status)
	assert.NotZero(t, entity.DateCreated)
}

func TestGetEntity(t *testing.T) {
	cfg.Setup("test")
	db.Setup(cfg)
	defer db.Close()

	entity, _ := createTestEntityStruct(1).Create()

	entity, err := entity.Get()
	assert.NoError(t, err)
	assert.NotZero(t, entity.ID)
	assert.Equal(t, "Name 1", entity.Name)
	assert.Equal(t, "Description 1", entity.Description)
	assert.Equal(t, "Kind 1", entity.Kind)
	assert.Equal(t, 1, entity.Importance)
	assert.Equal(t, 1, entity.Status)
	assert.NotZero(t, entity.DateCreated)
}

func TestSearchEntities(t *testing.T) {
	cfg.Setup("test")
	db.Setup(cfg)
	defer db.Close()

	entity, _ := createTestEntityStruct(1).Create()
	entity2, _ := createTestEntityStruct(1).Create()
	entity3, _ := createTestEntityStruct(1).Create()
	createTestEntityStruct(2).Create()

	params := &EntitiesSearchParameters{
		FilterKind: "1",
		FilterName: "",

		SortField:     "id",
		SortDirection: "DESC",

		Limit:  3,
		Offset: 0,
	}

	entities, err := entity.Search(params)
	assert.NoError(t, err)
	assert.Equal(t, params.Limit, len(entities))
	assert.NotZero(t, entities[0].ID)
	assert.Equal(t, entity3.Name, entities[0].Name)
	assert.Equal(t, entity2.Name, entities[1].Name)
	assert.Equal(t, entity.Name, entities[2].Name)
}

func TestUpdateEntity(t *testing.T) {
	cfg.Setup("test")
	db.Setup(cfg)
	defer db.Close()

	entity, _ := createTestEntityStruct(1).Create()

	entity.Name = "Name 2"
	entity.Description = "Description 2"
	entity.Importance = 2
	entity.Status = 2

	entity, err := entity.Update()
	assert.NoError(t, err)
	assert.NotZero(t, entity.ID)
	assert.Equal(t, "Name 2", entity.Name)
	assert.Equal(t, "Description 2", entity.Description)
	assert.Equal(t, 2, entity.Importance)
	assert.Equal(t, 2, entity.Status)
}
