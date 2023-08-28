<script lang="ts">
    import Sidepanel from "./sidepanel.svelte";
    import {isAuthedAdmin} from "$lib/auth";
    import {onMount} from "svelte";

    onMount(async () => {
        if (await isAuthedAdmin() === false) {
            window.location.href ="/login"
        }
    })
</script>

<Sidepanel></Sidepanel>
