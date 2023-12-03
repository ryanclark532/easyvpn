<script lang="ts">
	import { page } from '$app/stores';
	import { Section, SidebarBottomNav, SidebarBottomNavItem } from 'flowbite-svelte-blocks';
	import { Sidebar } from 'flowbite-svelte';
</script>

<Sidebar class=" border-r border-gray-200 w-96" style="height: 100vh">
	<div class=" border-b border-gray-200 p-3 mb-1">
		<div class="flex">
			<svg
				style="margin-top: 5px;"
				class="w-6 h-6 text-gray-500 dark:text-white"
				aria-hidden="true"
				xmlns="http://www.w3.org/2000/svg"
				viewBox="0 0 21 19"
			>
				<path
					stroke="currentColor"
					fill="transparent"
					stroke-linecap="round"
					stroke-linejoin="round"
					stroke-width="2"
					d="M11 4C5.5-1.5-1.5 5.5 4
					11l7 7 7-7c5.458-5.458-1.542-12.458-7-7Z"
				/></svg
			>
			<h2 class="text-2xl ml-2 mb-3">Server Status</h2>
		</div>
		<h2 class="text-xl rounded p-4 ml-4 hover:bg-gray-200">Statistics</h2>
		<h2 class="text-xl rounded p-4 ml-4 hover:bg-gray-200">Active Users</h2>
	</div>
	<div class=" border-b border-gray-200 p-3 mb-1">
		<div class="flex">
			<svg
				style="margin-top: 5px;"
				class="w-6 h-6 text-gray-500 dark:text-white"
				aria-hidden="true"
				xmlns="http://www.w3.org/2000/svg"
				fill="none"
				viewBox="0 0 20 18"
			>
				<path
					fill="transparent"
					stroke="currentColor"
					stroke-linecap="round"
					stroke-linejoin="round"
					stroke-width="2"
					d="M14 3a3 3 0 1 1-1.614 5.53M15 12a4 4 0 0 1 4 4v1h-3.348M10 4.5a3.5 3.5 0 1 1-7 0 3.5 3.5 0 0 1 7 0ZM5 11h3a4 4 0 0 1 4 4v2H1v-2a4 4 0 0 1 4-4Z"
				/>
			</svg>
			<h2 class="text-2xl ml-2 mb-3">User Management</h2>
		</div>
		<h2 class="text-xl rounded p-4 ml-4 hover:bg-gray-200">Users</h2>
		<h2 class="text-xl rounded p-4 ml-4 hover:bg-gray-200">Groups</h2>
	</div>
</Sidebar>
