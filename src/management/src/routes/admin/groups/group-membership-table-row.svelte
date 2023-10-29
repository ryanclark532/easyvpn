<script lang="ts">
	import { writable } from 'svelte/store';
	import type { User } from '../../../types/users';
	import { Checkbox } from 'flowbite-svelte';
	import { groupMembershipMasterCheckbox, selectedGroupMemberships } from '$lib/groups';
	import { onDestroy } from 'svelte';
	export let user: User;

	const checkbox = writable<boolean>();
const e =	groupMembershipMasterCheckbox.subscribe((e) => {
		checkbox.set(e);
        handleChange(e);
	});


    function handleChange(checked: boolean){
        checked ? 
			selectedGroupMemberships.update((prev) => {
				return [...prev, user.id];
			}) :
			selectedGroupMemberships.update((prev) => {
				return prev.filter((i) => i !== user.id);
			});
    } 

    onDestroy(e);
</script>

<tr class="border-b last:border-b-0">
	<td class="px-4 py-3"><Checkbox bind:checked={$checkbox} on:change={handleChange($checkbox)} /></td>
	<td class="px-4 py-3">{user.username}</td>
	<td class="px-4 py-3">{user.name}</td>
	<td class="px-4 py-3">today</td>
	<td class="px-4 py-3">{user.enabled}</td>
	<td class="px-4 py-3">{user.is_admin}</td>
</tr>
