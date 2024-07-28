<script lang="ts">
  import { onMount, onDestroy } from "svelte";

  interface Tool {
    tool: string;
    iconify_icon_name: string;
    experience_level: number;
    iconify_icon_height_factor: number;
    area: string;
  }

  let tools: Tool[] = [];
  let positions: { x: number; y: number; dx: number; dy: number }[] = [];
  let container: HTMLDivElement;

  onMount(async () => {
    const response = await fetch("http://localhost:8000/api");
    tools = await response.json();
    tools.sort((a, b) => {
      if (a.area < b.area) return -1;
      if (a.area > b.area) return 1;
      if (a.tool < b.tool) return -1;
      if (a.tool > b.tool) return 1;
      return b.experience_level - a.experience_level;
    });
    initializePositions();
    requestAnimationFrame(animate);
    window.addEventListener("resize", handleResize);
  });

  onDestroy(() => {
    window.removeEventListener("resize", handleResize);
  });

  function initializePositions() {
    positions = tools.map(() => ({
      x: Math.random() * container.clientWidth,
      y: Math.random() * container.clientHeight,
      dx: (Math.random() - 0.5) * 6,
      dy: (Math.random() - 0.5) * 6,
    }));
  }

  function handleResize() {
    positions = positions.map((pos) => {
      // Adjust positions to stay within the new container size
      return {
        ...pos,
        x: Math.min(pos.x, container.clientWidth - 56),
        y: Math.min(pos.y, container.clientHeight - 56),
      };
    });
  }

  function animate() {
    positions = positions.map((pos) => {
      let { x, y, dx, dy } = pos;
      x += dx;
      y += dy;

      if (x <= 0 || x >= container.clientWidth - 56) dx *= -1; // 56 is the icon size in px
      if (y <= 0 || y >= container.clientHeight - 56) dy *= -1;

      return { x, y, dx, dy };
    });
    requestAnimationFrame(animate);
  }
</script>

<div bind:this={container} class="relative w-full h-full">
  {#each tools as tool, i}
    <div
      class="absolute transition-transform duration-100 linear"
      style="transform: translate({positions[i].x}px, {positions[i].y}px);"
    >
      <h1>{tool.tool}</h1>
      <iconify-icon icon={tool.iconify_icon_name} height="3.5em"></iconify-icon>
    </div>
  {/each}
</div>
