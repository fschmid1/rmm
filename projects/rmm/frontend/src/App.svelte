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
  userStore, ws, deviceStore, device
} from './stores';
import jwtDecode from 'jwt-decode';
import { Websocket } from './lib/helper/ws';
import DevicePage from './lib/components/DevicePage.svelte';
import { Confirm } from 'svelte-lib';
import { SvelteToast } from '@zerodevx/svelte-toast';
import TokenOverview from './lib/components/TokenOverview.svelte';

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
	<Route primary={false} path="/tokens">
		<TokenOverview />
	</Route>
    <Route primary={false} path="/login">
        <Login />
    </Route>
</Router>

<Confirm></Confirm>

<SvelteToast />