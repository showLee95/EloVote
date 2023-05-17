package algorithm

import (
	"fmt"

	"github.com/bwmarrin/snowflake"
)

func GenID() int64 {
	node, err := snowflake.NewNode(1) // Replace 1 with your machine ID
	if err != nil {
		fmt.Println("Failed to create Snowflake node:", err)

	}
	id := node.Generate()
	return int64(id)
}
