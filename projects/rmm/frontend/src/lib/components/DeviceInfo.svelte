<script lang="ts">
import {
  Button,
  Hr,
  Table,
  TableBody,
  TableBodyCell,
  TableBodyRow
} from "flowbite-svelte";
import { customConfirm } from "../helper/functions";
import {
  callDeviceFunction
} from "../helper/http";
import type {
  Device
} from "../../types";

export let device: Device;
</script>

<div class="w-full">
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
    <Hr />
    <div class="flex mt-2">
        <Button color="red" class="mr-2" on:click={async () => {
						if (await customConfirm("Are you sure you want to reboot this device?")) {
							callDeviceFunction(device.deviceID, "reboot")
						}
						}}>Reboot</Button>
        <Button color="red" on:click={ async () => {
					if (await customConfirm("Are you sure you want to shutdown this device?")) {
						callDeviceFunction(device.deviceID, "shutdown")
					}
					}}>Shutdown</Button>
    </div>
</div>
