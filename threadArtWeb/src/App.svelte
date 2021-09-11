<script>
	export let name;
</script>

<main>
	<h1 class="">Hello {name}!</h1>
	<p>Visit the <a href="https://svelte.dev/tutorial">Svelte tutorial</a> to learn how to build Svelte apps.</p>
	<div class="columns">
		<div class="column"><p class="bd-notification is-primary">First column</p></div>
		<div class="column">Second column</div>
		<div class="column">Third column</div>
		<div class="column">Fourth column</div>
	</div>
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