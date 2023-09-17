import { checkToken } from '$lib/auth';
import {getUsers, users} from '$lib/users';
import {page} from "$app/stores";

export const ssr = false;

/** @type {import('../../../../.svelte-kit/types/src/routes').PageLoad} */
export async function load() {
	const token = await checkToken();
	const u = await getUsers(token);
	if(!u.data?.users){
		return
	}
	users.set(u.data.users);
}
