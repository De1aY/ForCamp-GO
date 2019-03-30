<template>
  <div class="layout">
    <header class="navbar"></header>
    <div class="menu">
      <div class="menu-action">
        <font-awesome-icon icon="bars" class="menu-icon" />
      </div>
      <div class="menu-content menu-content--main"
        v-if="activeMenu === 'main'">
        <router-link to="orgadmin" class="menu-item"
          v-bind:class="{'menu-item--active': $router.currentRoute.name === 'orgadmin'}"
          v-on:click.native="setActiveMenu('orgadmin')">
          <font-awesome-icon icon="sliders-h" class="menu-icon" />
        </router-link>
        <router-link to="profile" class="menu-item"
          v-bind:class="{'menu-item--active': $router.currentRoute.name === 'profile'}">
          <font-awesome-icon icon="user-alt" class="menu-icon" />
        </router-link>
        <router-link to="statistics" class="menu-item"
          v-bind:class="{ 'menu-item--active':$router.currentRoute.name === 'statistics'}">
          <font-awesome-icon icon="chart-bar" class="menu-icon" />
        </router-link>
      </div>
      <div class="menu-content menu-content--orgadmin"
         v-if="activeMenu === 'orgadmin'">
        <div class="menu-item" @click="setActiveMenu('main')">
          <font-awesome-icon icon="arrow-left" class="menu-icon" />
        </div>
        <div class="menu-item">
          <font-awesome-icon icon="university" class="menu-icon" />
        </div>
        <div class="menu-item">
          <font-awesome-icon icon="list-ul" class="menu-icon" />
        </div>
        <div class="menu-item">
          <font-awesome-icon icon="users" class="menu-icon" />
        </div>
        <div class="menu-item">
          <font-awesome-icon icon="user-graduate" class="menu-icon" />
        </div>
        <div class="menu-item">
          <font-awesome-icon icon="user-alt" class="menu-icon" />
        </div>
        <div class="menu-item">
          <font-awesome-icon icon="vote-yea" class="menu-icon" />
        </div>
        <div class="menu-item">
          <font-awesome-icon icon="history" class="menu-icon" />
        </div>
      </div>
    </div>
    <div class="content">
      <router-view/>
    </div>
  </div>
</template>

<script>

export default {
  name: 'Base',
  data() {
    return {
      activeMenu: this.$router.currentRoute.name === 'orgadmin' ? 'orgadmin' : 'main',
    };
  },
  methods: {
    setActiveMenu(menuName) {
      this.activeMenu = menuName;
    },
  },
};
</script>

<style lang="scss" scoped>
@import "../assets/scss/colors";

.layout {
  min-height: 100vh;
  display: grid;
  grid-template-rows: 70px auto;
  grid-template-columns: 70px auto;

  .navbar {
    grid-row-start: 1;
    grid-column-start: 2;
    background: #fff;
  }

  .menu {
    display: flex;
    flex-direction: column;
    justify-content: flex-start;
    align-items: stretch;
    grid-row-start: 1;
    grid-row-end: span 2;
    grid-column-start: 1;
    background: #fff;
    border-right: 1.5px solid $neutralLighter;

    &-content {
      display: flex;
      flex-direction: column;
      justify-content: flex-start;
      align-items: stretch;
    }

    &-action {
      height: 70px;
      display: flex;
      justify-content: center;
      align-items: center;

      &:hover {
        cursor: pointer;

        .menu-icon { color: $primary; }

      }

    }

    &-item {
      @extend .menu-action;
      border-right: 4px solid transparent;
      border-left: 4px solid transparent;
      transition: border 0.2s ease-in-out;

      &:hover { border-left: 4px solid $primary; }

      &--active {
        border-left: 4px solid $primary;

        .menu-icon { color: $primary; }

      }

    }

    &-icon {
      color: $neutralDarker;
      transition: color 0.2s ease-in-out;
    }

  }

  .content {
    grid-row-start: 2;
    grid-column-start: 2;
    background: $neutralLightest;
  }

}
</style>
