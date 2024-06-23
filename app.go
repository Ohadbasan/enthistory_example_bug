package main

import (
	"fmt"
	"log"

	"entgo.io/ent"
	"entgo.io/ent/entc/gen"

	"github.com/flume/enthistory"

	lala "entdemo/ent/schema"

	"entgo.io/ent/entc"
)

func main() {
	if err := enthistory.Generate("./ent/schema", []ent.Interface{
		// Add all the schemas you want history tracking on here
		lala.TestEntity{},
	},
		enthistory.WithUpdatedBy("userId", enthistory.ValueTypeInt),
		enthistory.WithHistoryTimeIndex(),
		enthistory.WithImmutableFields(),
		// Without this line, all triggers will be used as the default
		enthistory.WithTriggers(enthistory.OpTypeInsert),
	); err != nil {
		log.Fatal(fmt.Sprintf("running enthistory codegen: %v", err))
	}

	if err := entc.Generate("./ent/schema",
		&gen.Config{
			Features: []gen.Feature{gen.FeatureSnapshot},
		},
		entc.Extensions(
			enthistory.NewHistoryExtension(enthistory.WithAuditing()),
		),
	); err != nil {
		log.Fatal("running ent codegen:", err)
	}
}
