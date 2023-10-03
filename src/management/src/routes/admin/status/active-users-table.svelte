<script lang="ts">
	import {
		Table,
		TableHead,
		TableHeadCell,
		Checkbox,
		TableBody,
		TableBodyRow,
		TableBodyCell,
		Button
	} from 'flowbite-svelte';
</script>

<Table divClass="w-full p-5">
	<TableHead>
		<TableHeadCell padding="px-4 py-3" scope="col">Username</TableHeadCell>
		<TableHeadCell padding="px-4 py-3" scope="col">Session Length</TableHeadCell>
		<TableHeadCell padding="px-4 py-3" scope="col">Certificate</TableHeadCell>
		<TableHeadCell padding="px-4 py-3" scope="col">Disconnect</TableHeadCell>
	</TableHead>
	<TableBody>
		<TableBodyRow>
			<TableBodyCell tdClass="px-4 py-3">test</TableBodyCell>
			<TableBodyCell tdClass="px-4 py-3">test</TableBodyCell>
			<TableBodyCell tdClass="px-4 py-3">test</TableBodyCell>
			<TableBodyCell tdClass="px-4 py-3"><Button class="w-40">Disconnect</Button></TableBodyCell>
		</TableBodyRow>
	</TableBody>
</Table>
