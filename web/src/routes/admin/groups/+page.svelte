<script lang="ts">
	import Sidepanel from '../sidepanel.svelte';
	import type { PageData } from './$types';
	import GroupsTable from './groups-table.svelte';
	import { Heading } from 'flowbite-svelte';
	export let data: PageData;
</script>

<div class="flex">
	<Sidepanel />
	<div class="w-full p-5 ml-64">
		<div class="border-b-2 border-primary-700 mb-1">
			<Heading tag="h2" class="text-2xl mt-5 pb-1">User Groups</Heading>
		</div>
		<GroupsTable {data} />
	</div>
</div>
