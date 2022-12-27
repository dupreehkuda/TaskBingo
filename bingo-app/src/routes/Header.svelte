<script>
    import { Navbar, NavBrand, NavLi, NavUl, NavHamburger } from 'flowbite-svelte'
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
            <NavLi nonActiveClass="fonty" href="/login">Login</NavLi>
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