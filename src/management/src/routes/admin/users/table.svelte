<script lang="ts">
	import {
		TableBody,
		TableHead,
		TableHeadCell,
		Table,
		Checkbox
	} from 'flowbite-svelte';
	import {masterCheckbox, searchFilter, users} from '$lib/users';
	import TableBodyRow from './table-body-row.svelte';
	import {page} from "$app/stores";

	let userStore = $page.data.userStore

</script>

<Table divClass="w-full">
	<TableHead>
		<TableHeadCell padding="px-4 py-3" scope="col"><Checkbox bind:checked={$masterCheckbox} /></TableHeadCell>
		<TableHeadCell padding="px-4 py-3" scope="col">Name</TableHeadCell>
		<TableHeadCell padding="px-4 py-3" scope="col">Username</TableHeadCell>
		<TableHeadCell padding="px-4 py-3" scope="col">Last Logged In</TableHeadCell>
		<TableHeadCell padding="px-4 py-3" scope="col">Is Admin</TableHeadCell>
		<TableHeadCell padding="px-4 py-3" scope="col">Enabled</TableHeadCell>
	</TableHead>
	<TableBody>
		{#each $userStore as user}
			{#if !$searchFilter || user.Username.startsWith($searchFilter)}
				<TableBodyRow {user} />
			{/if}
		{/each}
	</TableBody>
</Table>
