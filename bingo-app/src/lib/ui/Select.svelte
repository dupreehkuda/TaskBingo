<script lang="ts">
    /**
     * Select — paper-styled native <select> with floating label and chevron.
     * Sticks to native semantics for free a11y and mobile UX.
     */
    export let label: string;
    export let value: string = '';
    export let items: { value: string; name: string }[] = [];
    export let name: string | undefined = undefined;
    export let id: string | undefined = undefined;
    export let required = false;
    export let disabled = false;

    let focused = false;
    $: filled = value !== '' && value !== undefined;
    $: lifted = focused || filled;
    $: selectId = id ?? `select-${name ?? Math.random().toString(36).slice(2, 8)}`;
</script>

<label class="field">
    <span class="label" class:lifted>{label}{required ? ' *' : ''}</span>
    <select
        id={selectId}
        {name}
        {required}
        {disabled}
        bind:value
        on:focus={() => (focused = true)}
        on:blur={() => (focused = false)}
        on:change
    >
        <option value="" disabled hidden></option>
        {#each items as item (item.value)}
            <option value={item.value}>{item.name}</option>
        {/each}
    </select>
    <svg class="chevron" viewBox="0 0 20 20" aria-hidden="true">
        <path d="M5 8l5 5 5-5" fill="none" stroke="currentColor" stroke-width="1.6" stroke-linecap="round" stroke-linejoin="round"/>
    </svg>
</label>

<style>
    .field { position: relative; display: flex; flex-direction: column; }

    .label {
        position: absolute;
        left: 1rem;
        top: 50%;
        transform: translateY(-50%);
        color: var(--ink-3);
        font-size: 0.95rem;
        pointer-events: none;
        transition: top var(--dur-fast) var(--ease),
                    transform var(--dur-fast) var(--ease),
                    font-size var(--dur-fast) var(--ease);
    }
    .label.lifted {
        top: 0.4rem;
        transform: translateY(0);
        font-size: 0.7rem;
        text-transform: uppercase;
        letter-spacing: 0.04em;
    }

    select {
        appearance: none;
        -webkit-appearance: none;
        background: var(--paper);
        border: 1px solid var(--hairline);
        border-radius: var(--radius-md);
        padding: 1.5rem 2.4rem 0.5rem 1rem;
        color: var(--ink);
        font-size: 0.95rem;
        min-height: 56px;
        outline: none;
        cursor: pointer;
        transition: border-color var(--dur-fast) var(--ease),
                    box-shadow var(--dur-fast) var(--ease);
    }
    select:focus {
        border-color: var(--accent-from);
        box-shadow: 0 0 0 3px var(--accent-ring);
    }

    .chevron {
        position: absolute;
        right: 1rem;
        top: 50%;
        transform: translateY(-50%);
        width: 18px;
        height: 18px;
        color: var(--ink-3);
        pointer-events: none;
    }
</style>
