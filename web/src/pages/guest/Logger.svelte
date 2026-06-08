<script lang="ts">
  import { getLogs, getLogDetail, deleteLog } from '../../api/logger'
  import { instance } from '../../lib/axiosLib'

  let selectedFile = $state<string | null>(null)
  let searchQuery = $state('')
  let selectedLevels = $state<Record<string, boolean>>({
    info: true,
    error: true,
  })
  let currentPage = $state(1)
  let expandedIndex = $state<number | null>(null)
  let searchTimeout: ReturnType<typeof setTimeout> | null = null

  const filesQuery = getLogs()

  const detailQuery = getLogDetail(
    () => ({
      param: { file_name: selectedFile ?? '' },
      query: {
        search: searchQuery,
        page: String(currentPage),
        limit: '50',
      },
    }),
    () => Object.keys(selectedLevels).filter(k => selectedLevels[k]).join(',')
  )

  function levelBadge(level: string): string {
    if (level === 'error' || level === 'emergency' || level === 'alert' || level === 'critical')
      return 'bg-red-600 text-white'
    if (level === 'warning') return 'bg-yellow-600 text-white'
    if (level === 'info' || level === 'notice') return 'bg-blue-600 text-white'
    if (level === 'sql') return 'bg-indigo-600 text-white'
    return 'bg-gray-600 text-white'
  }

  function formatSize(bytes: number) {
    if (bytes === 0) return '0 B'
    const k = 1024
    const sizes = ['B', 'KB', 'MB', 'GB']
    const i = Math.floor(Math.log(bytes) / Math.log(k))
    return parseFloat((bytes / Math.pow(k, i)).toFixed(1)) + ' ' + sizes[i]
  }

  function toggleLevel(level: string) {
    selectedLevels = { ...selectedLevels, [level]: !selectedLevels[level] }
    currentPage = 1
  }

  function allLevelsDisabled(): boolean {
    return Object.values(selectedLevels).every(v => !v)
  }

  function onSearchInput() {
    if (searchTimeout) clearTimeout(searchTimeout)
    searchTimeout = setTimeout(() => {
      currentPage = 1
    }, 300)
  }

  function clearSearch() {
    searchQuery = ''
    currentPage = 1
  }

  function selectFile(name: string) {
    selectedFile = name
    currentPage = 1
    expandedIndex = null
  }

  function toggleExpand(index: number) {
    expandedIndex = expandedIndex === index ? null : index
  }

  function totalPages(): number {
    if (!detailQuery.data?.meta) return 1
    return Math.ceil(detailQuery.data.meta.total / detailQuery.data.meta.limit) || 1
  }

  function pageNumbers(): (number | '...')[] {
    const tp = totalPages()
    if (tp <= 7) return Array.from({ length: tp }, (_, i) => i + 1)
    const pages: (number | '...')[] = [1]
    if (currentPage > 3) pages.push('...')
    const start = Math.max(2, currentPage - 1)
    const end = Math.min(tp - 1, currentPage + 1)
    for (let i = start; i <= end; i++) pages.push(i)
    if (currentPage < tp - 2) pages.push('...')
    if (tp > 1) pages.push(tp)
    return pages
  }

  async function handleDownload(name: string) {
    const url = instance.defaults.baseURL + '/log/' + encodeURIComponent(name) + '/download'
    const a = document.createElement('a')
    a.href = url
    a.download = name
    a.click()
  }

  async function handleDelete(name: string) {
    if (!confirm(`Delete ${name}?`)) return
    await deleteLog({ param: { file_name: name }, query: {} })
    if (selectedFile === name) selectedFile = null
    filesQuery.refetch?.()
  }
</script>

