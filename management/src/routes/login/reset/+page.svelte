<script lang="ts">
	import { Section, Register } from 'flowbite-svelte-blocks';
	import { Button, Label, Input, Spinner } from 'flowbite-svelte';
	import { passwordChangeResponse } from '$lib/auth';
</script>

<Section name="login">
	<Register>
		<div class="p-6 space-y-4 md:space-y-6 sm:p-8">
			<form class="flex flex-col space-y-6">
				<h3 class="text-xl font-medium text-gray-900 dark:text-white p-0">Change Password</h3>

				<Label class="space-y-2">
					<span>New Password</span>
					<Input name="newPassword" placeholder="•••••" required />
				</Label>
				<Label class="space-y-2">
					<span>Confirm New Password</span>
					<Input type="password" name="confirmPassword" placeholder="•••••" required />
				</Label>

				{#if $passwordChangeResponse.status === 'loading'}
					<Button>
						<Spinner class="mr-3" size="4" color="white" />Loading ...
					</Button>
				{:else}
					<Button type="submit" class="w-full1">Change Password</Button>
				{/if}
			</form>
		</div>
	</Register>
</Section>
