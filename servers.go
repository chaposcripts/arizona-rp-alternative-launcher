package main

import (
	"encoding/json"
	"io"
	"net/http"
	"sort"
)

var serverList []Server

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
					// fmt.Printf("SORTING %d %d %s, %d %d %s \n", a, servers.Arizona[a].Number, servers.Arizona[a].Name,)
					return servers.Arizona[a].Number < servers.Arizona[b].Number
				})
			}
		}
	}
	serverList = servers.Arizona
	return servers, err
}

// type Settings struct {
// 	Name   string
// 	Path   string
// 	Memory int
// }

// func BuildCommandLine(settings Settings, server Server, settingsStr string) (string, error) {
// 	var err error
// 	var line string = fmt.Sprintf(`start /d "%s" gta_sa.exe -c -h %s -p %d -n %s -mem %d -x %s`, settings.Path, server.IP, server.Port, settings.Name, settings.Memory, settingsStr)
// 	return line, err
// }
