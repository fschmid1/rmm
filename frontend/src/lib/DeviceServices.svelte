<script lang="ts">
import {
  onMount
} from "svelte";
import {
  fetchWithToken
} from "../http";
import type {
  Device
} from "../types";
import {
  apiBase
} from "../vars";
import Fa from 'svelte-fa';
import {
  faCheck,
  faX
} from '@fortawesome/free-solid-svg-icons';
import {
  Button,
  Li,
  List,
  Modal
} from "flowbite-svelte";
type Service = {
  name: string,
  enabled: boolean,
	status: string
}
export let device: Device;

let serviceList: Service[] = [];
let serviceModal = false;
let selectedService: Service | null = null;

onMount(async () => {
  const response = await fetchWithToken(`${apiBase}/devices/functions`, {
    method: "POST",
    body: JSON.stringify({
      id: device.deviceID,
      event: 'service-list',
      data: ''
    }),

  });

  serviceList = (await response.json()).data.split('\n').map((s: string) => {
    let service = {
      name: s.split('  ')[1],
      enabled: s.startsWith(' [ + ]'),
			status: ''
    };
    return service;
  });
});
</script>
<List tag="ul" class="spaces-y-1 h-96 overflow-y-scroll max-w-full" list="none">
    {#each serviceList as service}
    <Li icon class="cursor-pointer" >
        <span class="flex" on:click={async() => {
          	  selectedService = service;
							const response = await fetchWithToken(`${apiBase}/devices/functions`, {
								method: "POST",
								body: JSON.stringify({
									id: device.deviceID,
									event: 'service-status',
									data: service.name
								}),
							});
							selectedService.status = (await response.json()).data.replace(/\n/g, '<br />');
            	serviceModal = true;
            }}>
            <Fa icon={service.enabled ? faCheck : faX} class="mr-2" color="{service.enabled ? 'green': 'red'}" />
            {service.name}
        </span>
    </Li>
    {/each}
</List>

<Modal title="{selectedService?.name}" bind:open={serviceModal} autoclose size="lg">
		{@html selectedService?.status}
			{selectedService?.status}
    <svelte:fragment slot='footer'>
			{#if selectedService?.enabled}
        <Button color="red">Stop</Button>
			{:else}
				<Button color="green">Start</Button>
			{/if}
			<Button color="green">Restart</Button>
    </svelte:fragment>
  </Modal>
