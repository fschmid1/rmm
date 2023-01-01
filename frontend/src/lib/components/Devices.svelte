<script lang="ts">
import {
  Table,
  TableBody,
  TableBodyCell,
  TableBodyRow,
  TableHead,
  TableHeadCell
} from "flowbite-svelte";
import {
  onMount
} from "svelte";
import {
  Link, useNavigate
} from "svelte-navigator";
import {
  fetchWithToken
} from "../http";
import {
  deviceStore
} from "../stores";
import {
  apiBase
} from "../vars";
import ConnectionIndecator from "./ConnectionIndecator.svelte";

let navigate = useNavigate();

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
							<TableBodyRow class="cursor-pointer" on:click={() => navigate(`/devices/${device.id}/info`)}>
									<TableBodyCell><div class="flex"><ConnectionIndecator connected={device.connected} />{device.name}</div></TableBodyCell>
									<TableBodyCell>{device.systemInfo.ip}</TableBodyCell>
									<TableBodyCell>{device.systemInfo.cores}</TableBodyCell>
									<TableBodyCell>{device.systemInfo.memoryTotal.toFixed(2)} GB</TableBodyCell>
							</TableBodyRow>
            {/each}
        </TableBody>
    </Table>
</div>
