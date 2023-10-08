<script lang="ts">
	import Sidepanel from '../sidepanel.svelte';
	import Status from './status.svelte';
	import ActiveUsersTable from './active-users-table.svelte';
</script>

<div class="flex">
	<Sidepanel />
	<div class="w-full p-5 ml-64">
		<Status />
		<h2 class="text-2xl mt-5 mb-1">Active Server Connections</h2>
		<ActiveUsersTable />
	</div>
</div>
