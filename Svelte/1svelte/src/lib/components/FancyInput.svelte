<script>
	// $bindable with default value
	let { value = $bindable('Default Value'), ...props } = $props();
</script>

<!-- This is a simple input component that accepts a value and binds it to the input element. -->
<input type="text" bind:value {...props} />
