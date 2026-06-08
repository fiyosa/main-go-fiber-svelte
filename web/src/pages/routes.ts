import Home from './guest/Home.svelte'
import About from './guest/About.svelte'
import Logger from './guest/Logger.svelte'
import NotFound from './guest/NotFound.svelte'

const routes = {
  '/': Home,

  '/about': About,

  '/logger': Logger,

  '*': NotFound,
}

export default routes
