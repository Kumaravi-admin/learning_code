<script>
	let { onclose } = $props(); // receive onclose callback from parent component

	// simulate async data fetching
	async function fetchData() {
		await new Promise((f) => setTimeout(f, 2000)); // simulate delay
		return 'Data loaded successfully!';
	}
</script>

<!-- boundary handles loading and error states -->

<svelte:boundary>
	<!-- show while fetchData is resolving -->
	{#snippet pending()}
		<p>Loading data...</p>
	{/snippet}
	<!-- show if fetchData throws an error  -->
	{#snippet failed(error)}
		<p>Error: {error.message}</p>
	{/snippet}
	<!-- show  when fetchData resolves successfully -->
	<div>
		<p>{await fetchData()}</p>
		<button onclick={onclose}> Close</button>
	</div>
</svelte:boundary>
