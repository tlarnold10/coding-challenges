<script lang=ts>
  
  var showResults = false;
  var results = "";

  function onSubmit(e: any) {
    showResults = true;
    results = e.target[0].value;
    console.log(results);
  }
</script>

<style>
  .container {
    padding: 3em;
  }
</style>

<div class="container">
  <div class="form">
    <form on:submit|preventDefault={onSubmit}>
      <div class="mx-auto">
        <div>
          <div class="relative">
            <input type="text" id="cronPattern" placeholder=" " class="peer block w-full rounded-md border-black shadow-sm focus:border-primary-400 focus:ring focus:ring-primary-200 focus:ring-opacity-50 disabled:cursor-not-allowed disabled:bg-gray-50 disabled:text-gray-500" />
            <label for="cronPattern" class="peer-focus:base absolute left-2 top-0 z-10 -translate-y-2 transform bg-white px-1 text-xs text-gray-500 transition-all peer-placeholder-shown:translate-y-3 peer-placeholder-shown:text-sm peer-focus:-translate-y-2 peer-focus:text-xs peer-disabled:bg-transparent">Cron Pattern</label>
          </div>
        </div>
      </div>
    </form>
  </div>

  <br>

  {#if showResults}
    <div class="mx-auto">
      <input type="text" class="block w-full rounded-md border-black shadow-sm focus:border-primary-400 focus:ring focus:ring-primary-200 focus:ring-opacity-50 disabled:cursor-not-allowed disabled:bg-gray-50 disabled:text-gray-500" placeholder="{ results }" />
    </div> 
  {/if}
</div>
