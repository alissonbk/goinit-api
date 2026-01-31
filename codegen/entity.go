package codegen

import "fmt"

func GenerateExampleEntity() string {
	quotationMark := "`"

	return fmt.Sprintf(`
		package entity

		// An example entity
		type Example struct {
			BaseEntity
			Name string %[1]sdb:"name" json:"name"%[1]s
		}

	`, quotationMark)
}

func GenerateBaseEntity() string {
	quotationMark := "`"

	return fmt.Sprintf(`
		package entity

		import (
			"time"
		)

		type BaseEntity struct {
			ID        uint      %[1]sdb:"id"%[1]s
			CreatedAt time.Time %[1]sdb:"created_at"%[1]s
			UpdatedAt time.Time %[1]sdb:"updated_at"%[1]s
			DeletedAt time.Time %[1]sdb:"deleted_at"%[1]s
		}

	`, quotationMark)
}
