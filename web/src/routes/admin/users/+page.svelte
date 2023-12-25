<script lang="ts">
	import { Heading } from 'flowbite-svelte';
	import Sidepanel from '../sidepanel.svelte';
	import type { PageData } from './$types';
	import Table from './table.svelte';
	export let data: PageData;
</script>

<div class="flex">
	<Sidepanel username={data.username} />
	<div class="w-full p-5">
		<div class="border-b-2 border-primary-700 mb-3">
			<Heading tag="h2" class=" pb-1">Users</Heading>
		</div>
		<Table {data} />
	</div>
</div>
