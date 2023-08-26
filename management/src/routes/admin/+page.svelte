<script lang="ts">
    import Sidepanel from "./sidepanel.svelte";
    import {onMount} from "svelte";
    import {isAuthedAdmin, token} from "$lib/auth";
    import { goto } from '$app/navigation';

    onMount(()=>{
        if(!isAuthedAdmin($token)){
            goto("/login")
        }
    })
</script>

<Sidepanel></Sidepanel>
