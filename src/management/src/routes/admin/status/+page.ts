import { checkToken, getToken } from '$lib/auth';
import { redirect } from '@sveltejs/kit';
import {getVpnStatus, status} from "$lib/vpn";

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

    const response = await getVpnStatus(getToken())
    if(response.status !== "error"){
        status.set(response.data?.status ?? "unknown")
        return
    }
    status.set("unknown")

}
