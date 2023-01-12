<script lang="ts">
    import { Button, Li, List, Spinner } from 'flowbite-svelte';
    import { onMount } from 'svelte';
    import { Router, Route, Link } from 'svelte-navigator';
    import { fetchWithToken } from '../helper/http';
    import type { Device } from '../../types';
    import { apiBase } from '../../vars';
    import DeviceInfo from './DeviceInfo.svelte';
    import Fa from 'svelte-fa/src/fa.svelte';
    import { faArrowLeft } from '@fortawesome/free-solid-svg-icons';
    import DeviceHeader from './DeviceHeader.svelte';
    import DeviceServices from './DeviceServices.svelte';
    import DeviceProcesses from './DeviceProcesses.svelte';
    import DeviceUsage from './DeviceUsage.svelte';
    import DeviceRun from './DeviceRun.svelte';
    import { createQuery } from '@tanstack/svelte-query';

    export let id: string;

    const device = createQuery<Device, Error>({
        queryKey: ['devices', parseInt(id)],
        queryFn: () =>
            fetchWithToken(`${apiBase}/devices/${id}`, {
                method: 'GET',
            }).then((res) => res.json()),
    });
</script>

{#if $device.isLoading}
    <div class="text-center w-full"><Spinner /></div>
{:else}
    <div class="w-11/12 mx-auto mt-12">
        <Link to="/devices">
            <Button class="bg-gray-700 hover:bg-gray-600">
                <Fa icon={faArrowLeft} class="mr-2" />
                Back
            </Button>
        </Link>
        <div class="flex justify-between gap-4 mt-4">
            <div class="w-3/12 rounded-lg shadow-md border border-gray-700 ">
                <List tag="ul" class=" divide-gray-200 dark:divide-gray-700 py-2" list="none">
                    <Li class="py-2 pl-4">
                        <Link to="/devices/{$device.data.id}/info">
                            <span class="space-x-4">Info</span>
                        </Link>
                    </Li>
                    <Li class="py-2 pl-4">
                        <Link to="/devices/{$device.data.id}/services">
                            <span class="space-x-4">Services</span>
                        </Link>
                    </Li>
                    <Li class="py-2 pl-4">
                        <Link to="/devices/{$device.data.id}/processes">
                            <span class="space-x-4">Processes</span>
                        </Link>
                    </Li>
                    <Li class="py-2 pl-4">
                        <Link to="/devices/{$device.data.id}/usage">
                            <span class="space-x-4">Usage</span>
                        </Link>
                    </Li>
                    <Li class="py-2 pl-4">
                        <Link to="/devices/{$device.data.id}/run">
                            <span class="space-x-4">Run</span>
                        </Link>
                    </Li>
                </List>
            </div>
            <div class="flex w-9/12 flex-col">
                <DeviceHeader device={$device.data} />
                <div class="rounded-lg shadow-md border border-gray-700 p-4">
                    <Router>
                        <Route primary={false} path="/info">
                            <DeviceInfo device={$device.data} />
                        </Route>
                        <Route path="/services" primary={false}>
                            <DeviceServices device={$device.data} />
                        </Route>
                        <Route path="/processes" primary={false}>
                            <DeviceProcesses device={$device.data} />
                        </Route>
                        <Route path="/usage" primary={false}>
                            <DeviceUsage device={$device.data} />
                        </Route>
                        <Route path="/run" primary={false}>
                            <DeviceRun device={$device.data} />
                        </Route>
                    </Router>
                </div>
            </div>
        </div>
    </div>
{/if}
