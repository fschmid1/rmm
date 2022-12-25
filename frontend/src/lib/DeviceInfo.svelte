<script lang="ts">
import { Hr, Table, TableBody, TableBodyCell, TableBodyRow } from "flowbite-svelte";
import type {
  Device
} from "../types";
import ConnectionIndecator from "./ConnectionIndecator.svelte";

export let device: Device;
</script>

<div class="w-full">
    <div class="flex items-center justify-center">
        <ConnectionIndecator connected={device.connected} />
        {device.name}
    </div>
    <Table>
        <TableBody>
            <TableBodyRow>
                <TableBodyCell>Memory</TableBodyCell>
                <TableBodyCell class="text-end">{device.systemInfo.memoryTotal.toFixed(2)} GB</TableBodyCell>
            </TableBodyRow>
            <TableBodyRow>
                <TableBodyCell>CPU</TableBodyCell>
                <TableBodyCell class="text-end">{device.systemInfo.cpu} x {device.systemInfo.cores}</TableBodyCell>
            </TableBodyRow>
            <TableBodyRow>
                <TableBodyCell>GPU</TableBodyCell>
                <TableBodyCell class="text-end">{device.systemInfo.gpu}</TableBodyCell>
            </TableBodyRow>
            <TableBodyRow>
                <TableBodyCell>Disk</TableBodyCell>
                <TableBodyCell class="text-end">{device.systemInfo.diskTotal} GB</TableBodyCell>
            </TableBodyRow>
        </TableBody>
    </Table>
    <Hr />
    <Table>
        <TableBody>
            <TableBodyRow>
                <TableBodyCell>OS</TableBodyCell>
                <TableBodyCell class="text-end">{device.systemInfo.os}</TableBodyCell>
            </TableBodyRow>
            <TableBodyRow>
                <TableBodyCell>Hostname</TableBodyCell>
                <TableBodyCell class="text-end">{device.systemInfo.hostName}</TableBodyCell>
            </TableBodyRow>
            <TableBodyRow>
                <TableBodyCell>IP</TableBodyCell>
                <TableBodyCell class="text-end">{device.systemInfo.ip}</TableBodyCell>
            </TableBodyRow>
            <TableBodyRow>
                <TableBodyCell>MacAddress</TableBodyCell>
                <TableBodyCell class="text-end">{device.systemInfo.macAddress}</TableBodyCell>
            </TableBodyRow>
        </TableBody>
    </Table>
</div>
