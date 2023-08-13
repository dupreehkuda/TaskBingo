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

<div>
    <body>
<!--        <h1 class='spacer'>the odds are always in your favour</h1>-->

        <h1>Premise</h1>
        <div class='spacer premise'>
          <span>
            Two players agree upon a list of 16 tasks, which forms a 'pack'.
            Once a pack is set, the game commences.
            Players input the tasks they've completed, and these populate a 4x4 field.
<!--            <div class='spacer'></div>-->
            But here's the twist: tasks are scattered randomly across the field,
            introducing a delightful element of luck.
<!--            <div class='spacer'></div>-->
            The objective? Secure as many '4 in a row' combinations as possible.
            <span class='beta'>the odds are always in your favour</span>
          </span>
        </div>

        {#if show}
             <span>
                add friends on <a class='link' href='/people'>people tab</a> |
                create your own <a class='link' href='/newpack'>packs</a><br>
                <a class='link' href='/account'>create games</a> with your friends |
                challenge yourself
            </span>
        {:else }
            <span>
                <a class='link' href='/register'>register</a> or <a class='link' href='/login'>login</a>
            </span>
        {/if }
    </body>
</div>

<style>
    h1 {
        color: #ffffff;
        font-family: Prompt, serif;
        font-size: xx-large;
        font-weight: 300;
    }

    .beta {
      color: #7dffc6;
    }

    .premise {
      margin-left: 5rem;
      margin-right: 5rem;
    }

    span {
        display: inline-block;
        font-family: Prompt, serif;
        font-size: large;
        font-weight: 200;
        color: #ffffff;
        line-height: 40px;
        text-align: justify;
    }

    .spacer {
        margin-bottom: 3em;
    }

    .link {
        text-decoration: underline;
        color: #c8fff7;
    }
</style>