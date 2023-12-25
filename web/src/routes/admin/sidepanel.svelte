<script lang="ts">
	import { Sidebar } from 'flowbite-svelte';
	import UserInfo from '../../components/user-info.svelte';

	const username = localStorage.getItem('name') ?? '';
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
		<a href="/admin/status/statistics"
			><h2 class="text-xl rounded p-4 ml-4 hover:bg-gray-200 cursor-pointer">Statistics</h2></a
		>
		<a href="/admin/status/active-users"
			><h2 class="text-xl rounded p-4 ml-4 hover:bg-gray-200 cursor-pointer">Active Users</h2></a
		>
		<a href="/admin/status/logs"
			><h2 class="text-xl rounded p-4 ml-4 hover:bg-gray-200 cursor-pointer">Logs</h2></a
		>
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
		<a href="/admin/users">
			<h2 class="text-xl rounded p-4 ml-4 hover:bg-gray-200 cursor-pointer">Users</h2></a
		>
		<a href="/admin/users/groups"
			><h2 class="text-xl rounded p-4 ml-4 hover:bg-gray-200 cursor-pointer">Groups</h2></a
		>
	</div>
	<div class=" border-b border-gray-200 p-3 mb-1">
		<div class="flex">
			<svg
				style="margin-top: 5px;"
				class="w-6 h-6 text-gray-500 dark:text-white"
				aria-hidden="true"
				xmlns="http://www.w3.org/2000/svg"
				viewBox="0 0 20 20"
			>
				<g stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2">
					<path
						fill="transparent"
						d="M19 11V9a1 1 0 0 0-1-1h-.757l-.707-1.707.535-.536a1 1 0 0 0 0-1.414l-1.414-1.414a1 1 0 0 0-1.414 0l-.536.535L12 2.757V2a1 1 0 0 0-1-1H9a1 1 0 0 0-1 1v.757l-1.707.707-.536-.535a1 1 0 0 0-1.414 0L2.929 4.343a1 1 0 0 0 0 1.414l.536.536L2.757 8H2a1 1 0 0 0-1 1v2a1 1 0 0 0 1 1h.757l.707 1.707-.535.536a1 1 0 0 0 0 1.414l1.414 1.414a1 1 0 0 0 1.414 0l.536-.535L8 17.243V18a1 1 0 0 0 1 1h2a1 1 0 0 0 1-1v-.757l1.707-.708.536.536a1 1 0 0 0 1.414 0l1.414-1.414a1 1 0 0 0 0-1.414l-.535-.536.707-1.707H18a1 1 0 0 0 1-1Z"
					/>
					<path fill="transparent" d="M10 13a3 3 0 1 0 0-6 3 3 0 0 0 0 6Z" />
				</g>
			</svg>
			<h2 class="text-2xl ml-2 mb-3">Settings</h2>
		</div>
		<a href="/admin/settings/server"
			><h2 class="text-xl rounded p-4 ml-4 hover:bg-gray-200 cursor-pointer">Server Settings</h2></a
		>
		<a href="/admin/settings/client"
			><h2 class="text-xl rounded p-4 ml-4 hover:bg-gray-200 cursor-pointer">Client Settings</h2></a
		>

		<a href="/admin/settings/auth"
			><h2 class="text-xl rounded p-4 ml-4 hover:bg-gray-200 cursor-pointer">Authentication</h2></a
		>
		<a href="/admin/settings/config"
			><h2 class="text-xl rounded p-4 ml-4 hover:bg-gray-200 cursor-pointer">
				Edit Base Configuration
			</h2></a
		>
	</div>
	<div class="fixed bottom-0">
		<UserInfo {username} placement="top" />
	</div>
</Sidebar>
