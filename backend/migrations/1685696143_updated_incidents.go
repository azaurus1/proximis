package migrations

import (
	"encoding/json"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/daos"
	m "github.com/pocketbase/pocketbase/migrations"
	"github.com/pocketbase/pocketbase/models/schema"
)

func init() {
	m.Register(func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("nsrcim3c245u5ci")
		if err != nil {
			return err
		}

		// update
		edit_tags := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "wtncmspu",
			"name": "tags",
			"type": "relation",
			"required": false,
			"unique": false,
			"options": {
				"collectionId": "c5fqjrfhlr4m7tx",
				"cascadeDelete": false,
				"minSelect": null,
				"maxSelect": null,
				"displayFields": []
			}
		}`), edit_tags)
		collection.Schema.AddField(edit_tags)

		return dao.SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("nsrcim3c245u5ci")
		if err != nil {
			return err
		}

		// update
		edit_tags := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "wtncmspu",
			"name": "tags",
			"type": "relation",
			"required": false,
			"unique": false,
			"options": {
				"collectionId": "c5fqjrfhlr4m7tx",
				"cascadeDelete": false,
				"minSelect": null,
				"maxSelect": 1,
				"displayFields": []
			}
		}`), edit_tags)
		collection.Schema.AddField(edit_tags)

		return dao.SaveCollection(collection)
	})
}
