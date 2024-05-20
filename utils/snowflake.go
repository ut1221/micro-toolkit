package utils

import (
	sf "github.com/bwmarrin/snowflake"
	"strconv"
	"time"
)

var node *sf.Node

// Init
//
//	@Description: 初始化雪花算法节点
//	@Author PTJ 2024-05-14 17:57:10
//	@param startTime
//	@param machineID
//	@return err
func Init(startTime string, machineID int64) (err error) {
	var st time.Time
	st, err = time.Parse("2006-01-02", startTime)
	if err != nil {
		return
	}
	sf.Epoch = st.UnixNano() / 1000000
	node, err = sf.NewNode(machineID)
	return
}

// GetID
//
//	@Description: 生成ID
//	@Author PTJ 2024-05-14 17:57:44
//	@return string
func GetID() string {
	return strconv.FormatInt(node.Generate().Int64(), 10)
}
