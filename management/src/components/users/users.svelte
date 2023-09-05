<script lang="ts">
	import Tablehead from './tablehead.svelte';
	import Tablebody from './tablebody.svelte';
	import { onMount } from 'svelte';
	import { getUsers } from '$lib';
	onMount(() => {
		getUsers();
	});
</script>

<section class="w-full p-5">
	<div class="bg-white shadow-md sm:rounded-lg">
		<Tablehead />
		<Tablebody />
	</div>
</section>
