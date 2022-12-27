<script>
    import Header from "../Header.svelte";
    import { Timeline, TimelineItem } from 'flowbite-svelte';
</script>

<svelte:head>
    <link href="https://fonts.googleapis.com/css2?family=Josefin+Sans:wght@500&family=Unbounded&family=Yellowtail&display=swap" rel="stylesheet">
</svelte:head>


<title>About</title>
<div>
    <body>
        <h1>About</h1>
        <p>Task bingo is a game of bingo, but every number matches a task from the list.</p>
        <p>Just choose a task pack or create your own and play with your friend!</p>
    </body>
</div>
<div class="text">
    <Timeline>
        <TimelineItem customTimeClass="text" title="Application UI code in Tailwind CSS" date="February 2022">
          <p class="mb-4 text-base font-normal">
            Get access to over 20+ pages including a dashboard layout, charts, kanban board, calendar,
            and pre-order E-commerce & Marketing pages.
          </p>
        </TimelineItem>
        <TimelineItem title="Application UI code in Tailwind CSS" date="March 2022">
          <p class="text-base font-normal">
            All of the pages and components are first designed in Figma and we keep a parity between the
            two versions even as we update the project.
          </p>
        </TimelineItem>
        <TimelineItem title="Application UI code in Tailwind CSS" date="April 2022">
          <p class="text-base font-normal">
            Get started with dozens of web components and interactive elements built on top of Tailwind
            CSS.
          </p>
        </TimelineItem>
    </Timeline>
</div>


<style>
    .scrolling-wrapper {
  overflow-x: scroll;
  overflow-y: hidden;
  white-space: nowrap;
  
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
    }
</style>