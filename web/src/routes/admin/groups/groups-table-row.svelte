<script lang="ts">
	import {
		Table,
		Modal,
		Button,
		Checkbox,
		TableBodyRow,
		TableBodyCell,
		TableHead,
		TableHeadCell,
		TableBody
	} from 'flowbite-svelte';
	import type { Group } from '../../../types/groups';
	import {
		getGroupMembers,
		createGroupMembership,
		groupMembershipMasterCheckbox,
		selectedGroupMemberships,
		deleteGroupMembership,
		selectedGroups
	} from '$lib/groups';
	import GroupMembershipTableRow from './group-membership-table-row.svelte';
	import type { User } from '../../../types/users';
	export let group: Group;
	export let users: User[];
	let open: boolean;
	let addingUser: boolean;
	let selectedUsers: string[] = [];

	function toggleMembershipSelected(id: string) {
		const index = selectedUsers.findIndex((i) => i === id);
		if (index === -1) {
			selectedUsers.push(id);
		} else {
			selectedUsers.splice(index, 1);
		}
	}

	function toggleSelectedGroups() {
		selectedGroups.update((prev) => {
			const index = prev.findIndex((i) => i.id === group.id);
			if (index === -1) {
				prev.push(group);
			} else {
				prev.splice(index, 1);
			}
			return prev;
		});
	}
</script>

<TableBodyRow class="border-b last:border-b-0">
	<TableBodyCell class="px-4 py-3"><Checkbox on:change={toggleSelectedGroups} /></TableBodyCell>
	<TableBodyCell class="px-4 py-3">{group.name}</TableBodyCell>
	<TableBodyCell class="px-4 py-3">{group.member_count}</TableBodyCell>
	<TableBodyCell class="px-4 py-3">{group.is_admin}</TableBodyCell>
	<TableBodyCell class="px-4 py-3">{group.enabled}</TableBodyCell>
	<TableBodyCell class="px-4 py-3"
		><Button color="alternative" on:click={() => (open = !open)}>Add Users</Button></TableBodyCell
	>
</TableBodyRow>

<Modal title={`${group.name} Users`} bind:open class="min-w-full">
	{#if !addingUser}
		<Table>
			<TableHead class="bg-gray-200">
				<TableHeadCell scope="col" class="px-4 py-3"
					><Checkbox bind:checked={$groupMembershipMasterCheckbox} /></TableHeadCell
				>
				<TableHeadCell scope="col" class="px-4 py-3">Username</TableHeadCell>
				<TableHeadCell scope="col" class="px-4 py-3">Name</TableHeadCell>
				<TableHeadCell scope="col" class="px-4 py-3">Last Logged In</TableHeadCell>
				<TableHeadCell scope="col" class="px-4 py-3">Enabled</TableHeadCell>
				<TableHeadCell scope="col" class="px-4 py-3">Is Admin</TableHeadCell>
			</TableHead>
			{#await getGroupMembers(group.id)}
				<TableBody>
					<TableBodyCell class="px-4 py-3" />
					<TableBodyCell class="px-4 py-3" />
					<TableBodyCell class="px-4 py-3" />
					<TableBodyCell class="px-4 py-3" />
					<TableBodyCell class="px-4 py-3" />
					<TableBodyCell class="px-4 py-3" />
				</TableBody>
			{:then users}
				{#each users as user}
					<GroupMembershipTableRow {user} />
				{/each}
			{:catch error}
				<TableBody>
					<TableBodyCell class="px-4 py-3" />
					<TableBodyCell class="px-4 py-3" />
					<TableBodyCell class="px-4 py-3" />
					<TableBodyCell class="px-4 py-3" />
					<TableBodyCell class="px-4 py-3" />
					<TableBodyCell class="px-4 py-3" />
				</TableBody>
			{/await}
		</Table>
		{#if $selectedGroupMemberships.length === 0}
			<Button class="w-full" on:click={() => (addingUser = !addingUser)}>Add Users To Group</Button>
		{:else}
			<div class="flex">
				<Button class="w-1/2 m-1" on:click={() => (addingUser = !addingUser)}
					>Add Users To Group</Button
				>
				<Button
					class="w-1/2 m-1"
					on:click={() => {
						deleteGroupMembership($selectedGroupMemberships, group.id);
						selectedGroupMemberships.set([]);
						open = false;
					}}>Delete Selected Users</Button
				>
			</div>
		{/if}
	{:else}
		<Table divClass="text-black">
			<TableHead class="bg-gray-200">
				<TableHeadCell scope="col" class="px-4 py-3"><Checkbox /></TableHeadCell>
				<TableHeadCell scope="col" class="px-4 py-3">Username</TableHeadCell>
				<TableHeadCell scope="col" class="px-4 py-3">Name</TableHeadCell>
				<TableHeadCell scope="col" class="px-4 py-3">Last Logged In</TableHeadCell>
				<TableHeadCell scope="col" class="px-4 py-3">Enabled</TableHeadCell>
				<TableHeadCell scope="col" class="px-4 py-3">Is Admin</TableHeadCell>
			</TableHead>
			{#each users as user}
				<TableBodyRow>
					<TableBodyCell class="px-4 py-3 text-black"
						><Checkbox
							on:change={() => {
								toggleMembershipSelected(user.id.toString());
							}}
						/></TableBodyCell
					>
					<TableBodyCell class="px-4 py-3">{user.username}</TableBodyCell>
					<TableBodyCell class="px-4 py-3">{user.name}</TableBodyCell>
					<TableBodyCell class="px-4 py-3">today</TableBodyCell>
					<TableBodyCell class="px-4 py-3">{user.enabled}</TableBodyCell>
					<TableBodyCell class="px-4 py-3">{user.is_admin}</TableBodyCell>
				</TableBodyRow>
			{/each}
		</Table>
		<div class="flex">
			<Button class="w-1/2 m-1" on:click={() => (addingUser = !addingUser)}
				>Back To Memberships</Button
			>
			<Button
				class="w-1/2 m-1"
				on:click={() => {
					createGroupMembership(selectedUsers, group.id);
					open = false;
				}}>Add Selected Users</Button
			>
		</div>
	{/if}
</Modal>
