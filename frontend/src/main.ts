import './style.css';
import './app.css';
import { EventsOn, EventsEmit } from '../wailsjs/runtime'
import { config, loadConfig, saveConfig, parameterName } from './config';
import { StartGame } from '../wailsjs/go/main/App';

type Server = {
    number:               number,
	name:                 string,
	ip:                   string,
	port:                 number,
	online:               number,
	maxplayers:           number,
	password:             boolean,
	vk:                   string,
	tg:                   string,
	inst:                 string,
	icon:                 string,
	additionalIps:        string[],
	donateMultiplier:     number,
	experienceMultiplier: number,
	plotPoints:           {online: number, time: number}[]
}

var serversList: Server[];

function createServer(server: Server, id: number, isSelected = false) {
    const list = document.getElementById('servers-list');
    const serverDiv = document.createElement('div');
    serverDiv.classList.add('server');
    if (isSelected) serverDiv.classList.add('server-selected');
    serverDiv.id = id.toString();
    serverDiv.onclick = () => {
        document.querySelectorAll('.server').forEach((el) => {
            el.classList.remove('server-selected');
        });
        serverDiv.classList.add('server-selected');
        config.selectedServer = server.number;
        saveConfig();
    };
    const serverLogo = document.createElement('img');
    serverLogo.classList.add('server-logo')
    serverLogo.src = server.icon;

    const serverId = document.createElement('a');
    serverId.classList.add('server-id');
    serverId.textContent = `#${server.number}`;

    const serverName = document.createElement('a');
    serverName.classList.add('server-name');
    serverName.textContent = server.name;

    const serverPlayers = document.createElement('a');
    serverPlayers.classList.add('server-players');
    serverPlayers.textContent = `${server.online}/${server.maxplayers}`;
    
    serverDiv.appendChild(serverLogo);
    serverDiv.appendChild(serverId);
    serverDiv.appendChild(serverName);
    serverDiv.appendChild(serverPlayers);
    list?.appendChild(serverDiv);
}

addEventListener('DOMContentLoaded', () => {
    loadConfig();
    EventsEmit('servers:request');
    document.querySelectorAll('input[type=checkbox]').forEach((el) => el.addEventListener('click', () => {
        //@ts-ignore
        config.params[el.id] = el.checked
        saveConfig();
    }));
    //@ts-ignore
    document.getElementById('button-settings')?.addEventListener('click', () => document.getElementById('dialog-settings')?.showModal());
    const nameInput = document.getElementById('name');
    //@ts-ignore
    nameInput.addEventListener('input', () => config.name = nameInput.value);
    document.getElementById('path')?.addEventListener('click', () => EventsEmit('settings:requestFileDialog'))
    document.getElementById('button-play')?.addEventListener('click', () => {
        saveConfig();
        const selectedServer = serversList.find((s) => s.number == config.selectedServer) ?? serversList[0];
        const params: string[] = [
            '-c',
            '-arizona',
            `-h ${selectedServer.ip}`,
            `-p ${selectedServer.port}`,
            `-n ${config.name}`,
            `-mem ${config.memory ?? '2048'}`,
            '-referrer',
            '-userId undefined'
        ];
        for (const [param, value] of Object.entries(config.params)) {
            if (value) params.push(`-${parameterName[param] ?? "-UNKNOWN_" + param}`);
            console.log(param, value);
        }
        console.log(config)
        console.log(`Starting: ${params.join(' ')}`);
        
        StartGame(config.name, config.path, params);
    });

    loadConfig();
})

EventsOn('settings:fileDialogPathSelected', (path: string) => {
    config.path = path;
    // @ts-ignore
    document.getElementById('path').value = path;
});

EventsOn('servers:update', (servers: Server[]) => {
    serversList = servers;
    // alert(servers);
    servers.sort((a, b) => a.number - b.number).forEach((server) => {
        document.getElementById(server.number.toString())?.remove();
        createServer(server, server.number, config.selectedServer == server.number);
    });
    //@ts-ignore
    document.getElementById('servers-list').scrollTop = (config.selectedServer - 1) * 42
});

declare global {
    interface Window {
        startGame: () => void;
    }
}