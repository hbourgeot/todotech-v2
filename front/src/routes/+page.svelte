<script lang="ts">
	import { Icon } from '@steeze-ui/svelte-icon';
	import { BrandWhatsapp } from '@steeze-ui/tabler-icons';
</script>

<main>
	<header class="hero">
		<section class="bg-gray-900/50 h-full flex flex-col gap-y-5 w-full justify-center items-center">
			<h2 class="text-4xl md:text-8xl font-black">Somos Inversiones 7H</h2>
			<p class="text-xl md:text-2xl font-semibold">Todo a tu alcance.</p>
			<a
				target="_blank"
				href="https://wa.me/584145984147?text=Hola! Quisiera saber un poco más sobre "
				class="btn bg-white text-green-600 text-lg"
				><Icon src={BrandWhatsapp} class="h-5 w-5 mr-2" /> Contáctanos</a
			>
		</section>
	</header>
	<section class="container mx-auto my-5">
		<h2 class="text-4xl">Conoce nuestros productos</h2>
	</section>
</main>

<style lang="scss">
	.hero {
		background-image: url(https://cdn.pixabay.com/photo/2023/03/17/16/41/traffic-7859033_960_720.jpg);
		height: 400px;
		background-size: cover;
		color: white;
	}
</style>
