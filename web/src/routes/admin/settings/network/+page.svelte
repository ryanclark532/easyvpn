<script lang="ts">
	import { Badge, Button, Heading, Input, Label, P, Toggle, Tooltip } from 'flowbite-svelte';
	import Sidepanel from '../../sidepanel.svelte';
	import ConfirmationModal from '../../../../components/confirmation-modal.svelte';
	import type { NetworkSettings } from '../../../../types/settings';
	let open: boolean = false;
	export let data: NetworkSettings;
	let tcp = data.protocol === 'tcp';
	let udp = data.protocol === 'udp';
</script>

<div class="flex">
	<Sidepanel />
	<div class="w-full p-5">
		<div class="mb-4">
			<Heading tag="h2" class="mb-2">Network Settings</Heading>
			<P
				>This page contains the Network settings for the VPN Server, the Admin Web Server and the
				Client Web Server</P
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
			</svg><P
				>Changing the Hostname or Protocol after VPN clients are deployed will cause the existing
				clients to be unusable (until a new client configuration is downloaded)</P
			></Badge
		>
		<Heading tag="h4" class="mb-1">Hostname or IP Address</Heading>
		<Input placeholder="IP Address..." class="mb-4" bind:value={data.ip_address} />
		<Tooltip
			>This is the Public IP Address Or Hostname that VPN clients will use to connect to the Server</Tooltip
		>
		<Heading tag="h4" class="mb-1">Protocol</Heading>
		<div class="border border-gray-300 rounded mb-4">
			<div class="flex my-6 mx-2">
				<P class="justify-start w-1/2">TCP</P>
				<Toggle
					size="large"
					class="w-1/2 justify-end"
					bind:checked={tcp}
					on:click={() => {
						tcp = !tcp;
						udp = !tcp;
					}}
				/>
			</div>
			<div class="flex my-6 mx-2">
				<P tag="h5" class="justify-start w-1/2">UDP</P>
				<Toggle
					size="large"
					class="w-1/2 justify-end"
					bind:checked={udp}
					on:click={() => {
						udp = !udp;
						tcp = !udp;
					}}
				/>
			</div>
		</div>
		<Tooltip
			>The protocol used by the VPN server. TCP is used for reliability and security, UDP is used
			for speed</Tooltip
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
		<Input placeholder="IP Address..." class="mb-4" bind:value={data.web_server_port} />
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
	onConfirm={() => console.log('Confirmed')}
/>
