<script lang="ts">
	import { onDestroy } from 'svelte';
	import {checkboxMaster} from "../../stores/stores";

	export let username: string;
	export let name:string;
	export let IsAdmin: boolean;
	export let Enabled: boolean;

	let checkbox: boolean;

	const s = checkboxMaster.subscribe((checked) => {
		checkbox = checked;
	});

	onDestroy(s);
</script>

<tr class="border-b dark:border-gray-700">
	<th scope="col" class="px-4 py-3">
		<input
			id="slave-checkbox"
			type="checkbox"
			bind:checked={checkbox}
			class="w-4 h-4 text-blue-600 bg-gray-100 border-gray-300 rounded focus:ring-blue-500 dark:focus:ring-blue-600 dark:ring-offset-gray-800 focus:ring-2 dark:bg-gray-700 dark:border-gray-600"
		/>
	</th>
	<td class="px-4 py-3">{name}</td>
	<td class="px-4 py-3">{username}</td>
	<td class="px-4 py-3">1/09/2023</td>
	<td class="px-4 py-3">
		<label class="relative inline-flex items-center mb-4 cursor-pointer">
			<input type="checkbox" value="" class="sr-only peer" bind:checked={Enabled} />
			<div
					class="w-11 h-6 bg-gray-200 rounded-full peer peer-focus:ring-4 peer-focus:ring-blue-300 dark:peer-focus:ring-blue-800 dark:bg-gray-700 peer-checked:after:translate-x-full peer-checked:after:border-white after:content-[''] after:absolute after:top-0.5 after:left-[2px] after:bg-white after:border-gray-300 after:border after:rounded-full after:h-5 after:w-5 after:transition-all dark:border-gray-600 peer-checked:bg-blue-600"
			/>
		</label></td
	>
	<td class="px-4 py-3">
		<label class="relative inline-flex items-center mb-4 cursor-pointer">
			<input type="checkbox" value="" class="sr-only peer" bind:checked={IsAdmin} />
			<div
				class="w-11 h-6 bg-gray-200 rounded-full peer peer-focus:ring-4 peer-focus:ring-blue-300 dark:peer-focus:ring-blue-800 dark:bg-gray-700 peer-checked:after:translate-x-full peer-checked:after:border-white after:content-[''] after:absolute after:top-0.5 after:left-[2px] after:bg-white after:border-gray-300 after:border after:rounded-full after:h-5 after:w-5 after:transition-all dark:border-gray-600 peer-checked:bg-blue-600"
			/>
		</label></td
	>
</tr>
