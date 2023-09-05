<script lang="ts">
	import { Label, Input, Button, Modal, Textarea, Select, Toggle, Checkbox } from 'flowbite-svelte';
	import { userValidationMessage } from '../../stores/stores';
	import { createUser } from '$lib';

	let hidden: boolean = false;
	let confirmed: boolean = false;

	let name: string;
	let username: string;
	let password: string;
	let enabled: boolean = false;
	let admin: boolean = false;

	async function handleUserCreation() {
		hidden = !(await createUser(name, username, password, enabled, admin));
		name = '';
		username = '';
		password = '';
		enabled = false;
		admin = false;
	}
</script>

<div>
	<div class="flex justify-center m-5">
		<Button on:click={() => (hidden = true)}>Create User</Button>
	</div>
	<Modal title="Create New User" bind:open={hidden} autoclose class="min-w-full">
		{#if $userValidationMessage}
			<div class="bg-red-200 w-full py-1 text-center rounded-lg mb-6">
				<p class="text-sm text-red-600 font-medium">
					{$userValidationMessage}
				</p>
			</div>
		{/if}
		<form on:submit={handleUserCreation}>
			<div class="mb-6">
				<Label for="name">Name</Label>
				<Input type="text" id="name" placeholder="The users name..." required bind:value={name} />
			</div>
			<div class="mb-6">
				<Label for="username">Username</Label>
				<Input
					type="text"
					id="username"
					placeholder="A username to login with..."
					required
					bind:value={username}
				/>
			</div>
			<div class="mb-6">
				<Label for="password">Password</Label>
				<Input
					type="text"
					id="password"
					placeholder="Password for the user..."
					required
					bind:value={password}
				/>
			</div>
			<div class="grid gap-4 sm:grid-cols-2 mb-6">
				<Toggle id="enabled" bind:checked={enabled}>Enabled</Toggle>
				<Toggle id="admin" bind:checked={admin}>Admin</Toggle>
			</div>
			<div class="w-full bg-gray-200 p-3 flex items-center mb-6 rounded-lg">
				<Checkbox bind:checked={confirmed}>Default checkbox</Checkbox>
			</div>
			<Button disabled={!confirmed} on:click={() => handleUserCreation()} class="w-full">
				Create User</Button
			>
		</form>
	</Modal>
</div>
