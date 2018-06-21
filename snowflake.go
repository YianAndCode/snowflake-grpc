package main

import "time"

// EPOCH 时间偏移量
const EPOCH int64 = 1523980800000

// MAX41BIT 41 位的最大值
const MAX41BIT int64 = 2199023255551

// MAX5BIT 5 位的最大值
const MAX5BIT int64 = 31

// MAX12BIT 12 位的最大值
const MAX12BIT int64 = 4095

// seq 自增序列
var seq int64

// nodeID 节点号
var nodeID int64

// next 获取下一个唯一 ID
func next(serviceID int64) int64 {
	// 取得当前时间戳（毫秒），并减去偏移量
	var timestamp int64
	timestamp = time.Now().UnixNano()/int64(time.Millisecond) - EPOCH
	timestamp <<= 22
	// 节点号
	nodeID &= MAX5BIT
	nodeID <<= 17
	// 服务编号
	serviceID &= MAX5BIT
	serviceID <<= 12
	// 自增序列
	seq &= MAX12BIT
	seq++

	return timestamp | nodeID | serviceID | seq
}

// parse 从 next() 生成的 id 中解析出时间戳、节点号、服务编号和自增序列
func parse(id int64) (t time.Time, timestamp int64, nodeID int64, serviceID int64, seq int64) {
	/**
	 * 通过时间戳掩码从 id 取得时间戳部分，并右移操作后加上时间偏移量
	 */
	var tsMask int64
	tsMask = MAX41BIT << 22
	timestamp = ((id&tsMask)>>22 + EPOCH) / 1000
	t = time.Unix(timestamp/1000, (timestamp%1000)*100000)

	// 通过节点号掩码取得节点号
	var nodeIDMask int64
	nodeIDMask = MAX5BIT << 17
	nodeID = (id & nodeIDMask) >> 17

	// 通过服务编号掩码取得服务编号
	var serviceIDMask int64
	serviceIDMask = MAX5BIT << 12
	serviceID = (id & serviceIDMask) >> 12

	// 取得自增序列
	seqMask := MAX12BIT
	seq = id & seqMask
	return
}
