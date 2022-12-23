<script lang="ts">
import {
  onMount
} from 'svelte';
  import Header from './lib/Header.svelte';
	import { Router, Route } from 'svelte-navigator';
  import Redirects from './lib/Redirects.svelte';
  import Login from './lib/Login.svelte';
  import Devices from './lib/Devices.svelte';
  import { userStore } from './stores';
	import jwtDecode from 'jwt-decode';

onMount(async () => {
  document.documentElement.classList.add('dark');

	 if (location.pathname == '/login') return;
	 if (localStorage.getItem('token') == null) {
		 location.href = '/login';
		 return;
	 }
      userStore.set(jwtDecode(localStorage.getItem('token')));
});

</script>
<Header />

<Router>
	<Route primary={true} path="/devices">
		<Devices />
	</Route>
  <Route>
    <Redirects />
  </Route>
  <Route primary={false} path="/login">
    <Login />
  </Route>
</Router>
