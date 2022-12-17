<script>
    import { Navbar, NavBrand, NavLi, NavUl, NavHamburger } from 'flowbite-svelte'
</script>

<svelte:head>
    <link href="https://fonts.googleapis.com/css2?family=Josefin+Sans:wght@500&family=Unbounded&family=Yellowtail&display=swap" rel="stylesheet">
</svelte:head>

<main>
    <Navbar let:hidden let:toggle color="none">
        <NavBrand href="/">
            <img
                    src="https://flowbite.com/docs/images/logo.svg"
                    class="mr-3 h-6 sm:h-9"
                    alt="Flowbite Logo"
            />
            <span class="self-center whitespace-nowrap text-xl font-semibold dark:text-white">Sveltish Bingo</span>
        </NavBrand>
        <NavHamburger on:click={toggle}/>
        <NavUl {hidden}>
            <NavLi href="/">Home</NavLi>
            <NavLi href="/about">About</NavLi>
            <NavLi href="/login">Login</NavLi>
        </NavUl>
    </Navbar>
</main>

<style>
    main {
		background-color: #11202d;
		color: rgb(238, 237, 197);
	}

    /* text {
        color: rgb(238, 237, 197);
    } */

    span {
        font-family: Unbounded;
    }
</style>