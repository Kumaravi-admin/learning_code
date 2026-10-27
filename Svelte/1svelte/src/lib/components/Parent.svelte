<!-- Parent.svelte -->
<script>
	import { setUserContext } from './context';
	let { children } = $props();
	setUserContext({ name: 'world' });
</script>

{@render children()}
