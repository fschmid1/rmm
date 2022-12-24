<script lang="ts">
import { Table, TableBody, TableBodyCell, TableBodyRow, TableHead, TableHeadCell } from "flowbite-svelte";
import { onMount } from "svelte";
import { fetchWithToken } from "../http";
import { deviceStore } from "../stores";
import { apiBase } from "../vars";
import ConnectionIndecator from "./ConnectionIndecator.svelte";


	onMount(async () => {
		const response = await fetchWithToken(`${apiBase}/devices`, {
			method: "GET",
		});

		deviceStore.set(await response.json());
	});
</script>
<div class="w-9/12 mx-auto mt-12">


<Table striped={true}>
  <TableHead>
    <TableHeadCell>Name</TableHeadCell>
    <TableHeadCell>IP</TableHeadCell>
    <TableHeadCell>Cores</TableHeadCell>
    <TableHeadCell>Memory</TableHeadCell>
  </TableHead>
  <TableBody class="divide-y">
		{#each $deviceStore as device (device.id)}
			<TableBodyRow>
				<TableBodyCell><div class="flex"><ConnectionIndecator connected={device.connected} />{device.name}</div></TableBodyCell>
				<TableBodyCell>{device.systemInfo.ip}</TableBodyCell>
				<TableBodyCell>{device.systemInfo.cores}</TableBodyCell>
				<TableBodyCell>{device.systemInfo.memory}</TableBodyCell>
			</TableBodyRow>
		{/each}
  </TableBody>
</Table>
</div>