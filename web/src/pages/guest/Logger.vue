<script setup lang="ts">
  import { ref, computed, watch } from 'vue'
  import { useQuery, useMutation } from '@tanstack/vue-query'
  import instance from '../../lib/axios'

  const selectedFile = ref<string | null>(null)
  const searchQuery = ref('')
  const selectedLevels = ref<Record<string, boolean>>({
    info: true,
    error: true,
  })
  const currentPage = ref(1)
  const expandedIndex = ref<number | null>(null)

  type LogFile = { name: string; size: number }
  type LogEntry = { level: string; time: string; message: string }
  type PaginationMeta = { total: number; page: number; limit: number }
  type LogDetailResponse = { message: string; data: LogEntry[]; meta: PaginationMeta }

  const filesQuery = useQuery({
    queryKey: ['log'],
    queryFn: () => instance.get<{ data: LogFile[] }>('/log').then(r => r.data.data),
  })

  const levelsParam = computed(() =>
    Object.keys(selectedLevels.value).filter(k => selectedLevels.value[k]).join(',')
  )

  const detailQuery = useQuery({
    queryKey: computed(() => ['log', 'detail', selectedFile.value, searchQuery.value, currentPage.value, levelsParam.value]),
    queryFn: () => {
      const q = new URLSearchParams()
      if (searchQuery.value) q.set('search', searchQuery.value)
      if (currentPage.value) q.set('page', String(currentPage.value))
      q.set('limit', '50')
      if (levelsParam.value) q.set('levels', levelsParam.value)
      const qs = q.toString()
      return instance.get<LogDetailResponse>(`/log/${encodeURIComponent(selectedFile.value!)}${qs ? '?' + qs : ''}`).then(r => r.data)
    },
    enabled: computed(() => !!selectedFile.value),
  })

  const deleteMutation = useMutation({
    mutationFn: (fileName: string) => instance.delete(`/log/${encodeURIComponent(fileName)}`),
    onSuccess: (_, fileName) => {
      if (selectedFile.value === fileName) selectedFile.value = null
      filesQuery.refetch()
    },
  })

  let searchTimeout: ReturnType<typeof setTimeout> | null = null
  watch(searchQuery, () => {
    if (searchTimeout) clearTimeout(searchTimeout)
    searchTimeout = setTimeout(() => {
      currentPage.value = 1
    }, 300)
  })

  function levelBadge(level: string): string {
    if (level === 'error' || level === 'emergency' || level === 'alert' || level === 'critical') return 'bg-red-600 text-white'
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
    selectedLevels.value = { ...selectedLevels.value, [level]: !selectedLevels.value[level] }
    currentPage.value = 1
  }

  function clearSearch() {
    searchQuery.value = ''
    currentPage.value = 1
  }

  function selectFile(name: string) {
    selectedFile.value = name
    currentPage.value = 1
    expandedIndex.value = null
  }

  function toggleExpand(index: number) {
    expandedIndex.value = expandedIndex.value === index ? null : index
  }

  const totalPages = computed(() => {
    if (!detailQuery.data.value?.meta) return 1
    return Math.ceil(detailQuery.data.value.meta.total / detailQuery.data.value.meta.limit) || 1
  })

  const pageNumbers = computed((): (number | '...')[] => {
    const tp = totalPages.value
    if (tp <= 7) return Array.from({ length: tp }, (_, i) => i + 1)
    const pages: (number | '...')[] = [1]
    if (currentPage.value > 3) pages.push('...')
    const start = Math.max(2, currentPage.value - 1)
    const end = Math.min(tp - 1, currentPage.value + 1)
    for (let i = start; i <= end; i++) pages.push(i)
    if (currentPage.value < tp - 2) pages.push('...')
    if (tp > 1) pages.push(tp)
    return pages
  })

  function handleDownload(name: string) {
    const url = instance.defaults.baseURL + '/log/' + encodeURIComponent(name) + '/download'
    const a = document.createElement('a')
    a.href = url
    a.download = name
    a.click()
  }

  function handleDelete(name: string) {
    if (!confirm(`Delete ${name}?`)) return
    deleteMutation.mutate(name)
  }
</script>

