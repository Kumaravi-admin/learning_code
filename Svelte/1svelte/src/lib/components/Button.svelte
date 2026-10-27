<script>
	let { children, ...props } = $props();
</script>

<!-- `props` includes attachments -->
<button {...props}>{@render children?.()}</button>
