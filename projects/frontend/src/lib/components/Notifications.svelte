<script lang="ts">
    import { createQuery } from '@tanstack/svelte-query';
    import {
        Spinner,
        Table,
        TableBody,
        TableBodyCell,
        TableBodyRow,
        TableHead,
        TableHeadCell,
        Toggle,
    } from 'flowbite-svelte';
    import { onMount } from 'svelte';
    import type { Device } from '../../types';
    import { apiBase } from '../../vars';
    import { fetchWithToken } from '../helper/http';
    let notifications = createQuery<
        {
            user_id: number;
            device_id: number;
        }[],
        Error
    >({
        queryKey: ['notifications'],
        queryFn: () =>
            fetchWithToken(`${apiBase}/user/notifications`, {
                method: 'GET',
            }).then((res) => res.json()),
    });

    const devices = createQuery<
        (Device & {
            checked: boolean;
        })[],
        Error
    >({
        enabled: !!$notifications.data !== undefined,
        queryKey: ['notificationDevices'],
        queryFn: () =>
            fetchWithToken(`${apiBase}/devices`, {
                method: 'GET',
            })
                .then((res) => res.json())
                .then((data) =>
                    data.map((device) => {
                        return {
                            ...device,
                            checked: !!$notifications.data.find((el) => el.device_id === device.id),
                        };
                    }),
                ),
    });
    const toggleNotifications = async (deviceID: number) => {
        const reponse = await fetchWithToken(`${apiBase}/user/notifications/toggle`, {
            method: 'PATCH',
            body: JSON.stringify({
                deviceID,
            }),
        });
    };
</script>

{#if $devices.isLoading}
    <div class="text-center w-full"><Spinner /></div>
{:else}
    <Table>
        <TableHead>
            <TableHeadCell>Device</TableHeadCell>
            <TableHeadCell>Actions</TableHeadCell>
        </TableHead>
        <TableBody class="divide-y">
            {#each $devices.data as device (device.id)}
                <TableBodyRow>
                    <TableBodyCell>{device.name}</TableBodyCell>
                    <TableBodyCell>
                        <Toggle bind:checked={device.checked} on:change={() => toggleNotifications(device.id)} />
                    </TableBodyCell>
                </TableBodyRow>
            {/each}
        </TableBody>
    </Table>
{/if}
