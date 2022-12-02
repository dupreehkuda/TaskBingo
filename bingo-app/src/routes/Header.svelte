<script>
    import { Button } from 'flowbite-svelte';
    import { Navbar, NavBrand, NavLi, NavUl, NavHamburger } from 'flowbite-svelte'
</script>

<main>
    <Navbar let:hidden let:toggle rounded color="form">
        <NavBrand href="/">
            <img
                    src="https://flowbite.com/docs/images/logo.svg"
                    class="mr-3 h-6 sm:h-9"
                    alt="Flowbite Logo"
            />
            <span class="self-center whitespace-nowrap text-xl font-semibold dark:text-white">Sveltish Bingo</span>
        </NavBrand>
        <NavHamburger on:click={toggle} />
        <NavUl {hidden}>
            <NavLi href="/" active={true}>Home</NavLi>
            <NavLi href="/about">About</NavLi>
            <NavLi href="/login">Login</NavLi>
        </NavUl>
    </Navbar>

<!--    <Button href="/" gradient shadow="teal" color="teal">Task registry</Button>-->
<!--    <Button href="/" gradient shadow="teal" color="teal">Get Points</Button>-->
<!--    <Button href="/login" gradient shadow="teal" color="teal">Login</Button>-->
</main>

<style>
    a.clr-button{
        display:inline-block;
        padding:0.3em 1.2em;
        margin:0 0.1em 0.1em 0;
        border:0.16em solid rgba(255,255,255,0);
        border-radius:2em;
        box-sizing: border-box;
        text-decoration:none;
        font-family:'LoveFrom SJA',sans-serif;
        font-weight:300;
        color:#FFFFFF;
        background-color: #ED6F47;
        text-shadow: 0 0.04em 0.04em rgba(0,0,0,0.35);
        text-align:center;
        transition: all 0.2s;
    }
    a.clr-button:hover{
        border-color: rgba(255,255,255,1);
    }
    @media all and (max-width:30em){
        a.clr-button{
            display:block;
            margin:0.2em auto;
        }
    }
</style>