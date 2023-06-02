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

		// add
		new_title := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "oge9prve",
			"name": "title",
			"type": "text",
			"required": false,
			"unique": false,
			"options": {
				"min": null,
				"max": null,
				"pattern": ""
			}
		}`), new_title)
		collection.Schema.AddField(new_title)

		// add
		new_description := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "el5yw6xo",
			"name": "description",
			"type": "editor",
			"required": false,
			"unique": false,
			"options": {}
		}`), new_description)
		collection.Schema.AddField(new_description)

		// add
		new_priority := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "bdpqzx7r",
			"name": "priority",
			"type": "select",
			"required": false,
			"unique": false,
			"options": {
				"maxSelect": 1,
				"values": [
					"low",
					"medium",
					"high",
					"critical"
				]
			}
		}`), new_priority)
		collection.Schema.AddField(new_priority)

		// add
		new_categorisation := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "usr86vwb",
			"name": "categorisation",
			"type": "relation",
			"required": false,
			"unique": false,
			"options": {
				"collectionId": "tvoi3kqxen48eaq",
				"cascadeDelete": false,
				"minSelect": null,
				"maxSelect": 1,
				"displayFields": []
			}
		}`), new_categorisation)
		collection.Schema.AddField(new_categorisation)

		// add
		new_reporter := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "m1bl4san",
			"name": "reporter",
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
		}`), new_reporter)
		collection.Schema.AddField(new_reporter)

		// add
		new_assignee := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "wbgrwb1i",
			"name": "assignee",
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
		}`), new_assignee)
		collection.Schema.AddField(new_assignee)

		// add
		new_tags := &schema.SchemaField{}
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
		}`), new_tags)
		collection.Schema.AddField(new_tags)

		// add
		new_related_incidents := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "ppj1ncvl",
			"name": "related_incidents",
			"type": "relation",
			"required": false,
			"unique": false,
			"options": {
				"collectionId": "nsrcim3c245u5ci",
				"cascadeDelete": false,
				"minSelect": null,
				"maxSelect": 1,
				"displayFields": []
			}
		}`), new_related_incidents)
		collection.Schema.AddField(new_related_incidents)

		// add
		new_severity := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "zybsmk40",
			"name": "severity",
			"type": "select",
			"required": false,
			"unique": false,
			"options": {
				"maxSelect": 1,
				"values": [
					"low",
					"medium",
					"high",
					"critical"
				]
			}
		}`), new_severity)
		collection.Schema.AddField(new_severity)

		// add
		new_resolution := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "zsgp03a6",
			"name": "resolution",
			"type": "editor",
			"required": false,
			"unique": false,
			"options": {}
		}`), new_resolution)
		collection.Schema.AddField(new_resolution)

		// add
		new_comments := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "yqljzwzm",
			"name": "comments",
			"type": "relation",
			"required": false,
			"unique": false,
			"options": {
				"collectionId": "jqidajnzzsjl7my",
				"cascadeDelete": false,
				"minSelect": null,
				"maxSelect": 1,
				"displayFields": []
			}
		}`), new_comments)
		collection.Schema.AddField(new_comments)

		// add
		new_attachments := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "zmetm0tr",
			"name": "attachments",
			"type": "file",
			"required": false,
			"unique": false,
			"options": {
				"maxSelect": 99,
				"maxSize": 5242880,
				"mimeTypes": [],
				"thumbs": [],
				"protected": false
			}
		}`), new_attachments)
		collection.Schema.AddField(new_attachments)

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
					"assigned",
					"in_progress",
					"on_hold",
					"escalated",
					"pending",
					"resolved",
					"closed",
					"reopened",
					"reassigned",
					"duplicate",
					"invalid"
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

		// remove
		collection.Schema.RemoveField("oge9prve")

		// remove
		collection.Schema.RemoveField("el5yw6xo")

		// remove
		collection.Schema.RemoveField("bdpqzx7r")

		// remove
		collection.Schema.RemoveField("usr86vwb")

		// remove
		collection.Schema.RemoveField("m1bl4san")

		// remove
		collection.Schema.RemoveField("wbgrwb1i")

		// remove
		collection.Schema.RemoveField("wtncmspu")

		// remove
		collection.Schema.RemoveField("ppj1ncvl")

		// remove
		collection.Schema.RemoveField("zybsmk40")

		// remove
		collection.Schema.RemoveField("zsgp03a6")

		// remove
		collection.Schema.RemoveField("yqljzwzm")

		// remove
		collection.Schema.RemoveField("zmetm0tr")

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
	})
}
