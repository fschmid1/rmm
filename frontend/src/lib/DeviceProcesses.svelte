<script lang="ts">
import {
  onMount
} from "svelte";
import {
    callDeviceFunction,
  
} from "../http";
import type {
  Device
} from "../types";
import {
    A,
  Button,
  Input,
  Li,
  List,
  Modal,
  Spinner,
  Table,
  TableBody,
  TableBodyCell,
  TableBodyRow
} from "flowbite-svelte";
import { customConfirm } from "../functions";
type Process = {
  name: string,
	user: string,
	processID: string,
	options: string[],
}
export let device: Device;

let processList: Process[] = [];
let processModal = false;
let selectedProcess: Process | null = null;
let filter = '';
let loading = true;

onMount(async () => {
	await getProcesses();
});

async function getProcesses() {
  const response = await callDeviceFunction<string>(device.deviceID, 'process-list', '');

  processList = response.data.split('\n').map((s: string) => {
		let split = s.split(' ').filter(s => s !== '');
		let name = split[10];
		if (name?.split('/').length > 2) {
			const nameSplit = name.split('/');
			name = nameSplit[nameSplit.length - 1];
		}
    let process: Process = {
      name: name,
			user: split[0],
			processID: split[1],
			options: split.filter(s => s.startsWith('--')),
    };
    return process;
  });
	processList.splice(0, 1);
	processList.splice(processList.length - 1, 1);
	loading = false;
}
</script>
<Input id="search" bind:value={filter} placeholder="Search" size="md" class="mb-2">
    <svg slot="left" aria-hidden="true" class="w-6 h-6 text-gray-500 dark:text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z" /></svg>
</Input>
{#if loading}
	<div class="text-center">
		<Spinner size="12"></Spinner>
	</div>
{/if}
<List tag="ul" class="spaces-y-1 h-96 overflow-y-scroll min-w-full" list="none">
    {#each processList.filter(el => el.name.toLowerCase().includes(filter.toLowerCase())) as process}
    <Li  class="cursor-pointer" >
        <span class="flex" on:click={async() => {
          	  selectedProcess = process;
            	processModal = true;
            }}>
            {process.name}
        </span>
    </Li>
    {/each}
</List>

<Modal title="{selectedProcess?.name}" bind:open={processModal} size="lg">
	<Table>
		<TableBody>
			<TableBodyRow>
				<TableBodyCell>{selectedProcess.name}</TableBodyCell>
				<TableBodyCell>{selectedProcess.user}</TableBodyCell>
				<TableBodyCell>{selectedProcess.processID}</TableBodyCell>
				<TableBodyCell class="flex flex-col">
					{#each selectedProcess.options as option}
						<span>{option}</span>
					{/each}
				</TableBodyCell>
			</TableBodyRow>
		</TableBody>
	</Table>
    <svelte:fragment slot='footer'>
			<Button color="red" on:click={
				async () => {
					if (!(await customConfirm('Are you sure you want to kill this process?')))  return;
					await callDeviceFunction(device.deviceID, 'process-kill', selectedProcess.name);
					setTimeout(async () => {
						await getProcesses();
						filter = '';
					processModal = false;
					}, 50);
				}
			}>Kill</Button>
    </svelte:fragment>
  </Modal>
