<script lang="ts">
    import type { StatsBucket } from '../../routes/statsStore';

    /**
     * Spark line chart with numeric labels above each data point. Replaces the
     * previous bar chart that looked nice but was hard to read at a glance —
     * users couldn't tell whether a bar meant 1 or 5 bingos.
     *
     * `metric` collapses solo+duo bingos into a single value per bucket; the
     * tile row already shows the breakdown, so we keep this curve uncluttered.
     */
    export let buckets: StatsBucket[] = [];
    export let metric: 'bingo' | 'tasks' = 'bingo';
    export let title: string;

    const W = 480;
    const H = 100;
    const PAD_X = 22;
    const PAD_TOP = 22;
    const PAD_BOT = 18;

    $: values = buckets.map((b) => (metric === 'bingo' ? b.soloBingo + b.duoBingo : b.tasksClosed));
    $: max = Math.max(1, ...values);

    $: points = values.map((v, i) => {
        const innerW = W - 2 * PAD_X;
        const innerH = H - PAD_TOP - PAD_BOT;
        const x =
            values.length <= 1
                ? PAD_X + innerW / 2
                : PAD_X + (i / (values.length - 1)) * innerW;
        const y = PAD_TOP + (1 - v / max) * innerH;
        return { x, y, value: v, label: buckets[i]?.label ?? '' };
    });

    function smoothPath(pts: { x: number; y: number }[]): string {
        if (pts.length === 0) return '';
        if (pts.length === 1) return `M ${pts[0].x} ${pts[0].y}`;
        let d = `M ${pts[0].x} ${pts[0].y}`;
        for (let i = 1; i < pts.length; i++) {
            const p0 = pts[i - 1];
            const p1 = pts[i];
            const cpx = (p0.x + p1.x) / 2;
            d += ` C ${cpx} ${p0.y}, ${cpx} ${p1.y}, ${p1.x} ${p1.y}`;
        }
        return d;
    }

    $: linePath = smoothPath(points);
    $: areaPath = points.length
        ? `${linePath} L ${points[points.length - 1].x} ${H - PAD_BOT} L ${points[0].x} ${H - PAD_BOT} Z`
        : '';

    // Unique gradient id per instance so two charts on the same page don't clash.
    const gradientId = `spark-fill-${Math.random().toString(36).slice(2, 8)}`;
</script>

<section class="chart-section">
    <h3 class="title">{title}</h3>

    <svg
        class="chart"
        viewBox="0 0 {W} {H}"
        preserveAspectRatio="xMidYMid meet"
        role="img"
        aria-label={title}
    >
        <defs>
            <linearGradient id={gradientId} x1="0" x2="0" y1="0" y2="1">
                <stop offset="0%" stop-color="var(--accent-from)" stop-opacity="0.32" />
                <stop offset="100%" stop-color="var(--accent-from)" stop-opacity="0" />
            </linearGradient>
        </defs>

        {#if points.length > 0}
            <path class="area" d={areaPath} fill="url(#{gradientId})" />
            <path
                class="line"
                d={linePath}
                fill="none"
                stroke="var(--accent-from)"
                stroke-width="1.6"
                stroke-linecap="round"
                stroke-linejoin="round"
            />
            {#each points as p (p.x + ':' + p.label)}
                <circle class="dot" cx={p.x} cy={p.y} r="3" fill="var(--paper)" stroke="var(--accent-from)" stroke-width="1.5" />
                <text x={p.x} y={p.y - 7} text-anchor="middle" class="value">
                    {p.value === 0 ? '·' : p.value}
                </text>
                <text x={p.x} y={H - 4} text-anchor="middle" class="tick">{p.label}</text>
            {/each}
        {/if}
    </svg>
</section>

<style>
    .chart-section {
        display: flex;
        flex-direction: column;
        gap: var(--space-2);
    }

    .title {
        font-family: var(--font-display);
        font-size: 1rem;
        font-weight: 400;
        letter-spacing: -0.01em;
        color: var(--ink);
        margin: 0;
    }

    .chart {
        width: 100%;
        height: auto;
        display: block;
        overflow: visible;
    }

    .area {
        animation: fade-up 480ms var(--ease) both;
    }

    .line {
        stroke-dasharray: 800;
        stroke-dashoffset: 800;
        animation: draw 600ms var(--ease) forwards;
    }

    .dot {
        opacity: 0;
        animation: pop 280ms var(--ease) forwards;
    }
    /* Stagger the points along the curve. */
    .dot:nth-child(3) { animation-delay: 80ms; }
    .dot:nth-child(5) { animation-delay: 140ms; }
    .dot:nth-child(7) { animation-delay: 200ms; }
    .dot:nth-child(9) { animation-delay: 260ms; }
    .dot:nth-child(11) { animation-delay: 320ms; }
    .dot:nth-child(13) { animation-delay: 380ms; }
    .dot:nth-child(15) { animation-delay: 440ms; }

    .value {
        font-family: var(--font-display);
        font-size: 10px;
        font-weight: 500;
        fill: var(--ink);
        font-variant-numeric: tabular-nums;
        letter-spacing: -0.01em;
    }

    .tick {
        font-family: var(--font-body);
        font-size: 7.5px;
        fill: var(--ink-3);
        letter-spacing: 0.04em;
        text-transform: uppercase;
    }

    @keyframes draw {
        to { stroke-dashoffset: 0; }
    }

    @keyframes fade-up {
        from { opacity: 0; transform: translateY(6px); }
        to   { opacity: 1; transform: translateY(0); }
    }

    @keyframes pop {
        from { opacity: 0; transform: scale(0.6); transform-origin: center; }
        to   { opacity: 1; transform: scale(1); }
    }

    @media (prefers-reduced-motion: reduce) {
        .area, .line, .dot { animation: none; }
        .line { stroke-dashoffset: 0; }
        .dot { opacity: 1; }
    }
</style>
