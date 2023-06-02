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
		edit_status := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "ihcswsdz",
			"name": "status",
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
		}`), edit_status)
		collection.Schema.AddField(edit_status)

		return dao.SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("nsrcim3c245u5ci")
		if err != nil {
			return err
		}

		// update
		edit_status := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
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
		}`), edit_status)
		collection.Schema.AddField(edit_status)

		return dao.SaveCollection(collection)
	})
}
