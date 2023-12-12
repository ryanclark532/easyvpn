<script lang="ts">
	import Sidepanel from '../sidepanel.svelte';
	import Status from './status.svelte';
	import type { PageData } from './$types';
	import ActiveUserTable from './active-user-table.svelte';
	import {
		Table,
		TableHead,
		TableHeadCell,
		TableBody,
		TableBodyRow,
		TableBodyCell,
		Heading
	} from 'flowbite-svelte';
	export let data: PageData;
</script>

<div class="flex">
	<Sidepanel />
	<div class="w-full p-5">
		<Status {data} />

		<div class="mt-4">
			<Heading tag="h4">Active Configuration</Heading>
			<Table striped={true}>
				<TableBody class="divide-y">
					<TableBodyRow>
						<TableBodyCell>Server Name</TableBodyCell>
						<TableBodyCell>Sliver</TableBodyCell>
					</TableBodyRow>
					<TableBodyRow>
						<TableBodyCell>Server IP Address</TableBodyCell>
						<TableBodyCell>White</TableBodyCell>
					</TableBodyRow>
					<TableBodyRow>
						<TableBodyCell>Maximum Allowed Connections</TableBodyCell>
						<TableBodyCell>Black</TableBodyCell>
					</TableBodyRow>
					<TableBodyRow>
						<TableBodyCell>Current Active Connections</TableBodyCell>
						<TableBodyCell>Gray</TableBodyCell>
					</TableBodyRow>
					<TableBodyRow>
						<TableBodyCell>VPN Server Port</TableBodyCell>
						<TableBodyCell>Red</TableBodyCell>
					</TableBodyRow>
					<TableBodyRow>
						<TableBodyCell>Web Server Port</TableBodyCell>
						<TableBodyCell>Red</TableBodyCell>
					</TableBodyRow>
					<TableBodyRow>
						<TableBodyCell>Using Server As A Gateway</TableBodyCell>
						<TableBodyCell>Red</TableBodyCell>
					</TableBodyRow>
					<TableBodyRow>
						<TableBodyCell>Allow Clients Access To Private Network</TableBodyCell>
						<TableBodyCell>Red</TableBodyCell>
					</TableBodyRow>
				</TableBody>
			</Table>
		</div>
	</div>
</div>
