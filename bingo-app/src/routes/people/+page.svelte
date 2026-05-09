<script lang="ts">
    import type { PageData } from './$types';
    import Account from '../accountStore';
    import Page from '$lib/ui/Page.svelte';
    import Stack from '$lib/ui/Stack.svelte';
    import Cluster from '$lib/ui/Cluster.svelte';
    import Paper from '$lib/ui/Paper.svelte';
    import Input from '$lib/ui/Input.svelte';
    import Button from '$lib/ui/Button.svelte';
    import GameModal from '$lib/feature/GameModal.svelte';
    import { DeleteFriend, RequestFriend, AcceptFriend } from '../friendRequests';

    export let data: PageData;
    const { users } = data;

    let showModal = false;
    let selectedFriendID = '';
    let selectedPackID = '';
    let query = '';

    $: filtered = (users ?? []).filter((u: { userID: string; username: string; city?: string; bingo: number }) => {
        const q = query.trim().toLowerCase();
        if (!q) return true;
        return u.username.toLowerCase().includes(q) || (u.city ?? '').toLowerCase().includes(q);
    });

    function statusOf(userID: string) {
        return $Account?.friends.find((f) => f.userID === userID)?.status ?? 0;
    }
</script>

<svelte:head><title>People · taskbingo</title></svelte:head>

<Page width="normal">
    <Stack gap="l">
        <h1 class="title">People</h1>
        <Input label="Search by username or city" bind:value={query} type="search" />
        <Stack gap="s">
            {#each filtered as user (user.userID)}
                {@const s = statusOf(user.userID)}
                <Paper padding="sm">
                    <Cluster gap="m" align="center" justify="between">
                        <Stack gap="s">
                            <Cluster gap="s" align="baseline">
                                <span class="bingo">{user.bingo}</span>
                                <span class="username">{user.username}</span>
                            </Cluster>
                            <span class="city">{user.city}</span>
                        </Stack>
                        <Cluster gap="s">
                            {#if user.userID === $Account?.userID}
                                <Button variant="ghost" size="sm" href="/account">Account</Button>
                            {:else if s === 3}
                                <Button variant="accent" size="sm" on:click={() => { selectedFriendID = user.userID; selectedPackID = ''; showModal = true; }}>Play</Button>
                                <Button variant="ghost" size="sm" on:click={() => DeleteFriend(user.userID)}>Remove</Button>
                            {:else if s === 1}
                                <Button variant="ghost" size="sm" disabled>Sent</Button>
                                <Button variant="ghost" size="sm" on:click={() => DeleteFriend(user.userID)}>Cancel</Button>
                            {:else if s === 2}
                                <Button variant="accent" size="sm" on:click={() => AcceptFriend(user.userID)}>Accept</Button>
                                <Button variant="ghost" size="sm" on:click={() => DeleteFriend(user.userID)}>Decline</Button>
                            {:else}
                                <Button variant="accent" size="sm" on:click={() => RequestFriend(user.userID, user.username)}>Add friend</Button>
                            {/if}
                        </Cluster>
                    </Cluster>
                </Paper>
            {/each}
            {#if filtered.length === 0}
                <Paper padding="md">
                    <span class="empty">
                        {query ? 'No matches.' : 'No people yet.'}
                    </span>
                </Paper>
            {/if}
        </Stack>
    </Stack>

    <GameModal bind:showModal {selectedFriendID} {selectedPackID} />
</Page>

<style>
    .title {
        font-family: var(--font-display);
        font-size: 1.6rem;
        font-weight: 400;
        letter-spacing: -0.015em;
    }
    .username { font-weight: 500; color: var(--ink); }
    .bingo {
        font-family: var(--font-display);
        font-size: 1.05rem;
        color: var(--ink);
        min-width: 1.5em;
        text-align: right;
    }
    .city { font-size: 0.78rem; color: var(--ink-3); font-style: italic; }
    .empty { color: var(--ink-2); font-size: 0.9rem; font-style: italic; }
</style>
