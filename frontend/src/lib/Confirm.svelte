<script lang="ts">
import { Button, Modal } from "flowbite-svelte";
import { confirmModal } from "../stores";
</script>
<Modal bind:open={$confirmModal.open} size="xs">
  <div class="text-center">
      <svg aria-hidden="true" class="mx-auto mb-4 w-14 h-14 text-gray-400 dark:text-gray-200" fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4m0 4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z"></path></svg>
      <h3 class="mb-5 text-lg font-normal text-gray-500 dark:text-gray-400">{$confirmModal.text}</h3>
      <Button color="red" class="mr-2" on:click={ () => {
				confirmModal.set({
					open: false,
					text: "",
					desision: true
				});
			}}>Yes</Button>
      <Button color='alternative' on:click={ () => {
				confirmModal.set({
					open: false,
					text: "",
					desision: false
				});
			}}>No</Button>
  </div>
</Modal>