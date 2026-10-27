<!-- file: src/lib/components/Box.svelte -->

<!-- Using a custom property -->
<div class="box"></div>

<!-- Using a custom property -->
<style>
	.box {
		/* Use the variable with a fallback value */
		background-color: var(--box-color, #ccc);
		width: 100px;
		height: 100px;
	}
</style>
