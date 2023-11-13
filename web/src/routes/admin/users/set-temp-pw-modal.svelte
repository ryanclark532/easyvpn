<script lang="ts">
	import { changePassword } from '$lib/users';
	import { Input, DropdownItem, Label, Modal, Button } from 'flowbite-svelte';

	let open: boolean;

	export let userId: number;
</script>

<DropdownItem
	on:click={() => {
		open = !open;
	}}>Set Temporary Password</DropdownItem
>
<Modal bind:open title="Set Temporary Password">
	<div class="text-center">
		Users will need to login to the user panel and reset their password to continue using easypn
	</div>
	<form on:submit={(e)=>{changePassword(e, userId, true)}}>
		<Label for="password">Password</Label>
		<Input class="mb-4" id="password" name="password" placeholder="Password..." />
		<Label for="confirmPassword">Confirm Password</Label>
		<Input
			class="mb-4"
			id="confirm"
			name="confirm"
			placeholder="Confirm Password..."
		/>
		<Button class="w-full" type="submit">Set Temporary Password</Button>
	</form>
</Modal>
