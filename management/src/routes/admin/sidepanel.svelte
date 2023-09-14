<script>
	import {
		Sidebar,
		SidebarWrapper,
		SidebarBrand,
		SidebarItem,
		SidebarGroup
	} from 'flowbite-svelte';
	let spanClass = 'flex-1 ml-3 whitespace-nowrap text-l';
	let site = {
		name: 'Flowbite-Svelte',
		href: '/admin',
		img: '/images/flowbite-svelte-icon-logo.svg'
	};
</script>

<Sidebar style="height: 100vh; position: fixed;" class="bg-gray-50 h-full">
	<SidebarWrapper>
		<SidebarGroup>
			<SidebarBrand {site} />
			<SidebarItem label="Status" {spanClass} href="/admin/status">
				<svelte:fragment slot="icon">
					<svg
						aria-hidden="true"
						class="w-6 h-6 text-gray-400 transition duration-75 dark:text-gray-400 group-hover:text-gray-900 dark:group-hover:text-white"
						fill="currentColor"
						viewBox="0 0 20 20"
						xmlns="http://www.w3.org/2000/svg"
						><path d="M2 10a8 8 0 018-8v8h8a8 8 0 11-16 0z" /><path
							d="M12 2.252A8.014 8.014 0 0117.748 8H12V2.252z"
						/></svg
					></svelte:fragment
				>
			</SidebarItem>
			<SidebarItem label="Users" {spanClass} href="/admin/users">
				<svelte:fragment slot="icon"
					><svg
						aria-hidden="true"
						class="w-6 h-6 text-gray-400 transition duration-75 dark:text-gray-400 group-hover:text-gray-900 dark:group-hover:text-white"
						fill="currentColor"
						viewBox="0 0 20 20"
						xmlns="http://www.w3.org/2000/svg"
						><path d="M2 10a8 8 0 018-8v8h8a8 8 0 11-16 0z" /><path
							d="M12 2.252A8.014 8.014 0 0117.748 8H12V2.252z"
						/></svg
					></svelte:fragment
				>
			</SidebarItem>
			<SidebarItem label="Certificates" {spanClass} href="/admin/certificates">
				<svelte:fragment slot="icon"
					><svg
						aria-hidden="true"
						class="w-6 h-6 text-gray-400 transition duration-75 dark:text-gray-400 group-hover:text-gray-900 dark:group-hover:text-white"
						fill="currentColor"
						viewBox="0 0 20 20"
						xmlns="http://www.w3.org/2000/svg"
						><path d="M2 10a8 8 0 018-8v8h8a8 8 0 11-16 0z" /><path
							d="M12 2.252A8.014 8.014 0 0117.748 8H12V2.252z"
						/></svg
					></svelte:fragment
				>
			</SidebarItem>
			<SidebarItem label="Settings" href="/admin/settings">
				<svelte:fragment slot="icon"
					><svg
						aria-hidden="true"
						class="w-6 h-6 text-gray-400 transition duration-75 dark:text-gray-400 group-hover:text-gray-900 dark:group-hover:text-white"
						fill="currentColor"
						viewBox="0 0 20 20"
						xmlns="http://www.w3.org/2000/svg"
						><path d="M2 10a8 8 0 018-8v8h8a8 8 0 11-16 0z" /><path
							d="M12 2.252A8.014 8.014 0 0117.748 8H12V2.252z"
						/></svg
					></svelte:fragment
				>
			</SidebarItem>
		</SidebarGroup>
	</SidebarWrapper>
</Sidebar>
