<script lang="ts">
    import { faCopy, faHashtag, faPlus, faTrash } from '@fortawesome/free-solid-svg-icons';
    import { createQuery, useQueryClient } from '@tanstack/svelte-query';
    import { toast } from '@zerodevx/svelte-toast';
    import {
        Button,
        CloseButton,
        Drawer,
        Input,
        Label,
        Spinner,
        Table,
        TableBody,
        TableBodyCell,
        TableBodyRow,
        TableHead,
        TableHeadCell,
    } from 'flowbite-svelte';
    import Fa from 'svelte-fa/src/fa.svelte';
    import { customConfirm } from 'svelte-lib';
    import { sineIn } from 'svelte/easing';
    import type { DeviceToken } from '../../types';
    import { apiBase } from '../../vars';
    import { fetchWithToken } from '../helper/http';

    const queryClient = useQueryClient();

    let tokens = createQuery<DeviceToken[], Error>({
        queryKey: ['tokens'],
        queryFn: () =>
            fetchWithToken(`${apiBase}/devices/tokens`, {
                method: 'GET',
            }).then((res) => res.json()),
    });

    const copyToken = (token: string) => {
        navigator.clipboard.writeText(token);
        toast.push('Token copied to clipboard');
    };

    const createToken = async () => {
        const response = await fetchWithToken(`${apiBase}/devices/tokens`, {
            method: 'POST',
            body: JSON.stringify({
                name: tokenName,
            }),
        });
        const token = await response.json();
        drawerClosed = true;
        queryClient.setQueriesData(['tokens'], [...$tokens.data, token]);
        tokenName = '';
    };

    const deleteToken = async (token: DeviceToken) => {
        if (!(await customConfirm('Are you sure you want to delete this token?'))) return;
        const response = await fetchWithToken(`${apiBase}/devices/tokens/${token.id}`, {
            method: 'DELETE',
        });
        if (response.status === 200) {
            queryClient.setQueriesData(
                ['tokens'],
                $tokens.data.filter((el) => el.id !== token.id),
            );
            tokens = tokens;
        }
    };

    let drawerClosed = true;
    let tokenName = '';

    let transitionParams = {
        x: 320,
        duration: 200,
        easing: sineIn,
    };
</script>

<div class="flex w-full justify-end">
    <Button on:click={() => (drawerClosed = false)}>
        <Fa icon={faPlus} class="mr-2" />
        Create new token
    </Button>
</div>
{#if $tokens.isLoading}
    <div class="w-full text-center"><Spinner /></div>
{:else}
    <Table striped={true} class="mt-2">
        <TableHead>
            <TableHeadCell>Name</TableHeadCell>
            <TableHeadCell>Token</TableHeadCell>
            <TableHeadCell>Actions</TableHeadCell>
        </TableHead>
        <TableBody class="divide-y">
            {#each $tokens.data as token (token.id)}
                <TableBodyRow>
                    <TableBodyCell>{token.name}</TableBodyCell>
                    <TableBodyCell class="truncate" style="max-width: 10rem">{token.token}</TableBodyCell>
                    <TableBodyCell>
                        <button on:click={() => copyToken(token.token)} class="mr-2"
                            ><Fa class="cursor-pointer" size="lg" icon={faCopy} /></button
                        >
                        <button on:click={() => deleteToken(token)}
                            ><Fa class="cursor-pointer" size="lg" icon={faTrash} /></button
                        >
                    </TableBodyCell>
                </TableBodyRow>
            {/each}
        </TableBody>
    </Table>
{/if}
<Drawer transitionType="fly" placement="right" {transitionParams} bind:hidden={drawerClosed}>
    <div class="flex items-center">
        <h5
            id="drawer-label"
            class="inline-flex items-center mb-6 text-base font-semibold text-gray-500 uppercase dark:text-gray-400"
        >
            Create new token
        </h5>
        <CloseButton on:click={() => (drawerClosed = true)} class="mb-4 dark:text-white" />
    </div>
    <form action="#" class="mb-6" on:submit|preventDefault={createToken}>
        <div class="mb-6">
            <Label for="name" class="block mb-2">Name</Label>
            <Input id="name" name="name" bind:value={tokenName} required placeholder="server01" />
        </div>
        <Button type="submit" class="w-full">
            <Fa class="mr-2" icon={faHashtag} /> Create Token
        </Button>
    </form>
</Drawer>
