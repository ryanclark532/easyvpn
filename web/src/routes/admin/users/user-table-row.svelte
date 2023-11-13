<script lang="ts">
	import type { User } from '../../../types/users.js';
	import { selectedUsers } from '$lib/users';
	import { Checkbox, TableBodyCell, TableBodyRow } from 'flowbite-svelte';
	export let user: User;

	function toggleSelectedUsers() {
		selectedUsers.update((prev) => {
			const index = prev.findIndex((i) => i.id === user.id);
			if (index === -1) {
				prev.push(user);
			} else {
				prev.splice(index, 1);
			}
			return prev;
		});
	}
</script>

<TableBodyRow class="border-b dark:border-gray-700">
	<TableBodyCell scope="col" class="px-4 py-3">
		<TableBodyCell class="px-4 py-3"><Checkbox on:change={toggleSelectedUsers} /></TableBodyCell>
	</TableBodyCell>
	<TableBodyCell
		scope="row"
		class="px-4 py-3 font-medium text-gray-900 whitespace-nowrap dark:text-white"
	>
		{user.username}</TableBodyCell
	>
	<TableBodyCell class="px-4 py-3">{user.name}</TableBodyCell>
	<TableBodyCell class="px-4 py-3">Today</TableBodyCell>
	<TableBodyCell class="px-4 py-3">{user.enabled}</TableBodyCell>
	<TableBodyCell class="px-4 py-3">{user.is_admin}</TableBodyCell>
</TableBodyRow>
