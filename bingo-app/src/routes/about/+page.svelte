<script lang="ts">
    import Page from '$lib/ui/Page.svelte';
    import Stack from '$lib/ui/Stack.svelte';

    type Entry = { date: string; title: string; body: string; isFinal?: boolean };

    const entries: Entry[] = [
        { date: 'May 2020', title: 'Genesis', body: "In the throes of Covid lockdowns, when days melded into nights, my productivity levels were so low, they would make a sloth look like an overachiever. But then, amidst the haze of takeout containers and unwashed laundry, the concept of Task Bingo took its first breath." },
        { date: 'June 2020', title: 'Humble Beginnings', body: "Frankly, our initial version wasn't winning any beauty contests. It was a scrappy mess hastily pieced together. The graphics? Primitive. The UX? Don't even ask. But it worked, and that's what mattered." },
        { date: 'August 2020', title: 'Hiatus', body: "As quarantine faded, our lives began to resume some semblance of normality. Energized, productive, and even a tad athletic, we drifted away from our creation. The Task Bingo files were cast aside, lost in the chaos of other fleeting aspirations." },
        { date: 'December 2022', title: 'The Awakening', body: "Fast forward a bit. Motivation? Plummeted. Enthusiasm? Non-existent. As the days blurred, a ghost of an idea whispered, \"Remember Task Bingo?\" It wasn't so much an epiphany, more like stubbing your toe and finding a lost toy. However, this time, the bar was set higher." },
        { date: 'January 2023', title: 'The Grind', body: "Building the new version wasn't a walk in the park. More like trudging through a swamp with a backpack full of bricks. Between the coding hiccups and the occasional bout of good ol' procrastination, the journey had its fair share of bumps. But hey, nothing that a bit of elbow grease couldn't fix." },
        { date: 'August 2023', title: 'Endgame', isFinal: true, body: "Here we are, a tad later than anticipated — the beta version of taskbingo is now a reality. Still rough around the edges, maybe has a bug or twenty, but damn, isn't it a sight to behold? Plans for development are brewing, but for the moment, taskbingo is taking a short intermission. Stay tuned." },
    ];
</script>

<svelte:head><title>About · taskbingo</title></svelte:head>

<Page width="narrow">
    <Stack gap="xl">
        <header class="head">
            <h1 class="display">About</h1>
            <p class="meta">
                Made by <a href="https://t.me/dupreehkuda" target="_blank" rel="noopener noreferrer">@dupreehkuda</a>
                · <a href="https://github.com/dupreehkuda/taskbingo" target="_blank" rel="noopener noreferrer">repo</a>
            </p>
        </header>

        <ol class="timeline">
            {#each entries as e (e.date)}
                <li class="entry">
                    <span class="date">{e.date}</span>
                    <h3 class="title" class:final={e.isFinal}>{e.title}</h3>
                    <p class="body">{e.body}</p>
                </li>
            {/each}
        </ol>
    </Stack>
</Page>

<style>
    .head { text-align: center; }
    .display {
        font-family: var(--font-display);
        font-size: 2rem;
        font-weight: 400;
        letter-spacing: -0.018em;
    }
    .meta { color: var(--ink-3); font-size: 0.85rem; margin-top: var(--space-2); }
    .meta a { color: var(--ink-2); border-bottom: 1px solid rgba(40, 55, 95, 0.18); text-decoration: none; }
    .meta a:hover { color: var(--ink); }

    .timeline {
        list-style: none;
        margin: 0; padding: 0;
        display: flex; flex-direction: column;
        gap: var(--space-7);
        border-left: 1px solid var(--hairline);
        padding-left: var(--space-5);
        position: relative;
    }
    .timeline::before {
        content: '';
        position: absolute;
        left: -1px;
        top: 0;
        bottom: 0;
        width: 1px;
        background: var(--accent);
        opacity: 0.4;
    }

    .entry { display: flex; flex-direction: column; gap: var(--space-2); }
    .date {
        font-family: var(--font-display);
        font-style: italic;
        font-weight: 300;
        font-size: 0.85rem;
        color: var(--ink-3);
    }
    .title {
        font-family: var(--font-display);
        font-weight: 400;
        font-size: 1.1rem;
        color: var(--ink);
        margin: 0;
    }
    .title.final {
        background: var(--accent);
        -webkit-background-clip: text;
        background-clip: text;
        color: transparent;
    }
    .body {
        font-size: 0.9rem;
        line-height: 1.6;
        color: var(--ink-2);
    }
</style>
