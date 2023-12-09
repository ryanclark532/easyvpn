<script lang="ts">
	import { Badge, Button, Checkbox, Heading, Input, P, Toggle, Tooltip } from 'flowbite-svelte';
	import Sidepanel from '../../sidepanel.svelte';
	import ConfirmationModal from '../../../../components/confirmation-modal.svelte';
	import type { VpnSettings } from '../../../../types/settings';
	export let data: VpnSettings;
	let open = false;
</script>

<div class="flex">
	<Sidepanel />
	<div class="w-full p-5">
		<div class="mb-4">
			<Heading tag="h2" class="mb-2">VPN Settings</Heading>
			<P
				>VPN Settings control how your vpn server functions. And what settings the server serves to
				clients.</P
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
			</svg><P>Changing VPN settings will require the VPN server to restart</P></Badge
		>
		<Heading tag="h4" class="mb-1">VPN Subnet</Heading>
		<div class="flex mb-4">
			<Input placeholder="IP Address..." class="mr-4 w-3/4" bind:value={data.vpn_subnet} />
			<Input placeholder="Subnet Mask..." class="w-1/4"  bind:value={data.vpn_subnet_mask}/>
		</div>
		<Tooltip>VPN clients will be assigned ip addresses in this subnet</Tooltip>
		<Heading tag="h4" class="mb-1">DNS Servers</Heading>
		<div class="flex mb-4">
			<Input placeholder="DNS Server 1" class="mr-2 w-1/2" bind:value={data.dnsserver1} />
			<Input placeholder="DNS Server 2" class="ml-2 w-1/2" bind:value={data.dnsserver2} />
		</div>
		<Tooltip>DNS Servers that will be pushed to VPN clients</Tooltip>
<Heading tag="h4" class="mb-1">Private Network</Heading>
		<div class="flex mb-4 p-2 rounded  border border-gray-200">
			<P class="w-1/2" >Use VPN Server As Gateway</P>
			<Toggle class="w-1/2 justify-end" bind:checked={data.use_as_gateway}/>
			</div>
		<Tooltip>Whether the VPN server will be used as the gateway for all client traffic. This is required for geolocation hiding</Tooltip>
		<div class="flex mb-4 p-2 rounded  border border-gray-200">
			<P class="w-1/2" >Allow VPN Clients Access to Private Network</P>
			<Toggle class="w-1/2 justify-end" bind:checked={data.private_access}/>
			</div>
		<Tooltip>Whether VPN clients will be able to discover and comminucate with other devices connected to the VPN servers network</Tooltip>
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
	onConfirm={() => console.log('Confirmed')}
/>
