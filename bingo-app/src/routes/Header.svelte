<script lang="ts">
    import { Navbar, NavBrand, NavLi, NavUl, NavHamburger } from 'flowbite-svelte'
    import {onMount} from "svelte";

    onMount(() => {
        isAuthorized()
    })

    function parseJwt(token: string) {
        var base64Url = token.split('.')[1];
        var base64 = base64Url.replace(/-/g, '+').replace(/_/g, '/');
        var jsonPayload = decodeURIComponent(window.atob(base64).split('').map(function(c) {
            return '%' + ('00' + c.charCodeAt(0).toString(16)).slice(-2);
        }).join(''));

        return JSON.parse(jsonPayload);
    }

    $: show = false
    function isAuthorized() {
        var cookiestring=RegExp("auth=[^;]+").exec(document.cookie)
        const data = parseJwt(decodeURIComponent(!!cookiestring ? cookiestring.toString().replace(/^[^=]+./,"") : ""))
        show = data.user != ''
    }
</script>

<svelte:head>
    <link href="https://fonts.googleapis.com/css2?family=Josefin+Sans:wght@500&family=Unbounded:wght@200;300;400;500&family=Yellowtail&display=swap" rel="stylesheet">
</svelte:head>

<main>
    <Navbar let:hidden let:toggle color="none">
        <NavBrand href="/">
            <span class="self-center whitespace-nowrap text-xl font-semibold dark:text-white">Task Bingo</span>
        </NavBrand>
        <NavHamburger on:click={toggle}/>
        <NavUl {hidden}>
            <NavLi nonActiveClass="fonty" href="/">Home</NavLi>
            <NavLi nonActiveClass="fonty" href="/about">About</NavLi>
            {#if show}
                <NavLi nonActiveClass="fonty" href="/account">Account</NavLi>
            {:else}
                <NavLi nonActiveClass="fonty" href="/login">Login</NavLi>
            {/if}
        </NavUl>
    </Navbar>
</main>

<style>
    main {
		background-color: #11202d;
		color: rgb(238, 237, 197);
	}

    span {
        font-weight: 400;
    }
</style>