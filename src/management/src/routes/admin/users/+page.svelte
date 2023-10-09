<script lang="ts">
	import Sidepanel from '../sidepanel.svelte';
	import Table from './table.svelte';
	import { masterCheckbox, selectedUsers } from '$lib/users';
	import ConfirmationModal from '../../../components/confirmation-modal.svelte';
	import { _userStore } from '../../+layout';

	let deleteModal = false;
	let tempPasswordModal = false;
	let toggleAdmin = false;
	let toggleEnabled = false;
</script>

<div class="flex">
	<Sidepanel />
	<div class="w-full p-5 ml-64">
		<h2 class="text-2xl mt-5 mb-1">Users</h2>
		<Table />
	</div>
</div>

<ConfirmationModal
	open={deleteModal}
	title="Confirm User Updates"
	onConfirm={_userStore.delete}
	data={$selectedUsers}
/>

<ConfirmationModal
	open={tempPasswordModal}
	title="Confirm User Updates"
	onConfirm={_userStore.setTempPw}
	data={$selectedUsers}
/>

<ConfirmationModal
	open={toggleAdmin}
	title="Confirm User Updates"
	onConfirm={_userStore.updateUser}
	data={$selectedUsers.map((u) => {
		return {
			...u,
			admin: !u.IsAdmin
		};
	})}
/>

<ConfirmationModal
	on:actiontriggered={() => {
		selectedUsers.set([]);
		masterCheckbox.set(false);
	}}
	open={toggleEnabled}
	title="Confirm User Updates"
	onConfirm={_userStore.updateUser}
	data={$selectedUsers.map((u) => {
		return {
			...u,
			enabled: !u.Enabled
		};
	})}
/>
