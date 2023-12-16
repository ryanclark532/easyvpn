<script lang="ts">
	import { Badge, Button, Heading, Input, Label, P, Toggle, Tooltip } from 'flowbite-svelte';
	import Sidepanel from '../../sidepanel.svelte';
	import ConfirmationModal from '../../../../components/confirmation-modal.svelte';
	import type { PageData } from './$types';
	import { setSettings } from '$lib/settings';
	import type { Settings } from '../../../../types/settings';
	let open: boolean = false;
	export let data: Settings;
</script>

<div class="flex">
	<Sidepanel />
	<div class="w-full p-5">
		<div class="mb-4">
			<Heading tag="h2" class="mb-2">Server Settings</Heading>
			<P
				>Server settings dictate how your server behaves, as well as how the admin panel is served
				to you.</P
			>
		</div>

		<Badge color="yellow" class="mb-2 w-full p-3">
			<svg
				class="ml-2 mr-2 w-6 h-6 text-gray-800 dark:text-white"
				aria-hidden="true"
				xmlns="http://www.w3.org/2000/svg"
				fill="none"
				viewBox="0 0 20 20"
			>
				<path
					stroke="currentColor"
					stroke-linecap="round"
					stroke-linejoin="round"
					stroke-width="2"
					d="M10 11V6m0 8h.01M19 10a9 9 0 1 1-18 0 9 9 0 0 1 18 0Z"
				/>
			</svg><P>Changing Server settings will require the VPN server to be restarted.</P></Badge
		><Heading tag="h4" class="mb-1">VPN Subnet</Heading>
		<div class="flex mb-4">
			<Input placeholder="IP Address..." class="mr-4 w-3/4" bind:value={data.server.vpn_subnet} />
			<Input placeholder="Subnet Mask..." class="w-1/4" bind:value={data.server.vpn_subnet_mask} />
		</div>
		<Heading tag="h4" class="mb-1">Hostname or IP Address</Heading>
		<Input placeholder="IP Address..." class="mb-4" bind:value={data.server.ip_address} />
		<Tooltip
			>This is the Public IP Address Or Hostname that VPN clients will use to connect to the Server</Tooltip
		>
		<Heading tag="h4" class="mb-1">VPN Server Port</Heading>
		<Input placeholder="IP Address..." class="mb-4" bind:value={data.server.port} />
		<Tooltip
			>This is the port the VPN server will listen on. This will need to be port forwarded on your
			router</Tooltip
		>
		<Badge color="yellow" class="mb-2 w-full p-3">
			<svg
				class="ml-2 mr-2 w-6 h-6 text-gray-800 dark:text-white"
				aria-hidden="true"
				xmlns="http://www.w3.org/2000/svg"
				fill="none"
				viewBox="0 0 20 20"
			>
				<path
					stroke="currentColor"
					stroke-linecap="round"
					stroke-linejoin="round"
					stroke-width="2"
					d="M10 11V6m0 8h.01M19 10a9 9 0 1 1-18 0 9 9 0 0 1 18 0Z"
				/>
			</svg><P
				>Changing the web server port will cause the admin panel to become unresponsive. You will
				need to navigate to the new port to regain access
			</P></Badge
		>
		<Heading tag="h4" class="mb-1">Web Server Port</Heading>
		<Input placeholder="IP Address..." class="mb-4" bind:value={data.server.web_server_port} />
		<Tooltip>The port that the web server is served on</Tooltip>
		<Button
			class="mt-4 w-full"
			on:click={() => {
				open = !open;
			}}>Save Settings</Button
		>
	</div>
</div>

<ConfirmationModal
	title="Confirm Settings Changes"
	subtext="Changing settings will cause the vpn server to restart, and disconnect any connected clients. Please confirm this action"
	{open}
	onConfirm={() => setSettings(data)}
/>
