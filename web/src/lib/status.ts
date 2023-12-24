import { writable } from 'svelte/store';

export const vpnLog = writable<string[][]>([[]]);

export function getVpnLogs() {
	const socket = new WebSocket('ws://localhost:8080/vpn/log');

	socket.addEventListener('message', (event: { data: string }) => {
		const lines = event.data.split('\n');
		const formattedLines = lines.map((line) => {
			return splitOnFirstLetter(line);
		});
		vpnLog.set(formattedLines);
	});
}

export function splitOnFirstLetter(inputString: string) {
	const match = RegExp(/[a-zA-Z]/).exec(inputString);
	if (match) {
		const firstLetterIndex = match.index;
		return [inputString.substring(0, firstLetterIndex), inputString.substring(firstLetterIndex!)];
	} else {
		return [inputString];
	}
}
