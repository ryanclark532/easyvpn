<script lang="ts">
	import Sidepanel from '../sidepanel.svelte';
	import CreateUserModal from './create-user-modal.svelte';
	import Table from './table.svelte';
	import { Button, Dropdown, DropdownItem, Heading, Input, Modal } from 'flowbite-svelte';
	import { deleteUsers, searchFilter, selectedUsers } from '$lib/users';

	let deleteModal = false;
</script>

<div class="flex">
	<Sidepanel />
	<div class="w-full p-5">
		<Heading tag={'h2'}>Users</Heading>
		<div class="my-6 w-100 float-left">
			<Input
				id="default-input"
				placeholder="Search By Username..."
				style="height: 40px; width: 400px"
				bind:value={$searchFilter}
			/>
		</div>
		<div class="float-right mt-6">
			{#if $selectedUsers.length > 0}
				<Button class="w-40">Actions</Button>
				<Dropdown>
					<DropdownItem
						on:click={() => {
							deleteModal = true;
						}}>Delete User</DropdownItem
					>
					{#if $selectedUsers.length === 1}
						<DropdownItem>Update User</DropdownItem>
					{:else}
						<DropdownItem>Toggle Admin</DropdownItem>
						<DropdownItem>Toggle Enabled</DropdownItem>
					{/if}
					<DropdownItem>Set Temporary Password</DropdownItem>
				</Dropdown>
			{/if}
			<CreateUserModal />
		</div>
		<Table />
	</div>
</div>

<Modal title="" bind:open={deleteModal} autoclose size="sm" class="w-full">
	<svg
		class="text-gray-400 dark:text-gray-500 w-11 h-11 mb-3.5 mx-auto"
		aria-hidden="true"
		fill="currentColor"
		viewBox="0 0 20 20"
		xmlns="http://www.w3.org/2000/svg"
		><path
			fill-rule="evenodd"
			d="M9 2a1 1 0 00-.894.553L7.382 4H4a1 1 0 000 2v10a2 2 0 002 2h8a2 2 0 002-2V6a1 1 0 100-2h-3.382l-.724-1.447A1 1 0 0011 2H9zM7 8a1 1 0 012 0v6a1 1 0 11-2 0V8zm5-1a1 1 0 00-1 1v6a1 1 0 102 0V8a1 1 0 00-1-1z"
			clip-rule="evenodd"
		/></svg
	>
	<p class="mb-4 text-gray-500 dark:text-gray-300 text-center">
		Are you sure you want to delete selected users?
	</p>
	<div class="flex justify-center items-center space-x-4">
		<Button
			color="light"
			on:click={() => {
				deleteModal = false;
			}}>No, cancel</Button
		>
		<Button
			color="red"
			on:click={() => {
				deleteUsers($selectedUsers);
			}}>Yes, I'm sure</Button
		>
	</div>
</Modal>
