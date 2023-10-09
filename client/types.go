package client

type Result struct {
    NumClients uint
    NumSuccess []uint
    AvgTimeElapsedMs []int64
}

type ClientResult struct {
    success []bool
    timeElapsedMs []int64
}