<template>
  <div class="flex h-screen min-h-0">
    <aside class="w-60 shrink-0 border-r border-gray-800 bg-gray-950 flex flex-col">
      <div class="p-3 border-b border-gray-800">
        <h2 class="text-xs font-semibold text-gray-400 uppercase tracking-wider">Log Files</h2>
      </div>
      <div class="flex-1 overflow-y-auto">
        <div v-if="filesQuery.isLoading.value" class="p-4 text-sm text-gray-600">Loading...</div>
        <div v-else-if="filesQuery.data.value && filesQuery.data.value.length > 0">
          <button
            v-for="file in filesQuery.data.value"
            :key="file.name"
            :class="['w-full text-left px-3 py-2 text-sm border-b border-gray-900 hover:bg-gray-800/50 transition', selectedFile === file.name ? 'bg-gray-800 text-white' : 'text-gray-400']"
            @click="selectFile(file.name)"
          >
            <div class="flex items-center justify-between gap-2">
              <span class="truncate">{{ file.name }}</span>
              <span class="text-xs text-gray-600 shrink-0">{{ formatSize(file.size) }}</span>
            </div>
          </button>
        </div>
        <div v-else class="p-4 text-sm text-gray-600">No log files found</div>
      </div>
    </aside>

    <main class="flex-1 flex flex-col min-w-0 bg-gray-950">
      <template v-if="selectedFile">
        <div class="flex items-center gap-2 p-3 border-b border-gray-800 bg-gray-950 flex-wrap">
          <div class="relative flex-1 min-w-[200px]">
            <input
              type="text"
              placeholder="Search logs..."
              class="w-full bg-gray-900 border border-gray-700 rounded px-3 py-1.5 text-sm text-gray-200 placeholder-gray-600 focus:outline-none focus:border-blue-500"
              v-model="searchQuery"
            />
            <button
              v-if="searchQuery"
              class="absolute right-2 top-1/2 -translate-y-1/2 text-gray-500 hover:text-gray-300"
              @click="clearSearch">✕</button>
          </div>
          <div class="flex items-center gap-1 flex-wrap">
            <button
              :class="['text-xs mr-1.5 px-3 py-1.5 rounded font-semibold uppercase transition bg-blue-700 text-white', selectedLevels['info'] ? 'ring-2 ring-white/40' : 'ring-0 opacity-60']"
              @click="toggleLevel('info')">info</button>
            <button
              :class="['text-xs px-3 py-1.5 rounded font-semibold uppercase transition bg-red-700 text-white', selectedLevels['error'] ? 'ring-2 ring-white/40' : 'ring-0 opacity-60']"
              @click="toggleLevel('error')">error</button>
          </div>
          <div class="flex items-center gap-1 ml-auto">
            <button
              class="text-xs px-3 py-1.5 rounded bg-gray-800 hover:bg-gray-700 text-gray-300 transition"
              @click="detailQuery.refetch()">⟳</button>
            <button
              class="text-xs px-3 py-1.5 rounded bg-gray-800 hover:bg-gray-700 text-gray-300 transition"
              @click="handleDownload(selectedFile!)">Download</button>
            <button
              class="text-xs px-3 py-1.5 rounded bg-red-800 hover:bg-red-700 text-white transition"
              @click="handleDelete(selectedFile!)">Delete</button>
          </div>
        </div>

        <div class="flex-1 overflow-y-auto">
          <div v-if="detailQuery.isLoading.value" class="p-6 text-sm text-gray-500">Loading...</div>
          <div v-else-if="detailQuery.isError.value" class="p-6 text-sm text-red-400">Failed to load log file</div>
          <div v-else-if="(detailQuery.data.value?.data ?? []).length > 0">
            <div v-for="(entry, i) in detailQuery.data.value?.data ?? []" :key="i" class="border-b border-gray-800">
              <button
                class="w-full flex items-start gap-2 px-4 py-2.5 text-left hover:bg-gray-800/30 transition"
                @click="toggleExpand(i)"
              >
                <span class="shrink-0 w-3 mt-0.5 text-xs text-gray-700">{{ expandedIndex === i ? '▼' : '▶' }}</span>
                <span :class="['shrink-0 text-xs font-mono px-2 py-0.5 rounded font-semibold uppercase', levelBadge(entry.level)]">{{ entry.level }}</span>
                <span class="text-xs text-gray-500 font-mono shrink-0 w-32">{{ entry.time }}</span>
                <span class="text-sm text-gray-300 break-all line-clamp-2">{{ entry.message }}</span>
              </button>
              <div v-if="expandedIndex === i" class="px-12 pb-3 pt-1 text-xs text-gray-400 font-mono bg-gray-900/30">
                <pre class="whitespace-pre-wrap break-all text-gray-400 pl-3 border-l border-gray-700">{{ entry.message }}</pre>
              </div>
            </div>
          </div>
          <div v-else class="p-6 text-sm text-gray-500">No log entries match your filters</div>
        </div>

        <div v-if="detailQuery.data.value?.meta && totalPages > 1" class="flex items-center justify-between px-4 py-2 border-t border-gray-800 bg-gray-950">
          <span class="text-xs text-gray-500">{{ detailQuery.data.value.meta.total }} entries</span>
          <div class="flex items-center gap-1">
            <button
              class="px-3 py-1 text-xs rounded bg-gray-800 hover:bg-gray-700 text-gray-300 transition disabled:opacity-40 disabled:cursor-not-allowed"
              :disabled="currentPage <= 1"
              @click="currentPage = Math.max(1, currentPage - 1)">Prev</button>
            <template v-for="p in pageNumbers" :key="typeof p === 'string' ? p : p">
              <span v-if="p === '...'" class="px-1 text-xs text-gray-600">...</span>
              <button
                v-else
                :class="['px-3 py-1 text-xs rounded transition', currentPage === p ? 'bg-gray-700 text-white' : 'bg-gray-800 text-gray-400 hover:bg-gray-700']"
                @click="currentPage = p">{{ p }}</button>
            </template>
            <button
              class="px-3 py-1 text-xs rounded bg-gray-800 hover:bg-gray-700 text-gray-300 transition disabled:opacity-40 disabled:cursor-not-allowed"
              :disabled="currentPage >= totalPages"
              @click="currentPage = Math.min(totalPages, currentPage + 1)">Next</button>
          </div>
        </div>
      </template>
      <div v-else class="flex-1 flex items-center justify-center text-gray-600 text-sm">Select a log file to view</div>
    </main>
  </div>
</template>
