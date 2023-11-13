<script lang="ts">
	import Sidepanel from '../sidepanel.svelte';
	import type { PageData } from './$types';
	import Table from './table.svelte';
	export let data: PageData;
</script>

<div class="flex">
	<Sidepanel />
	<div class="w-full p-5 ml-64">
		<div class="border-b-2 border-primary-700 mb-1">
			<h2 class="text-2xl mt-5 pb-1">Users</h2>
		</div>
		<Table {data} />
	</div>
</div>
