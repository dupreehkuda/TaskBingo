<script lang="ts">
    import { onDestroy, onMount } from 'svelte';
    import { browser } from '$app/environment';
    import { page } from '$app/stores';
    import { Auth } from '$lib/stores/auth';

    let scrolled = false;

    function onScroll() {
        scrolled = window.scrollY > 64;
    }

    onMount(() => {
        if (!browser) return;
        onScroll();
        window.addEventListener('scroll', onScroll, { passive: true });
    });

    onDestroy(() => {
        if (!browser) return;
        window.removeEventListener('scroll', onScroll);
    });

    $: pathname = $page.url.pathname;
    $: isActive = (href: string) => pathname === href || (href !== '/' && pathname.startsWith(href));
    $: authed = $Auth;
</script>

<div class="header" class:scrolled>
    <div class="bar">
        <a class="brand" href="/">taskbingo</a>
        <nav class="nav" aria-label="Primary">
            {#if authed}
                <a class:active={isActive('/account')} href="/account" data-sveltekit-prefetch>play</a>
                <a class:active={isActive('/packs')} href="/packs" data-sveltekit-prefetch>packs</a>
                <a class:active={isActive('/people')} href="/people" data-sveltekit-prefetch>people</a>
            {/if}
            <a class:active={isActive('/about')} href="/about">about</a>
            {#if authed}
                <a class="avatar" href="/account" aria-label="Account" data-sveltekit-prefetch></a>
            {:else}
                <a class:active={isActive('/login')} href="/login">login</a>
            {/if}
        </nav>
    </div>
</div>

<style>
    .header {
        position: sticky;
        top: 0;
        z-index: 10;
        padding: var(--space-3) var(--space-5);
        transition: padding var(--dur-base) var(--ease);
    }
    .header.scrolled {
        padding: var(--space-2) var(--space-5);
    }

    /*
     * Single bar that *transitions* between plain and glass states.
     * Keeping the same DOM tree across states is what enables the smooth
     * backdrop-filter / background ramp; an {#if}…{:else} swap would unmount
     * one tree and mount the other, killing the animation.
     */
    .bar {
        max-width: 80rem;
        margin: 0 auto;
        padding: var(--space-3) var(--space-2);
        display: flex;
        align-items: center;
        justify-content: space-between;
        gap: var(--space-5);

        background: rgba(252, 253, 255, 0);
        border: 1px solid rgba(255, 255, 255, 0);
        border-radius: var(--radius-lg);
        backdrop-filter: blur(0px) saturate(100%);
        -webkit-backdrop-filter: blur(0px) saturate(100%);
        box-shadow: 0 0 0 0 rgba(40, 55, 95, 0),
                    inset 0 1px 0 rgba(255, 255, 255, 0);

        transition: background var(--dur-base) var(--ease),
                    border-color var(--dur-base) var(--ease),
                    backdrop-filter var(--dur-base) var(--ease),
                    -webkit-backdrop-filter var(--dur-base) var(--ease),
                    box-shadow var(--dur-base) var(--ease),
                    padding var(--dur-base) var(--ease);
    }

    .header.scrolled .bar {
        padding: var(--space-2) var(--space-4);
        background: rgba(252, 253, 255, 0.55);
        border-color: rgba(255, 255, 255, 0.65);
        backdrop-filter: blur(28px) saturate(130%);
        -webkit-backdrop-filter: blur(28px) saturate(130%);
        box-shadow: 0 22px 44px -18px rgba(40, 55, 95, 0.16),
                    inset 0 1px 0 rgba(255, 255, 255, 0.7);
    }

    .brand {
        font-family: var(--font-display);
        font-size: 1.15rem;
        color: var(--ink);
        text-decoration: none;
        letter-spacing: -0.012em;
    }

    .nav {
        display: flex;
        align-items: center;
        gap: var(--space-5);
        font-size: 0.78rem;
        font-weight: 500;
        color: var(--ink-2);
    }
    .nav a {
        text-decoration: none;
        color: inherit;
        position: relative;
        padding-bottom: 2px;
        transition: color var(--dur-base) var(--ease);
    }
    .nav a:hover { color: var(--ink); }
    .nav a.active {
        color: var(--ink);
    }
    .nav a.active::after {
        content: '';
        position: absolute;
        left: 0; right: 0; bottom: -8px; height: 1px;
        background: var(--accent);
        border-radius: 1px;
    }
    .nav a:not(.active)::after {
        content: '';
        position: absolute;
        left: 50%; right: 50%; bottom: -8px; height: 1px;
        background: var(--accent);
        border-radius: 1px;
        opacity: 0;
        transition: left var(--dur-base) var(--ease),
                    right var(--dur-base) var(--ease),
                    opacity var(--dur-base) var(--ease);
    }
    .nav a:not(.active):hover::after {
        left: 0;
        right: 0;
        opacity: 0.5;
    }

    .avatar {
        width: 28px;
        height: 28px;
        border-radius: var(--radius-pill);
        background: var(--accent);
        box-shadow: 0 0 0 1px rgba(255, 255, 255, 0.7);
        transition: transform var(--dur-base) var(--ease),
                    box-shadow var(--dur-base) var(--ease);
    }
    .avatar:hover {
        transform: scale(1.05);
        box-shadow: 0 0 0 1px rgba(255, 255, 255, 0.7),
                    0 6px 14px -4px var(--accent-shadow);
    }

    @media (max-width: 640px) {
        .bar { gap: var(--space-3); }
        .nav { gap: var(--space-3); font-size: 0.72rem; }
    }
</style>
