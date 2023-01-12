<script lang="ts">
    import { onMount } from 'svelte';
    import { useNavigate } from 'svelte-navigator';
    const navigate = useNavigate();
    onMount(() => {
        navigate('/devices');
    });
</script>
