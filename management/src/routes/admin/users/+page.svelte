<script lang="ts">
	import Sidepanel from '../sidepanel.svelte';
	import CreateUserModal from './create-user-modal.svelte';
	import Table from './table.svelte';
	import { Button, Dropdown, DropdownItem, Heading, Input, Modal } from 'flowbite-svelte';
	import { deleteUsers, searchFilter, selectedUsers, setTemporaryPassword } from '$lib/users';
	import ConfirmationModal from '../../../components/confirmation-modal.svelte';
	let deleteModal = false;
	let tempPasswordModal = false;
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
					<DropdownItem
						on:click={() => {
							tempPasswordModal = true;
						}}>Set Temporary Password</DropdownItem
					>
				</Dropdown>
			{/if}
			<CreateUserModal />
		</div>
		<Table />
	</div>
</div>

<ConfirmationModal
	open={deleteModal}
	title="Confirm User Deletion"
	onConfirm={deleteUsers}
	data={$selectedUsers}
/>

<ConfirmationModal
	open={tempPasswordModal}
	title="Confirm User Updates"
	onConfirm={setTemporaryPassword}
	data={$selectedUsers}
/>
