<script lang="ts">
    import { Avatar, Dropdown, DropdownHeader, DropdownItem, Navbar } from 'flowbite-svelte';
    import { useNavigate } from 'svelte-navigator';
    import { userStore } from '../../stores';
    import { apiBase } from '../../vars';
    const navigate = useNavigate();
</script>

<Navbar>
    <div class="flex ml-auto items-center md:order-2 cursor-pointer" id="avatar-menu">
        {#if $userStore}
            <Avatar />
            <span class="block ml-2 text-sm">{$userStore?.name}</span>
            <Dropdown placement="bottom" triggeredBy="#avatar-menu">
                <DropdownHeader>
                    <DropdownItem on:click={() => navigate('/settings/profile')}>Settings</DropdownItem>
                    <DropdownItem on:click={() => navigate('/settings/tokens')}>Tokens</DropdownItem>
                    <DropdownItem on:click={() => navigate('/settings/notifications')}>Notifications</DropdownItem>
                </DropdownHeader>
                <DropdownItem
                    on:click={async () => {
                        await fetch(`${apiBase}/auth/logout`, {
                            method: 'POST',
                        });
                        location.pathname = '/login';
                    }}>Sign out</DropdownItem
                >
            </Dropdown>
        {/if}
        {#if !$userStore}
            <Avatar size="md" class="pace-y-2.5 animate-pulse w-8" />
            <span class="pace-y-2.5 w-24 h-4 ml-2 rounded bg-gray-500 animate-pulse" />
        {/if}
    </div>
</Navbar>
