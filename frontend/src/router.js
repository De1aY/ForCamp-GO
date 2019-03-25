import Vue from 'vue';
import Router from 'vue-router';

import Layout from '@/layouts/Base.vue';

import Profile from '@/views/portal/Profile.vue';
import OrgAdmin from '@/views/portal/orgadmin/Main.vue';
import Statistics from '@/views/portal/Statistics.vue';

Vue.use(Router);

export default new Router({
  mode: 'history',
  base: process.env.BASE_URL,
  routes: [
    {
      path: '/',
      name: 'index',
    },
    {
      path: '/portal',
      name: 'base',
      component: Layout,
      redirect: { name: 'profile' },
      children: [
        {
          path: 'orgadmin',
          name: 'orgadmin',
          component: OrgAdmin,
        },
        {
          path: 'profile',
          name: 'profile',
          component: Profile,
        },
        {
          path: 'statistics',
          name: 'statistics',
          component: Statistics,
        },
      ],
    },
  ],
});
