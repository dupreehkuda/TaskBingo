<script lang="ts">
    /**
     * Toggle — pill switch. Uses native <input type="checkbox"> under
     * the hood so forms read it as a normal field.
     */
    export let label: string;
    export let checked = false;
    export let name: string | undefined = undefined;
    export let disabled = false;
</script>

<label class="toggle" class:disabled>
    <input type="checkbox" {name} bind:checked {disabled} on:change />
    <span class="track" aria-hidden="true">
        <span class="thumb"></span>
    </span>
    <span class="label-text">{label}</span>
</label>

<style>
    .toggle {
        display: inline-flex;
        align-items: center;
        gap: var(--space-3);
        cursor: pointer;
        user-select: none;
    }
    .toggle.disabled { cursor: not-allowed; opacity: 0.5; }

    input {
        position: absolute;
        opacity: 0;
        pointer-events: none;
    }

    .track {
        position: relative;
        width: 44px;
        height: 26px;
        border-radius: var(--radius-pill);
        background: var(--paper-soft);
        border: 1px solid var(--hairline);
        transition: background var(--dur-fast) var(--ease);
    }

    .thumb {
        position: absolute;
        left: 2px;
        top: 2px;
        width: 20px;
        height: 20px;
        border-radius: 999px;
        background: white;
        box-shadow: 0 1px 2px rgba(40, 55, 95, 0.18);
        transition: transform var(--dur-fast) var(--ease);
    }

    input:checked + .track {
        background: var(--accent);
        border-color: transparent;
    }
    input:checked + .track .thumb {
        transform: translateX(18px);
    }

    input:focus-visible + .track {
        outline: 2px solid var(--accent-ring);
        outline-offset: 2px;
    }

    .label-text {
        font-size: 0.9rem;
        color: var(--ink-2);
    }
</style>
