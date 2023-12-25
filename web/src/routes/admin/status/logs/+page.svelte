<script lang="ts">
	import { getVpnLogs, splitOnFirstLetter, vpnLog } from '$lib/status';
	import { onMount } from 'svelte';
	import Sidepanel from '../../sidepanel.svelte';
	import {
		Badge,
		Input,
		Label,
		P,
		Select,
		Table,
		TableBody,
		TableBodyCell,
		TableBodyRow,
		TableHead,
		TableHeadCell,
		Textarea
	} from 'flowbite-svelte';
	import Layout from '../../../+layout.svelte';
	import type { PageData } from './$types';
	let selected;
	export let data: PageData;
	let countries = [
		{ value: 'us', name: 'United States' },
		{ value: 'ca', name: 'Canada' },
		{ value: 'fr', name: 'France' }
	];
	onMount(() => {
		getVpnLogs();
	});
</script>

<div class="flex">
	<Sidepanel username={data.username} />
	<div class="w-full p-5">
		<h2 class="text-2xl mt-5 mb-1">VPN Server Logs</h2>
		<Badge color="yellow" class="mb-2 w-full p-3">
			<svg
				class="ml-2 mr-2 w-6 h-6 text-gray-800 dark:text-white"
				aria-hidden="true"
				xmlns="http://www.w3.org/2000/svg"
				fill="none"
				viewBox="0 0 20 20"
			>
				<path
					stroke="currentColor"
					stroke-linecap="round"
					stroke-linejoin="round"
					stroke-width="2"
					d="M10 11V6m0 8h.01M19 10a9 9 0 1 1-18 0 9 9 0 0 1 18 0Z"
				/>
			</svg><P
				>These are the unedited logs from OpenVPN. These should be used to dianose connection issues</P
			></Badge
		>
		<div class="flex my-2">
			<Input class="w-1/2 mr-1" placeholder="Filter By Event Text..." />
			<Select
				class="w-1/2 ml-1"
				placeholder="Filter By Time Period"
				items={countries}
				bind:value={selected}
			/>
		</div>
		<Table>
			<TableHead>
				<TableHeadCell>Event Time Stamp</TableHeadCell>
				<TableHeadCell>Event Text</TableHeadCell>
			</TableHead>
			<TableBody>
				{#each $vpnLog as log}
					<TableBodyRow>
						<TableBodyCell>{log[0]}</TableBodyCell>
						<TableBodyCell tdClass="overflow-y-hidden ">{log[1]}</TableBodyCell>
					</TableBodyRow>
				{/each}
			</TableBody>
		</Table>
	</div>
</div>
