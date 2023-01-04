<script lang="ts">
import { Table, TableBody, TableBodyCell, TableBodyRow, TableHead, TableHeadCell, Toggle } from "flowbite-svelte";
import { onMount } from "svelte";
import type { Device } from "../../types";
import { apiBase } from "../../vars";
import { fetchWithToken } from "../helper/http";
	let notifications: {user_id: number, device_id: number}[] = [];
	let devices: (Device & {checked: boolean})[] = []
	let drawerClosed = false;
	onMount(async () => {
		let response = await fetchWithToken(`${apiBase}/user/notifications`, {method: 'GET'});
		notifications = await response.json();

		response = await fetchWithToken(`${apiBase}/devices`, {method: 'GET'});
		devices = await response.json();
		devices = devices.map(device => {
			return {...device, checked: !!notifications.find(el => el.device_id === device.id)}
		});
	})
	const toggleNotifications = async (deviceID: number) => {
		const reponse = await fetchWithToken(`${apiBase}/user/notifications/toggle`, {
			method: 'PATCH',
			body: JSON.stringify({deviceID})
		})
	}
</script>
<Table>
	<TableHead>
		<TableHeadCell>Device</TableHeadCell>
		<TableHeadCell>Actions</TableHeadCell>
	</TableHead>
	<TableBody class="divide-y">
		{#each devices as device (device.id)}
		<TableBodyRow>
			<TableBodyCell>{device.name}</TableBodyCell>
			<TableBodyCell>
				<Toggle bind:checked={device.checked} on:change={() => toggleNotifications(device.id)}></Toggle>
			</TableBodyCell>
		</TableBodyRow>
		{/each}
	</TableBody>
</Table>