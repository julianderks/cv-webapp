<script>
	import { page } from '$app/stores';
	import cat from '$lib/assets/sorry.png';
</script>

<div class="relative flex h-screen flex-col items-center justify-center overflow-clip">
	<h1 class="basis-1/3 text-5xl font-bold">HTTP {$page.status} - {$page.error.message}</h1>
	<img src={cat} alt="Error Cat" class="pop-up absolute bottom-0 left-1/4 basis-2/3" />
</div>

<style>
	@keyframes pop-up {
		from {
			transform: translateY(100%);
		}
		to {
			transform: translateY(10%);
		}
	}
	.pop-up {
		animation: pop-up 1s ease-out forwards;
	}
</style>
