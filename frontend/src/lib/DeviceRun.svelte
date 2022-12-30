<script lang="ts">
import { Input} from "flowbite-svelte";
import { callDeviceFunction } from "../http";
import type { Device } from "../types";


	export let device: Device;
	let commands: string[] = [];
	let commandIndex = -1;
	let command = '';
	let output = '';

	const handleKey = async (event) => {
	if (event.key === "Enter") {
		output = (await callDeviceFunction<string>(device.deviceID, 'run', command)).data.replace(/\n/g, '<br />');
		commands = [command, ...commands];
		command = '';
		commandIndex = -1;
	}  else if (event.key === "ArrowUp") {
		event.preventDefault();
		if (commandIndex < commands.length - 1  || commands.length === -1) {
			
			commandIndex = commandIndex + 1;
			command = commands[commandIndex];
		}
	} else if (event.key === "ArrowDown") {
		event.preventDefault();
		if (commandIndex > 0) {
			commandIndex = commandIndex - 1;
			command = commands[commandIndex];
		} else if (commandIndex === 0) {
			commandIndex = -1;
			command = '';
		}
	}
}

</script>
<div class="mb-2 h-96 overflow-y-scroll">
	{@html output}
</div>

<Input type="text" bind:value={command} id="command" placeholder="" required on:keydown={handleKey} />