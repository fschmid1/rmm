<script lang="ts">
	export let connected = false;
</script>


<div class="w-4 h-4 rounded-full {connected ? 'animate-pulse' : ''} mr-2 {connected ? 'bg-green-400': 'bg-red-600'}" />