package entity

import (
	"encoding/json"
	"fmt"

	"github.com/mitchellh/mapstructure"
	"github.com/lendlord/lendlord-server/tools"
)

//type BaseModelFields = struct {
//	ID        string
//	CreatedAt string
//	UpdatedAt string
//	DeletedAt string
//}
//
//var baseModelFields = BaseModelFields{
//	"id",
//	"created_at",
//	"updated_at",
//	"deleted_at",
//}

func InitModelFields(tableName string, modelFiled interface{}) interface{} {
	// marshal struct
	b, err := json.Marshal(&modelFiled)
	if err != nil {
		panic(err)
	}

	// translate to map
	var m map[string]string
	err = json.Unmarshal(b, &m)
	if err != nil {
		panic(err)
	}

	// splice table name column name
	for k, _ := range m {
		m[k] = fmt.Sprintf("%s.%s", tableName, tools.CamelCaseToUnderscore(k))
	}

	// translate map to struct
	if err := mapstructure.Decode(m, &modelFiled); err != nil {
		panic(err)
	}

	return modelFiled
}
