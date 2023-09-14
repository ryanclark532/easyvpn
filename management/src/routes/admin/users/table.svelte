<script lang="ts">
	import {
		TableBody,
		TableBodyCell,
		TableBodyRow,
		TableHead,
		TableHeadCell,
		Table,
		Checkbox
	} from 'flowbite-svelte';
	import { users } from '$lib/users';
</script>

<Table divClass="w-full">
	<TableHead>
		<TableHeadCell padding="px-4 py-3" scope="col"><Checkbox /></TableHeadCell>
		<TableHeadCell padding="px-4 py-3" scope="col">Name</TableHeadCell>
		<TableHeadCell padding="px-4 py-3" scope="col">Username</TableHeadCell>
		<TableHeadCell padding="px-4 py-3" scope="col">Last Logged In</TableHeadCell>
		<TableHeadCell padding="px-4 py-3" scope="col">Is Admin</TableHeadCell>
		<TableHeadCell padding="px-4 py-3" scope="col">Enabled</TableHeadCell>
	</TableHead>
	<TableBody>
		{#each $users as user}
			<TableBodyRow>
				<TableBodyCell tdClass="px-4 py-3"><Checkbox /></TableBodyCell>
				<TableBodyCell tdClass="px-4 py-3">{user.Name}</TableBodyCell>
				<TableBodyCell tdClass="px-4 py-3">{user.Username}</TableBodyCell>
				<TableBodyCell tdClass="px-4 py-3">Yesterday</TableBodyCell>
				<TableBodyCell tdClass="px-4 py-3">{user.IsAdmin}</TableBodyCell>
				<TableBodyCell tdClass="px-4 py-3">{user.Enabled}</TableBodyCell>
			</TableBodyRow>
		{/each}
	</TableBody>
</Table>
