<script>
	import Header from '../components/Header.svelte';
	import '../app.postcss';
</script>

<svelte:head>
	<link rel="icon" type="image/svg" href="../static/favicon.svg" />
</svelte:head>

<Header/>
<slot />