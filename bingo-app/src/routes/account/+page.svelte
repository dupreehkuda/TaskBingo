<script lang="ts">
    import { Button, Card } from "flowbite-svelte";

    const data = {
        "name":"dupreeh",
        "wins":7,
        "lose":3,
        "scoreboard":1,
        "friends": [
            "zonedberg",
            "estepina",
            "hokage",
            "mrak",
        ],
        "packs": [
            {
                "name":"First Ever Pack",
                "tasks": [
                    "30s guitar solo",
                    "30 lines of code",
                    "Learn new dish",
                    "Room clean up",
                    "Topic in music theory",
                    "Read 5 chapters of fiction",
                    "200 push ups",
                    "Make complicated origami",
                    "Make a crossword",
                    "Practice math exam",
                    "Write an essay",
                    "30min sport",
                    "Topic in finances/economy",
                    "Make some art",
                    "Solve linguistic task",
                    "Learn by heart 30 geo objects",
                ]
            }
        ]
    }
</script>

<main>
    <div class="leftspace spacer">
        <!-- <h3 class="mb-1">Profile</h3> -->
        <h1>{data.name}
            <span class="leftspace win">{data.wins}</span>
            <span>:</span>
            <span class="lose">{data.lose}</span>
        </h1>
        <div class="leftspace">
            <span class="leftspace">Winrate: {(data.wins / (data.wins + data.lose)) * 100}%</span>
            <span>Scoreboard: {data.scoreboard}</span>
        </div>
    </div>

    <div class="leftspaceheading"><h3 class="mb-2">Friends</h3></div>
    <div class="scrolling-wrapper spacer">
        {#each data.friends as friend}
            <Card padding="sm" class="card">
                <div class="flex flex-col items-center pb-4">
                    <h5 class="mb-1">{friend}</h5>
                    <span class="text-sm text-gray-300">7/1</span>
                    <div class="flex mt-4 space-x-3 lg:mt-6">
                        <Button color="light" class="fonty dark:text-white">Play</Button>
                    </div>
                </div>
            </Card>
        {/each}
    </div>

    <div class="leftspaceheading"><h3 class="mb-2">Packs</h3></div>
    <div class="scrolling-wrapper spacer packs">
        {#each data.packs as pack}
          <Card size="md" padding="xl" class="card">
            <h5 class="mb-4 text-xl">{pack.name}</h5>
            <ul class="my-7 space-y-2">
                {#each pack.tasks as task, i}
                    <li class="flex space-x-2">
                        <span class="leading-tight">{i+1}</span>
                        <span class="leading-tight text-gray-300">{task}</span>
                    </li>
                {/each}
            <Button color="light" class="fonty  w-full">Choose pack</Button>
          </Card>
          {/each}
    </div>
</main>

<style>
    span {
        text-align: left;
        font-weight: 300;
    }

    h1 {
        text-align: left;
        font-family: Unbounded;
        font-weight: 400;
        margin-left: 0.35em;
        font-size: xx-large;
    }

    h3 {
        text-align: left;
        margin-left: 0.35em;
        font-family: Unbounded;
    }

    .scrolling-wrapper {
        display: flex;
        flex-wrap: nowrap;
    }

    .spacer {
        margin-bottom: 2em;
    }

    .win {
        font-weight: 300;
        color: #36e804;
    }

    .lose {
        font-weight: 300;
        color: #eb1a0f;
    }

    .leftspace {
        margin-left: 0.35em;
        text-align: left;
    }

    .leftspaceheading {
        margin-left: 0.7em;
        text-align: left;
    }

    .packs {
        max-width: min-content;
    }

    main {
        min-width: 50%;
        text-align: center;
        max-width: max-content;
        margin: 0 auto;
    }

    div {
        -ms-overflow-style: none; /* for Internet Explorer, Edge */
        scrollbar-width: none; /* for Firefox */
        overflow-y: scroll; 
    }

    div::-webkit-scrollbar {
        display: none; /* for Chrome, Safari, and Opera */
    }
</style>
