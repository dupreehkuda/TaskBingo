<script lang="ts">
    import { goto } from '$app/navigation';
    import Page from '$lib/ui/Page.svelte';
    import Stack from '$lib/ui/Stack.svelte';
    import Cluster from '$lib/ui/Cluster.svelte';
    import Paper from '$lib/ui/Paper.svelte';
    import Input from '$lib/ui/Input.svelte';
    import Toggle from '$lib/ui/Toggle.svelte';
    import Button from '$lib/ui/Button.svelte';
    import TaskListEditor from '$lib/feature/TaskListEditor.svelte';
    import { _CreatePack } from './+page';

    let name = '';
    let tasks: string[] = [];
    let isPrivate = false;
    let submitting = false;
    let error = '';

    $: canSave = name.trim().length > 0 && tasks.length === 16 && !submitting;

    async function submit(e: Event) {
        e.preventDefault();
        if (!canSave) return;
        submitting = true;
        error = '';
        const status = await _CreatePack({ name: name.trim(), tasks, isPrivate });
        submitting = false;
        if (status === 200) {
            goto('/account');
        } else {
            error = 'Could not save the pack. Try again.';
        }
    }
</script>

<svelte:head><title>New pack · taskbingo</title></svelte:head>

<Page width="normal">
    <Stack gap="l">
        <h1 class="title">New pack</h1>

        <Paper padding="lg">
            <form on:submit={submit}>
                <Stack gap="l">
                    <Input label="Pack name" name="name" bind:value={name} required />

                    <div>
                        <span class="section-label">Tasks</span>
                        <TaskListEditor bind:tasks />
                    </div>

                    <Cluster gap="m" align="center" justify="between">
                        <Toggle label="Private (only you can see and play)" bind:checked={isPrivate} />
                        <Button type="submit" variant="accent" disabled={!canSave}>
                            {submitting ? 'Saving…' : 'Create pack'}
                        </Button>
                    </Cluster>

                    {#if error}<span class="error">{error}</span>{/if}
                </Stack>
            </form>
        </Paper>
    </Stack>
</Page>

<style>
    .title {
        font-family: var(--font-display);
        font-size: 1.6rem;
        font-weight: 400;
        letter-spacing: -0.015em;
    }
    .section-label {
        display: block;
        font-size: 0.65rem;
        text-transform: uppercase;
        letter-spacing: 0.16em;
        color: var(--ink-3);
        margin-bottom: var(--space-3);
    }
    .error { color: #c44e4e; font-size: 0.85rem; }
</style>
