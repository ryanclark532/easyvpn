<script lang="ts">
	import Sidepanel from '../sidepanel.svelte';
	import type { PageData } from './$types';
	import GroupsTable from './groups-table.svelte';
	export let data: PageData;


</script>

<div class="flex">
	<Sidepanel />
	<div class="w-full p-5 ml-64">
		<h2 class="text-2xl mt-5 mb-1">User Groups</h2>
		<GroupsTable {data} />
	</div>
</div>
