<script lang="ts">
    import { onMount } from "svelte";

    let data = $state("Loading...");

    function loadHealth() {
        data = "Loading...";
        fetch("/api/health").then((received) => {
            console.log(received);
            return received.json();
        }).then((finalData) => {
            if (Object.hasOwn(finalData, "status") && Object.hasOwn(finalData, "data") && finalData.status === "success") {
                data = finalData.data;
            } else {
                data = finalData;
            }
        }).catch((e) => {
            console.log("Some error, think...", e);
            data = "error :(";
        });
    }

    onMount(loadHealth);
</script>

<h2>Here is health of my server:</h2>
<p>{JSON.stringify(data)}</p>
<button class="mt-6" onclick={loadHealth}>Refresh</button>
