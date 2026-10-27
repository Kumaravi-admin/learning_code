<script>
	let { InputValue = $bindable(), ...props } = $props();
</script>

<input type="text" bind:value={InputValue} {...props} />
