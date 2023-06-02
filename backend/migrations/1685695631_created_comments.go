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
			"id": "jqidajnzzsjl7my",
			"created": "2023-06-02 08:47:11.019Z",
			"updated": "2023-06-02 08:47:11.019Z",
			"name": "comments",
			"type": "base",
			"system": false,
			"schema": [
				{
					"system": false,
					"id": "tdh57ian",
					"name": "content",
					"type": "editor",
					"required": false,
					"unique": false,
					"options": {}
				},
				{
					"system": false,
					"id": "avvatlec",
					"name": "commenter",
					"type": "relation",
					"required": false,
					"unique": false,
					"options": {
						"collectionId": "_pb_users_auth_",
						"cascadeDelete": false,
						"minSelect": null,
						"maxSelect": 1,
						"displayFields": []
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

		collection, err := dao.FindCollectionByNameOrId("jqidajnzzsjl7my")
		if err != nil {
			return err
		}

		return dao.DeleteCollection(collection)
	})
}
