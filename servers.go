package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/user"
	"sort"
)

type ServerPlotPoint struct {
	Online int `json:"online"`
	Time   int `json:"time"`
}

type Server struct {
	Number     int    `json:"number"`
	Name       string `json:"name"`
	IP         string `json:"ip"`
	Port       int    `json:"port"`
	Online     int    `json:"online"`
	MaxPlayers int    `json:"maxplayers"`
	// Password             bool              `json:"password"`
	// VK                   string            `json:"vk"`
	// TG                   string            `json:"tg"`
	// Inst                 string            `json:"inst"`
	Icon string `json:"icon"`
	// AdditionalIps        []string          `json:"additionalIps"`
	// DonateMultiplier     int               `json:"donateMultiplier"`
	// ExperienceMultiplier int               `json:"experienceMultiplier"`
	// PlotPoints           []ServerPlotPoint `json:"plotPoints"`
}

type ArizonaServerInfo struct {
	Arizona       []Server `json:"arizona"`
	ArizonaMobile []Server `json:"arizonaMobile"`
	// ArizonaStaging []Server `json:"arizona_staging"`
	// Rodina         []Server `json:"rodina"`
	// Village        []Server `json:"village"`
	// Arizonav       []Server `json:"arizonav"`
}

const (
	ArizonaServersListURL string = "https://api.arizona-five.com/launcher/servers"
)

func LoadServers() (ArizonaServerInfo, error) {
	var servers ArizonaServerInfo
	var err error
	var jsonBytes []byte
	var listFileReaded = false
	winUser, err := user.Current()
	if err == nil {
		listPath := winUser.HomeDir + "\\Documents\\alt-launcher-servers.json"
		fmt.Println("Custom servers list path=", listPath)
		if _, err := os.Stat(listPath); err == nil {
			jsonBytes, err = os.ReadFile(listPath)
			fmt.Println("List file:", string(jsonBytes), err)
			if err == nil {
				listFileReaded = true
				fmt.Println("Reading servers list from file")
			}
		}
	}
	if !listFileReaded {
		response, err := http.Get(ArizonaServersListURL)
		if err != nil {
			return servers, err
		}
		jsonBytes, err = io.ReadAll(response.Body)
		if err != nil {
			return servers, err
		}
	}
	// fmt.Println(string(jsonBytes))
	if err == nil {
		err = json.Unmarshal(jsonBytes, &servers)
		if err == nil {
			sort.Slice(servers.Arizona, func(a, b int) bool {
				return servers.Arizona[a].Number < servers.Arizona[b].Number
			})
		}
	}

	return servers, err
}
