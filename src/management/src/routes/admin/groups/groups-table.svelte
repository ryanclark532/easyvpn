<script lang="ts">
	import { deleteGroup, groupsFilter, selectedGroups } from '$lib/groups';
	import { Button, Dropdown, DropdownItem, Input, Table } from 'flowbite-svelte';
	import type { PageData } from './$types';
	import AddGroupModal from './add-group-modal.svelte';
	import GroupsTableRow from './groups-table-row.svelte';
	import ConfirmationModal from '../../../components/confirmation-modal.svelte';
	import UpdateGroupModal from './update-group-modal.svelte';
	export let data: PageData;

	let deleteConfirmation: boolean;
</script>

<section class="bg-gray-50 dark:bg-gray-900">
	<div class="bg-white dark:bg-gray-800 relative shadow-md sm:rounded-lg overflow-hidden">
		<div
			class="flex flex-col md:flex-row items-center justify-between space-y-3 md:space-y-0 md:space-x-4 p-4"
		>
			<div class="w-full md:w-1/2">
				<form class="flex items-center">
					<label for="simple-search" class="sr-only">Search</label>
					<div class="relative w-full">
						<div class="absolute inset-y-0 left-0 flex items-center pl-3 pointer-events-none">
							<svg
								aria-hidden="true"
								class="w-5 h-5 text-gray-500 dark:text-gray-400"
								fill="currentColor"
								xmlns="http://www.w3.org/2000/svg"
							>
								<path
									fill-rule="evenodd"
									d="M8 4a4 4 0 100 8 4 4 0 000-8zM2 8a6 6 0 1110.89 3.476l4.817 4.817a1 1 0 01-1.414 1.414l-4.816-4.816A6 6 0 012 8z"
									clip-rule="evenodd"
								/>
							</svg>
						</div>
						<input
							type="text"
							id="simple-search"
							class="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-primary-500 focus:border-primary-500 block w-full pl-10 p-2 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-primary-500 dark:focus:border-primary-500"
							placeholder="Search"
							bind:value={$groupsFilter}
						/>
					</div>
				</form>
			</div>
			<div
				class="w-full md:w-auto flex flex-col md:flex-row space-y-2 md:space-y-0 items-stretch md:items-center justify-end md:space-x-3 flex-shrink-0"
			>
				{#if $selectedGroups.length === 1}
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
				{/if}
				<AddGroupModal />
			</div>
		</div>

		<div class="overflow-x-auto">
			<table class="w-full text-sm text-left text-gray-500 dark:text-gray-400">
				<thead
					class="text-xs text-gray-700 uppercase bg-gray-50 dark:bg-gray-700 dark:text-gray-400"
				>
					<tr>
						<th scope="col" class="px-4 py-3" />
						<th scope="col" class="px-4 py-3">Name</th>
						<th scope="col" class="px-4 py-3">Number Of Members</th>
						<th scope="col" class="px-4 py-3">Is Admin</th>
						<th scope="col" class="px-4 py-3">Is Enabled</th>
						<th scope="col" class="px-4 py-3">Add Users To Group</th>
					</tr>
				</thead>
				<tbody>
					{#each data.groups as group}
						{#if !$groupsFilter || group.name.startsWith($groupsFilter)}
							<GroupsTableRow {group} {data} />
						{/if}
					{/each}
				</tbody>
			</table>
		</div>
	</div>
</section>

<ConfirmationModal
	title="Confirm Group Deletion"
	subtext="Deleting this group is permanent and cannot be recovered. Please confirm the deletion."
	onConfirm={() => deleteGroup($selectedGroups[0].id)}
	open={deleteConfirmation}
/>
