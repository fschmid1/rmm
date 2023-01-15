<script lang="ts">
    import { faCopy } from '@fortawesome/free-solid-svg-icons';
    import { createQuery } from '@tanstack/svelte-query';
    import { toast } from '@zerodevx/svelte-toast';
    import { Spinner, Table, TableBody, TableBodyCell, TableBodyRow, TableHead, TableHeadCell } from 'flowbite-svelte';
    import Fa from 'svelte-fa/src/fa.svelte';
    import { useNavigate } from 'svelte-navigator';
    import type { Device } from '../../types';
    import { apiBase } from '../../vars';
    import { fetchWithToken } from '../helper/http';
    import ConnectionIndecator from './ConnectionIndecator.svelte';
    let navigate = useNavigate();

    const devices = createQuery<Device[], Error>({
        queryKey: ['devices'],
        queryFn: () =>
            fetchWithToken(`${apiBase}/devices`, {
                method: 'GET',
            }).then((res) => res.json()),
    });

    const handleClick = (event, device) => {
        if (event.target.classList.contains('ip')) {
            toast.push('IP copied to clipboard');
            navigator.clipboard.writeText(device.systemInfo.ip);
            return;
        }
        navigate(`/devices/${device.id}/info`);
    };
</script>

<div class="w-9/12 mx-auto mt-12">
    {#if $devices.isLoading}
        <div class="text-center w-full"><Spinner /></div>
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
                    <TableBodyRow class="cursor-pointer" on:click={(event) => handleClick(event, device)}>
                        <TableBodyCell>
                            <div class="flex">
                                <ConnectionIndecator connected={device.connected} />{device.name}
                            </div>
                        </TableBodyCell>
                        <TableBodyCell>
                            {#if device.systemInfo.ip}
                                <div class="ip flex">
                                    {device.systemInfo.ip}
                                    <div class="ml-2"><Fa icon={faCopy} class="" /></div>
                                </div>
                            {/if}
                        </TableBodyCell>
                        <TableBodyCell>{device.systemInfo.cores}</TableBodyCell>
                        <TableBodyCell>{device.systemInfo.memoryTotal.toFixed(2)} GB</TableBodyCell>
                    </TableBodyRow>
                {/each}
            </TableBody>
        </Table>
    {/if}
</div>

<style>
    .hover-trigger .hover-target {
        display: none;
    }

    .hover-trigger:hover .hover-target {
        display: block;
    }
</style>
