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
        <a class="link" href="https://github.com/dupreehkuda/taskbingo">Github</a>
    </body>
</div>
<div class="text">
    <Timeline>
      <TimelineItem customTimeClass="text" title="Rebirth" date="January 2023">
        <p class="mb-4 text-base font-normal">
          Yep. This is a perfect pet-project
        </p>
      </TimelineItem>
      <TimelineItem customTimeClass="text" title="Rebirth" date="December 2022">
        <p class="mb-4 text-base font-normal">
          Wait a minuite, this app can help a lot of people, maybe this is the perfect pet-project...
        </p>
      </TimelineItem>
      <TimelineItem title="End of task bingo" date="August 2020">
        <p class="mb-4 text-base font-normal">
          The quarantine ended. We were focused, productive, sporty and happy. 
          That was time to move on and files of original bingo was lost in abbys of other projects, plans and routine.
        </p>
      </TimelineItem>
      <TimelineItem title="First steps" date="June 2020">
        <p class="text-base font-normal">
          Very quirky version was built quick and dirty, 
          so two people could at least start playing and be more productive
        </p>
      </TimelineItem>
      <TimelineItem title="The birth of the idea" date="May 2020">
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
        color: #2ee6e3;
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
        font-weight: 300;
    }
</style>