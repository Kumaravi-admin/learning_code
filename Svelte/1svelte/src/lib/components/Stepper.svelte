<!-- Register the custom element -->
<svelte:options customElement="my-stepper" />

<script>
	// Dispatch custom events
	function dispatch(type) {
		$host().dispatchEvent(new CustomEvent(type));
	}
</script>

<!-- Render the buttons -->
<button onclick={() => dispatch('decrement')}>decrement</button>
<button onclick={() => dispatch('increment')}>increment</button>
