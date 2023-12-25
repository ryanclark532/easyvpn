<script lang="ts">
	import { Badge, Button, Checkbox, Heading, Input, P, Toggle, Tooltip } from 'flowbite-svelte';
	import Sidepanel from '../../sidepanel.svelte';
	import ConfirmationModal from '../../../../components/confirmation-modal.svelte';
	import { setSettings } from '$lib/settings';
	import type { Settings } from '../../../../types/settings';
	import type { PageData } from './$types';
	export let data: PageData;
	let open = false;
</script>

<div class="flex">
	<Sidepanel username={data.username} />
	<div class="w-full p-5">
		<div class="mb-4">
			<Heading tag="h2" class="mb-2">Client Settings</Heading>
			<P>Client Settings control how your VPN clients behave and what they have acess to.</P>
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
				>Changing some client settings may require the VPN server to restart. Clients will also need
				to download and import a new config file from the user panel to have changes take effect.</P
			></Badge
		>

		<Tooltip>VPN clients will be assigned ip addresses in this subnet</Tooltip>
		<Heading tag="h4" class="mb-1">DNS Servers</Heading>
		<div class="flex mb-4">
			<Input placeholder="DNS Server 1" class="mr-2 w-1/2" bind:value={data.settings.dnsserver1} />
			<Input placeholder="DNS Server 2" class="ml-2 w-1/2" bind:value={data.settings.dnsserver2} />
		</div>
		<Tooltip>DNS Servers that will be pushed to VPN clients</Tooltip>
		<Heading tag="h4" class="mb-1">Private Network</Heading>
		<div class="flex mb-4 p-2 rounded border border-gray-200">
			<P class="w-1/2">Use VPN Server As Gateway</P>
			<Toggle class="w-1/2 justify-end" bind:checked={data.settings.use_as_gateway} />
		</div>
		<Tooltip
			>Whether the VPN server will be used as the gateway for all client traffic. This is required
			for geolocation hiding</Tooltip
		>
		<div class="flex mb-4 p-2 rounded border border-gray-200">
			<P class="w-1/2">Allow VPN Clients Access to Private Network</P>
			<Toggle class="w-1/2 justify-end" bind:checked={data.settings.private_access} />
		</div>
		<Tooltip
			>Whether VPN clients will be able to discover and comminucate with other devices connected to
			the VPN servers network</Tooltip
		>
		<Button
			class="w-full"
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
	onConfirm={() => setSettings(data.settings)}
/>
