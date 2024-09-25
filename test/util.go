package test

import (
	"github.com/bwmarrin/snowflake"
)

func GeneratID() snowflake.ID {
	node, _ := snowflake.NewNode(1)
	return node.Generate()
}

func GeneratIDString() string {
	return GeneratID().String()
}
