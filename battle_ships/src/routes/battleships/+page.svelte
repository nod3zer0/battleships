<script>
    import BattleshipsGrid from "$lib/BattleshipsGrid.svelte";
    export let data;
</script>

<h1>user session ID</h1>
<!-- <div>{data.user.username}</div> -->

<h1>Games</h1>
{#each data.games as game}
    <div>
        {game.username}
        <a href="battleships/{game.id}">{game.name}</a>
    </div>
{/each}
<!-- <div class="parent">
    <div class="child"><BattleshipsGrid /></div>
    <div class="child"><BattleshipsGrid /></div>
</div> -->
