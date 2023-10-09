<script lang="ts">
	import { Checkbox, TableBodyCell, TableBodyRow } from 'flowbite-svelte';
	import type { User } from '../../../types/users.js';
	import { masterCheckbox, selectedUsers } from '$lib/users.js';
	import { onDestroy } from 'svelte';
	import { writable } from 'svelte/store';

	export let user: User;
	const checked = writable<boolean>();
	const s = masterCheckbox.subscribe((e) => {
		checked.set(e);
	});
	const x = checked.subscribe((e) => addSelectedUser(e, user));

	function addSelectedUser(checked: boolean, user: User) {
		selectedUsers.update((prev) => {
			if (checked) {
				return [user, ...prev];
			}
			return prev.filter((u) => u.ID !== user.ID);
		});
	}

	onDestroy(s);
	onDestroy(x);
</script>

<tr class="border-b dark:border-gray-700">
	<th scope="col" class="px-4 py-3">
		<input
			id="default-checkbox"
			type="checkbox"
			bind:checked={$checked}
			class="w-4 h-4 text-blue-600 bg-gray-100 border-gray-300 rounded focus:ring-blue-500 dark:focus:ring-blue-600 dark:ring-offset-gray-800 focus:ring-2 dark:bg-gray-700 dark:border-gray-600"
		/></th
	>
	<th scope="row" class="px-4 py-3 font-medium text-gray-900 whitespace-nowrap dark:text-white">
		{user.Username}</th
	>
	<td class="px-4 py-3">{user.Name}</td>
	<td class="px-4 py-3">Today</td>
	<td class="px-4 py-3">{user.Enabled}</td>
	<td class="px-4 py-3">{user.IsAdmin}</td>
</tr>
