package main

import (
	"encoding/json"
	"io"
	"net/http"
	"sort"
)

type ServerPlotPoint struct {
	Online int `json:"online"`
	Time   int `json:"time"`
}

type Server struct {
	Number               int               `json:"number"`
	Name                 string            `json:"name"`
	IP                   string            `json:"ip"`
	Port                 int               `json:"port"`
	Online               int               `json:"online"`
	MaxPlayers           int               `json:"maxplayers"`
	Password             bool              `json:"password"`
	VK                   string            `json:"vk"`
	TG                   string            `json:"tg"`
	Inst                 string            `json:"inst"`
	Icon                 string            `json:"icon"`
	AdditionalIps        []string          `json:"additionalIps"`
	DonateMultiplier     int               `json:"donateMultiplier"`
	ExperienceMultiplier int               `json:"experienceMultiplier"`
	PlotPoints           []ServerPlotPoint `json:"plotPoints"`
}

type ArizonaServerInfo struct {
	Arizona        []Server `json:"arizona"`
	ArizonaMobile  []Server `json:"arizonaMobile"`
	ArizonaStaging []Server `json:"arizona_staging"`
	Rodina         []Server `json:"rodina"`
	Village        []Server `json:"village"`
	Arizonav       []Server `json:"arizonav"`
}

const ArizonaServersList string = "https://api.arizona-five.com/launcher/servers"

func LoadServers() (ArizonaServerInfo, error) {
	var servers ArizonaServerInfo
	var err error
	response, err := http.Get(ArizonaServersList)
	if err == nil {
		bytes, err := io.ReadAll(response.Body)
		if err == nil {
			err = json.Unmarshal(bytes, &servers)
			if err == nil {
				sort.Slice(servers.Arizona, func(a, b int) bool {
					return servers.Arizona[a].Number < servers.Arizona[b].Number
				})
			}
		}
	}
	return servers, err
}
