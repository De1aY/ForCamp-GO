import Vue from 'vue';
import Router from 'vue-router';

import Layout from '@/layouts/Base.vue';

import Profile from '@/views/portal/Profile.vue';
import Statistics from '@/views/portal/Statistics.vue';

// OrgAdmin
import Teams from '@/views/portal/orgadmin/Teams.vue';
import Actions from '@/views/portal/orgadmin/Actions.vue';
import Reasons from '@/views/portal/orgadmin/Reasons.vue';
import Employees from '@/views/portal/orgadmin/Employees.vue';
import Dashboard from '@/views/portal/orgadmin/Dashboard.vue';
import Categories from '@/views/portal/orgadmin/Categories.vue';
import Participants from '@/views/portal/orgadmin/Participants.vue';

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
          component: null,
          redirect: { name: 'orgadmin/dashboard' },
        },
        {
          path: 'orgadmin/dashboard',
          name: 'orgadmin/dashboard',
          component: Dashboard,
        },
        {
          path: 'orgadmin/categories',
          name: 'orgadmin/categories',
          component: Categories,
        },
        {
          path: 'orgadmin/teams',
          name: 'orgadmin/teams',
          component: Teams,
        },
        {
          path: 'orgadmin/participants',
          name: 'orgadmin/participants',
          component: Participants,
        },
        {
          path: 'orgadmin/employees',
          name: 'orgadmin/employees',
          component: Employees,
        },
        {
          path: 'orgadmin/reasons',
          name: 'orgadmin/reasons',
          component: Reasons,
        },
        {
          path: 'orgadmin/actions',
          name: 'orgadmin/actions',
          component: Actions,
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
