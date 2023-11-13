export type ServerStatus = 'notRunning' | 'running' | 'starting' | 'unknown';

export type ServerStatusResponse = {
	status: ServerStatus;
};

export const ServerStatusMapping: Record<ServerStatus, string> = {
	running: 'The VPN Server Is Running',
	notRunning: 'The VPN Server Is Not Running',
	starting: 'The VPN Server Is Starting',
	unknown: 'We failed to get the Server Status, please try again'
};

export type Connection = {
	CommonName: string;
	Address: string;
	BytesRec: string;
	BytesSent: string;
	ConnectedSince: Date;
};

export type ActiveConnectionsResponse = {
	connections: Connection[];
};
