package main

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"time"

	"github.com/gilperopiola/frutils"
	"github.com/gilperopiola/lyfe-companyon-backend/utils"
)

type Entity struct {
	ID          int
	Name        string
	Description string
	Kind        string
	Importance  int
	Status      int //1 = enabled, 2 = disabled
	DateCreated time.Time
}

//EntityActions are external actions that the controllers can call
type EntityActions interface {
	Create() (*Entity, error)
	Search() ([]*Entity, error)
	Get() (*Entity, error)
	Update() (*Entity, error)
	Delete() (*Entity, error)
}

//EntityTestingActions are functions that aid in testing
type EntityTestingActions interface {
	GenerateTestRequest(token, method, url string) *httptest.ResponseRecorder
	GenerateJSONBody() string
	generateSearchURLString() string
}

type EntitiesSearchParameters struct {
	FilterKind string
	FilterName string

	SortField     string
	SortDirection string

	Limit  int
	Offset int
}

//Create inserts a new entity on the database
func (entity *Entity) Create() (*Entity, error) {
	result, err := db.DB.Exec(`INSERT INTO entities (name, description, kind, importance) VALUES (?, ?, ?, ?)`, entity.Name, entity.Description, entity.Kind, entity.Importance)
	if err != nil {
		return &Entity{}, err
	}

	entity.ID = frutils.GetID(result)

	return entity.Get()
}

//Search retrieves all the entities that match certain parameters
func (entity *Entity) Search(params *EntitiesSearchParameters) ([]*Entity, error) {
	orderByString := "id ASC"
	if params.SortField != "" && params.SortDirection != "" {
		orderByString = params.SortField + " " + params.SortDirection
	}

	query := fmt.Sprintf(`SELECT id, name, description, kind, importance, status, dateCreated 
						  FROM entities WHERE kind LIKE ? AND name LIKE ? AND status = 1
						  ORDER BY %s LIMIT ? OFFSET ?`, orderByString)

	params.FilterKind = "%" + params.FilterKind + "%"
	params.FilterName = "%" + params.FilterName + "%"

	rows, err := db.DB.Query(query, params.FilterKind, params.FilterName, params.Limit, params.Offset)

	defer rows.Close()
	if err != nil {
		return []*Entity{}, err
	}

	entities := []*Entity{}
	for rows.Next() {
		tempEntity := &Entity{}
		err = rows.Scan(&tempEntity.ID, &tempEntity.Name, &tempEntity.Description, &tempEntity.Kind,
			&tempEntity.Importance, &tempEntity.Status, &tempEntity.DateCreated)
		if err != nil {
			return []*Entity{}, err
		}

		entities = append(entities, tempEntity)
	}

	return entities, nil
}

//Get retrieves a single entity
func (entity *Entity) Get() (*Entity, error) {
	err := db.DB.QueryRow(`SELECT name, description, kind, importance, status, dateCreated FROM entities WHERE id = ?`, entity.ID).
		Scan(&entity.Name, &entity.Description, &entity.Kind, &entity.Importance, &entity.Status, &entity.DateCreated)
	if err != nil {
		return &Entity{}, err
	}

	return entity, nil
}

//Update modifies an entity
func (entity *Entity) Update() (*Entity, error) {
	_, err := db.DB.Exec(`UPDATE entities SET name = ?, description = ?, kind = ?, importance = ?, status = ? WHERE id = ?`,
		entity.Name, entity.Description, entity.Kind, entity.Importance, entity.Status, entity.ID)
	if err != nil {
		return &Entity{}, err
	}

	return entity.Get()
}

//UpdateImportance modifies an entity's importance
func (entity *Entity) UpdateImportance() (*Entity, error) {
	_, err := db.DB.Exec(`UPDATE entities SET importance = ? WHERE id = ?`, entity.Importance, entity.ID)
	if err != nil {
		return &Entity{}, err
	}

	return entity.Get()
}

//Disable disables an entity
func (entity *Entity) Disable() (*Entity, error) {
	_, err := db.DB.Exec(`UPDATE entities SET status = ? WHERE id = ?`, 2, entity.ID)
	if err != nil {
		return &Entity{}, err
	}

	return entity.Get()
}

//GenerateTestRequest sends an HTTP request to the given URL with the given method
func (entity *Entity) GenerateTestRequest(method, url string) *httptest.ResponseRecorder {
	w := httptest.NewRecorder()
	body := entity.GetJSONBody()
	req, _ := http.NewRequest(method, "/Entity"+url, bytes.NewReader([]byte(body)))
	rtr.ServeHTTP(w, req)
	return w
}

//GetJSONBody returns an entity as the JSON necessary to send an HTTP request
func (entity *Entity) GetJSONBody() string {
	body := `{
		"name": "` + entity.Name + `",
		"description": "` + entity.Description + `",
		"kind": "` + entity.Kind + `",
		"importance": ` + utils.ToString(entity.Importance) + `,
		"status": ` + utils.ToString(entity.Status) + `
	}`
	return body
}

//generateSearchURLString generates the last part of the URL for the search endpoint
func (params *EntitiesSearchParameters) generateSearchURLString() string {
	return fmt.Sprintf("?kind=%s&name=%s&sortField=%s&sortDirection=%s&limit=%d&offset=%d",
		params.FilterKind, params.FilterName, params.SortField, params.SortDirection, params.Limit, params.Offset)
}