<div class="flex h-screen min-h-0">
  <aside class="w-60 shrink-0 border-r border-gray-800 bg-gray-950 flex flex-col">
    <div class="p-3 border-b border-gray-800">
      <h2 class="text-xs font-semibold text-gray-400 uppercase tracking-wider">Log Files</h2>
    </div>
    <div class="flex-1 overflow-y-auto">
      {#if filesQuery.isLoading}
        <div class="p-4 text-sm text-gray-600">Loading...</div>
      {:else if filesQuery.data && filesQuery.data.length > 0}
        {#each filesQuery.data as file}
          <button
            class="w-full text-left px-3 py-2 text-sm border-b border-gray-900 hover:bg-gray-800/50 transition {selectedFile ===
            file.name
              ? 'bg-gray-800 text-white'
              : 'text-gray-400'}"
            onclick={() => selectFile(file.name)}
          >
            <div class="flex items-center justify-between gap-2">
              <span class="truncate">{file.name}</span>
              <span class="text-xs text-gray-600 shrink-0">{formatSize(file.size)}</span>
            </div>
          </button>
        {/each}
      {:else}
        <div class="p-4 text-sm text-gray-600">No log files found</div>
      {/if}
    </div>
  </aside>

  <main class="flex-1 flex flex-col min-w-0 bg-gray-950">
    {#if selectedFile}
      <div class="flex items-center gap-2 p-3 border-b border-gray-800 bg-gray-950 flex-wrap">
        <div class="relative flex-1 min-w-[200px]">
          <input
            type="text"
            placeholder="Search logs..."
            class="w-full bg-gray-900 border border-gray-700 rounded px-3 py-1.5 text-sm text-gray-200 placeholder-gray-600 focus:outline-none focus:border-blue-500"
            bind:value={searchQuery}
            oninput={onSearchInput}
          />
          {#if searchQuery}
            <button
              class="absolute right-2 top-1/2 -translate-y-1/2 text-gray-500 hover:text-gray-300"
              onclick={clearSearch}>✕</button
            >
          {/if}
        </div>
        <div class="flex items-center gap-1 flex-wrap">
          <button
            class="text-xs mr-1.5 px-3 py-1.5 rounded font-semibold uppercase transition bg-blue-700 text-white {selectedLevels[
              'info'
            ] || allLevelsDisabled()
              ? 'ring-2 ring-white/40'
              : 'ring-0 opacity-60'}"
            onclick={() => toggleLevel('info')}>info</button
          >
          <button
            class="text-xs px-3 py-1.5 rounded font-semibold uppercase transition bg-red-700 text-white {selectedLevels[
              'error'
            ] || allLevelsDisabled()
              ? 'ring-2 ring-white/40'
              : 'ring-0 opacity-60'}"
            onclick={() => toggleLevel('error')}>error</button
          >
        </div>
        <div class="flex items-center gap-1 ml-auto">
          <button
            class="text-xs px-3 py-1.5 rounded bg-gray-800 hover:bg-gray-700 text-gray-300 transition"
            onclick={() => detailQuery.refetch()}>⟳</button
          >
          <button
            class="text-xs px-3 py-1.5 rounded bg-gray-800 hover:bg-gray-700 text-gray-300 transition"
            onclick={() => handleDownload(selectedFile!)}>Download</button
          >
          <button
            class="text-xs px-3 py-1.5 rounded bg-red-800 hover:bg-red-700 text-white transition"
            onclick={() => handleDelete(selectedFile!)}>Delete</button
          >
        </div>
      </div>

      <div class="flex-1 overflow-y-auto">
        {#if detailQuery.isLoading}
          <div class="p-6 text-sm text-gray-500">Loading...</div>
        {:else if detailQuery.isError}
          <div class="p-6 text-sm text-red-400">Failed to load log file</div>
        {:else if (detailQuery.data?.data ?? []).length > 0}
          {#each detailQuery.data?.data ?? [] as entry, i}
            <div class="border-b border-gray-800">
              <button
                class="w-full flex items-start gap-2 px-4 py-2.5 text-left hover:bg-gray-800/30 transition"
                onclick={() => toggleExpand(i)}
              >
                <span class="shrink-0 w-3 mt-0.5 text-xs text-gray-700">{expandedIndex === i ? '▼' : '▶'}</span>
                <span
                  class="shrink-0 text-xs font-mono px-2 py-0.5 rounded font-semibold uppercase {levelBadge(
                    entry.level
                  )}">{entry.level}</span
                >
                <span class="text-xs text-gray-500 font-mono shrink-0 w-32">{entry.time}</span>
                <span class="text-sm text-gray-300 break-all line-clamp-2">{entry.message}</span>
              </button>
              {#if expandedIndex === i}
                <div class="px-12 pb-3 pt-1 text-xs text-gray-400 font-mono bg-gray-900/30">
                  <pre
                    class="whitespace-pre-wrap break-all text-gray-400 pl-3 border-l border-gray-700">{entry.message}</pre>
                </div>
              {/if}
            </div>
          {/each}
        {:else}
          <div class="p-6 text-sm text-gray-500">No log entries match your filters</div>
        {/if}
      </div>

      {#if detailQuery.data?.meta && totalPages() > 1}
        <div class="flex items-center justify-between px-4 py-2 border-t border-gray-800 bg-gray-950">
          <span class="text-xs text-gray-500">{detailQuery.data.meta.total} entries</span>
          <div class="flex items-center gap-1">
            <button
              class="px-3 py-1 text-xs rounded bg-gray-800 hover:bg-gray-700 text-gray-300 transition disabled:opacity-40 disabled:cursor-not-allowed"
              disabled={currentPage <= 1}
              onclick={() => (currentPage = Math.max(1, currentPage - 1))}>Prev</button
            >
            {#each pageNumbers() as p}
              {#if p === '...'}
                <span class="px-1 text-xs text-gray-600">...</span>
              {:else}
                <button
                  class="px-3 py-1 text-xs rounded transition {currentPage === p
                    ? 'bg-gray-700 text-white'
                    : 'bg-gray-800 text-gray-400 hover:bg-gray-700'}"
                  onclick={() => (currentPage = p)}>{p}</button
                >
              {/if}
            {/each}
            <button
              class="px-3 py-1 text-xs rounded bg-gray-800 hover:bg-gray-700 text-gray-300 transition disabled:opacity-40 disabled:cursor-not-allowed"
              disabled={currentPage >= totalPages()}
              onclick={() => (currentPage = Math.min(totalPages(), currentPage + 1))}>Next</button
            >
          </div>
        </div>
      {/if}
    {:else}
      <div class="flex-1 flex items-center justify-center text-gray-600 text-sm">Select a log file to view</div>
    {/if}
  </main>
</div>
