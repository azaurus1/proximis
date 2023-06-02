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
			"id": "tvoi3kqxen48eaq",
			"created": "2023-06-02 08:42:17.767Z",
			"updated": "2023-06-02 08:42:17.767Z",
			"name": "category",
			"type": "base",
			"system": false,
			"schema": [
				{
					"system": false,
					"id": "ocki1q8z",
					"name": "title",
					"type": "text",
					"required": false,
					"unique": false,
					"options": {
						"min": null,
						"max": null,
						"pattern": ""
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

		collection, err := dao.FindCollectionByNameOrId("tvoi3kqxen48eaq")
		if err != nil {
			return err
		}

		return dao.DeleteCollection(collection)
	})
}
