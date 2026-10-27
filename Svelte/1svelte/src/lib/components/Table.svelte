<script>
	let { data, header, row } = $props(); // Get the props passed to the component
</script>

<div>
	<table>
		<!-- Render the header -->
		{#if header}
			<thead>
				<tr> {@render header()}</tr>
			</thead>
		{/if}

		<!-- Render the rows -->
		{#if data}
			<tbody>
				{#each data as d, i (i)}
					<tr>{@render row(d)}</tr>
				{/each}
			</tbody>
		{/if}
	</table>
</div>

<style>
	table {
		text-align: left;
		border-spacing: 0;
	}
	tbody tr:nth-child(2n + 1) {
		background: #f0f0f0;
	}
	table :global(th),
	table :global(td) {
		padding: 0.5em;
	}
</style>
