package migrations

import (
	"encoding/json"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/daos"
	m "github.com/pocketbase/pocketbase/migrations"
	"github.com/pocketbase/pocketbase/models"
)

func init() {
	m.Register(func(db dbx.Builder) error {
		jsonData := `{
			"id": "nsrcim3c245u5ci",
			"created": "2023-06-02 08:30:18.358Z",
			"updated": "2023-06-02 08:30:18.358Z",
			"name": "incidents",
			"type": "base",
			"system": false,
			"schema": [
				{
					"system": false,
					"id": "ihcswsdz",
					"name": "field",
					"type": "select",
					"required": false,
					"unique": false,
					"options": {
						"maxSelect": 1,
						"values": [
							"new",
							"in_progress",
							"resolved"
						]
					}
				}
			],
			"indexes": [],
			"listRule": null,
			"viewRule": null,
			"createRule": null,
			"updateRule": null,
			"deleteRule": null,
			"options": {}
		}`

		collection := &models.Collection{}
		if err := json.Unmarshal([]byte(jsonData), &collection); err != nil {
			return err
		}

		return daos.New(db).SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("nsrcim3c245u5ci")
		if err != nil {
			return err
		}

		return dao.DeleteCollection(collection)
	})
}
