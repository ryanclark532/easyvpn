import { checkToken, getToken } from '$lib/auth';
import { getUsers, users } from '$lib/users';
import { redirect } from '@sveltejs/kit';
import {getVpnStatus} from "$lib/vpn";

export const ssr = false;

/** @type {import('../../../../.svelte-kit/types/src/routes').PageLoad} */
export async function load() {
    const token = await checkToken();
    if (!token.passwordValid && token.valid) {
        throw redirect(307, '/user/reset');
    }
    if (!token.admin && token.valid) {
        throw redirect(307, '/user');
    }

    if (!token.valid) {
        throw redirect(307, 'login');
    }

    const status = await getVpnStatus(getToken())
    return {
       serverStatus: status
    }

}
