import { redirect } from '@sveltejs/kit';
import { getVpnStatus, status } from '$lib/vpn';

export const ssr = false;

/** @type {import('../../../../.svelte-kit/types/src/routes').PageLoad} */
export async function load() {}
