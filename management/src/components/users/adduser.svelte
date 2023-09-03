<script lang="ts">
	import { userValidationMessage } from '../../stores/stores';
	import { createUser } from '$lib';

	let hidden: boolean = false;
	let confirmed: boolean = false;

	let name: string;
	let username: string;
	let password: string;
	let enabled: boolean = false;
	let admin: boolean = false;

	async function handleUserCreation() {
		hidden = !(await createUser(name, username, password, enabled, admin));
		name = '';
		username = '';
		password = '';
		enabled = false;
		admin = false;
	}
</script>

<div class="flex justify-center m-5">
	<button
		id="defaultModalButton"
		data-modal-toggle="defaultModal"
		class="block text-white bg-primary-700 hover:bg-primary-800 focus:ring-4 focus:outline-none focus:ring-primary-300 font-medium rounded-lg text-sm px-5 py-2.5 text-center dark:bg-primary-600 dark:hover:bg-primary-700 dark:focus:ring-primary-800"
		type="button"
		on:click={() => {
			hidden = !hidden;
		}}
	>
		Create product
	</button>
</div>

{#if hidden}
	<div class="fixed top-0 left-0 flex items-center justify-center w-full h-full z-50">
		<div
			id="defaultModal"
			tabindex="-1"
			aria-hidden="true"
			class="bg-white mx-auto overflow-y-auto overflow-x-hidden rounded-lg shadow-lg w-11/12 sm:w-3/4 md:w-2/3 lg:w-1/2 xl:w-1/3"
		>
			<div class="relative p-4 w-full max-w-2xl h-full md:h-auto">
				<!-- Modal content -->
				<div class="relative p-4 bg-white rounded-lg shadow dark:bg-gray-800 sm:p-5">
					<!-- Modal header -->
					<div
						class="flex justify-between items-center pb-4 mb-4 rounded-t border-b sm:mb-5 dark:border-gray-600"
					>
						<h3 class="text-lg font-semibold text-gray-900 dark:text-white">Add Users</h3>
						<button
							type="button"
							class="text-gray-400 bg-transparent hover:bg-gray-200 hover:text-gray-900 rounded-lg text-sm p-1.5 ml-auto inline-flex items-center dark:hover:bg-gray-600 dark:hover:text-white"
							data-modal-toggle="defaultModal"
							on:click={() => {
								hidden = !hidden;
							}}
						>
							<svg
								aria-hidden="true"
								class="w-5 h-5"
								fill="currentColor"
								viewBox="0 0 20 20"
								xmlns="http://www.w3.org/2000/svg"
								><path
									fill-rule="evenodd"
									d="M4.293 4.293a1 1 0 011.414 0L10 8.586l4.293-4.293a1 1 0 111.414 1.414L11.414 10l4.293 4.293a1 1 0 01-1.414 1.414L10 11.414l-4.293 4.293a1 1 0 01-1.414-1.414L8.586 10 4.293 5.707a1 1 0 010-1.414z"
									clip-rule="evenodd"
								/></svg
							>
							<span class="sr-only">Close modal</span>
						</button>
					</div>
					{#if $userValidationMessage}
						<div class="bg-red-200 w-full py-1 text-center rounded-lg mb-6">
							<p class="text-sm text-red-600 font-medium">
								{$userValidationMessage}
							</p>
						</div>
					{/if}
					<!-- Modal body -->
					<div>
						<div class="mb-6">
							<label for="name" class="block mb-2 text-sm font-medium text-gray-900 dark:text-white"
								>Name</label
							>
							<input
								bind:value={name}
								type="text"
								id="name"
								class="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500"
							/>
						</div>
						<div class="mb-6">
							<label
								for="username"
								class="block mb-2 text-sm font-medium text-gray-900 dark:text-white">Username</label
							>
							<input
								bind:value={username}
								type="text"
								id="username"
								class="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500"
							/>
						</div>
						<div class="mb-6">
							<label
								for="password"
								class="block mb-2 text-sm font-medium text-gray-900 dark:text-white">Password</label
							>
							<input
								bind:value={password}
								type="text"
								id="password"
								class="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500"
							/>
						</div>
						<div class="grid grid-cols-2 mb-6">
							<label class="relative inline-flex items-center cursor-pointer">
								<input type="checkbox" class="sr-only peer" bind:checked={admin} />
								<div
									class="w-11 h-6 bg-gray-200 rounded-full peer peer-focus:ring-4 peer-focus:ring-blue-300 dark:peer-focus:ring-blue-800 dark:bg-gray-700 peer-checked:after:translate-x-full peer-checked:after:border-white after:content-[''] after:absolute after:top-0.5 after:left-[2px] after:bg-white after:border-gray-300 after:border after:rounded-full after:h-5 after:w-5 after:transition-all dark:border-gray-600 peer-checked:bg-blue-600"
								/>
								<span class="ml-3 text-sm font-medium text-gray-900 dark:text-gray-300"
									>Admin user</span
								>
							</label>
							<label class="relative inline-flex items-center cursor-pointer">
								<input type="checkbox" bind:checked={enabled} class="sr-only peer" />
								<div
									class="w-11 h-6 bg-gray-200 rounded-full peer peer-focus:ring-4 peer-focus:ring-blue-300 dark:peer-focus:ring-blue-800 dark:bg-gray-700 peer-checked:after:translate-x-full peer-checked:after:border-white after:content-[''] after:absolute after:top-0.5 after:left-[2px] after:bg-white after:border-gray-300 after:border after:rounded-full after:h-5 after:w-5 after:transition-all dark:border-gray-600 peer-checked:bg-blue-600"
								/>
								<span class="ml-3 text-sm font-medium text-gray-900 dark:text-gray-300"
									>Enabled</span
								>
							</label>
						</div>
						<div class=" w-full bg-gray-200 p-3 flex items-center mb-6 rounded-lg">
							<input
								bind:checked={confirmed}
								id="checkbox-1"
								type="checkbox"
								value=""
								class="w-4 h-4 text-blue-600 bg-gray-100 border-gray-300 rounded focus:ring-blue-500 dark:focus:ring-blue-600 dark:ring-offset-gray-800 dark:focus:ring-offset-gray-800 focus:ring-2 dark:bg-gray-700 dark:border-gray-600"
							/>
							<label
								for="checkbox-1"
								class="ml-2 text-sm font-medium text-gray-900 dark:text-gray-300"
							>
								Confirm User Options
							</label>
						</div>

						<button
							on:click={() => handleUserCreation()}
							class="w-full text-white inline-flex items-center bg-primary-700 hover:bg-primary-800 focus:ring-4 focus:outline-none focus:ring-primary-300 font-medium rounded-lg text-sm px-3 py-2.5 text-center dark:bg-primary-600 dark:hover:bg-primary-700 dark:focus:ring-primary-800 {!confirmed
								? 'opacity-50 cursor-not-allowed'
								: ''}"
						>
							<svg
								class="mr-1 -ml-1 w-6 h-6"
								fill="currentColor"
								viewBox="0 0 20 20"
								xmlns="http://www.w3.org/2000/svg"
							>
								<path
									fill-rule="evenodd"
									d="M10 5a1 1 0 011 1v3h3a1 1 0 110 2h-3v3a1 1 0 11-2 0v-3H6a1 1 0 110-2h3V6a1 1 0 011-1z"
									clip-rule="evenodd"
								/>
							</svg>
							Add new User
						</button>
					</div>
				</div>
			</div>
		</div>
	</div>
{/if}
