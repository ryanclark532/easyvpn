import { handleRedirects } from '$lib/auth';
import type { Connection, ServerStatus } from '../../../types/vpn';

export async function load({ cookies }) {
	const connectionsResponse = await fetch('http://localhost:8080/vpn/connections', {
		headers: {
			JWT: cookies.get('JWT')
		},
		credentials: 'include',
		method: 'GET'
	}).then((response) => response.json());
return {
connections: (connectionsResponse.connections as Connection[]) || []
}
}


