import { handleRedirects } from '$lib/auth';

export const ssr = false;

/** @type {import('../../../../.svelte-kit/types/src/routes').PageLoad} */
export async function load({ fetch, cookies, locals }) {
	const headers = new Headers();
	headers.append('JWT', cookies.get('JWT') ?? '');
	const username = locals.user.name;
	await fetch(`http://localhost:8080/user/config/${username}`, {
		method: 'POST',
		headers,
		credentials: 'include'
	});

	return {
		username
	};
}
