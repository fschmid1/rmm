<script lang="ts">
    import { Button, Input, Label } from 'flowbite-svelte';
    import { useNavigate } from 'svelte-navigator';
    import { userStore } from '../../stores';
    import { apiBase } from '../../vars';

    let email = '';
    let password = '';

    const navigate = useNavigate();

    const login = async () => {
        const response = await fetch(`${apiBase}/auth/login`, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({
                email,
                password,
            }),
        });
        const data = await response.json();
        if (data.user) {
            navigate('/');
            location.reload();
            userStore.set(data.user);
        }
    };
</script>

<div class="flex w-full justify-center mt-12">
    <div class="w-9/12 md:w-1/2 mt-12 rounded-lg p-8 border border-gray-700 grid">
        <Label class="space-y-2">
            <span>Email</span>
            <Input type="email" bind:value={email} placeholder="email" size="md" />
        </Label>
        <Label class="space-y-2 mt-4">
            <span>Password</span>
            <Input
                type="password"
                bind:value={password}
                class="w-full"
                placeholder=""
                size="md"
                on:keypress={(event) => {
                    if (event.key === 'Enter') login();
                }}
            />
        </Label>
        <Button on:click={login} class="mt-6" size="lg">Login</Button>
    </div>
</div>
