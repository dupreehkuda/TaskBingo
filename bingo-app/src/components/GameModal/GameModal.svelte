<script lang="ts">
	import Account from "../../routes/accountStore";
	import { Button, Select } from "flowbite-svelte";
	import { CreateGame } from "./newGame";

	export let showModal: boolean;
	export let selectedPackID: string;
    export let selectedFriendID: string;

	let dialog: HTMLDialogElement; 

	function close() {
		selectedPackID = ""
		selectedFriendID = ""
		dialog.close()
	}

	function createGame() {
		CreateGame(selectedFriendID, selectedPackID)
		selectedPackID = ""
		selectedFriendID = ""
		dialog.close()
	}

	$: if (dialog && showModal) dialog.showModal();
</script>

<!-- svelte-ignore a11y-click-events-have-key-events a11y-no-noninteractive-element-interactions -->
<dialog
	bind:this={dialog}
	on:close={() => (showModal = false)}
	on:click|self={() => close()}
>
	<!-- svelte-ignore a11y-no-static-element-interactions -->
	<div class="flex flex-col" on:click|stopPropagation>
		<h5>Create new game</h5>

		<div class="container">
			<form on:submit={createGame}>
				<Select required class="mt-2" items={$Account.friends
				.filter(friend => friend.status === 3)
				.map((friend) => ({value: friend.userID, name: friend.username}))
				} bind:value={selectedFriendID} />

				<Select required class="mt-2" items={$Account.likedPacks.map((pack) => ({
					value: pack.id,
					name: pack.pack.title,
				}))} bind:value={selectedPackID} />

				<Button type="submit" class="fonty mt-3">Create</Button>
			  </form>
		</div>
	</div>
</dialog>

<style>
	dialog {
		max-width: 32em;
		min-width: 20em;
		border-radius: 10px; 
		border: none;
		padding: 0;
	}

	dialog::backdrop {
		background: rgba(0, 0, 0, 0.3);
	}

	dialog > div {
		padding: 1em;
	}

	dialog[open] {
		animation: zoom 0.3s cubic-bezier(0.34, 1.56, 0.64, 1);
	}

	@keyframes zoom {
		from {
			transform: scale(0.95);
		}
		to {
			transform: scale(1);
		}
	}

	dialog[open]::backdrop {
		animation: fade 0.2s ease-out;
	}

	@keyframes fade {
		from {
			opacity: 0;
		}
		to {
			opacity: 1;
		}
	}

	h5 {
        font-weight: 400;
        font-size: large;
    }
</style>
