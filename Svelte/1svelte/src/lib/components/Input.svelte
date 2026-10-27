<script>
	// initialize state for duration, currentTime and paused using $state
	let src = 'https://sveltejs.github.io/assets/caminandes-llamigos.mp4';
	let poster = 'https://sveltejs.github.io/assets/caminandes-llamigos.jpg';
	let duration = $state(0);
	let currentTime = $state(0);
	let paused = $state(true);
	// helper function to format time in mm:ss
	function format(time) {
		const minutes = Math.floor(time / 60);
		const seconds = Math.floor(time % 60)
			.toString()
			.padStart(2, '0');
		return `${minutes}:${seconds}`;
	}
</script>

<!-- bind duration, currentTime and paused -->
<video {src} {poster} bind:duration bind:currentTime bind:paused>
	<track kind="captions" />
</video>
<!-- toggle play/pause -->
<button onclick={() => (paused = !paused)}> {paused ? 'Play' : 'Pause'}</button>
<!-- format duration and currentTime -->
<p>{format(currentTime)} / {format(duration)}</p>
<!-- bind currentTime -->
<input type="range" min="0" max={duration} bind:value={currentTime} />
