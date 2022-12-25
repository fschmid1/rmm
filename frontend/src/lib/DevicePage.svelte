<script lang="ts">
import {  Button, Li, List, } from "flowbite-svelte";
import { onMount } from "svelte";
import { Router, Route, Link } from "svelte-navigator";
import { fetchWithToken } from "../http";
import type { Device } from "../types";
import { apiBase } from "../vars";
import DeviceInfo from "./DeviceInfo.svelte";
import Fa from 'svelte-fa/src/fa.svelte'
import { faArrowLeft } from '@fortawesome/free-solid-svg-icons'


	export let id: string;

	let device: Device;


onMount(async () => {
  const response = await fetchWithToken(`${apiBase}/devices/${id}`, {
    method: "GET",
  });
	device = await response.json();
});
</script>

{#if device}
<div class="w-11/12 mx-auto mt-12">
	<Link to="/devices">
		<Button class="bg-gray-700 hover:bg-gray-600">
			<Fa icon={faArrowLeft} class="mr-2" />
			Back
		</Button>
	</Link>
	<div class="flex  justify-between gap-4 mt-4">
		<div class="w-3/12 rounded-lg shadow-md border border-gray-700 ">
			<List tag="ul" class=" divide-gray-200 dark:divide-gray-700 py-2" list="none">
				<Li class="py-2 pl-4">
					<Link to="/devices/{device.id}/info">
						<span class="space-x-4">Info</span>
					</Link>
				</Li>
				<Li class="py-2 pl-4">
					<span class="space-x-4">Test</span>
				</Li>
				<Li class="py-2 pl-4">
					<span class="space-x-4">Test</span>
				</Li>
				<Li class="py-2 pl-4">
					<span class="space-x-4">Test</span>
				</Li>
			</List>
		</div>
		<div class="w-9/12 rounded-lg shadow-md border border-gray-700 p-4">
			<Router>
				<Route path="/info" >
					<DeviceInfo device={device} />
				</Route>
			</Router>
		</div>
	</div>
	</div>
{/if}