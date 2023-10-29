<script lang="ts">
	import { Table, Modal, Button, Checkbox } from 'flowbite-svelte';
	import type { Group } from '../../../types/groups';
	import {
		getGroupMembers,
		createGroupMembership,
		groupMembershipMasterCheckbox,
selectedGroupMemberships,
deleteGroupMembership
	} from '$lib/groups';
	import type { PageData } from '../../$types';
	import GroupMembershipTableRow from './group-membership-table-row.svelte';
	export let group: Group;
	export let data: PageData;
	let open: boolean;
	let addingUser: boolean;
	let selectedUsers: string[] = [];

	function toggleSelected(id: string) {
		const index = selectedUsers.findIndex((i) => i === id);
		if (index === -1) {
			selectedUsers.push(id);
		} else {
			selectedUsers.splice(index, 1);
		}
	}
</script>

<tr class="border-b last:border-b-0" on:click={() => (open = !open)}>
	<td class="px-4 py-3">{group.name}</td>
	<td class="px-4 py-3">{group.member_count}</td>
	<td class="px-4 py-3">{group.is_admin}</td>
	<td class="px-4 py-3">{group.enabled}</td>
</tr>

<Modal title={`${group.name} Users`} bind:open class="min-w-full">
	{#if !addingUser}
		<Table>
			<thead class="text-xs text-gray-700 uppercase bg-gray-50 dark:bg-gray-700 dark:text-gray-400">
				<tr>
					<th scope="col" class="px-4 py-3"
						><Checkbox bind:checked={$groupMembershipMasterCheckbox} /></th
					>
					<th scope="col" class="px-4 py-3">Username</th>
					<th scope="col" class="px-4 py-3">Name</th>
					<th scope="col" class="px-4 py-3">Last Logged In</th>
					<th scope="col" class="px-4 py-3">Enabled</th>
					<th scope="col" class="px-4 py-3">Is Admin</th>
				</tr>
			</thead>
			{#await getGroupMembers(group.id)}
				<tr class="border-b last:border-b-0">
					<td class="px-4 py-3"><Checkbox /></td>
					<td class="px-4 py-3" />
					<td class="px-4 py-3" />
					<td class="px-4 py-3" />
					<td class="px-4 py-3" />
					<td class="px-4 py-3" />
				</tr>
			{:then users}
				{#each users as user}
					<GroupMembershipTableRow {user} />
				{/each}
			{:catch}
				<tr class="border-b last:border-b-0">
					<td class="px-4 py-3"><Checkbox /></td>
					<td class="px-4 py-3" />
					<td class="px-4 py-3" />
					<td class="px-4 py-3" />
					<td class="px-4 py-3" />
					<td class="px-4 py-3" />
				</tr>
			{/await}
		</Table>
		{#if $selectedGroupMemberships.length === 0}
			<Button class="w-full" on:click={() => (addingUser = !addingUser)}>Add Users To Group</Button>
		{:else}
			<div class="flex">
				<Button class="w-1/2 m-1" on:click={() => (addingUser = !addingUser)}
					>Add Users To Group</Button
				>
				<Button class="w-1/2 m-1" on:click={()=>deleteGroupMembership($selectedGroupMemberships, group.id)}>Delete Selected Users</Button>
			</div>
		{/if}
	{:else}
		<Table>
			<thead class="text-xs text-gray-700 uppercase bg-gray-50 dark:bg-gray-700 dark:text-gray-400">
				<tr>
					<th scope="col" class="px-4 py-3"> <Checkbox /></th>
					<th scope="col" class="px-4 py-3">Username</th>
					<th scope="col" class="px-4 py-3">Name</th>
					<th scope="col" class="px-4 py-3">Last Logged In</th>
					<th scope="col" class="px-4 py-3">Enabled</th>
					<th scope="col" class="px-4 py-3">Is Admin</th>
				</tr>
			</thead>
			{#each data.users as user}
				<tr>
					<td class="px-4 py-3"
						><Checkbox
							on:change={() => {
								toggleSelected(user.id);
							}}
						/></td
					>
					<td class="px-4 py-3">{user.username}</td>
					<td class="px-4 py-3">{user.name}</td>
					<td class="px-4 py-3">today</td>
					<td class="px-4 py-3">{user.enabled}</td>
					<td class="px-4 py-3">{user.is_admin}</td>
				</tr>
			{/each}
		</Table>
		<Button class="w-1/2 m-1" on:click={() => createGroupMembership(selectedUsers, group.id)}
			>Add Selected Users</Button
		>
	{/if}
</Modal>
