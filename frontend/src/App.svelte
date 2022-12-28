<script lang="ts">
import {
  onMount
} from 'svelte';
import Header from './lib/Header.svelte';
import {
  Router,
  Route
} from 'svelte-navigator';
import Redirects from './lib/Redirects.svelte';
import Login from './lib/Login.svelte';
import Devices from './lib/Devices.svelte';
import {
  userStore, ws, deviceStore, device
} from './stores';
import jwtDecode from 'jwt-decode';
import { Websocket } from './ws';
import DevicePage from './lib/DevicePage.svelte';
import Confirm from './lib/Confirm.svelte';
import { SvelteToast } from '@zerodevx/svelte-toast';

onMount(async () => {
  document.documentElement.classList.add('dark');

  if (location.pathname == '/login') return;
  if (localStorage.getItem('token') == null) {
    location.href = '/login';
    return;
  }
  userStore.set((jwtDecode(localStorage.getItem('token')) as any).user);

	ws.set(new Websocket($userStore.id, localStorage.getItem('token')));

	$ws.on('device-connection', (data: {id: number, connected: boolean}) => {
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
    <Route>
        <Redirects />
    </Route>
    <Route primary={false} path="/login">
        <Login />
    </Route>
</Router>

<Confirm></Confirm>

<SvelteToast />