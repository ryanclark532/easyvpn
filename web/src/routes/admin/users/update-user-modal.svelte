<script lang="ts">
	import { DropdownItem, Modal, Input, Label, Button, Checkbox, Tooltip } from 'flowbite-svelte';
	import type { User } from '../../../types/users';
	import { updateUser } from '$lib/users';

	let open: boolean;
	export let user: User;
	let response: Error | undefined;
</script>

<DropdownItem
	on:click={() => {
		open = !open;
	}}>Update User</DropdownItem
>
<Modal title="Add User" bind:open class="min-w-full">
	<form
		on:submit={async (e) => {
			e.preventDefault();
			response = await updateUser(user);
		}}
	>
		<div>
			{#if response}
				<div class="bg-red-300 p-1 rounded text-center">
					<p class="text-base text-red-600">{response.message}</p>
				</div>
			{/if}
			<div class="mb-4">
				<Label for="name" class="mb-2">Name</Label>
				<Input
					bind:value={user.name}
					type="text"
					id="name"
					name="name"
					placeholder="Name..."
					required
				/>
			</div>
			<div class="mb-4">
				<Label for="username" class="mb-2">Username</Label>
				<Input
					bind:value={user.username}
					type="text"
					id="username"
					name="username"
					placeholder="Username the user will login with..."
					required
				/>
			</div>
			<div class="grid grid-cols-2 mb-4">
				<Checkbox bind:checked={user.is_admin} id="is_admin" name="is_admin">Admin User</Checkbox>
				<Tooltip>Admin users will have access to the admin panel</Tooltip>
				<Checkbox bind:checked={user.enabled} id="enabled" name="enabled">User Enabled</Checkbox>
				<Tooltip>User is enabled and can be used</Tooltip>
			</div>
			<Button type="submit" class="w-full">Update User</Button>
		</div>
	</form>
</Modal>
