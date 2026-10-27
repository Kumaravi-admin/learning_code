<script>
	let count = $state(0); // Initialize count state
	setInterval(() => {
		count += 1; // Increment count every second
	}, 1000); // Increment count every second
</script>

<h1>{count}</h1>
<!-- Render the count -->
