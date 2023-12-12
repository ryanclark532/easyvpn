<script lang="ts">
	import type { PageData } from '../$types';
	import Sidepanel from '../../sidepanel.svelte';
	import ActiveUserTable from "../active-user-table.svelte";
	export let data: PageData;

    console.log(data); 
</script>

<div class="flex">
	<Sidepanel />
	<div class="w-full p-5">
		<h2 class="text-2xl mt-5 mb-1">Active Server Connections</h2>
		<ActiveUserTable {data} />
	</div>
</div>