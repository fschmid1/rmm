<script lang="ts">
import {
    onMount
} from 'svelte';
import Header from './lib/components/Header.svelte';
import {
    Router,
    Route
} from 'svelte-navigator';
import Redirects from './lib/components/Redirects.svelte';
import Login from './lib/components/Login.svelte';
import Devices from './lib/components/Devices.svelte';
import {
    userStore,
    ws,
    deviceStore,
    device
} from './stores';
import {
    Websocket
} from './lib/helper/ws';
import DevicePage from './lib/components/DevicePage.svelte';
import {
    Confirm
} from 'svelte-lib';
import {
    SvelteToast
} from '@zerodevx/svelte-toast';
import SettingsPage from './lib/components/SettingsPage.svelte';
import {
    fetchWithToken
} from './lib/helper/http';
import {
    apiBase
} from './vars';
import type {
    User
} from './types';

onMount(async () => {
    document.documentElement.classList.add('dark');

    if (location.pathname == '/login') return;
    if (localStorage.getItem('token') == null) {
        location.href = '/login';
        return;
    }

    const user = await (await (fetchWithToken(`${apiBase}/user`, {
        method: 'GET'
    }))).json();
    userStore.set(user as User);

    ws.set(new Websocket($userStore.id, localStorage.getItem('token')));

    $ws.on('device-connection', (data: {
        id: number,
        connected: boolean
    }) => {
        $deviceStore = $deviceStore.map(device => {
            if (device.id == data.id) {
                device.connected = data.connected;
            }
            return device;
        });
        if ($device?.id == data.id) {
            $device.connected = data.connected;
        }
    });
});
</script>
<Header />
<Router>
	<Route primary={true} path="/devices">
		<Devices />
	</Route>
	<Route primary={true} path="/devices/:id/*" let:params >
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

<Confirm></Confirm>

<SvelteToast />
