<script lang="ts">
  import CronForm from "../components/CronForm.svelte";
</script>

<div class="container mx-auto">
<h1 class="text-3xl font-bold underline">Trevor Arnold's Crontab Tool</h1>

<br>

<p>Enter the cron pattern in the first box and a human readable interpretation will be presented in the Cron Pattern Meaning box</p>

<br>

<CronForm />

<br>

<div>
  <table class="w-full border-collapse bg-white text-left text-sm text-gray-500">
    <thead class="bg-gray-50">
      <tr>
        <th scope="col" class="px-6 py-4 font-medium text-gray-900">Symbol</th>
        <th scope="col" class="px-6 py-4 font-medium text-gray-900">Symbol Name</th>
        <th scope="col" class="px-6 py-4 font-medium text-gray-900">Meaning</th>
      </tr>
    </thead>
    <tbody class="divide-y divide-gray-100 border-t border-gray-100">
      <tr class="odd:bg-white even:bg-gray-50">
        <td class="px-6 py-4">*</td>
        <td class="px-6 py-4">star</td>
        <td class="px-6 py-4">any value (wildcard)</td>
      </tr>
      <tr class="odd:bg-white even:bg-gray-50">
        <td class="px-6 py-4">,</td>
        <td class="px-6 py-4">comma</td>
        <td class="px-6 py-4">list separator (0, 15, 30, 45)</td>
      </tr>
      <tr class="odd:bg-white even:bg-gray-50">
        <td class="px-6 py-4">-</td>
        <td class="px-6 py-4">hyphen</td>
        <td class="px-6 py-4">range separator (1-5)</td>
      </tr>
      <tr class="odd:bg-white even:bg-gray-50">
        <td class="px-6 py-4">/</td>
        <td class="px-6 py-4">forward slash</td>
        <td class="px-6 py-4">step value (1/10)</td>
      </tr>
    </tbody>
  </table>
</div>
</div>
