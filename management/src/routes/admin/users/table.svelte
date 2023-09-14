<script lang="ts">
	import {
		TableBody,
		TableBodyCell,
		TableBodyRow,
		TableHead,
		TableHeadCell,
		Table
	} from 'flowbite-svelte';
</script>

<Table style="width: calc(100% - 256px); margin-left: 256px;">
	<TableHead>
		<TableHeadCell padding="px-4 py-3" scope="col">Product name</TableHeadCell>
		<TableHeadCell padding="px-4 py-3" scope="col">Brand</TableHeadCell>
		<TableHeadCell padding="px-4 py-3" scope="col">Category</TableHeadCell>
		<TableHeadCell padding="px-4 py-3" scope="col">Price</TableHeadCell>
	</TableHead>
	<TableBody>
		<TableBodyRow>
			<TableBodyCell tdClass="px-4 py-3">test</TableBodyCell>
			<TableBodyCell tdClass="px-4 py-3">test</TableBodyCell>
			<TableBodyCell tdClass="px-4 py-3">test</TableBodyCell>
			<TableBodyCell tdClass="px-4 py-3">test</TableBodyCell>
		</TableBodyRow>
	</TableBody>
</Table>
