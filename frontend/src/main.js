import { library } from '@fortawesome/fontawesome-svg-core';
import {
  faBars,
  faUsers,
  faListUl,
  faSadTear,
  faHistory,
  faVoteYea,
  faUserAlt,
  faChartBar,
  faSlidersH,
  faArrowLeft,
  faUniversity,
  faUserGraduate,
} from '@fortawesome/free-solid-svg-icons';
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome';

import Vue from 'vue';
import App from './App.vue';
import router from './router';
import store from './store';

// Main menu icons
library.add(faBars);
library.add(faUserAlt);
library.add(faChartBar);
library.add(faSlidersH);

// OrgAdmin menu icons
library.add(faUsers);
library.add(faListUl);
library.add(faHistory);
library.add(faVoteYea);
library.add(faArrowLeft);
library.add(faUniversity);
library.add(faUserGraduate);

// OrgAdmin/Dashboard
library.add(faSadTear);

Vue.component('font-awesome-icon', FontAwesomeIcon);

Vue.config.productionTip = false;

new Vue({
  router,
  store,
  render: h => h(App),
}).$mount('#app');
