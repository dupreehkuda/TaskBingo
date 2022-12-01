<script lang="ts">
	export let name: string;

	const newResp = {
		userID: 'dupreeh'
	}

	let userID = 'dupreeh'
	let result = null

	async function doPost() {
		const res = await fetch('http://localhost:8081/api', {
			method: 'POST',
			body: JSON.stringify(newResp)
		})

		const json = await res.json()
		console.log(JSON.stringify(json))
		result = JSON.stringify(json)
	}
</script>

<main>
	<h1>Hello {name}!</h1>
	<button type="button" on:click={doPost}>
		get points
	</button>
	<p>Result: </p>
	<pre>{result}</pre>
</main>

<style>
	main {
		text-align: center;
		padding: 1em;
		max-width: 240px;
		margin: 0 auto;
	}

	h1 {
		color: #ff3e00;
		text-transform: uppercase;
		font-size: 4em;
		font-weight: 100;
	}

	@media (min-width: 640px) {
		main {
			max-width: none;
		}
	}
</style>