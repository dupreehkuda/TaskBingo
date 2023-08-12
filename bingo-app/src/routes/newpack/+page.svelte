<script lang="ts">
    import { Label, Input, Button } from "flowbite-svelte";
    import { _Submit } from "./+page";
    import { goto } from '$app/navigation';

    function redirectOnAccount() {
        goto('/account');
    }

    async function submit(event: any) {
        const status = await _Submit(event)
        if (status == 200) { redirectOnAccount }
    }
</script>

<svelte:head>
    <title>New Task Pack</title>
</svelte:head>

<body>
    <main>
        <form on:submit={submit}>
            <div class="mb-4">
                <Label color="disabled" for="input-group-1" class="block mb-2">Task Name</Label>
                <Input label="Task Name" id="name" name="name" required/>
            </div>

            <div class="flex flex-col mb-4">
                <Label color="disabled" for="input-group-1" class="block mb-2">Tasks</Label>
                <Input size="sm" defaultClass="spacer05" label="Task 1" id="task1" name="task1" required/>
                <Input size="sm" defaultClass="spacer05" label="Task 2" id="task2" name="task2" required/>
                <Input size="sm" defaultClass="spacer05" label="Task 3" id="task3" name="task3" required/>
                <Input size="sm" defaultClass="spacer05" label="Task 4" id="task4" name="task4" required/>
                <Input size="sm" defaultClass="spacer05" label="Task 5" id="task5" name="task5" required/>
                <Input size="sm" defaultClass="spacer05" label="Task 6" id="task6" name="task6" required/>
                <Input size="sm" defaultClass="spacer05" label="Task 7" id="task7" name="task7" required/>
                <Input size="sm" defaultClass="spacer05" label="Task 8" id="task8" name="task8" required/>
                <Input size="sm" defaultClass="spacer05" label="Task 9" id="task9" name="task9" required/>
                <Input size="sm" defaultClass="spacer05" label="Task 10" id="task10" name="task10" required/>
                <Input size="sm" defaultClass="spacer05" label="Task 11" id="task11" name="task11" required/>
                <Input size="sm" defaultClass="spacer05" label="Task 12" id="task12" name="task12" required/>
                <Input size="sm" defaultClass="spacer05" label="Task 13" id="task13" name="task13" required/>
                <Input size="sm" defaultClass="spacer05" label="Task 14" id="task14" name="task14" required/>
                <Input size="sm" defaultClass="spacer05" label="Task 15" id="task15" name="task15" required/>
                <Input size="sm" defaultClass="spacer05" label="Task 16" id="task16" name="task16" required/>
            </div>

            <Button class="mb-3" type="submit" color="light">Create pack</Button>
        </form>
    </main>
</body>


<style>
    main {
        text-align: center;
        padding: 1em;
        max-width: 440px;
        margin: 0 auto;
    }
</style>