<script lang="ts">
    import Sidepanel from '../sidepanel.svelte';
    import Status from "./status.svelte"
    import ActiveUsersTable from './active-users-table.svelte';
</script>

<div class="flex">
    <Sidepanel />
    <div class="w-full p-5">
        <Status />
                <ActiveUsersTable />

    </div>
   
</div>


