<script lang="ts">
	import { Table, Modal } from 'flowbite-svelte';
	import type { Group } from '../../../types/groups';
	import { getGroupMembers } from '$lib/groups';
	export let group: Group;
	let open: boolean;
</script>

<tr class="border-b last:border-b-0" on:click={() => (open = !open)}>
	<td class="px-4 py-3">{group.name}</td>
	<td class="px-4 py-3">{group.member_count}</td>
	<td class="px-4 py-3">{group.is_admin}</td>
	<td class="px-4 py-3">{group.enabled}</td>
</tr>

<Modal title={`${group.name} Users`} bind:open autoclose class="min-w-full">
	<Table>
		<thead class="text-xs text-gray-700 uppercase bg-gray-50 dark:bg-gray-700 dark:text-gray-400">
			<tr>
				<th scope="col" class="px-4 py-3">Username</th>
				<th scope="col" class="px-4 py-3">Name</th>
				<th scope="col" class="px-4 py-3">Last Logged In</th>
				<th scope="col" class="px-4 py-3">Enabled</th>
				<th scope="col" class="px-4 py-3">Is Admin</th>
			</tr>
		</thead>
		{#await getGroupMembers(group.id)}
			Loading....
		{:then users}
			{#each users as user}
				<tr class="border-b last:border-b-0">
					<td class="px-4 py-3">{user.username}</td>
					<td class="px-4 py-3">{user.name}</td>
					<td class="px-4 py-3">today</td>
					<td class="px-4 py-3">{user.enabled}</td>
					<td class="px-4 py-3">{user.is_admin}</td>
				</tr>
			{/each}
		{/await}
	</Table>
</Modal>
