<script lang="ts">
    import { Navbar, NavBrand, NavLi, NavUl, NavHamburger } from 'flowbite-svelte'
    import { onMount } from "svelte";

    onMount(() => {
        isAuthorized()
    })

    function parseJwt(token: string) {
        let base64Url = token.split('.')[1];
        let base64 = base64Url.replace(/-/g, '+').replace(/_/g, '/');
        let jsonPayload = decodeURIComponent(window.atob(base64).split('').map(function(c) {
            return '%' + ('00' + c.charCodeAt(0).toString(16)).slice(-2);
        }).join(''));

        return JSON.parse(jsonPayload);
    }

    $: show = false

    function isAuthorized() {
        let cookiestring=RegExp("auth=[^;]+").exec(document.cookie)
        const data = parseJwt(decodeURIComponent(!!cookiestring ? cookiestring.toString().replace(/^[^=]+./,"") : ""))
        show = data.user != ''
    }
</script>

<svelte:head>
    <link href="https://fonts.googleapis.com/css2?family=Abril+Fatface:wght@200;300;400;500&family=Josefin+Sans:wght@200;300;400;500&family=Montserrat:wght@200;300;400;500&family=Oswald:wght@200;300;400;500&family=Prompt:wght@200;300;400;500&family=Righteous:wght@200;300;400;500&family=Unbounded:wght@200;300;400;500&family=Yellowtail:wght@200;300;400;500&display=swap&family=Meow+Script" rel="stylesheet">
</svelte:head>

<main>
    <Navbar let:hidden let:toggle color="none" class='dark'>
        <NavBrand href="/">
            <span class="self-center whitespace-nowrap text-xl font-semibold dark:text-white">taskbingo<span class='beta'>beta</span></span>
        </NavBrand>
        <NavHamburger on:click={toggle}/>
        <NavUl {hidden}>
            {#if show}
                <NavLi nonActiveClass="fonty" href="/people" data-sveltekit-prefetch>people</NavLi>
                <NavLi nonActiveClass="fonty" href="/packs" data-sveltekit-prefetch>packs</NavLi>
            {/if}
            <NavLi nonActiveClass="fonty" href="/about">about</NavLi>
            {#if show}
                <NavLi nonActiveClass="fonty" href="/account" data-sveltekit-prefetch>account</NavLi>
            {:else}
                <NavLi nonActiveClass="fonty" href="/login">login</NavLi>
            {/if}
        </NavUl>
    </Navbar>
</main>

<style>
    main {
        color: #f1eebc;
        background-color: #07417b;
	  }

    .beta {
        color: #7dffc6;
        /*font-family: Meow Script;*/
        font-size: x-large;
    }

    span {
        font-weight: 300;
        font-size: xx-large;
        color: #ffffff;
    }
</style>