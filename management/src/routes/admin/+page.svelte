<script lang="ts">
	import Sidepanel from './sidepanel.svelte';
	import { isAuthedAdmin } from '$lib/auth';
	import { onDestroy, onMount } from 'svelte';
	import { page } from '$app/stores';
	import Status from '../../components/status/status.svelte';
	import Settings from '../../components/settings/settings.svelte';
	import Certificates from '../../components/certificates/certificates.svelte';
	import Users from '../../components/users/users.svelte';

	let subpage = 'status';
	page.subscribe((e) => (subpage = e.url.searchParams.get('subpage') ?? 'status'));

	onMount(async () => {
		if ((await isAuthedAdmin()) === false) {
			window.location.href = '/login';
		}
	});
</script>

<Sidepanel />
<div class="admin-container bg-gray-50">
	{#if subpage === 'status'}
		<Status />
	{/if}
	{#if subpage === 'users'}
		<Users />
	{/if}
	{#if subpage === 'settings'}
		<Settings />
	{/if}

	{#if subpage === 'certificates'}
		<Certificates />
	{/if}
</div>

<style>
	.admin-container {
		margin-left: 256px;
	}
</style>
