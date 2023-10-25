<script lang="ts">
	import { Label, Input, Button, Modal, Checkbox, Tooltip } from 'flowbite-svelte';
	import { createGroup } from '$lib/groups';
	let defaultModal = false;
</script>

<div>
	<Button on:click={() => (defaultModal = true)}>Create Group</Button>
</div>
<Modal title="Add Group" bind:open={defaultModal} autoclose class="min-w-full">
	<form on:submit={createGroup}>
		<div>
			<div class="mb-4">
				<Label for="name" class="mb-2">Group Name</Label>
				<Input type="text" id="name" placeholder="Group Name..." required />
			</div>
			<div class="w-full flex mb-4">
				<div class="w-1/2">
					<Checkbox id="enabled">Users in group enabled</Checkbox>
					<Tooltip>Roles set per user take priority over group roles</Tooltip>
				</div>
				<div class="w-1/2">
					<Checkbox id="is_admin">Users in group are admins</Checkbox>
					<Tooltip>Roles set per user take priority over group roles</Tooltip>
				</div>
			</div>
			<Button class="w-full" type="submit">Create Group</Button>
		</div>
	</form>
</Modal>
