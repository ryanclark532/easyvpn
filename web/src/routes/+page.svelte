<script lang="ts">
	import { Heading, P, Tooltip } from 'flowbite-svelte';
	import { onDestroy } from 'svelte';
	import type { PageData } from './$types';
	import { deleteUserConfig } from '$lib/users';
	import { beforeNavigate } from '$app/navigation';

	export let data: PageData;

	beforeNavigate(() => {
		deleteUserConfig(data.username);
	});
</script>

<div class="h-screen flex items-center justify-center flex-col">
	<Heading tag="h3" class="mb-4 pb-2 border-b-2 border-primary-500 w-1/3 text-center"
		>EasyVPN User Center</Heading
	>
	<a
		class="text-2xl text-center text-blue-500 flex mb-4"
		href="http://localhost:8080/tmp/{data.username}.ovpn"
		>OpenVPN Config File<svg
			class="ml-2 mt-1 w-6 h-6 text-blue-500 dark:text-white"
			aria-hidden="true"
			xmlns="http://www.w3.org/2000/svg"
			fill="none"
			viewBox="0 0 16 18"
		>
			<path
				stroke="currentColor"
				stroke-linecap="round"
				stroke-linejoin="round"
				stroke-width="2"
				d="M8 1v11m0 0 4-4m-4 4L4 8m11 4v3a2 2 0 0 1-2 2H3a2 2 0 0 1-2-2v-3"
			/>
		</svg></a
	>
	<Tooltip>Downloads your unique VPN config to connect to this VPN server</Tooltip>
	<a
		class="text-2xl text-center text-blue-500 flex"
		href="https://openvpn.net/client/"
		target="_blank"
		>OpenVPN Connect Client <svg
			class="ml-2 mt-1 w-6 h-6 text-blue-500 dark:text-white"
			aria-hidden="true"
			xmlns="http://www.w3.org/2000/svg"
			fill="none"
			viewBox="0 0 18 15"
		>
			<path
				stroke="currentColor"
				stroke-linecap="round"
				stroke-linejoin="round"
				stroke-width="2"
				d="M1 7.5h11m0 0L8 3.786M12 7.5l-4 3.714M12 1h3c.53 0 1.04.196 1.414.544.375.348.586.82.586 1.313v9.286c0 .492-.21.965-.586 1.313A2.081 2.081 0 0 1 15 14h-3"
			/>
		</svg></a
	>
	<Tooltip
		>This will navigate to an external website to select the correct version of OpenVPN Connect for
		your system.</Tooltip
	>
</div>
