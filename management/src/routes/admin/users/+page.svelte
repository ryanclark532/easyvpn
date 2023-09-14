<script lang="ts">
	import { page } from '$app/stores';
	import Sidepanel from '../sidepanel.svelte';
	import CreateUserModal from './create-user-modal.svelte';
	import Table from './table.svelte';
	import { users } from '$lib/users';
	import {Input, Label} from "flowbite-svelte";

	users.set($page.data.data.users);

</script>

<div class="flex">
	<Sidepanel />
	<div class="w-full p-5">
			<div class="mb-6 w-100 float-left">
				<Label for="default-input" class="mb-2">Default input</Label>
				<Input id="default-input" placeholder="Default input" style="height: 40px; width: 400px"/>
			</div>
			<CreateUserModal/>
		<Table />
	</div>
</div>
