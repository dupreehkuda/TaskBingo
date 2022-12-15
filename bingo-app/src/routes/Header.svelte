<script>
    import { Navbar, NavBrand, NavLi, NavUl, NavHamburger } from 'flowbite-svelte'
</script>

<main>
    <Navbar let:hidden let:toggle rounded color="dark" fluid>
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