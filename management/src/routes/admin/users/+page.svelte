<script lang="ts">
	import { page } from '$app/stores';
	import Sidepanel from "../sidepanel.svelte"
	import Table from "./table.svelte"

	page.subscribe((e)=> console.log(e.data))
</script>
<Sidepanel />

<div class="ml-5">
	<Table />
</div>
