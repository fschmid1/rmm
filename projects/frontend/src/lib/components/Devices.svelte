<script lang="ts">
import {
    Spinner,
    Table,
    TableBody,
    TableBodyCell,
    TableBodyRow,
    TableHead,
    TableHeadCell
} from "flowbite-svelte";
import {
    Link,
    useNavigate
} from "svelte-navigator";
import {
    fetchWithToken
} from "../helper/http";
import {
    apiBase
} from "../../vars";
import ConnectionIndecator from "./ConnectionIndecator.svelte";
import {
    createQuery
} from "@tanstack/svelte-query";
import type {
    Device
} from "../../types";
let navigate = useNavigate();

const devices = createQuery <
    Device[],
    Error >
    ({
        queryKey: ['devices'],
        queryFn: () => fetchWithToken(`${apiBase}/devices`, {
            method: "GET",
        }).then(res => res.json()),
    })
</script>
<div class="w-9/12 mx-auto mt-12">

    {#if $devices.isLoading}
    <div class="text-center w-full"><Spinner/></div>
        {:else}
        <Table striped={true}>
            <TableHead>
                <TableHeadCell>Name</TableHeadCell>
                <TableHeadCell>IP</TableHeadCell>
                <TableHeadCell>Cores</TableHeadCell>
                <TableHeadCell>Memory</TableHeadCell>
            </TableHead>
            <TableBody class="divide-y">
                {#each $devices.data as device (device.id)}
                <TableBodyRow class="cursor-pointer" on:click={() => navigate(`/devices/${device.id}/info`)}>
                    <TableBodyCell><div class="flex"><ConnectionIndecator connected={device.connected} />{device.name}</div></TableBodyCell>
                    <TableBodyCell>{device.systemInfo.ip}</TableBodyCell>
                    <TableBodyCell>{device.systemInfo.cores}</TableBodyCell>
                    <TableBodyCell>{device.systemInfo.memoryTotal.toFixed(2)} GB</TableBodyCell>
                </TableBodyRow>
                {/each}
            </TableBody>
        </Table>
        {/if}

        </div>
