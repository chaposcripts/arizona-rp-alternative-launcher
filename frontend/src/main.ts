import './style.css';
import './app.css';
import { EventsOn, EventsEmit } from '../wailsjs/runtime'
import { config, loadConfig, saveConfig, parameterName } from './config';
import { StartGame, UpdateServerInfo } from '../wailsjs/go/main/App';

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

var serversList: Server[] = [];

function createServer(server: Server, id: number, isSelected = false, hideNumber = false) {
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
        UpdateServerInfo(server.ip);
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
    serverPlayers.id = `players-count-${server.number}`
    serverPlayers.classList.add('server-players');
    serverPlayers.textContent = `${server.online}/${server.maxplayers}`;
    
    serverDiv.appendChild(serverLogo);
    if (!hideNumber)
        serverDiv.appendChild(serverId);
    serverDiv.appendChild(serverName);
    serverDiv.appendChild(serverPlayers);
    list?.appendChild(serverDiv);
}

addEventListener('DOMContentLoaded', () => {
    loadConfig();
    EventsEmit('servers:request');
    setInterval(() => {
        console.log('Requesting servers update...');
        EventsEmit('servers:request');
    }, 10000);
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
            `-mem ${config.memory.toString().length > 0 ? config.memory : '2048'}`,
            '-referrer',
            '-userId undefined',
            config.pass.length > 0 ? `-z ${config.pass}` : '',
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

EventsOn('servers:update', (servers: Server[], mobileServers: Server[]) => {
    console.log('servers:update');
    const scrollAfterCreation = serversList.length == 0;

    mobileServers.sort((a, b) => a.name.length - b.name.length).forEach((ms) => {
        ms.number += servers.length;
        servers.push(ms)
    });

    serversList = servers;
    // alert(servers);
    servers.sort((a, b) => a.number - b.number).forEach((server) => {
        document.getElementById(server.number.toString())?.remove();
        createServer(server, server.number, config.selectedServer == server.number, server.number > servers.length - mobileServers.length);
    });
    //@ts-ignore
    if (scrollAfterCreation) document.getElementById('servers-list').scrollTop = (config.selectedServer - 1) * 42
});

EventsOn('server:update_players', (host: string, players: number, maxplayers: number) => {
    const server = serversList.find((s) => s.ip == host);
    if (server) {
        const playersCount = document.getElementById(`players-count-${server.number}`);
        if (playersCount) playersCount.textContent = `${players}/${maxplayers}`;
    }
    console.log('server:update_players', host, players, maxplayers);
})