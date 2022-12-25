<script lang="ts">
import type {
  Device
} from "../types";
import ConnectionIndecator from "./ConnectionIndecator.svelte";
export let device: Device;
</script>

<div class="flex items-center justify-center rounded-lg shadow-md border border-gray-700 p-4 mb-4">
    <ConnectionIndecator connected={device.connected} />
    {device.name}
</div>
