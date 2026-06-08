<script lang="ts">
  import { getLogs } from '../../api/logger/getLogs'
  import { getLogDetail } from '../../api/logger/getLogDetail'

  let selectedFile = $state<string | null>(null)

  const filesQuery = getLogs()

  const detailQuery = getLogDetail(() => selectedFile)

  function levelClass(level: string) {
    if (level === 'error' || level === 'fatal') return 'text-red-400'
    if (level === 'warn') return 'text-yellow-400'
    return 'text-green-400'
  }
</script>

<div class="p-6 max-w-4xl mx-auto">
  <h1 class="text-2xl font-bold mb-4">Logger</h1>

  <div class="flex gap-2 mb-4 flex-wrap">
    <button
      class="px-3 py-1 rounded text-sm transition"
      class:bg-blue-600:text-white:bg-gray-700:text-gray-300={selectedFile === null}
      onclick={() => {
        selectedFile = null
      }}>All Files</button
    >
    {#each filesQuery.data ?? [] as f}
      <button
        class="px-3 py-1 rounded text-sm transition"
        class:bg-blue-600:text-white:bg-gray-700:text-gray-300={selectedFile === f}
        onclick={() => {
          selectedFile = f
        }}
      >
        {f}
      </button>
    {/each}
    {#if (filesQuery.data ?? []).length === 0 && !filesQuery.isLoading}
      <span class="text-gray-500 text-sm self-center">No log files found</span>
    {/if}
  </div>

  {#if detailQuery.isLoading}
    <p class="text-gray-400">Loading...</p>
  {:else if detailQuery.isError}
    <p class="text-red-400">Failed to load log file</p>
  {:else if detailQuery.data && detailQuery.data.logs.length > 0}
    <div class="overflow-x-auto">
      <table class="w-full text-sm font-mono">
        <thead>
          <tr class="text-gray-500 border-b border-gray-700">
            <th class="text-left py-2 pr-4 w-40">Time</th>
            <th class="text-left py-2 pr-4 w-20">Level</th>
            <th class="text-left py-2">Message</th>
          </tr>
        </thead>
        <tbody>
          {#each detailQuery.data.logs as entry}
            <tr class="border-b border-gray-800 hover:bg-gray-800/50">
              <td class="py-1.5 pr-4 text-gray-400 whitespace-nowrap">{entry.time}</td>
              <td class="py-1.5 pr-4 {levelClass(entry.level)}">{entry.level}</td>
              <td class="py-1.5 text-gray-200 break-all">{entry.message}</td>
            </tr>
          {/each}
        </tbody>
      </table>
    </div>
  {:else if selectedFile}
    <p class="text-gray-500">No log entries</p>
  {:else}
    <p class="text-gray-500">Select a log file above</p>
  {/if}
</div>
