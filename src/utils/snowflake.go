package utils

import "github.com/bwmarrin/snowflake"

var node *snowflake.Node

func InitializeSnowFlake() error {
	// Skip if node is already initialized
	if node != nil {
		return nil
	}

	var err error
	node, err = snowflake.NewNode(1)
	if err != nil {
		return err
	}
	return nil
}

func GenerateSnowflakeID() int64 {
	return node.Generate().Int64()
}
