<script>
    import BattleshipsGrid from "$lib/BattleshipsGrid.svelte";
    export let data;
</script>

<h1>user session ID</h1>
<div>{data.user.username}</div>

<div class="parent">
    <div class="child"><BattleshipsGrid /></div>
    <div class="child"><BattleshipsGrid /></div>
</div>
