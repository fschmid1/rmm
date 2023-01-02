<script lang="ts">
import {
    Button,
    CloseButton,
    Drawer,
    Input,
    Label,
    Table,
    TableBody,
    TableBodyCell,
    TableBodyRow,
    TableHead,
    TableHeadCell
} from "flowbite-svelte";
import {
    onMount
} from "svelte";
import Fa from "svelte-fa/src/fa.svelte";
import {
    faArrowLeft,
    faCopy,
    faPlus,
	faHashtag,
    faTrash
} from '@fortawesome/free-solid-svg-icons';
import type {
    DeviceToken
} from "../../types";
import {
    apiBase
} from "../../vars";
import {
    fetchWithToken
} from "../helper/http";
import {
    toast
} from "@zerodevx/svelte-toast";
import { Link } from "svelte-navigator";
import { sineIn } from 'svelte/easing';
import { customConfirm } from "../helper/functions";

let tokens: DeviceToken[] = [];

onMount(async () => {
    const response = await fetchWithToken(`${apiBase}/devices/tokens`, {
        method: "GET",
    });
    tokens = await response.json();
});

const copyToken = (token: string) => {
    navigator.clipboard.writeText(token);
    toast.push("Token copied to clipboard");
}

const createToken = async () => {
	const response = await fetchWithToken(`${apiBase}/devices/tokens`, {
		method: "POST",
		body: JSON.stringify({
			name: tokenName
		})
	});
	const token = await response.json();
	drawerClosed = true;
	tokens = [...tokens, token];
	tokenName = "";
}

const deleteToken = async (token: DeviceToken) => {
	if (!(await customConfirm("Are you sure you want to delete this token?")))
		return;
	const response = await fetchWithToken(`${apiBase}/devices/tokens/${token.id}`, {
		method: "DELETE"
	});
	if (response.status === 200) {
		tokens = tokens.filter(el => el.id !== token.id);
	}
}

let drawerClosed = true;
let tokenName = "";

let transitionParams = {
    x: 320,
    duration: 200,
    easing: sineIn
  };
</script>

<div class="w-9/12 mx-auto mt-12">
	<div class="flex w-full justify-between">
		<Link to="-1">
			<Button class="bg-gray-700 hover:bg-gray-600">
				<Fa icon={faArrowLeft} class="mr-2" />
				Back
			</Button>
		</Link>
		<Button on:click={() => drawerClosed = false}>
			<Fa icon={faPlus} class="mr-2" />
			Create new token
		</Button>
	</div>
    <Table striped={true} class="mt-2">
        <TableHead>
            <TableHeadCell>Name</TableHeadCell>
            <TableHeadCell>Token</TableHeadCell>
            <TableHeadCell>Actions</TableHeadCell>
        </TableHead>
        <TableBody class="divide-y">
            {#each tokens as token (token.id)}
            <TableBodyRow>
                <TableBodyCell>{token.name}</TableBodyCell>
                <TableBodyCell class="truncate" style="max-width: 10rem">{token.token}</TableBodyCell>
                <TableBodyCell>
					<button on:click={() => copyToken(token.token)} class="mr-2"><Fa class="cursor-pointer" size="lg" icon={faCopy} /></button>
					<button on:click={() => deleteToken(token)}><Fa class="cursor-pointer" size="lg" icon={faTrash} /></button>
				</TableBodyCell>
            </TableBodyRow>
            {/each}
        </TableBody>
    </Table>
</div>
<Drawer transitionType="fly" placement="right" {transitionParams} bind:hidden={drawerClosed}>
	<div class='flex items-center'>
		<h5 id="drawer-label" class="inline-flex items-center mb-6 text-base font-semibold text-gray-500 uppercase dark:text-gray-400">Create new token</h5>
		<CloseButton on:click={() => (drawerClosed = true)} class='mb-4 dark:text-white'/>
	  </div>
	   <form action="#" class="mb-6">
		<div class="mb-6">
		  <Label for='name' class='block mb-2'>Name</Label>
		  <Input id='name' name='name' bind:value={tokenName} required placeholder="server01" />
		</div>
			<Button type="submit" class="w-full" on:click={createToken}>
				<Fa class="mr-2" icon={faHashtag} /> Create Token
			</Button>
	   </form>
</Drawer>
