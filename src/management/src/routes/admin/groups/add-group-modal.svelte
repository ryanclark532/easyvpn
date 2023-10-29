<script lang="ts">
	import { Label, Input, Button, Modal, Checkbox, Tooltip } from 'flowbite-svelte';
	import { createGroup, createGroupResponse } from '$lib/groups';
	let defaultModal = false;
</script>

<div>
	<Button on:click={() => (defaultModal = true)}>Create Group</Button>
</div>
<Modal title="Add Group" bind:open={defaultModal} class="min-w-full">
	{#if $createGroupResponse}
		<p class="text-base color-red-600">{$createGroupResponse}</p>
	{/if}
	<form on:submit={createGroup}>
		<div class="mb-4">
			<Label for="name" class="mb-2">Group Name</Label>
			<Input type="text" name="name" id="name" placeholder="Group Name..." required />
		</div>
		<div class="w-full flex mb-4">
			<div class="w-1/2">
				<Checkbox id="enabled" name="enabled">Users in group enabled</Checkbox>
				<Tooltip>Roles set per user take priority over group roles</Tooltip>
			</div>
			<div class="w-1/2">
				<Checkbox id="is_admin" name="is_admin">Users in group are admins</Checkbox>
				<Tooltip>Roles set per user take priority over group roles</Tooltip>
			</div>
		</div>
		<Button class="w-full" type="submit">Create Group</Button>
	</form>
</Modal>
