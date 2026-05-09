<script lang="ts">
    export let label: string;
    export let value: string = '';
    export let name: string | undefined = undefined;
    export let id: string | undefined = undefined;
    export let placeholder: string = '';
    export let rows = 3;
    export let required = false;
    export let disabled = false;
    export let error: string | undefined = undefined;

    let focused = false;
    let textareaEl: HTMLTextAreaElement | undefined;
    $: filled = value !== '' && value !== undefined && value !== null;
    $: lifted = focused || filled;
    $: areaId = id ?? `textarea-${name ?? Math.random().toString(36).slice(2, 8)}`;

    function onInput(e: Event) {
        const t = e.currentTarget as HTMLTextAreaElement;
        value = t.value;
    }

    /** Programmatic focus — used e.g. to keep cursor on the field after submit. */
    export function focus(): void {
        textareaEl?.focus();
    }
</script>

<label class="field" class:has-error={error}>
    <span class="label" class:lifted>{label}{required ? ' *' : ''}</span>
    <textarea
        bind:this={textareaEl}
        id={areaId}
        {name}
        {rows}
        {placeholder}
        {required}
        {disabled}
        {value}
        aria-invalid={error ? 'true' : 'false'}
        on:input={onInput}
        on:focus={() => (focused = true)}
        on:blur={() => (focused = false)}
    ></textarea>
    {#if error}<span class="msg msg--error">{error}</span>{/if}
</label>

<style>
    .field { display: flex; flex-direction: column; position: relative; }

    .label {
        position: absolute;
        left: 1rem;
        top: 1rem;
        color: var(--ink-3);
        font-size: 0.95rem;
        pointer-events: none;
        transition: top var(--dur-fast) var(--ease),
                    font-size var(--dur-fast) var(--ease);
    }
    .label.lifted {
        top: 0.4rem;
        font-size: 0.7rem;
        letter-spacing: 0.04em;
        text-transform: uppercase;
    }

    textarea {
        background: var(--paper);
        border: 1px solid var(--hairline);
        border-radius: var(--radius-md);
        padding: 1.5rem 1rem 0.75rem;
        color: var(--ink);
        font-size: 0.95rem;
        font-family: inherit;
        line-height: 1.5;
        resize: vertical;
        min-height: 96px;
        outline: none;
        transition: border-color var(--dur-fast) var(--ease),
                    box-shadow var(--dur-fast) var(--ease);
    }
    textarea:focus {
        border-color: var(--accent-from);
        box-shadow: 0 0 0 3px var(--accent-ring);
    }

    .has-error textarea { border-color: #c44e4e; }

    .msg {
        font-size: 0.75rem;
        color: var(--ink-3);
        margin: 0.25rem 0 0 1rem;
    }
    .msg--error { color: #c44e4e; }
</style>
