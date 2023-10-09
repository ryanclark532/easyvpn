<script lang="ts">
	import { Label, Input, Button, Modal, Checkbox, Tooltip } from 'flowbite-svelte';
	import { page } from '$app/stores';
	let defaultModal = false;
	let createUserConfirmation = false;
	let response: string;
	let userStore = $page.data.userStore;
</script>

<button
	on:click={() => (defaultModal = true)}
	type="button"
	class="flex items-center justify-center text-white bg-primary-700 hover:bg-primary-800 focus:ring-4 focus:ring-primary-300 font-medium rounded-lg text-sm px-4 py-2 dark:bg-primary-600 dark:hover:bg-primary-700 focus:outline-none dark:focus:ring-primary-800"
>
	<svg
		class="h-3.5 w-3.5 mr-2"
		fill="currentColor"
		xmlns="http://www.w3.org/2000/svg"
		aria-hidden="true"
	>
		<path
			clip-rule="evenodd"
			fill-rule="evenodd"
			d="M10 3a1 1 0 011 1v5h5a1 1 0 110 2h-5v5a1 1 0 11-2 0v-5H4a1 1 0 110-2h5V4a1 1 0 011-1z"
		/>
	</svg>
	Add User
</button>
<Modal
	title="User Creation Successful"
	bind:open={createUserConfirmation}
	autoclose
	size="sm"
	class="w-full"
/>

<Modal title="Add User" bind:open={defaultModal} class="min-w-full">
	{#if response}
		<h5 class="text-l font-medium text-red-600 p-2 bg-red-300 w-full rounded-lg border-red-600">
			{response}
		</h5>
	{/if}
	<form
		on:submit={async (e) => {
			response = await userStore.create(e);
			defaultModal = !!response;
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
