import Home from './guest/Home.svelte'
import About from './guest/About.svelte'
import NotFound from './guest/NotFound.svelte'

const routes = {
  '/': Home,

  '/about': About,

  '*': NotFound,
}

export default routes
