package common

const (
	FaultsAllowed        int     = 1
	LEADER               string  = "leader"
	FOLLOWER             string  = "follower"
	MaxReconnectAttempts float64 = 3
)

var ServerPortMap = map[int]int{
	1: 8001,
	2: 8002,
	3: 8003,
}
