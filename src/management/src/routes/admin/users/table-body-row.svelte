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

<TableBodyRow>
	<TableBodyCell tdClass="px-4 py-3"
		><Checkbox
			bind:checked={$checked}
		/></TableBodyCell
	>
	<TableBodyCell tdClass="px-4 py-3">{user.Name}</TableBodyCell>
	<TableBodyCell tdClass="px-4 py-3">{user.Username}</TableBodyCell>
	<TableBodyCell tdClass="px-4 py-3">Yesterday</TableBodyCell>
	<TableBodyCell tdClass="px-4 py-3">{user.IsAdmin}</TableBodyCell>
	<TableBodyCell tdClass="px-4 py-3">{user.Enabled}</TableBodyCell>
</TableBodyRow>
