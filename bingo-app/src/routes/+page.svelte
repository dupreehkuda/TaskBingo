<script lang="ts">
    import { browser } from '$app/environment';

    let show: boolean
    if ( browser ) {
        show = RegExp("auth=[^;]+").exec(document.cookie) !== null
    }
</script>

<svelte:head>
    <title>taskbingo</title>
</svelte:head>

<body>
  <h1>Premise</h1>
  <div class='spacer flex flex-col inline-block container mx-auto px-4'>
    <span>
      Two players agree upon a list of 16 tasks, which forms a 'pack'.
      Once a pack is set, the game commences.
      Players input the tasks they've completed, and these populate a 4x4 field.
      But here's the twist: tasks are scattered randomly across the field,
      introducing a delightful element of luck.
      The objective? Secure as many '4 in a row' combinations as possible.
    </span>
    <span class='beta'>the odds are always in your favour</span>
  </div>

  {#if show}
    <span>
      add friends on <a class='link' href='/people'>people tab</a><br>
      create your own <a class='link' href='/newpack'>packs</a><br>
      <a class='link' href='/account'>create games</a> with friends
    </span>
  {:else }
    <span>
      <a class='link' href='/register'>register</a> or <a class='link' href='/login'>login</a>
    </span>
  {/if }
</body>

<style>
    h1 {
        color: #ffffff;
        font-family: Prompt, serif;
        font-size: xx-large;
        font-weight: 300;
    }

    .beta {
      color: #7dffc6;
      font-weight: 400;
      margin-top: 0.75em;
    }

    span {
        display: inline-block;
        font-family: Prompt, serif;
        font-size: large;
        font-weight: 200;
        color: #ffffff;
        line-height: 30px;
        text-align: center;
    }

    .spacer {
        margin-bottom: 3em;
    }

    .link {
        text-decoration: underline;
        color: #c8fff7;
        font-weight: 400;
    }
</style>