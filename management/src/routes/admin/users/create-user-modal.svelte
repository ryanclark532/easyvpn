<script lang="ts">
	import { Label, Input, Button, Modal, Checkbox, Tooltip } from 'flowbite-svelte';
	import { createUser, createUserResponse } from '$lib/users';
	import { onDestroy } from 'svelte';
	let defaultModal = false;
	let createUserConfirmation = false;
	const s = createUserResponse.subscribe((e) => (createUserConfirmation = e.status === 'ready'));

	onDestroy(s);
</script>

<div class="float-right mt-7">
	<Button on:click={() => (defaultModal = true)}>Create User</Button>
</div>
<Modal title="User Creation Successful" bind:open={createUserConfirmation} autoclose size="sm" class="w-full">
	<p class="mb-4 text-gray-500 dark:text-gray-300 text-center">
        {$createUserResponse.data}
	</p>
</Modal>

<Modal title="Add User" bind:open={defaultModal} class="min-w-full">
	{#if $createUserResponse.status === 'error'}
		<h5 class="text-l font-medium text-red-600 p-2 bg-red-300 w-full rounded-lg border-red-600">
			{$createUserResponse.data}
		</h5>
	{/if}
	<form
		on:submit={(e) => {
			createUser(e);
			defaultModal = false;
		}}
	>
		<div>
			<div class="mb-4">
				<Label for="name" class="mb-2">Name</Label>
				<Input type="text" id="name" name="name" placeholder="Name..." required />
			</div>
			<div class="mb-4">
				<Label for="username" class="mb-2">Username</Label>
				<Input
					type="text"
					id="username"
					name="username"
					placeholder="Username the user will login with..."
					required
				/>
			</div>
			<div class="mb-4">
				<Label for="password" class="mb-2">Password</Label>
				<Input
					type="text"
					id="password"
					name="password"
					placeholder="Secure password for the user..."
					required
				/>
			</div>
			<div class="grid grid-cols-2 mb-4">
				<Checkbox id="is_admin" name="is_admin">Admin User</Checkbox>
				<Tooltip>Admin users will have access to the admin panel</Tooltip>
				<Checkbox id="enabled" name="enabled">User Enabled</Checkbox>
				<Tooltip>User is enabled and can be used</Tooltip>
			</div>

			<Button type="submit" class="w-full">
				<svg
					class="mr-1 -ml-1 w-6 h-6"
					fill="currentColor"
					viewBox="0 0 20 20"
					xmlns="http://www.w3.org/2000/svg"
					><path
						fill-rule="evenodd"
						d="M10 5a1 1 0 011 1v3h3a1 1 0 110 2h-3v3a1 1 0 11-2 0v-3H6a1 1 0 110-2h3V6a1 1 0 011-1z"
						clip-rule="evenodd"
					/></svg
				>
				Add New User
			</Button>
		</div>
	</form>
</Modal>
