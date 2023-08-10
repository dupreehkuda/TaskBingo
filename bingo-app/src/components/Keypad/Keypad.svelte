<script>
    import { onMount } from 'svelte';
    import { Button, Label, Input } from 'flowbite-svelte';
    import { createEventDispatcher } from 'svelte';
  
    let numbers = [1, 2, 3, 4, 5, 6, 7, 8, 9, 0];
    let enteredCode = '';
  
    /**
	 * @param {string | number} number
	 */
    function handleClick(number) {
        enteredCode += number;
    }

    function handleDelete() {
        enteredCode = enteredCode.slice(0, -1);
    }

    const dispatch = createEventDispatcher();

    function handleSubmit() {
        dispatch('submit', Number(enteredCode));
        enteredCode = ''
    }
  </script>
  
  <main>
    <div class="flex grid grid-cols-5 place-content-stretch gap-1.5">
        <div class="col-span-3"><Input id='enteredCode' value={enteredCode} placeholder="" readonly /></div>
        <div>
            <Button on:click={() => handleDelete()}>
                <svg aria-hidden="true" class="w-5 h-5" fill="currentColor" viewBox="0 0 20 20" xmlns="http://www.w3.org/2000/svg">
                    <path fill-rule="evenodd" d="M14.707 5.293a1 1 0 010 1.414L11.414 10l3.293 3.293a1 1 0 01-1.414 1.414L10 11.414l-3.293 3.293a1 1 0 01-1.414-1.414L8.586 10 5.293 6.707a1 1 0 011.414-1.414L10 8.586l3.293-3.293a1 1 0 011.414 0z" clip-rule="evenodd"></path>
                </svg>
            </Button>
        </div>
        <div>
            <Button on:click={() => handleSubmit()}>
                <svg aria-hidden="true" class="w-5 h-5" fill="currentColor" viewBox="0 0 20 20" xmlns="http://www.w3.org/2000/svg">
                    <path fill-rule="evenodd" d="M10.293 3.293a1 1 0 011.414 0l6 6a1 1 0 010 1.414l-6 6a1 1 0 01-1.414-1.414L14.586 11H3a1 1 0 110-2h11.586l-4.293-4.293a1 1 0 010-1.414z" clip-rule="evenodd"></path>
                </svg>
            </Button>
        </div>
        {#each numbers as number}
            <div><Button class="w-full" on:click={() => handleClick(number)}>{number}</Button></div>
        {/each}
    </div>
</main>

<style>
  main {
    display: flex;
    flex-direction: column;
    align-items: center;
  }
</style>