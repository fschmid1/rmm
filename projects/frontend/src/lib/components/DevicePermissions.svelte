<script lang="ts">
    import { createQuery, useQueryClient } from '@tanstack/svelte-query';
    import { sineIn } from 'svelte/easing';
    import { Drawer, Spinner, Accordion, Li, AccordionItem, Toggle, Button, CloseButton, List } from 'flowbite-svelte';
    import type { User, Device, DevicePermission } from '../../types';
    import { userStore } from '../../stores';
    import { apiBase } from '../../vars';
    import { fetchWithToken } from '../helper/http';
    import Fa from 'svelte-fa';
    import { faPlus, faTrash } from '@fortawesome/free-solid-svg-icons';
    import { customConfirm } from 'svelte-lib';
    export let device: Device;

    let drawerClosed = true;

    const permissions = createQuery<DevicePermission[], Error>({
        queryKey: ['permissions', device.id],
        queryFn: () =>
            fetchWithToken(`${apiBase}/devices/${device.id}/permissions`, {
                method: 'GET',
            }).then((res) => res.json()),
    });

    const users = createQuery<User[], Error>({
        queryKey: ['users'],
        queryFn: () =>
            fetchWithToken(`${apiBase}/user/all`, {
                method: 'GET',
            }).then((res) => res.json()),
    });

    const handleToggle = (permission: DevicePermission) => {
        setTimeout(async () => {
            const res = await fetchWithToken(`${apiBase}/devices/${device.id}/permissions`, {
                method: 'PATCH',
                body: JSON.stringify(permission),
            });
            if (res.status !== 200) {
                $permissions.refetch();
            }
        }, 0);
    };

    const handleDelete = async (permission: DevicePermission) => {
        if ((await customConfirm('Are you sure you want to delete this permission?')) === false) return;
        const res = await fetchWithToken(`${apiBase}/devices/${device.id}/permissions/${permission.id}`, {
            method: 'DELETE',
        });
        if (res.status === 200) {
            client.invalidateQueries(['permissions', device.id]);
        }
    };

    const client = useQueryClient();

    let transitionParams = {
        x: 320,
        duration: 200,
        easing: sineIn,
    };

    const addPermissions = async (user: User) => {
        const res = await fetchWithToken(`${apiBase}/devices/${device.id}/permissions`, {
            method: 'PATCH',
            body: JSON.stringify({
                deviceID: device.id,
                userID: user.id,
                run: false,
                kill: false,
                reboot: false,
                shutdown: false,
                processList: false,
                changePermissions: false,
                serviceList: false,
                serviceStart: false,
                serviceStop: false,
                serviceRestart: false,
                serviceStatus: false,
                serviceLogs: false,
            }),
        });
        client.invalidateQueries(['permissions', device.id]);
        drawerClosed = true;
    };

    $: filterdUsers = $users.data?.filter((user) => {
        const permission = $permissions.data?.find((permission) => permission.user.id === user.id);
        return permission === undefined && user.id !== $userStore.id;
    });

    $: canChange = $permissions.data?.find((permission) => permission.user.id === $userStore.id)?.changePermissions;
</script>

{#if canChange}
    <div class="flex w-full justify-end mb-2">
        <Button on:click={() => (drawerClosed = false)}>
            <Fa icon={faPlus} class="mr-2" />
            Add
        </Button>
    </div>
{/if}
{#if $permissions.isLoading}
    <div class="text-center">
        <Spinner size="12" />
    </div>
{/if}
{#if $permissions.data}
    <Accordion>
        {#each $permissions.data as permission (permission.id)}
            {#if permission.user.id !== $userStore.id}
                <AccordionItem class="flex justify-between items-center">
                    <span slot="header">
                        {permission.user.name}
                    </span>
                    {#if canChange}
                        <div class="flex flex-row relative">
                            <div class="absolute right-2 cursor-pointer" on:click={() => handleDelete(permission)}>
                                <Fa icon={faTrash} />
                            </div>
                            <div class="w-1/2">
                                <Toggle
                                    class="my-1"
                                    on:click={() => handleToggle(permission)}
                                    bind:checked={permission.run}>Run</Toggle
                                >
                                <Toggle
                                    class="my-1"
                                    on:click={() => handleToggle(permission)}
                                    bind:checked={permission.kill}>Kill</Toggle
                                >
                                <Toggle
                                    class="my-1"
                                    on:click={() => handleToggle(permission)}
                                    bind:checked={permission.reboot}>Reboot</Toggle
                                >
                                <Toggle
                                    class="my-1"
                                    on:click={() => handleToggle(permission)}
                                    bind:checked={permission.shutdown}>Shutdown</Toggle
                                >
                                <Toggle
                                    class="my-1"
                                    on:click={() => handleToggle(permission)}
                                    bind:checked={permission.processList}>ProcessList</Toggle
                                >
                                <Toggle
                                    class="my-1"
                                    on:click={() => handleToggle(permission)}
                                    bind:checked={permission.changePermissions}>ChangePermission</Toggle
                                >
                            </div>
                            <div class="w-1/2">
                                <Toggle
                                    class="my-1"
                                    on:click={() => handleToggle(permission)}
                                    bind:checked={permission.serviceList}>ServiceList</Toggle
                                >
                                <Toggle
                                    class="my-1"
                                    on:click={() => handleToggle(permission)}
                                    bind:checked={permission.serviceStart}>ServiceStart</Toggle
                                >
                                <Toggle
                                    class="my-1"
                                    on:click={() => handleToggle(permission)}
                                    bind:checked={permission.serviceStop}>ServiceStop</Toggle
                                >
                                <Toggle
                                    class="my-1"
                                    on:click={() => handleToggle(permission)}
                                    bind:checked={permission.serviceRestart}>ServiceRestart</Toggle
                                >
                                <Toggle
                                    class="my-1"
                                    on:click={() => handleToggle(permission)}
                                    bind:checked={permission.serviceStatus}>ServiceStatus</Toggle
                                >
                                <Toggle
                                    class="my-1"
                                    on:click={() => handleToggle(permission)}
                                    bind:checked={permission.serviceLogs}>ServiceLogs</Toggle
                                >
                            </div>
                        </div>
                    {/if}
                </AccordionItem>
            {/if}
        {/each}
    </Accordion>
{/if}
<Drawer transitionType="fly" placement="right" {transitionParams} bind:hidden={drawerClosed}>
    <div class="flex items-center">
        <h5
            id="drawer-label"
            class="inline-flex items-center mb-6 text-base font-semibold text-gray-500 uppercase dark:text-gray-400"
        >
            Add permissions
        </h5>
        <CloseButton on:click={() => (drawerClosed = true)} class="mb-4 dark:text-white" />
    </div>
    <List>
        {#if filterdUsers.length > 0}
            {#each filterdUsers as user (user.id)}
                <Li class="flex justify-between items-center my-2">
                    <span>{user.name}</span>
                    <Button on:click={() => addPermissions(user)} class="ml-2">
                        <Fa icon={faPlus} class="mr-2" />
                        Add
                    </Button>
                </Li>
            {/each}
        {/if}
    </List>
</Drawer>
