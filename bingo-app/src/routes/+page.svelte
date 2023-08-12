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
        <h1 class='spacer'>taskbingo transforms routine tasks into a game</h1>
        {#if show}
             <span>
                add friends on <a class='link' href='/people'>people tab</a><br>
                create your own <a class='link' href='/newpack'>packs</a><br>
                <a class='link' href='/account'>create games</a> with your friends<br>
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

    span {
        font-family: Prompt, serif;
        font-size: x-large;
        font-weight: 200;
        color: #ffffff;
        line-height: 40px;
    }

    .spacer {
        margin-bottom: 3em;
    }

    .link {
        text-decoration: underline;
        color: #c8fff7;
    }
</style>