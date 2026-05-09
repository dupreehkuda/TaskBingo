<script lang="ts">
    /**
     * BingoCell — single 4×4 board cell. Shared between the live game grid
     * (BingoGrid) and the home-page demo (HeroDemo) so the visual states
     * (empty, marked, winning) stay identical across both surfaces.
     *
     * - Empty: paper-soft tile, faint dot in --ink-4.
     * - Marked (on): periwinkle accent gradient, white numerals, pop-in.
     * - Winning (on + win): sage gradient, dark numerals, soft halo.
     *
     * The component owns its own enter animations. Parents control layout.
     */
    export let n: number = 0;
    export let on: boolean = false;
    export let win: boolean = false;
</script>

<div class="cell" class:on class:win>
    <span class="num">{on ? n : ''}</span>
</div>

<style>
    .cell {
        aspect-ratio: 1;
        border-radius: var(--radius-md);
        display: flex;
        align-items: center;
        justify-content: center;
        background: var(--paper-soft);
        color: var(--ink-4);
        position: relative;
        box-shadow: 0 0 0 0 transparent;
        transform: scale(1);
        transition: background 320ms var(--ease),
                    color 320ms var(--ease),
                    box-shadow 480ms var(--ease),
                    transform 320ms var(--ease);
    }

    .cell.on {
        background: var(--accent);
        color: var(--on-accent);
        box-shadow: 0 4px 10px -3px var(--accent-shadow),
                    inset 0 1px 0 rgba(255, 255, 255, 0.25);
        animation: pop 360ms var(--ease) both;
    }

    /* Winning-line cell: gradient + ink text + soft outer halo. */
    .cell.win {
        background: var(--win);
        color: var(--ink);
        box-shadow: 0 4px 10px -3px var(--win-shadow),
                    0 0 0 2px var(--win-ring),
                    inset 0 1px 0 rgba(255, 255, 255, 0.35);
        animation: glow 700ms var(--ease) forwards;
    }

    .num {
        font-family: var(--font-display);
        font-weight: 600;
        font-size: clamp(0.95rem, 2.4vw, 1.35rem);
        line-height: 1;
        font-variant-numeric: tabular-nums;
        opacity: 0;
        transition: opacity 220ms var(--ease) 80ms;
    }
    .cell.on .num { opacity: 1; }

    @keyframes pop {
        0%   { transform: scale(0.92); }
        55%  { transform: scale(1.03); }
        100% { transform: scale(1); }
    }

    @keyframes glow {
        0% {
            box-shadow: 0 4px 10px -3px var(--win-shadow),
                        0 0 0 0 transparent,
                        inset 0 1px 0 rgba(255, 255, 255, 0.35);
        }
        55% {
            box-shadow: 0 4px 10px -3px var(--win-shadow),
                        0 0 0 8px var(--win-ring-soft),
                        inset 0 1px 0 rgba(255, 255, 255, 0.35);
        }
        100% {
            box-shadow: 0 4px 10px -3px var(--win-shadow),
                        0 0 0 2px var(--win-ring),
                        inset 0 1px 0 rgba(255, 255, 255, 0.35);
        }
    }

    @media (prefers-reduced-motion: reduce) {
        .cell, .cell.on, .cell.win { animation: none !important; transition: none; }
        .num { opacity: 1; transition: none; }
        .cell.on .num { opacity: 1; }
    }
</style>
