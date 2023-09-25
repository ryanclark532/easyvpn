<script lang="ts">
    import type {serverStatus} from "../../../types/vpn";
    import {ServerStatusMapping} from "../../../types/vpn";
    import {Button} from "flowbite-svelte";
    import {page} from "$app/stores";

    let status: serverStatus = $page.data.serverStatus.data.status

    let background : string;
    let containerClass: string
    $: {
        background = status === "running"
            ? "bg-green-300 border-green-600"
            : status === "notRunning"
                ? "bg-red-300 border-red-600"
                : "bg-orange-300 border-orange-600"

       containerClass =`text-center text-l text font-medium p-2  w-full rounded-lg ${background}`}
</script>

<div class={containerClass}>
    <h2 class="mt-5 text-2xl">{ServerStatusMapping[status]}</h2>
    <div class="mt-7 mb-7">
    {#if status === "running"}
        <Button class="w-40">
          Active Connections
        </Button>
        <Button class="w-40">
            Stop Server
        </Button>
        <Button class="w-40">
            Restart Server
        </Button>
        {/if}
    {#if status === "starting"}
        <Button class="w-40">
            See Users
        </Button>
        <Button class="w-40">
            Restart Server
        </Button>
    {/if}
    {#if status === "notRunning"}
        <Button class="w-40">
            See Users
        </Button>
        <Button class="w-40">
            Start Server
        </Button>
    {/if}
    </div>
</div>