<script lang="ts">
    import Page from '$lib/ui/Page.svelte';
    import Cluster from '$lib/ui/Cluster.svelte';
    import Button from '$lib/ui/Button.svelte';
    import HeroDemo from '$lib/feature/HeroDemo.svelte';
    import { Auth } from '$lib/stores/auth';

    $: authed = $Auth;
</script>

<svelte:head>
    <title>taskbingo</title>
</svelte:head>

<Page width="normal">
    <div class="hero">
        <div class="hero__text">
            <span class="eyebrow">premise</span>
            <h1 class="display">Sixteen tasks.<br/><em>One quiet board.</em></h1>
            <p class="lead">
                Agree on sixteen tasks, mark them off as you go,
                and collect 4-in-a-row lines.
            </p>
            <p class="kicker">the odds are always in your favour.</p>
            <Cluster gap="m" align="center">
                {#if authed}
                    <Button variant="accent" size="lg" href="/account">Continue playing</Button>
                    <Button variant="quiet" href="/packs">browse packs</Button>
                {:else}
                    <Button variant="accent" size="lg" href="/register">Start a game</Button>
                    <Button variant="quiet" href="/login">log in</Button>
                {/if}
            </Cluster>
        </div>

        <div class="hero__demo">
            <HeroDemo />
        </div>
    </div>
</Page>

<style>
    .hero {
        display: grid;
        grid-template-columns: 1.1fr 0.9fr;
        gap: var(--space-6);
        align-items: center;
        min-height: calc(100vh - 12rem);
    }
    @media (max-width: 900px) {
        .hero {
            grid-template-columns: 1fr;
            gap: var(--space-7);
            min-height: 0;
        }
    }

    .hero__text {
        display: flex;
        flex-direction: column;
        gap: var(--space-4);
        max-width: 28rem;
    }

    .hero__demo {
        display: flex;
        justify-content: center;
        align-items: center;
    }

    .eyebrow {
        font-size: 0.65rem;
        letter-spacing: 0.18em;
        text-transform: uppercase;
        color: var(--ink-3);
    }

    .display {
        font-family: var(--font-display);
        font-weight: 400;
        font-size: clamp(2.2rem, 4vw, 2.8rem);
        line-height: 1.05;
        letter-spacing: -0.022em;
        color: var(--ink);
    }
    .display em { font-style: italic; color: var(--ink-3); font-weight: 300; }

    .lead {
        font-size: 0.95rem;
        line-height: 1.6;
        color: var(--ink-2);
        max-width: 36ch;
    }

    .kicker {
        font-family: var(--font-display);
        font-style: italic;
        font-weight: 400;
        font-size: 1.1rem;
        line-height: 1.3;
        letter-spacing: -0.01em;
        color: var(--accent-from);
        margin-top: calc(var(--space-1) * -1);
    }
</style>
