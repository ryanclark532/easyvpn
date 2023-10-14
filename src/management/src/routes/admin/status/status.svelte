<script lang="ts">
	import { ServerStatusMapping } from '../../../types/vpn';
	import { Button } from 'flowbite-svelte';
	import { redirect } from '@sveltejs/kit';
	import ConfirmationModal from '../../../components/confirmation-modal.svelte';
	import { page } from '$app/stores';
	import { invalidate } from '$app/navigation';
	import { vpnOperation } from '$lib/vpn';
	import type { PageData } from './$types';

	let background: string;
	let containerClass: string;

	export let data: PageData;
	$: {
		background =
			data.status === 'running'
				? 'bg-green-300 border-green-600'
				: data.status === 'notRunning'
				? 'bg-red-300 border-red-600'
				: 'bg-orange-300 border-orange-600';

		containerClass = `text-center text-l text font-medium p-2  w-full rounded-lg ${background}`;
	}

	let stop: boolean;
	let start: boolean;
	let restart: boolean;
</script>

<div class={containerClass}>
	<div class="flex justify-center">
		<h2 class="mt-5 text-2xl">{ServerStatusMapping[$page.data.status]}</h2>
		<button class="ml-2 mt-6" on:click={() => invalidate('http://localhost:8080/vpn')}>
			<svg xmlns="http://www.w3.org/2000/svg" height="2em" viewBox="0 0 512 512"
				><style>
					svg {
						fill: #000000;
					}
				</style><path
					d="M142.9 142.9c62.2-62.2 162.7-62.5 225.3-1L327 183c-6.9 6.9-8.9 17.2-5.2 26.2s12.5 14.8 22.2 14.8H463.5c0 0 0 0 0 0H472c13.3 0 24-10.7 24-24V72c0-9.7-5.8-18.5-14.8-22.2s-19.3-1.7-26.2 5.2L413.4 96.6c-87.6-86.5-228.7-86.2-315.8 1C73.2 122 55.6 150.7 44.8 181.4c-5.9 16.7 2.9 34.9 19.5 40.8s34.9-2.9 40.8-19.5c7.7-21.8 20.2-42.3 37.8-59.8zM16 312v7.6 .7V440c0 9.7 5.8 18.5 14.8 22.2s19.3 1.7 26.2-5.2l41.6-41.6c87.6 86.5 228.7 86.2 315.8-1c24.4-24.4 42.1-53.1 52.9-83.7c5.9-16.7-2.9-34.9-19.5-40.8s-34.9 2.9-40.8 19.5c-7.7 21.8-20.2 42.3-37.8 59.8c-62.2 62.2-162.7 62.5-225.3 1L185 329c6.9-6.9 8.9-17.2 5.2-26.2s-12.5-14.8-22.2-14.8H48.4h-.7H40c-13.3 0-24 10.7-24 24z"
				/></svg
			>
		</button>
	</div>
	<div class="mt-7 mb-7">
		{#if $page.data.status === 'running'}
			<Button class="w-40" on:click={() => (stop = !stop)}>Stop Server</Button>
			<Button class="w-40" on:click={() => (restart = !restart)}>Restart Server</Button>
		{/if}
		{#if $page.data.status === 'starting'}
			<Button
				class="w-40"
				on:click={() => {
					throw redirect(307, '/admin/users');
				}}
			>
				See Users
			</Button>
			<Button class="w-40">Restart Server</Button>
		{/if}
		{#if $page.data.status === 'notRunning'}
			<Button
				class="w-40"
				on:click={() => {
					throw redirect(307, '/admin/users');
				}}
			>
				See Users
			</Button>
			<Button class="w-40" on:click={() => (start = !start)}>Start Server</Button>
		{/if}
	</div>
</div>

<ConfirmationModal open={stop} title="Stop VPN Server" onConfirm={vpnOperation} data={'stop'} />

<ConfirmationModal
	open={restart}
	title="Restart VPN Server"
	onConfirm={vpnOperation}
	data={'restart'}
/>
<ConfirmationModal open={start} title="Start VPN Server" onConfirm={vpnOperation} data={'start'} />
