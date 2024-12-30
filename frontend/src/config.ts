import { ReadConfig, SaveConfig } from '../wailsjs/go/main/App';
export type Config = {
    name: string,
    path: string,
    memory: number,
    selectedServer: number,
    pass: string,
    params: {
        wideScreen:    boolean,
        autoLogin:     boolean,
        preload:       boolean,
        windowed:      boolean,
        seasons:       boolean,
        graphics:      boolean,
        shitPc:        boolean,
        cefDirtyRects: boolean,
        authCef:       boolean,
        grass:         boolean,
        oldResolution: boolean,
        hdrResolution: boolean,
    },
}
;
export var config: Config = {
    name: 'Nick_Name',
    path: '',
    memory: 4096,
    selectedServer: 1,
    pass: '',
    params: {
        wideScreen:    false,
        autoLogin:     false,
        preload:       false,
        windowed:      false,
        seasons:       false,
        graphics:      false,
        shitPc:        false,
        cefDirtyRects: false,
        authCef:       false,
        grass:         false,
        oldResolution: false,
        hdrResolution: false,
    }
}

export const parameterName: Record<string, string> = {
    "windowed": "window",
    "autoLogin": "x",
    "wideScreen": "widescreen",
    "preload": "ldo",
    "seasons": "seasons",
    "graphics": "graphics",
    "shitPc": "t",
    "cefDirtyRects": "cef_dirty_rects",
    "authCef": "auth_cef_enable",
    "grass": "enable_grass",
    "oldResolution": "16bpp",
    "hdrResolution": "allow_hdr"
}

export function loadConfig() {
    ReadConfig().then((cfg) => {
        console.log('Config loaded:', cfg);
        config = JSON.parse(cfg) as Config;
        // @ts-ignore
        document.getElementById('name').value = config.name;
        // @ts-ignore
        document.getElementById('memory').value = config.memory;
        // @ts-ignore
        document.getElementById('path').value = config.path;
        // @ts-ignore
        document.getElementById('pass').value = config.pass ?? '';

        // @ts-ignore
        for (const [k, v] of Object.entries(config.params)) document.getElementById(k).checked = v;
    });
}

export function saveConfig() {
    // @ts-ignore
    config.name = document.getElementById('name').value;
    // @ts-ignore
    config.memory = document.getElementById('memory').value;
    // @ts-ignore
    config.path = document.getElementById('path').value;
    // @ts-ignore
    config.pass = document.getElementById('pass').value;

    
    for (const name of Object.keys(parameterName)) {
        // @ts-ignore
        config.params[name] = document.getElementById(name).checked ?? false;
    }
    console.log(config);
    console.log(Object.keys(config.params), JSON.stringify(config));
    SaveConfig(JSON.stringify(config));
}
