<script lang="ts">
	import { Avatar, Button, Dropdown, DropdownItem, P } from 'flowbite-svelte';
	import ConfirmationModal from './confirmation-modal.svelte';
	import { setToken } from '$lib/auth';
	import { goto } from '$app/navigation';

	export let username: string;
	export let placement: string = 'bottom';

	let signOutOpen = false;
</script>

<div class="m-4 cursor-pointer">
	<div class="flex">
		<Avatar>{username[0].toUpperCase()}</Avatar>
		<P class="mt-2 ml-2 text-lg">{username}</P>
	</div>
	<Dropdown {placement} class="w-64">
		<DropdownItem>
			<P class="text-lg ml-2 ">My Profile</P>
		</DropdownItem>
		<DropdownItem on:click={() => (signOutOpen = !signOutOpen)}>
			<P class="text-lg ml-2 ">Sign Out</P>
		</DropdownItem>
	</Dropdown>
</div>

<ConfirmationModal
	open={signOutOpen}
	title="Sign Out"
	subtext="Do you want to sign out of your account?"
	onConfirm={() => {
		setToken('');
		goto('/login');
	}}
/>
