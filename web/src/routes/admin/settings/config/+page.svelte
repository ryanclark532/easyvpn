<script lang="ts">
	import { Heading, P, Textarea } from 'flowbite-svelte';
	import Sidepanel from '../../sidepanel.svelte';
</script>

<div class="flex">
	<Sidepanel />
	<div class="w-full p-5">
<Heading tag="h2" class="mb-2">Network Settings</Heading>
			<P>This page contains the raw OpenVPN config file. Allowing As much customization as needed</P
			>
<Textarea placeholder="Your message" name="message" />
</div>

</div> 
