<script lang="ts">
	import { Badge, Button, Heading, P, Textarea } from 'flowbite-svelte';
	import Sidepanel from '../../sidepanel.svelte';
	import type { PageData } from './$types';
	import ConfirmationModal from '../../../../components/confirmation-modal.svelte';

	export let data: PageData;
	let open;
</script>

<div class="flex">
	<Sidepanel />
	<div class="w-full p-5">
		<Heading tag="h2" class="mb-2">Network Settings</Heading>
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
				>Changing the raw VPN config file is only suggested for advanced users. Changes to this file
				could prevent clients from being able to connect to this server. Changes to this file will
				require the server to restart to take effect</P
			></Badge
		>
		<Textarea
			class="h-3/4 mb-4"
			placeholder="VPN Config..."
			name="message"
			bind:value={data.config}
		/>
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
