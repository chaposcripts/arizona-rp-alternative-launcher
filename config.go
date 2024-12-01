package main

type ConfigParam struct {
	WideScreen    bool `json:"wideScreen"`
	AutoLogin     bool `json:"autoLogin"`
	Preload       bool `json:"preload"`
	Windowed      bool `json:"windowed"`
	Seasons       bool `json:"seasons"`
	Graphics      bool `json:"graphics"`
	ShitPc        bool `json:"shitPc"`
	CefDirtyRects bool `json:"cefDirtyRects"`
	AuthCef       bool `json:"authCef"`
	Grass         bool `json:"grass"`
	OldResolution bool `json:"oldResolution"`
	HdrResolution bool `json:"hdrResolution"`
}

type Config struct {
	Name           string      `json:"name"`
	Path           string      `json:"path"`
	Memory         int         `json:"memory"`
	SelectedServer int         `json:"selectedServer"`
	Params         ConfigParam `json:"params"`
}

/*
var defaultConfig Config = Config{
	Name:           "",
	Path:           "",
	Memory:         4096,
	SelectedServer: 1,
	Params: ConfigParam{
		WideScreen:    true,
		AutoLogin:     false,
		Preload:       true,
		Windowed:      false,
		Seasons:       false,
		Graphics:      false,
		ShitPc:        false,
		CefDirtyRects: false,
		AuthCef:       true,
		Grass:         false,
		OldResolution: false,
		HdrResolution: false,
	},
}
*/
