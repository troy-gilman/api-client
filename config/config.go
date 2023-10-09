package config

import (
    "os"
    "io/ioutil"
    "encoding/json"
)

func LoadConfig(path string) (Config, error) {
    var cnfg Config

    jsonFile, err := os.Open(path)
    if err != nil {
        return cnfg, err
    }
    defer jsonFile.Close()

    bytes, err := ioutil.ReadAll(jsonFile)
    if err != nil {
        return cnfg, err
    }

    json.Unmarshal(bytes, &cnfg)

    return cnfg, nil
}
