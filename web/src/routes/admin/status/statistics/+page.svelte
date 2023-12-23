<script lang="ts">
	import Status from './status.svelte';
	import type { PageData } from './$types';
	import { Table, TableBody, TableBodyRow, TableBodyCell, Heading } from 'flowbite-svelte';
	import Sidepanel from '../../sidepanel.svelte';
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
						<TableBodyCell>{data.hostname}</TableBodyCell>
					</TableBodyRow>
					<TableBodyRow>
						<TableBodyCell>Server IP Address</TableBodyCell>
						<TableBodyCell>{data.settings.ip_address}</TableBodyCell>
					</TableBodyRow>
					<TableBodyRow>
						<TableBodyCell>Maximum Allowed Connections</TableBodyCell>
						<TableBodyCell>{data.settings.max_connections}</TableBodyCell>
					</TableBodyRow>
					<TableBodyRow>
						<TableBodyCell>Current Active Connections</TableBodyCell>
						<TableBodyCell>{data.connections}</TableBodyCell>
					</TableBodyRow>
					<TableBodyRow>
						<TableBodyCell>VPN Server Port</TableBodyCell>
						<TableBodyCell>{data.settings.port}</TableBodyCell>
					</TableBodyRow>
					<TableBodyRow>
						<TableBodyCell>Web Server Port</TableBodyCell>
						<TableBodyCell>{data.settings.web_server_port}</TableBodyCell>
					</TableBodyRow>
					<TableBodyRow>
						<TableBodyCell>Using Server As A Gateway</TableBodyCell>
						<TableBodyCell>{data.settings.use_as_gateway ? 'Yes' : 'No'}</TableBodyCell>
					</TableBodyRow>
					<TableBodyRow>
						<TableBodyCell>Allow Clients Access To Private Network</TableBodyCell>
						<TableBodyCell>{data.settings.private_access ? 'Yes' : 'No'}</TableBodyCell>
					</TableBodyRow>
				</TableBody>
			</Table>
		</div>
	</div>
</div>
