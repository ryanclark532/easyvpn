<script lang="ts">
	import { Button, Modal } from 'flowbite-svelte';

	export let title: string;
	export let subtext: string;
	export let onConfirm: (...args: any) => Promise<any> | any;
	export let open: boolean;
</script>

<Modal {title} bind:open autoclose>
	<div>{subtext}</div>
	<div class="flex">
		<Button class="w-1/2 m-1">Close</Button>
		<Button class="w-1/2 m-1" on:click={onConfirm}>Confirm</Button>
	</div>
</Modal>
