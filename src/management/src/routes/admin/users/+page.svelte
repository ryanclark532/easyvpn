<script lang="ts">
	import Sidepanel from '../sidepanel.svelte';
	import CreateUserModal from './create-user-modal.svelte';
	import Table from './table.svelte';
	import { Button, Dropdown, DropdownItem, Heading, Input } from 'flowbite-svelte';
	import { masterCheckbox, searchFilter, selectedUsers } from '$lib/users';
	import ConfirmationModal from '../../../components/confirmation-modal.svelte';
	import { _userStore } from '../../+layout';

	let deleteModal = false;
	let tempPasswordModal = false;
	let toggleAdmin = false;
	let toggleEnabled = false;
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
							deleteModal = !deleteModal;
						}}>Delete User</DropdownItem
					>
					{#if $selectedUsers.length === 1}
						<DropdownItem>Update User</DropdownItem>
					{:else}
						<DropdownItem
							on:click={() => {
								toggleAdmin = !toggleAdmin;
							}}>Toggle Admin</DropdownItem
						>
						<DropdownItem
							on:click={() => {
								toggleEnabled = !toggleEnabled;
							}}>Toggle Enabled</DropdownItem
						>
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
	title="Confirm User Updates"
	onConfirm={_userStore.delete}
	data={$selectedUsers}
/>

<ConfirmationModal
	open={tempPasswordModal}
	title="Confirm User Updates"
	onConfirm={_userStore.setTempPw}
	data={$selectedUsers}
/>

<ConfirmationModal
	open={toggleAdmin}
	title="Confirm User Updates"
	onConfirm={_userStore.updateUser}
	data={$selectedUsers.map((u) => {
		return {
			...u,
			admin: !u.IsAdmin
		};
	})}
/>

<ConfirmationModal
	on:actiontriggered={() => {
		selectedUsers.set([]);
		masterCheckbox.set(false);
	}}
	open={toggleEnabled}
	title="Confirm User Updates"
	onConfirm={_userStore.updateUser}
	data={$selectedUsers.map((u) => {
		return {
			...u,
			enabled: !u.Enabled
		};
	})}
/>
