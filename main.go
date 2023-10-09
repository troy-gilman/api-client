package main

import (
    "os"
    "fmt"
    "api-client/client"
    "api-client/config"
)

func getConfigFilePath() string {
    args := os.Args
    if len(args) < 2 {
        fmt.Println("Must specify a configuration file path")
        os.Exit(1)
    }
    return args[1]
}

func printResult(configuration config.Config, result client.Result) {
    numClients := result.NumClients
    for requestIdx, _ := range configuration.Requests {
        numSuccess := result.NumSuccess[requestIdx]
        avgTimeElapsedMs := result.AvgTimeElapsedMs[requestIdx]
        fmt.Printf("Request #%d - successful (%d/%d) - avg time elapsed (%d ms)\n", requestIdx, numSuccess, numClients, avgTimeElapsedMs)   
    }
}

func main() {
    configPath := getConfigFilePath()
    configuration, err := config.LoadConfig(configPath)
    if err != nil {
        fmt.Println(err.Error())
        os.Exit(1)
    }

    fmt.Printf("Using configuration file: %s\n", configPath)
    fmt.Printf("Launching %d clients\n", configuration.NumClients)

    //for _, r := range configuration.Requests {
    //    fmt.Println(r)
    //}

    result := client.LaunchClients(configuration)
    printResult(configuration, result)
}




