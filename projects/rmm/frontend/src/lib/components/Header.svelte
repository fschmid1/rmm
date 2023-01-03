<script lang="ts">
  import { Navbar, DarkMode, Dropdown, DropdownItem, DropdownHeader, Skeleton, Avatar  } from 'flowbite-svelte';
  import { userStore } from '../../stores';
</script>

<Navbar>
  <div class="flex ml-auto items-center md:order-2 cursor-pointer" id="avatar-menu">
    {#if $userStore}
      <Avatar />
      <span class="block ml-2 text-sm">{$userStore?.name}</span>
	  <Dropdown placement="bottom" triggeredBy="#avatar-menu">
		<DropdownHeader>
			<DropdownItem href="/tokens">Tokens</DropdownItem>
		</DropdownHeader>
		<DropdownItem
		  on:click={async () => {
					localStorage.clear();
			location.reload();
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