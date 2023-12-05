<script lang="ts">
	import { deleteGroup, groupsFilter, selectedGroups } from '$lib/groups';
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
	import type { PageData } from './$types';
	import AddGroupModal from './add-group-modal.svelte';
	import GroupsTableRow from './groups-table-row.svelte';
	import UpdateGroupModal from './update-group-modal.svelte';
	import ConfirmationModal from '../../../../components/confirmation-modal.svelte';
	export let data: PageData;

	let deleteConfirmation: boolean;
</script>

<div class="rounded relative shadow-md overflow-hidden">
	<div class="flex p-4">
		<div class="w-1/2">
			<Input bind:value={$groupsFilter} placeholder="Search Groups..." />
		</div>
		<div class="w-1/2 flex justify-end">
			{#if $selectedGroups.length === 1}
				<div class="mr-2">
					<Button
						>Actions &nbsp;
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
					</Button>
					<Dropdown>
						<UpdateGroupModal group={$selectedGroups[0]} />
						<DropdownItem on:click={() => (deleteConfirmation = !deleteConfirmation)}
							>Delete Group</DropdownItem
						>
					</Dropdown>
				</div>
			{/if}
			<AddGroupModal />
		</div>
	</div>
	<Table>
		<TableHead class="bg-gray-200">
			<TableHeadCell scope="col" class="px-4 py-3" />
			<TableHeadCell scope="col" class="px-4 py-3">Name</TableHeadCell>
			<TableHeadCell scope="col" class="px-4 py-3">Number Of Members</TableHeadCell>
			<TableHeadCell scope="col" class="px-4 py-3">Is Admin</TableHeadCell>
			<TableHeadCell scope="col" class="px-4 py-3">Is Enabled</TableHeadCell>
			<TableHeadCell scope="col" class="px-4 py-3">Add Users To Group</TableHeadCell>
		</TableHead>
		<TableBody class="divide-y">
			{#each data.groups as group}
				{#if !$groupsFilter || group.name.startsWith($groupsFilter)}
					<GroupsTableRow {group} users={data.users} />
				{/if}
			{/each}
		</TableBody>
	</Table>
</div>

<ConfirmationModal
	title="Confirm Group Deletion"
	subtext="Deleting this group is permanent and cannot be recovered. Please confirm the deletion."
	onConfirm={() => deleteGroup($selectedGroups[0].id)}
	open={deleteConfirmation}
/>
