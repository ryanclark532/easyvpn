<script lang="ts">
	import { Section, Register } from 'flowbite-svelte-blocks';
	import { Button, Label, Input } from 'flowbite-svelte';
	import { changePassword } from '$lib/users';
	import { getID } from '$lib/auth'
	import { goto } from '$app/navigation';
	let response: Error | undefined
	async function handleSubmit(e: Event){
		response = await changePassword(e, Number(getID() ?? ""), false)
		if(!response){
			goto("/login")
		}
	}

</script>

<Section name="login">
	<Register href="/">
		<div class="p-6 space-y-4 md:space-y-6 sm:p-8">
			<form class="flex flex-col space-y-6" on:submit={handleSubmit}>
				<h3 class="text-xl font-medium text-gray-900 dark:text-white p-0">Change Password</h3>

	{#if response}
		<h5 class="text-l font-medium text-red-600 p-2 bg-red-300 w-full rounded-lg border-red-600">
			{response}
		</h5>
	{/if}
				<Label class="space-y-2">
					<span> Password</span>
					<Input type="password" name="password" placeholder="•••••" required />
				</Label>
				<Label class="space-y-2">
					<span>Confirm Password</span>
					<Input type="password" name="confirm" placeholder="•••••" required />
				</Label>
				<Button type="submit" class="w-full1">Change Password</Button>
			</form>
		</div>
	</Register>
</Section>
