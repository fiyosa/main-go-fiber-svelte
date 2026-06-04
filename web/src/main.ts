import { mount } from 'svelte'
import { setHashRoutingEnabled, setBasePath } from '@keenmate/svelte-spa-router'
import './app.css'
import App from './App.svelte'

if (location.hash.startsWith('#/')) history.replaceState(null, '', location.hash.slice(1))
setHashRoutingEnabled(false)
setBasePath('/')

const app = mount(App, { target: document.getElementById('app')! })

export default app
