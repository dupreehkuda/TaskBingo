<script lang="ts">
    import Account from '../../routes/accountStore';
    import Stack from '$lib/ui/Stack.svelte';
    import Cluster from '$lib/ui/Cluster.svelte';
    import Paper from '$lib/ui/Paper.svelte';
    import Select from '$lib/ui/Select.svelte';
    import Button from '$lib/ui/Button.svelte';
    import { CreateGame } from '$lib/data/newGame';

    export let showModal: boolean;
    export let selectedPackID = '';
    export let selectedFriendID = '';

    let dialog: HTMLDialogElement;
    let exists = false;

    function close() {
        selectedPackID = '';
        selectedFriendID = '';
        exists = false;
        dialog?.close();
    }

    function checkIfExists(): boolean {
        for (const g of $Account.games) {
            if ((g.user1Id === selectedFriendID || g.user2Id === selectedFriendID) && g.status !== 3) {
                exists = true;
                return true;
            }
        }
        exists = false;
        return false;
    }

    function create(e: Event) {
        e.preventDefault();
        if (checkIfExists()) return;
        if (!selectedFriendID || !selectedPackID) return;
        CreateGame(selectedFriendID, selectedPackID);
        close();
    }

    $: if (dialog && showModal) dialog.showModal();
</script>

<dialog
    bind:this={dialog}
    on:close={() => (showModal = false)}
    on:click|self={close}
    aria-labelledby="gm-title"
>
    <Paper padding="lg">
        <Stack gap="m">
            <h2 id="gm-title">Create a game</h2>
            <form on:submit={create}>
                <Stack gap="m">
                    <Select
                        label="Friend"
                        bind:value={selectedFriendID}
                        items={$Account.friends.filter((f) => f.status === 3).map((f) => ({ value: f.userID, name: f.username }))}
                        on:change={checkIfExists}
                        required
                    />
                    <Select
                        label="Pack"
                        bind:value={selectedPackID}
                        items={$Account.likedPacks.map((p) => ({ value: p.id, name: p.pack.title }))}
                        required
                    />
                    {#if exists}
                        <span class="warn">An active game with this friend already exists.</span>
                    {/if}
                    <Cluster gap="s" justify="end">
                        <Button type="button" variant="ghost" on:click={close}>Cancel</Button>
                        <Button type="submit" variant="accent" disabled={exists || !selectedFriendID || !selectedPackID}>Create</Button>
                    </Cluster>
                </Stack>
            </form>
        </Stack>
    </Paper>
</dialog>

<style>
    dialog {
        max-width: 26rem;
        min-width: 18rem;
        width: calc(100% - 2rem);
        border: none;
        background: transparent;
        padding: 0;
        margin: auto;
    }
    dialog::backdrop {
        background: rgba(28, 37, 64, 0.35);
        backdrop-filter: blur(8px);
        -webkit-backdrop-filter: blur(8px);
    }
    dialog[open] { animation: zoom var(--dur-base) var(--ease); }

    h2 {
        font-family: var(--font-display);
        font-size: 1.25rem;
        font-weight: 400;
        letter-spacing: -0.012em;
        margin: 0;
    }

    .warn {
        font-size: 0.85rem;
        color: #c44e4e;
    }

    @keyframes zoom {
        from { transform: scale(0.96); opacity: 0; }
        to { transform: scale(1); opacity: 1; }
    }
</style>
