<script lang="ts">
	import { Button, TableBodyCell, TableBodyRow } from 'flowbite-svelte';
	import type { Connection } from '../../../types/vpn.js';
	export let connection: Connection;
</script>

<TableBodyRow class="border-b dark:border-gray-700">
	<TableBodyCell scope="col" class="px-4 py-3" />
	<TableBodyCell
		scope="row"
		class="px-4 py-3 font-medium text-gray-900 whitespace-nowrap dark:text-white"
	>
		{connection.CommonName}</TableBodyCell
	>
	<TableBodyCell class="px-4 py-3">{connection.Address}</TableBodyCell>
	<TableBodyCell class="px-4 py-3">{connection.BytesSent}</TableBodyCell>
	<TableBodyCell class="px-4 py-3">{connection.BytesRec}</TableBodyCell>
	<TableBodyCell class="px-4 py-3">{connection.ConnectedSince}</TableBodyCell>
	<TableBodyCell class="px-4 py-3"><Button color="alternative">Disconnect</Button></TableBodyCell>
</TableBodyRow>
