<script lang="ts">
    import { Button, Card } from "flowbite-svelte";
</script>

<main>
    <div class="leftspace spacer">
        <!-- <h3 class="mb-1">Profile</h3> -->
        <h1>dupreeh
            <span class="leftspace win">7</span>
            <span>:</span>
            <span class="lose">3</span>
        </h1>
        <div class="leftspace">
            <span class="leftspace">Winrate: 70%</span>
            <span>Scoreboard: 1</span>
        </div>
    </div>

    <div class="leftspaceheading"><h3 class="mb-2">Friends</h3></div>
    <div class="scrolling-wrapper spacer">
        <Card padding="sm" class="card">
            <div class="flex flex-col items-center pb-4">
                <h5 class="mb-1">zonedberg</h5>
                <span class="text-sm text-gray-300">7/1</span>
                <div class="flex mt-4 space-x-3 lg:mt-6">
                    <Button color="light" class="fonty dark:text-white">Play</Button>
                </div>
            </div>
        </Card>
  
        <Card padding="sm" class="card">
            <div class="flex flex-col items-center pb-4">
                <h5 class="mb-1">estepina</h5>
                <span class="text-sm text-gray-300">1/5</span>
                <div class="flex mt-4 space-x-3 lg:mt-6">
                    <Button color="light" class="fonty dark:text-white">Play</Button>
                </div>
            </div>
        </Card>   

        <Card padding="sm" class="card">
            <div class="flex flex-col items-center pb-4">
                <h5 class="mb-1">dupreeh</h5>
                <span class="text-sm text-gray-300">2/1</span>
                <div class="flex mt-4 space-x-3 lg:mt-6">
                    <Button color="light" class="fonty dark:text-white">Play</Button>
                </div>
            </div>
        </Card>
        
        <Card padding="sm" class="card">
            <div class="flex flex-col items-center pb-4">
                <h5 class="mb-1">zonedberg</h5>
                <span class="text-sm text-gray-300">7/1</span>
                <div class="flex mt-4 space-x-3 lg:mt-6">
                    <Button color="light" class="fonty dark:text-white">Play</Button>
                </div>
            </div>
        </Card>
  
        <Card padding="sm" class="card">
            <div class="flex flex-col items-center pb-4">
                <h5 class="mb-1">estepina</h5>
                <span class="text-sm text-gray-300">1/5</span>
                <div class="flex mt-4 space-x-3 lg:mt-6">
                    <Button color="light" class="fonty dark:text-white">Play</Button>
                </div>
            </div>
        </Card>   

        <Card padding="sm" class="card">
            <div class="flex flex-col items-center pb-4">
                <h5 class="mb-1">dupreeh</h5>
                <span class="text-sm text-gray-300">2/1</span>
                <div class="flex mt-4 space-x-3 lg:mt-6">
                    <Button color="light" class="fonty dark:text-white">Play</Button>
                </div>
            </div>
        </Card>
        
    </div>

    <div class="leftspaceheading"><h3 class="mb-2">Packs</h3></div>
    <div class="scrolling-wrapper spacer packs">
          <Card size="md" padding="xl" class="card">
            <h5 class="mb-4 text-xl">First Ever Edition</h5>
            <ul class="my-7 space-y-2">
                <li class="flex space-x-2">
                    <span class="leading-tight">1</span>
                    <span class="leading-tight text-gray-300">30s guitar solo</span>
                </li>
                <li class="flex space-x-2">
                    <span class="leading-tight">2</span>
                    <span class="leading-tight text-gray-300">30 lines of code</span>
                </li>
                <li class="flex space-x-2">
                    <span class="leading-tight">3</span>
                    <span class="leading-tight text-gray-300">Learn new dish</span>
                </li>
                <li class="flex space-x-2">
                    <span class="leading-tight">4</span>
                    <span class="leading-tight text-gray-300">Room clean up</span>
                </li>
                <li class="flex space-x-2">
                    <span class="leading-tight">5</span>
                    <span class="leading-tight text-gray-300">Topic in music theory</span>
                </li>
                <li class="flex space-x-2">
                    <span class="leading-tight">6</span>
                    <span class="leading-tight text-gray-300">Read 5 chapters of fiction</span>
                </li>
                <li class="flex space-x-2">
                    <span class="leading-tight">7</span>
                    <span class="leading-tight text-gray-300">200 push ups</span>
                </li>
                <li class="flex space-x-2">
                    <span class="leading-tight">8</span>
                    <span class="leading-tight text-gray-300">Make complicated origami</span>
                </li>
                <li class="flex space-x-2">
                    <span class="leading-tight">9</span>
                    <span class="leading-tight text-gray-300">Make a crossword</span>
                </li>
                <li class="flex space-x-2">
                    <span class="leading-tight">10</span>
                    <span class="leading-tight text-gray-300">Practice math exam</span>
                </li>
                <li class="flex space-x-2">
                    <span class="leading-tight">11</span>
                    <span class="leading-tight text-gray-300">Write an essay</span>
                </li>
                <li class="flex space-x-2">
                    <span class="leading-tight">12</span>
                    <span class="leading-tight text-gray-300">30min sport</span>
                </li>
                <li class="flex space-x-2">
                    <span class="leading-tight">13</span>
                    <span class="leading-tight text-gray-300">Topic in finances/economy</span>
                </li>
                <li class="flex space-x-2">
                    <span class="leading-tight">14</span>
                    <span class="leading-tight text-gray-300">Make some art</span>
                </li>
                <li class="flex space-x-2">
                    <span class="leading-tight">15</span>
                    <span class="leading-tight text-gray-300">Solve linguistic task</span>
                </li>
                <li class="flex space-x-2">
                    <span class="leading-tight">16</span>
                    <span class="leading-tight text-gray-300">Learn by heart 30 geo objects</span>
                </li>
            </ul>
            <Button color="light" class="fonty  w-full">Choose pack</Button>
          </Card>
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
