<script lang="ts">
	import { writable } from 'svelte/store';
	import type { User } from '../../../types/users';
	import { Checkbox, TableBodyCell, TableBodyRow } from 'flowbite-svelte';
	import { groupMembershipMasterCheckbox, selectedGroupMemberships } from '$lib/groups';
	import { onDestroy } from 'svelte';
	export let user: User;

	const checkbox = writable<boolean>();
	const e = groupMembershipMasterCheckbox.subscribe((e) => {
		checkbox.set(e);
		handleChange(e);
	});

	function handleChange(checked: boolean) {
		checked
			? selectedGroupMemberships.update((prev) => {
					return [...prev, user.id.toString()];
			  })
			: selectedGroupMemberships.update((prev) => {
					return prev.filter((i) => i !== user.id.toString());
			  });
	}

	onDestroy(e);
</script>

<TableBodyRow>
	<TableBodyCell tdClass="px-4 py-3"
		><Checkbox bind:checked={$checkbox} on:change={() => handleChange($checkbox)} /></TableBodyCell
	>
	<TableBodyCell tdClass="px-4 py-3">{user.username}</TableBodyCell>
	<TableBodyCell tdClass="px-4 py-3">{user.name}</TableBodyCell>
	<TableBodyCell tdClass="px-4 py-3">today</TableBodyCell>
	<TableBodyCell tdClass="px-4 py-3">{user.enabled}</TableBodyCell>
	<TableBodyCell tdClass="px-4 py-3">{user.is_admin}</TableBodyCell>
</TableBodyRow>
