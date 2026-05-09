<script lang="ts">
    import { goto } from '$app/navigation';
    import Page from '$lib/ui/Page.svelte';
    import Stack from '$lib/ui/Stack.svelte';
    import Input from '$lib/ui/Input.svelte';
    import Button from '$lib/ui/Button.svelte';
    import { _Submit } from './+page';
    import { markAuthed } from '$lib/stores/auth';

    let username = '';
    let email = '';
    let city = '';
    let password = '';
    let showPassword = false;
    let error = '';
    let submitting = false;

    async function submit(e: Event) {
        e.preventDefault();
        error = '';
        submitting = true;
        const status = await _Submit(e);
        submitting = false;

        if (status === 200) {
            markAuthed();
            goto('/account');
        } else if (status === 409) {
            error = 'Username or email already in use.';
        } else {
            error = 'Something went wrong. Try again.';
        }
    }
</script>

<svelte:head><title>Register · taskbingo</title></svelte:head>

<Page width="narrow">
    <form on:submit={submit} class="form">
        <Stack gap="l" align="stretch">
            <header class="head">
                <span class="eyebrow">register</span>
                <h1 class="title">Create an account.</h1>
            </header>

            <Stack gap="m">
                <Input label="Username" name="username" bind:value={username} required />
                <Input label="Email" name="email" type="email" bind:value={email} required />
                <Input label="City" name="city" bind:value={city} required />
                <div class="pw">
                    <Input
                        label="Password"
                        name="password"
                        type={showPassword ? 'text' : 'password'}
                        bind:value={password}
                        required
                    />
                    <button type="button" class="eye" on:click={() => (showPassword = !showPassword)} aria-label={showPassword ? 'Hide password' : 'Show password'}>
                        {#if showPassword}
                            <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.6" stroke-linecap="round" stroke-linejoin="round" width="20" height="20"><path d="M2 12s3.5-7 10-7 10 7 10 7-3.5 7-10 7-10-7-10-7z"/><circle cx="12" cy="12" r="3"/></svg>
                        {:else}
                            <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.6" stroke-linecap="round" stroke-linejoin="round" width="20" height="20"><path d="M3 3l18 18"/><path d="M2 12s3.5-7 10-7c2 0 3.7.6 5.1 1.5"/><path d="M9.5 9.5a3 3 0 0 0 4.2 4.2"/><path d="M22 12s-3.5 7-10 7c-1.4 0-2.6-.3-3.7-.8"/></svg>
                        {/if}
                    </button>
                </div>
                {#if error}
                    <span class="error" role="alert">{error}</span>
                {/if}
                <Button type="submit" variant="accent" fullWidth disabled={submitting}>
                    {submitting ? 'Creating…' : 'Register'}
                </Button>
            </Stack>

            <p class="alt">Have an account? <a href="/login">Log in</a></p>
        </Stack>
    </form>
</Page>

<style>
    .form { width: 100%; }

    .head {
        display: flex;
        flex-direction: column;
        gap: var(--space-2);
        align-items: center;
        text-align: center;
    }
    .eyebrow {
        font-size: 0.65rem;
        letter-spacing: 0.18em;
        text-transform: uppercase;
        color: var(--ink-3);
    }
    .title {
        font-family: var(--font-display);
        font-size: 1.8rem;
        font-weight: 400;
        letter-spacing: -0.018em;
        color: var(--ink);
        margin: 0;
    }

    .pw { position: relative; }
    .eye {
        position: absolute;
        right: 0.75rem;
        top: calc(50% - 0.6rem);
        background: transparent;
        border: none;
        color: var(--ink-3);
        cursor: pointer;
        padding: 0.5rem;
        border-radius: var(--radius-md);
        transition: color var(--dur-base) var(--ease);
    }
    .eye:hover { color: var(--ink); }
    .eye:focus-visible { outline: 2px solid var(--accent-ring); outline-offset: 2px; }

    .error { color: #c44e4e; font-size: 0.85rem; text-align: center; }
    .alt { text-align: center; font-size: 0.85rem; color: var(--ink-2); }
    .alt a {
        color: var(--ink);
        border-bottom: 1px solid rgba(40, 55, 95, 0.25);
        text-decoration: none;
        transition: color var(--dur-base) var(--ease), border-bottom-color var(--dur-base) var(--ease);
    }
    .alt a:hover {
        color: var(--accent-from);
        border-bottom-color: var(--accent-from);
    }
</style>
