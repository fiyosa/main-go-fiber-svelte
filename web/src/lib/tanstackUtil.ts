import { createQuery, createMutation, QueryClient } from '@tanstack/svelte-query'
import type {
  CreateQueryOptions,
  CreateMutationOptions,
  CreateQueryResult,
  CreateMutationResult,
  DefaultError,
  QueryKey,
} from '@tanstack/svelte-query'

export type { CreateQueryOptions }

const queryClient = new QueryClient({
  defaultOptions: {
    queries: {
      refetchOnWindowFocus: false,
      retry: false,
      staleTime: 1000 * 60 * 1,
    },
  },
})

export function useQuery<
  TQueryFnData = unknown,
  TError = DefaultError,
  TData = TQueryFnData,
  TQueryKey extends QueryKey = QueryKey,
>(options: () => CreateQueryOptions<TQueryFnData, TError, TData, TQueryKey>): CreateQueryResult<TData, TError> {
  return createQuery(options, () => queryClient)
}

export function useMutation<TData = unknown, TError = DefaultError, TVariables = void, TContext = unknown>(
  options: () => CreateMutationOptions<TData, TError, TVariables, TContext>
): CreateMutationResult<TData, TError, TVariables, TContext> {
  return createMutation(options, () => queryClient)
}
