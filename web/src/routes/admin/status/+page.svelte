<script lang="ts">
	import Sidepanel from '../sidepanel.svelte';
	import Status from './status.svelte';
	import type { PageData } from './$types';
	import  ActiveUserTable from "./active-user-table.svelte"
	export let data: PageData;
</script>

<div class="flex">
	<Sidepanel />
	<div class="w-full p-5 ">
		<Status {data} />
		<h2 class="text-2xl mt-5 mb-1">Active Server Connections</h2>
		<ActiveUserTable {data} />
	</div>
</div>
