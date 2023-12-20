<script lang="ts">
	import { Badge, Button, Checkbox, Heading, Input, P, Toggle, Tooltip } from 'flowbite-svelte';
	import Sidepanel from '../../sidepanel.svelte';
	import ConfirmationModal from '../../../../components/confirmation-modal.svelte';
	import type { Settings } from '../../../../types/settings';
	import { setSettings } from '$lib/settings';
	export let data: Settings;
	let open = false;
</script>

<div class="flex">
	<Sidepanel />
	<div class="w-full p-5">
		<div class="mb-4">
			<Heading tag="h2" class="mb-2">Authentication Settings</Heading>
			<P>Authentication Settings control how VPN clients authenticate with the server.</P>
		</div>
		<Heading tag="h4" class="mb-1">Maximum Authentication Attempts</Heading>
		<Input placeholder="5..." class="mb-4" bind:value={data.max_auth_attempts} />
		<Tooltip>After the maximum amount of attempts a users account will be locked</Tooltip><Heading
			tag="h4"
			class="mb-1">Lockout Timeout</Heading
		>
		<Input placeholder="500000..." class="mb-4" bind:value={data.lockout_timeout} />
		<Tooltip
			>The amount of time a user will be locked out after incorrect password attempts. Expressed in
			miliseconds</Tooltip
		>
		<div class="flex mb-4 p-2 rounded border border-gray-200">
			<P class="w-1/2">Allow Users to change their own password</P>
			<Toggle class="w-1/2 justify-end" bind:checked={data.allow_change_pw} />
		</div>
		<Tooltip>Whether VPN users will be able to change their own password via the user panel</Tooltip
		>
		<div class="flex mb-4 p-2 rounded border border-gray-200">
			<P class="w-1/2">Enforce Strong Passwords</P>
			<Toggle class="w-1/2 justify-end" bind:checked={data.enforce_strong_pw} />
		</div>
		<Tooltip
			>Whether VPN users will be forced to set strong passwords. A strong password must be 8
			characters and contain at least one capitol, number and special symbol</Tooltip
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
	onConfirm={() => setSettings(data)}
/>
