
import Vue from 'vue';

import Buefy from 'buefy'
import App from './App.vue'
import 'buefy/dist/buefy.css'

Vue.use(Buefy)

new Vue({
    el: "#app",
    render: h=> h(App)
})

//literalmente nada
