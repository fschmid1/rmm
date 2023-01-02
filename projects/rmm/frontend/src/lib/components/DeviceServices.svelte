<script lang="ts">
import {
  onMount
} from "svelte";
import {
  callDeviceFunction,
} from "../helper/http";
import type {
  Device
} from "../../types";
import Fa from 'svelte-fa';
import {
  faCheck,
  faX
} from '@fortawesome/free-solid-svg-icons';
import {
  Button,
  Input,
  Li,
  List,
  Modal,
  Spinner
} from "flowbite-svelte";
import { customConfirm } from "../helper/functions";
type Service = {
  name: string,
  enabled: boolean,
  status: string
}
export let device: Device;
let serviceList: Service[] = [];
let serviceModal = false;
let selectedService: Service | null = null;
let filter = '';
let loading = true;

onMount(async () => {
  const response = await callDeviceFunction < string > (device.deviceID, 'service-list', '');

  serviceList = response.data.split('\n').map((s: string) => {
    let service: Service = {
      name: s.split('  ')[1],
      enabled: s.startsWith(' [ + ]'),
      status: ''
    };
    return service;
  });
  loading = false;
});

const openServiceModal = async (service: Service) => {
  selectedService = service;
  const response = await callDeviceFunction < string > (device.deviceID, 'service-status', service.name);
  selectedService.status = response.data.replace(/\n/g, '<br />');
  serviceModal = true;
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
    {#each serviceList.filter(el => el.name?.toLowerCase().includes(filter.toLowerCase())) as service}
    <Li icon class="cursor-pointer" >
        <!-- svelte-ignore a11y-click-events-have-key-events -->
        <!-- svelte-ignore missing-declaration -->
        <span class="flex" on:click={()=> {
            openServiceModal(service);
            }
            }>
            <Fa icon={service.enabled ? faCheck : faX} class="mr-2" color="{service.enabled ? 'green': 'red'}" />
            {service.name}
        </span>
    </Li>
    {/each}
</List>

<Modal title="{selectedService?.name}" bind:open={serviceModal} size="lg">
    {@html selectedService?.status}
    {selectedService?.status}
    <svelte:fragment slot='footer'>
        {#if selectedService?.enabled}
        <Button color="red" on:click={
						async () => {
								if (!(await customConfirm('Are you sure you want to stop this service?'))) {
									return;
								}
								await callDeviceFunction(device.deviceID, 'service-stop', selectedService?.name);
								serviceList = serviceList.map((s) => {
									if (s.name === selectedService?.name) {
										s.enabled = false;
									}
									return s;
								});
								await openServiceModal(selectedService);
						}
				}>Stop</Button>
        {:else}
        <Button color="green" on:click={
						async () => {
								await callDeviceFunction(device.deviceID, 'service-start', selectedService?.name);
								serviceList = serviceList.map((s) => {
									if (s.name === selectedService?.name) {
										s.enabled = true;
									}
									return s;
								});
								await openServiceModal(selectedService);
						}
				}>Start</Button>
        {/if}
        <Button color="green" on:click={
						async () => {
								if (!(await customConfirm('Are you sure you want to restart this service?'))) {
									return;
								}
								await callDeviceFunction(device.deviceID, 'service-restart', selectedService?.name);
								await openServiceModal(selectedService);
						}
				}>Restart</Button>
        </svelte:fragment>
        </Modal>
