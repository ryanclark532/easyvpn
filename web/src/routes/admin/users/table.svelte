<script lang="ts">
	import { deleteUser, selectedUsers, userFilter } from '$lib/users';
	import type { PageData } from './$types';
	import UserTableRow from './user-table-row.svelte';
	import CreateUserModal from './create-user-modal.svelte';
	import {
		Button,
		Dropdown,
		DropdownItem,
		Input,
		Table,
		TableBody,
		TableHead,
		TableHeadCell
	} from 'flowbite-svelte';
	import UpdateUserModal from './update-user-modal.svelte';
	import SetTempPwModal from './set-temp-pw-modal.svelte';
	import ConfirmationModal from '../../../components/confirmation-modal.svelte';

	export let data: PageData;

	let deleteConfirmation: boolean = false;
</script>

<div class="rounded relative shadow-md overflow-hidden">
	<div class="flex p-4">
		<div class="w-1/2">
			<Input bind:value={$userFilter} placeholder="Search Groups..." />
		</div>
		<div class="w-1/2 flex justify-end">
			{#if $selectedUsers.length === 1}
				<div class="mr-2">
					<Button>
						<svg
							class="-ml-1 mr-1.5 w-5 h-5"
							fill="currentColor"
							xmlns="http://www.w3.org/2000/svg"
							aria-hidden="true"
						>
							<path
								clip-rule="evenodd"
								fill-rule="evenodd"
								d="M5.293 7.293a1 1 0 011.414 0L10 10.586l3.293-3.293a1 1 0 111.414 1.414l-4 4a1 1 0 01-1.414 0l-4-4a1 1 0 010-1.414z"
							/>
						</svg>
						Actions &nbsp;
					</Button>
					<Dropdown>
						<SetTempPwModal userId={$selectedUsers[0].id} />
						<UpdateUserModal user={$selectedUsers[0]} />
						<DropdownItem on:click={() => {deleteConfirmation = !deleteConfirmation}}>Delete User</DropdownItem>
					</Dropdown>
				</div>
			{/if}
			<CreateUserModal />
		</div>
	</div>
	<Table>
		<TableHead class="bg-gray-200">
			<TableHeadCell scope="col" class="px-4 py-3" />
			<TableHeadCell scope="col" class="px-4 py-3">Username</TableHeadCell>
			<TableHeadCell scope="col" class="px-4 py-3">Name</TableHeadCell>
			<TableHeadCell scope="col" class="px-4 py-3">Last Logged In</TableHeadCell>
			<TableHeadCell scope="col" class="px-4 py-3">Enabled</TableHeadCell>
			<TableHeadCell scope="col" class="px-4 py-3">Admin User</TableHeadCell>
		</TableHead>
		<TableBody>
			{#each data.users as user}
				{#if !$userFilter || user.username.startsWith($userFilter)}
					<UserTableRow {user} />
				{/if}
			{/each}
		</TableBody>
	</Table>
</div>


<ConfirmationModal
	title="Confirm User Deletion"
	subtext="Deleting this user is permanent and cannot be recovered. Please confirm the deletion."
	onConfirm={() => deleteUser($selectedUsers[0].id)}
	open={deleteConfirmation}
/>
