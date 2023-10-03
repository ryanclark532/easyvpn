<script lang="ts">
	import { Section, Register } from 'flowbite-svelte-blocks';
	import { Button, Label, Input, Spinner } from 'flowbite-svelte';
	import { _authStore } from '../../+layout';
	let response: Error | undefined;
</script>

<Section name="login">
	<Register>
		<div class="p-6 space-y-4 md:space-y-6 sm:p-8">
			<form
				class="flex flex-col space-y-6"
				on:submit={async (e) => {
					response = await _authStore.changePassword(e);
				}}
			>
				<h3 class="text-xl font-medium text-gray-900 dark:text-white p-0">Change Password</h3>
				{#if response}
					<h5
						class="text-l font-medium text-red-600 p-2 bg-red-300 w-full rounded-lg border-red-600"
					>
						{response}
					</h5>
				{/if}
				<Label class="space-y-2">
					<span>New Password</span>
					<Input name="password" placeholder="•••••" required />
				</Label>
				<Label class="space-y-2">
					<span>Confirm New Password</span>
					<Input type="password" name="confirmPassword" placeholder="•••••" required />
				</Label>
				<Button type="submit">test</Button>
			</form>
		</div>
	</Register>
</Section>
