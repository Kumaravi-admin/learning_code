<script>
	import { onDestroy } from 'svelte'; // import onDestroy

	// onDestroy is called when the component is being destroyed
	onDestroy(() => {
		console.log('the component is being destroyed');
	});
</script>

<!-- Render the component -->
<h1>Hello Svelte learner</h1>
