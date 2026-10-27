<script>
  let count = $state(0)
  let src = "hello";
  let alt = "hello";
</script>

<button class="counter" onclick={() => count++}>Count is {count}</button>

<img src={src} alt={alt}>