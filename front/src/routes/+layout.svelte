<script>
	// Your selected Skeleton theme:
	import '@skeletonlabs/skeleton/themes/theme-skeleton.css';

	// This contains the bulk of Skeletons required styles:
	// NOTE: this will be renamed skeleton.css in the v2.x release.
	import '@skeletonlabs/skeleton/styles/skeleton.css';

	// Finally, your application's global stylesheet (sometimes labeled 'app.css')
	import '../app.postcss';
	import { AppShell } from '@skeletonlabs/skeleton';
	import { Icon } from '@steeze-ui/svelte-icon';
	import { Menu2 } from '@steeze-ui/tabler-icons';
	import { AppBar } from '@skeletonlabs/skeleton';
</script>

<AppShell>
	<!-- Router Slot -->

	<slot />

	<svelte:fragment slot="footer">Footer</svelte:fragment>
</AppShell>
