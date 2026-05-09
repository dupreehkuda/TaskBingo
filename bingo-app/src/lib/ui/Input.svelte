<script lang="ts">
    /**
     * Input — text input with floating label. Label rests inside the input
     * until focus or value, then slides up.
     */
    export let label: string;
    export let value: string = '';
    export let type: 'text' | 'email' | 'password' | 'search' | 'url' = 'text';
    export let name: string | undefined = undefined;
    export let id: string | undefined = undefined;
    export let placeholder: string = '';
    export let required = false;
    export let disabled = false;
    export let error: string | undefined = undefined;
    export let hint: string | undefined = undefined;

    let focused = false;
    $: filled = value !== '' && value !== undefined && value !== null;
    $: lifted = focused || filled;
    $: inputId = id ?? `input-${name ?? Math.random().toString(36).slice(2, 8)}`;

    function onInput(e: Event) {
        const t = e.currentTarget as HTMLInputElement;
        value = t.value;
    }
</script>

<label class="field" class:has-error={error}>
    <span class="label" class:lifted>{label}{required ? ' *' : ''}</span>
    <input
        id={inputId}
        {type}
        {name}
        {placeholder}
        {required}
        {disabled}
        {value}
        aria-invalid={error ? 'true' : 'false'}
        aria-describedby={error || hint ? `${inputId}-msg` : undefined}
        on:input={onInput}
        on:focus={() => (focused = true)}
        on:blur={() => (focused = false)}
        on:keydown
        on:keyup
        on:change
    />
    {#if error}
        <span id="{inputId}-msg" class="msg msg--error">{error}</span>
    {:else if hint}
        <span id="{inputId}-msg" class="msg">{hint}</span>
    {/if}
</label>

<style>
    .field {
        display: flex;
        flex-direction: column;
        position: relative;
    }

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
                    font-size var(--dur-fast) var(--ease),
                    color var(--dur-fast) var(--ease);
    }

    .label.lifted {
        top: 0.4rem;
        transform: translateY(0);
        font-size: 0.7rem;
        color: var(--ink-3);
        letter-spacing: 0.04em;
        text-transform: uppercase;
    }

    input {
        background: var(--paper);
        border: 1px solid var(--hairline);
        border-radius: var(--radius-md);
        padding: 1.5rem 1rem 0.5rem;
        color: var(--ink);
        font-size: 0.95rem;
        min-height: 56px;
        transition: border-color var(--dur-fast) var(--ease),
                    box-shadow var(--dur-fast) var(--ease);
        outline: none;
    }

    input:focus {
        border-color: var(--accent-from);
        box-shadow: 0 0 0 3px var(--accent-ring);
    }

    .has-error input {
        border-color: #c44e4e;
    }
    .has-error input:focus {
        box-shadow: 0 0 0 3px rgba(196, 78, 78, 0.25);
    }

    .msg {
        font-size: 0.75rem;
        color: var(--ink-3);
        margin-top: 0.25rem;
        margin-left: 1rem;
    }
    .msg--error {
        color: #c44e4e;
    }

    input:disabled {
        opacity: 0.6;
        cursor: not-allowed;
    }
</style>
