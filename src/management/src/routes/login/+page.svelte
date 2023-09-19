<script lang="ts">
	import { Section, Register } from 'flowbite-svelte-blocks';
	import {Button, Checkbox, Label, Input, Spinner} from 'flowbite-svelte';
	import { handleLogin, loginResponse } from '$lib/auth';
</script>

<Section name="login">
	<Register>
		<div class="p-6 space-y-4 md:space-y-6 sm:p-8">
			<form class="flex flex-col space-y-6" on:submit={handleLogin}>
				<h3 class="text-xl font-medium text-gray-900 dark:text-white p-0">Sign In</h3>
				{#if $loginResponse.status === 'error'}
					<h5
						class="text-l font-medium text-red-600 p-2 bg-red-300 w-full rounded-lg border-red-600"
					>
						{$loginResponse.data}
					</h5>
				{/if}

				<Label class="space-y-2">
					<span>Your email</span>
					<Input name="username" placeholder="name@company.com" required />
				</Label>
				<Label class="space-y-2">
					<span>Your password</span>
					<Input type="password" name="password" placeholder="•••••" required />
				</Label>
				<div class="flex items-start">
					<Checkbox>Remember me</Checkbox>
					<a href="/" class="ml-auto text-sm text-blue-700 hover:underline dark:text-blue-500"
						>Forgot password?</a
					>
				</div>
				{#if $loginResponse.status ==='loading'}
					<Button>
						<Spinner class="mr-3" size="4" color="white" />Loading ...
					</Button>
					{:else}
					<Button type="submit" class="w-full1">Sign in</Button>
				{/if}
			</form>
		</div>
	</Register>
</Section>
