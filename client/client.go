package client

import (
	"api-client/config"
	"io/ioutil"
	"net/http"
	"time"
)

func httpGet(url string) error {
    resp, err := http.Get(url)
    if err != nil {
        return err
    }

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return err
    }

    _ = string(body)
    //log.Printf(sb)
    return nil
}

func handleRequest(request config.Request) error {
    var err error
    switch request.Type {
    case "GET":
        err = httpGet(request.URI)
    }
    return err
}

func startClient(resultChan chan ClientResult, configuration config.Config) {
    numRequests := len(configuration.Requests)

    var result ClientResult
    result.success = make([]bool, numRequests)
    result.timeElapsedMs = make([]int64, numRequests)

    for i, request := range configuration.Requests {
        startTime := time.Now()
        err := handleRequest(request)
        result.success[i] = err == nil
        result.timeElapsedMs[i] = time.Since(startTime).Milliseconds()
    }
    
    resultChan <- result
}

func processResults(clientResults []ClientResult, numRequests int) Result {
    var result Result
    result.NumClients = uint(len(clientResults))
    result.NumSuccess = make([]uint, numRequests)
    result.AvgTimeElapsedMs = make([]int64, numRequests)

    for _, cResult := range clientResults {
        for requestIdx, success := range cResult.success { 
            if success {
                result.NumSuccess[requestIdx]++
                result.AvgTimeElapsedMs[requestIdx] += cResult.timeElapsedMs[requestIdx]
            }
        }
    }

    for requestIdx, numSuccess := range result.NumSuccess {
        if numSuccess == 0 {
            result.AvgTimeElapsedMs[requestIdx] = 0
        } else {
            result.AvgTimeElapsedMs[requestIdx] /= int64(numSuccess)
        }
    }

    return result
}

func LaunchClients(configuration config.Config) Result {
    channels := make([]chan ClientResult, configuration.NumClients)
    for i := range channels {
        channel := make(chan ClientResult)
        channels[i] = channel
        go startClient(channel, configuration)
    }
    
    results := make([]ClientResult, len(channels))
    for i, channel := range channels {
        select {
        case result := <-channel:
            results[i] = result
            close(channel)
        }
    }

    return processResults(results, len(configuration.Requests))
}


