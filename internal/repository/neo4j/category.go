package neo4j

import (
	"context"
	"log"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

func IncreaseCategoryFreq(cat string) error {
	//searches category . If found increases its frequency
	ctx := context.Background()
	sesion := neo4jDB.NewSession(ctx, neo4j.SessionConfig{
		AccessMode: neo4j.AccessModeWrite,
	})
	smt := "MATCH (c:Category {name : $category}) SET c.frequency = c.frequency + 1"
	_, err := sesion.Run(ctx, smt, map[string]interface{}{
		"category": cat,
	})
	if err != nil {
		log.Println(err.Error())
		return err
	}
	return nil
}

func CreateNewCategory(cat string) error {
	//creates a new category in the db

	ctx := context.Background()
	sesion := neo4jDB.NewSession(ctx, neo4j.SessionConfig{
		AccessMode: neo4j.AccessModeWrite,
	})
	smt := "MERGE (c:Category {name : $category}) ON MATCH SET c.frequency = c.frequency + 1 ON CREATE SET c.frequency = $frequency"
	_, err := sesion.Run(ctx, smt, map[string]interface{}{
		"category":  cat,
		"frequency": 1,
	})
	if err != nil {
		log.Fatalf(err.Error())
		return err
	}
	return nil
}
