<script lang="ts">
    import CurrentGame from "../../routes/currentGame";
    import Account from "../../routes/accountStore";
	  import { onDestroy, onMount } from "svelte";
	  import type { Unsubscriber } from "svelte/store";

    let usersNumbers: Number[] = $CurrentGame.user1ID === $Account.userID ? $CurrentGame.user1Numbers : $CurrentGame.user2Numbers;
    export let numbers: Number[] = [];

    let unsubscribe: Unsubscriber;

    onMount(() => {
      unsubscribe = CurrentGame.subscribe(value => {
        usersNumbers = value.user1ID === $Account.userID ? value.user1Numbers : value.user2Numbers;
      });
    });

    onDestroy(() => {
      unsubscribe();
    });
  </script>
  

  <div class="rounded-rect-grid">
    {#each numbers as number (number)}
      {#if usersNumbers.includes(number)}
        <div class="grid-item"><p>{number}</p></div>
      {:else}
        <div class="grid-item empty"><p></p></div>
      {/if}
    {/each}
  </div>

  
  <style lang="postcss">
    @import 'tailwindcss/base';
    @import 'tailwindcss/components';
    @import 'tailwindcss/utilities';
  
    .rounded-rect-grid {
      display: grid;
      grid-template-columns: repeat(4, 1fr);
      grid-gap: 10px;
    }
  
    .grid-item {
      display: flex;
      align-items: center;
      justify-content: center;
      background-color: #db7676;
      width: 100%;
      padding-top: 100%;
      border-radius: 18%;
      position: relative;
    }

    .grid-item.empty {
      background-color: #f1eebc;
    }

    .grid-item p {
      margin: 0;
      text-align: center;
      color: white;
      font-weight: bold;
      font-size: x-large;
      align-self: center;
      position: absolute;
      top: 50%;           
      left: 50%;          
      transform: translate(-50%, -50%); 
    }
  </style>
  