<script lang="ts">
    import { QueryClient, QueryClientProvider } from '@tanstack/svelte-query';
    import { SvelteToast } from '@zerodevx/svelte-toast';
    import { onMount } from 'svelte';
    import { Confirm } from 'svelte-lib';
    import { Route, Router } from 'svelte-navigator';
    import DevicePage from './lib/components/DevicePage.svelte';
    import Devices from './lib/components/Devices.svelte';
    import Header from './lib/components/Header.svelte';
    import Login from './lib/components/Login.svelte';
    import Redirects from './lib/components/Redirects.svelte';
    import SettingsPage from './lib/components/SettingsPage.svelte';
    import { fetchWithToken } from './lib/helper/http';
    import { Websocket } from './lib/helper/ws';
    import { userStore, ws } from './stores';
    import type { Device, User } from './types';
    import { apiBase } from './vars';

    const queryClient = new QueryClient();

    onMount(async () => {
        document.documentElement.classList.add('dark');

        if (location.pathname == '/login') return;
        if (localStorage.getItem('token') == null) {
            location.href = '/login';
            return;
        }

        const user = await (
            await fetchWithToken(`${apiBase}/user`, {
                method: 'GET',
            })
        ).json();
        userStore.set(user as User);

        ws.set(new Websocket($userStore.id, localStorage.getItem('token')));

        $ws.on('device-connection', (data: { id: number; connected: boolean }) => {
            const newData = (queryClient.getQueryData(['devices']) as Device[]).map((device: any) => {
                if (device.id == data.id) {
                    device.connected = data.connected;
                }
                return device;
            });
            queryClient.setQueryData(['devices'], newData);
        });
    });
</script>

<QueryClientProvider client={queryClient}>
    <Router>
        <Header />
        <Route primary={true} path="/devices">
            <Devices />
        </Route>
        <Route primary={true} path="/devices/:id/*" let:params>
            <DevicePage id={params.id} />
        </Route>
        <Route primary={false} path="/settings/*">
            <SettingsPage />
        </Route>
        <Route>
            <Redirects />
        </Route>
        <Route primary={false} path="/login">
            <Login />
        </Route>
    </Router>
    <Confirm />
    <SvelteToast />
</QueryClientProvider>
