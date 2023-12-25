import type { Handle } from '@sveltejs/kit';
import type { AuthUser } from './types/auth';

export const handle: Handle = async ({ event, resolve }) => {
	console.log(event.url.pathname);
	const url = event.url.pathname;
	if (url === '/login' || url === '/login/reset' || url === '/') {
		return await resolve(event);
	}
	const { headers } = event.request;
	const jwt = headers.get('cookie')?.substring(4).split(';')[0];
	if (!jwt) {
		return await resolve(event);
	}
	const authcheck = await event
		.fetch('http://localhost:8080/auth/user', {
			headers: {
				JWT: jwt
			},
			credentials: 'include'
		})
		.then((response) => response.json());
	if ('error' in authcheck) {
		return new Response('Redirect', { status: 303, headers: { Location: '/login' } });
	}
	if (new Date(authcheck.attrs.password_expiry) < new Date()) {
		return new Response('Redirect', { status: 303, headers: { Location: '/login/reset' } });
	}
	event.locals.user = authcheck as AuthUser;
	return await resolve(event);
};
