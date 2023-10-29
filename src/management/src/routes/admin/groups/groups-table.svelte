<script lang="ts">
	import { groupsFilter } from '$lib/groups';
	import type { PageData } from './$types';
	import AddGroupModal from './add-group-modal.svelte';
	import GroupsTableRow from './groups-table-row.svelte';
	export let data: PageData;
</script>

<section class="bg-gray-50 dark:bg-gray-900">
	<div class="bg-white dark:bg-gray-800 relative shadow-md sm:rounded-lg overflow-hidden">
		<div
			class="flex flex-col md:flex-row items-center justify-between space-y-3 md:space-y-0 md:space-x-4 p-4"
		>
			<div class="w-full md:w-1/2">
				<form class="flex items-center">
					<label for="simple-search" class="sr-only">Search</label>
					<div class="relative w-full">
						<div class="absolute inset-y-0 left-0 flex items-center pl-3 pointer-events-none">
							<svg
								aria-hidden="true"
								class="w-5 h-5 text-gray-500 dark:text-gray-400"
								fill="currentColor"
								xmlns="http://www.w3.org/2000/svg"
							>
								<path
									fill-rule="evenodd"
									d="M8 4a4 4 0 100 8 4 4 0 000-8zM2 8a6 6 0 1110.89 3.476l4.817 4.817a1 1 0 01-1.414 1.414l-4.816-4.816A6 6 0 012 8z"
									clip-rule="evenodd"
								/>
							</svg>
						</div>
						<input
							type="text"
							id="simple-search"
							class="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-primary-500 focus:border-primary-500 block w-full pl-10 p-2 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-primary-500 dark:focus:border-primary-500"
							placeholder="Search"
							bind:value={$groupsFilter}
						/>
					</div>
				</form>
			</div>
			<div
				class="w-full md:w-auto flex flex-col md:flex-row space-y-2 md:space-y-0 items-stretch md:items-center justify-end md:space-x-3 flex-shrink-0"
			>
				<AddGroupModal />
			</div>
		</div>

		<div class="overflow-x-auto">
			<table class="w-full text-sm text-left text-gray-500 dark:text-gray-400">
				<thead
					class="text-xs text-gray-700 uppercase bg-gray-50 dark:bg-gray-700 dark:text-gray-400"
				>
					<tr>
						<th scope="col" class="px-4 py-3 w-1/6">Name</th>
						<th scope="col" class="px-4 py-3 w-1/6">Number Of Members</th>
						<th scope="col" class="px-4 py-3 w-1/6">Is Admin</th>
						<th scope="col" class="px-4 py-3 w-1/6">Is Enabled</th>
					</tr>
				</thead>
				<tbody>
					{#each data.groups as group}
						{#if !$groupsFilter || group.name.startsWith($groupsFilter)}
							<GroupsTableRow {group} {data} />
						{/if}
					{/each}
				</tbody>
			</table>
		</div>
		<nav
			class="flex flex-col md:flex-row justify-between items-start md:items-center space-y-3 md:space-y-0 p-4"
			aria-label="Table navigation"
		>
			<span class="text-sm font-normal text-gray-500 dark:text-gray-400">
				Showing
				<span class="font-semibold text-gray-900 dark:text-white">1-10</span>
				of
				<span class="font-semibold text-gray-900 dark:text-white">{data.groups.length}</span>
			</span>
			<ul class="inline-flex items-stretch -space-x-px">
				<li>
					<a
						href="#"
						class="flex items-center justify-center h-full py-1.5 px-3 ml-0 text-gray-500 bg-white rounded-l-lg border border-gray-300 hover:bg-gray-100 hover:text-gray-700 dark:bg-gray-800 dark:border-gray-700 dark:text-gray-400 dark:hover:bg-gray-700 dark:hover:text-white"
					>
						<span class="sr-only">Previous</span>
						<svg
							class="w-5 h-5"
							aria-hidden="true"
							fill="currentColor"
							xmlns="http://www.w3.org/2000/svg"
						>
							<path
								fill-rule="evenodd"
								d="M12.707 5.293a1 1 0 010 1.414L9.414 10l3.293 3.293a1 1 0 01-1.414 1.414l-4-4a1 1 0 010-1.414l4-4a1 1 0 011.414 0z"
								clip-rule="evenodd"
							/>
						</svg>
					</a>
				</li>
				<li>
					<a
						href="#top"
						class="flex items-center justify-center text-sm py-2 px-3 leading-tight text-gray-500 bg-white border border-gray-300 hover:bg-gray-100 hover:text-gray-700 dark:bg-gray-800 dark:border-gray-700 dark:text-gray-400 dark:hover:bg-gray-700 dark:hover:text-white"
						>1</a
					>
				</li>
				<li>
					<a
						href="#top"
						class="flex items-center justify-center text-sm py-2 px-3 leading-tight text-gray-500 bg-white border border-gray-300 hover:bg-gray-100 hover:text-gray-700 dark:bg-gray-800 dark:border-gray-700 dark:text-gray-400 dark:hover:bg-gray-700 dark:hover:text-white"
						>2</a
					>
				</li>
				<li>
					<a
						href="#top"
						aria-current="page"
						class="flex items-center justify-center text-sm z-10 py-2 px-3 leading-tight text-primary-600 bg-primary-50 border border-primary-300 hover:bg-primary-100 hover:text-primary-700 dark:border-gray-700 dark:bg-gray-700 dark:text-white"
						>3</a
					>
				</li>
				<li>
					<a
						href="#"
						class="flex items-center justify-center text-sm py-2 px-3 leading-tight text-gray-500 bg-white border border-gray-300 hover:bg-gray-100 hover:text-gray-700 dark:bg-gray-800 dark:border-gray-700 dark:text-gray-400 dark:hover:bg-gray-700 dark:hover:text-white"
						>...</a
					>
				</li>
				<li>
					<a
						href="#"
						class="flex items-center justify-center text-sm py-2 px-3 leading-tight text-gray-500 bg-white border border-gray-300 hover:bg-gray-100 hover:text-gray-700 dark:bg-gray-800 dark:border-gray-700 dark:text-gray-400 dark:hover:bg-gray-700 dark:hover:text-white"
						>100</a
					>
				</li>
				<li>
					<a
						href="#"
						class="flex items-center justify-center h-full py-1.5 px-3 leading-tight text-gray-500 bg-white rounded-r-lg border border-gray-300 hover:bg-gray-100 hover:text-gray-700 dark:bg-gray-800 dark:border-gray-700 dark:text-gray-400 dark:hover:bg-gray-700 dark:hover:text-white"
					>
						<span class="sr-only">Next</span>
						<svg
							class="w-5 h-5"
							aria-hidden="true"
							fill="currentColor"
							xmlns="http://www.w3.org/2000/svg"
						>
							<path
								fill-rule="evenodd"
								d="M7.293 14.707a1 1 0 010-1.414L10.586 10 7.293 6.707a1 1 0 011.414-1.414l4 4a1 1 0 010 1.414l-4 4a1 1 0 01-1.414 0z"
								clip-rule="evenodd"
							/>
						</svg>
					</a>
				</li>
			</ul>
		</nav>
	</div>
</section>
