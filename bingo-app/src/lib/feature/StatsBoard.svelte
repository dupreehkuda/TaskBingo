<script lang="ts">
    import { browser } from '$app/environment';
    import Stats, { type StatsResult } from '../../routes/statsStore';
    import { API_URL, WEB_URL } from '../../routes/temporary';
    import Cluster from '$lib/ui/Cluster.svelte';
    import TabBar from '$lib/ui/TabBar.svelte';
    import StatTile from './StatTile.svelte';
    import BarChart from './BarChart.svelte';
    import TopPacksList from './TopPacksList.svelte';

    let days: string = '7';
    let loading = false;

    async function fetchStats(d: number) {
        if (!browser) return;
        loading = true;
        try {
            const res = await fetch(`${API_URL}/api/user/stats?days=${d}`, {
                method: 'GET',
                headers: { Origin: WEB_URL },
                credentials: 'include',
            });
            if (res.ok) {
                const data: StatsResult = await res.json();
                Stats.set(data);
            }
        } finally {
            loading = false;
        }
    }

    $: if (browser) fetchStats(Number(days));

    function pct(part: number, whole: number) {
        if (whole === 0) return '—';
        return `${Math.round((part / whole) * 100)}%`;
    }
</script>

<section class="board" class:loading>
    <Cluster gap="m" align="baseline" justify="between">
        <h2 class="heading">Recent activity</h2>
        <TabBar
            bind:value={days}
            items={[
                { value: '7', label: '7 days' },
                { value: '30', label: '30 days' },
            ]}
            ariaLabel="Stats window"
        />
    </Cluster>

    {#if $Stats}
        {@const s = $Stats.summary}
        {@const totalDuoGames = s.wins + s.losses}
        <div class="tiles">
            <StatTile label="games" value={s.games} />
            <StatTile
                label="win rate"
                value={pct(s.wins, totalDuoGames)}
                hint={totalDuoGames > 0 ? `${s.wins} / ${totalDuoGames}` : undefined}
            />
            <StatTile label="bingos" value={s.totalBingo} hint={`solo ${s.soloBingo} · duo ${s.duoBingo}`} />
            <StatTile label="tasks closed" value={s.tasksClosed} />
        </div>

        <div class="charts">
            <BarChart buckets={$Stats.buckets} metric="bingo" title="Bingos" />
            <BarChart buckets={$Stats.buckets} metric="tasks" title="Tasks closed" />
        </div>

        <TopPacksList title="Most played packs" items={$Stats.topPacks} />

        {#if s.games === 0}
            <p class="empty">No games yet — your first board awaits.</p>
        {/if}
    {/if}
</section>

<style>
    .board {
        display: flex;
        flex-direction: column;
        gap: var(--space-4);
        transition: opacity var(--dur-base) var(--ease);
    }
    .board.loading { opacity: 0.6; }

    .heading {
        font-family: var(--font-display);
        font-size: 1.1rem;
        font-weight: 400;
        letter-spacing: -0.01em;
        color: var(--ink);
        margin: 0;
    }

    .tiles {
        display: grid;
        grid-template-columns: repeat(4, 1fr);
        gap: var(--space-3);
    }
    @media (max-width: 640px) {
        .tiles { grid-template-columns: repeat(2, 1fr); }
    }

    .charts {
        display: grid;
        grid-template-columns: 1fr 1fr;
        gap: var(--space-5);
    }
    @media (max-width: 640px) {
        .charts { grid-template-columns: 1fr; }
    }

    .empty {
        font-family: var(--font-display);
        font-style: italic;
        color: var(--ink-3);
        text-align: center;
        margin: var(--space-3) 0 0;
    }
</style>
