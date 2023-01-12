<script lang="ts">
    import { Button, Input, Label } from 'flowbite-svelte';
    import { onDestroy, onMount } from 'svelte';
    import { userStore } from '../../stores';
    import type { User } from '../../types';
    import { apiBase } from '../../vars';
    import { fetchWithToken } from '../helper/http';

    let user: User;
    let unsub;

    onMount(async () => {
        user = $userStore;
        unsub = userStore.subscribe((u) => {
            user = u;
        });
    });

    onDestroy(() => {
        unsub();
    });

    const saveUser = async () => {
        const res = await fetchWithToken(`${apiBase}/user`, {
            method: 'PATCH',
            body: JSON.stringify(user),
        });
        if (res.status == 200) {
            $userStore = user;
        }
        userStore.set((await res.json()) as User);
    };
</script>

{#if user}
    <form on:submit|preventDefault={saveUser}>
        <div class="grid gap-6 mb-6 md:grid-cols-2">
            <div>
                <Label for="name" class="mb-2">Name</Label>
                <Input type="text" bind:value={user.name} id="name" placeholder="" required />
            </div>
            <div>
                <Label for="email" class="mb-2">Email</Label>
                <Input type="text" id="email" value={user.email} placeholder="" readonly required />
            </div>
        </div>
        <div class="mb-6">
            <Label for="pushtoken" class="mb-2">PushOver Token</Label>
            <Input type="text" bind:value={user.pushToken} id="pushtoken" placeholder="" required />
        </div>
        <Button type="submit">Save</Button>
    </form>
{/if}
