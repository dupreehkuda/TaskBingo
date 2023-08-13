<script lang="ts">
    import { Timeline, TimelineItem } from 'flowbite-svelte';
</script>

<svelte:head>
    <link href="https://fonts.googleapis.com/css2?family=Josefin+Sans:wght@500&family=Unbounded&family=Yellowtail&display=swap" rel="stylesheet">
</svelte:head>


<title>About</title>
<div>
    <body>
        <h1>About</h1>
        <p>Made by <a class='link' href='https://t.me/dupreehkuda' target="_blank" rel="noopener noreferrer">@dupreehkuda</a></p>
        <p>Repo on <a class="link" href="https://github.com/dupreehkuda/taskbingo" target="_blank" rel="noopener noreferrer">github</a></p>
    </body>
</div>
<div class="text">
    <Timeline>
      <TimelineItem date="May 2020">
        <h3>Genesis</h3>
        <p class="text-base font-normal">
          In the throes of Covid lockdowns, when days melded into nights, my productivity levels were so low,
          they would make a sloth look like an overachiever. But then, amidst the haze of takeout containers
          and unwashed laundry, the concept of Task Bingo took its first breath.
        </p>
      </TimelineItem>
      <TimelineItem date="June 2020">
        <h3>Humble Beginnings</h3>
        <p class="text-base font-normal">
          Frankly, our initial version wasn't winning any beauty contests. It was a scrappy mess hastily pieced together.
          The graphics? Primitive. The UX? Don’t even ask. But it worked, and that's what mattered.
        </p>
      </TimelineItem>
      <TimelineItem date="August 2020">
        <h3>Hiatus</h3>
        <p class="mb-4 text-base font-normal">
          As quarantine faded, our lives began to resume some semblance of normality.
          Energized, productive, and even a tad athletic, we drifted away from our creation.
          The Task Bingo files were cast aside, lost in the chaos of other fleeting aspirations.
        </p>
      </TimelineItem>
      <TimelineItem date="December 2022">
        <h3>The Awakening</h3>
        <p class="mb-4 text-base font-normal">
          Fast forward a bit. Motivation? Plummeted. Enthusiasm? Non-existent.
          As the days blurred, a ghost of an idea whispered, "Remember Task Bingo?"
          It wasn't so much an epiphany, more like stubbing your toe and finding a lost toy.
          However, this time, the bar was set higher.
        </p>
      </TimelineItem>
      <TimelineItem date="January 2023">
        <h3>The Grind</h3>
        <p class="mb-4 text-base font-normal">
          Building the new version wasn't a walk in the park. More like trudging through a swamp with
          a backpack full of bricks. Between the coding hiccups and the occasional bout of good ol'
          procrastination, the journey had its fair share of bumps. But hey, nothing that a bit of elbow grease couldn't fix.
        </p>
      </TimelineItem>
      <TimelineItem date="August 2023">
        <h3 class='beta'>Endgame</h3>
        <p class="mb-4 text-base font-normal">
          Here we are, a tad later than anticipated - the beta version of taskbingo is now a reality.
          Still rough around the edges, maybe has a bug or twenty, but damn, isn't it a sight to behold?
          Plans for development are brewing, but for the moment, taskbingo is taking a short intermission. Stay tuned.
        </p>
      </TimelineItem>
    </Timeline>
</div>


<style>
    .link {
        font-family: Prompt;
        color: #ff7cb2;
        font-weight: 300;
    }

    .beta {
      color: #7dffc6;
    }

    .text {
        text-align: left;
        padding: 1em;
        max-width: 740px;
        margin: 0 auto;
        font-family: Prompt;
        color: #11202d;
    }

    h1 {
        font-family: Prompt;
        font-size: x-large;
    }

    p {
        padding-inline: 1em;
        font-family: Prompt;
        font-size: medium;
        font-weight: 200;
        color: #e9e8d5;
    }

    h3 {
      font-weight: 400;
      font-size: large;
      color: #ffffff;
      font-style: italic;
    }
</style>