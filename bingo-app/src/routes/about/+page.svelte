<script>
    import { Timeline, TimelineItem } from 'flowbite-svelte';
</script>

<svelte:head>
    <link href="https://fonts.googleapis.com/css2?family=Josefin+Sans:wght@500&family=Unbounded&family=Yellowtail&display=swap" rel="stylesheet">
</svelte:head>


<title>About</title>
<div>
    <body>
        <h1>About</h1>
        <p>Made by @dupreehkuda</p>
        <p>Repo on <a class="link" href="https://github.com/dupreehkuda/taskbingo">Github</a></p>
    </body>
</div>
<div class="text">
    <Timeline>
      <TimelineItem date="January 2023">
        <h3>New Era</h3>
        <p class="mb-4 text-base font-normal">
          Yep. This is a perfect pet-project to work with SvelteKit, gRPC, Redis, microservices and also help
          some people to be more productive and happy.
        </p>
      </TimelineItem>
      <TimelineItem date="December 2022">
        <h3>Rebirth</h3>
        <p class="mb-4 text-base font-normal">
          Wait a minuite, this app can help a lot of people, maybe I can make a pet-project out of it to learn nowadays
          industry technologies, concepts, standarts...
        </p>
      </TimelineItem>
      <TimelineItem date="August 2020">
        <h3>End of task bingo</h3>
        <p class="mb-4 text-base font-normal">
          The quarantine ended. We were focused, productive, sporty and happy. 
          That was time to move on and files of the original bingo was lost in abbys of other projects, plans and routine.
        </p>
      </TimelineItem>
      <TimelineItem date="June 2020">
        <h3>First steps</h3>
        <p class="text-base font-normal">
          Very quirky version was built quick and dirty, 
          so two people could at least start playing and be more productive.
        </p>
      </TimelineItem>
      <TimelineItem date="May 2020">
        <h3>The birth of the idea</h3>
        <p class="text-base font-normal">
          Covid lockdowns. Decreased productivity. Feeling depressed.
          The idea of task bingo was born and implemented then. Yeah.
        </p>
      </TimelineItem>
    </Timeline>
</div>


<style>
    .link {
        font-family: Unbounded;
        color: #23ebe7;
        font-weight: 400;
    }

    .text {
        text-align: left;
        padding: 1em;
        max-width: 740px;
        margin: 0 auto;
        font-family: Unbounded;
        color: #EEEDC5;
    }
    h1 {
        font-family: Unbounded;
        font-size: x-large;
    }

    p {
        padding-inline: 1em;
        font-family: Unbounded;
        font-size: medium;
        font-weight: 200;
        color: aliceblue;
    }

    h3 {
      font-weight: 400;
      color: aliceblue;
    }
</style>