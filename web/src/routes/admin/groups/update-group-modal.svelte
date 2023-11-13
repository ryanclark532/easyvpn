<script lang="ts">
	import { Label, Input, Button, Modal, Checkbox, Tooltip, DropdownItem } from 'flowbite-svelte';
	import { createGroup, createGroupResponse, updateGroup } from '$lib/groups';
	import type { Group } from '../../../types/groups';
	let defaultModal = false;

	export let group: Group;

	let updateResponse: Error | undefined;
</script>

<div>
	<DropdownItem on:click={() => (defaultModal = true)}>Update Group</DropdownItem>
</div>
<Modal title="Update Group" bind:open={defaultModal} class="min-w-full">
	{#if updateResponse}
		<div class="bg-red-300 p-1 rounded text-center">
			<p class="text-base text-red-600">{updateResponse.message}</p>
		</div>
	{/if}
	<form
		on:submit={async (e) => {
			updateResponse = await updateGroup(group.id, e);
		}}
	>
		<div class="mb-4">
			<Label for="name" class="mb-2">Group Name</Label>
			<Input
				type="text"
				name="name"
				id="name"
				placeholder="Group Name..."
				requiredi
				bind:value={group.name}
			/>
		</div>
		<div class="w-full flex mb-4">
			<div class="w-1/2">
				<Checkbox id="enabled" name="enabled" bind:checked={group.enabled}
					>Users in group enabled</Checkbox
				>
				<Tooltip>Roles set per user take priority over group roles</Tooltip>
			</div>
			<div class="w-1/2">
				<Checkbox id="is_admin" name="is_admin" bind:checked={group.is_admin}
					>Users in group are admins</Checkbox
				>
				<Tooltip>Roles set per user take priority over group roles</Tooltip>
			</div>
		</div>
		<Button class="w-full" type="submit">Update Group</Button>
	</form>
</Modal>
