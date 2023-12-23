<script lang="ts">
	import type { PageData } from './$types';
	import { Input, Table, TableBody, TableHead, TableHeadCell } from 'flowbite-svelte';
	import ActiveUserTableRow from './active-user-table-row.svelte';
	import { activeConnectionsFilter } from '$lib/vpn';
	export let data: PageData;
</script>

<div class="rounded relative shadow-md overflow-hidden">
	<div class="flex p-4">
		<div class="w-1/2">
			<Input bind:value={$activeConnectionsFilter} placeholder="Search Active Connections..." />
		</div>
	</div>
	<Table>
		<TableHead class="bg-gray-200">
			<TableHeadCell scope="col" class="px-4 py-3" />
			<TableHeadCell scope="col" class="px-4 py-3">Username</TableHeadCell>
			<TableHeadCell scope="col" class="px-4 py-3">Address</TableHeadCell>
			<TableHeadCell scope="col" class="px-4 py-3">Bytes Sent</TableHeadCell>
			<TableHeadCell scope="col" class="px-4 py-3">Bytes Received</TableHeadCell>
			<TableHeadCell scope="col" class="px-4 py-3">Connected Since</TableHeadCell>
			<TableHeadCell scope="col" class="px-4 py-3">Disconnect</TableHeadCell>
		</TableHead>
		<TableBody>
			{#each data.connections as connection}
				{#if !$activeConnectionsFilter || connection.CommonName.startsWith($activeConnectionsFilter)}
					<ActiveUserTableRow {connection} />
				{/if}
			{/each}
		</TableBody>
	</Table>
</div>
