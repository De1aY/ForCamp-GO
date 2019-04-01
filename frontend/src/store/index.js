import Vue from 'vue';
import Vuex from 'vuex';

import global from './global';

Vue.use(Vuex);

const store = {
  state: {},
  modules: {
    global,
  },
};

export default new Vuex.Store(store);
